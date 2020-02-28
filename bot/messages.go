package bot

import (
	"fmt"
	"github.com/alexcom/tba/telegram"
	"github.com/sirupsen/logrus"
)

type MessageLevel uint8

const (
	_ MessageLevel = iota
	INFO
	WARN
	ERROR
)

func (bot Bot) Report(chatId int, level MessageLevel, msg string) {
	req := telegram.SendMessageRequest{}
	req.ChatID = chatId
	req.Text = fmt.Sprintf("%s %s", icon(level), msg)
	if _, err := bot.Telegram.SendMessage(req); err != nil {
		logrus.WithError(err).Error("reporting")
	}
}

func icon(level MessageLevel) string {
	switch level {
	case INFO:
		return "ℹ"
	case WARN:
		return "⚠"
	case ERROR:
		return "‼"
	default:
		logrus.Errorf("unexpected Message Level %d", level)
		return ""
	}
}
