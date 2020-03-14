package bot

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Status struct {
	lastUpdate int
	changed    bool
}

type jsonStatus struct {
	LastUpdate int `json:"last_update"`
}

const statusFileName = "status.json"

func LoadStatus(workingDir string) (*Status, error) {
	filename := filepath.Join(workingDir, statusFileName)
	logrus.Infof("loading status from file: %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			logrus.Info("no status file, will create new one")
			return &Status{}, nil
		}
		return nil, err
	}
	decoder := json.NewDecoder(file)
	jStatus := jsonStatus{}
	err = decoder.Decode(&jStatus)
	if err != nil {
		return nil, err
	}
	logrus.Info("status reading success")
	return &Status{
		lastUpdate: jStatus.LastUpdate,
	}, nil
}

func (s Status) Changed() bool {
	return s.changed
}

func (s *Status) SetUpdate(update int) {
	if s.lastUpdate == update {
		return
	}
	if s.lastUpdate > update {
		logrus.Warnf("new update number is %d, but last update is %d, ignoring\n", update, s.lastUpdate)
		return
	}
	s.lastUpdate = update
	s.changed = true
}

func (s Status) LastUpdate() int {
	return s.lastUpdate
}

func (s *Status) Save() error {
	if !s.changed {
		return nil
	}
	file, err := os.OpenFile(statusFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(&jsonStatus{LastUpdate: s.lastUpdate})
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	s.changed = false
	logrus.Info("status saving success")
	return nil
}
