package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"text/template"
)

const templates = `
{{define "bot_url"}}https://api.telegram.org/bot{{.Token}}/{{.Method}}{{end}}
{{define "file_dl_url"}}https://api.telegram.org/file/bot{{.Token}}/{{.Path}}{{end}}
`

func NewClient(apiToken string, longPollingTimeout int) *BaseClient {
	return &BaseClient{
		tmpl:               template.Must(template.New("templates").Parse(templates)),
		httpClient:         http.Client{},
		apiToken:           apiToken,
		longPollingTimeout: longPollingTimeout,
	}
}

type BaseClient struct {
	tmpl               *template.Template
	httpClient         http.Client
	apiToken           string
	longPollingTimeout int
}

type params struct {
	Token  string
	Method string
}

func (c BaseClient) makeUrl(method string) (string, error) {
	urlBuffer := bytes.Buffer{}
	err := c.tmpl.ExecuteTemplate(&urlBuffer, "bot_url", params{
		Token:  c.apiToken,
		Method: method,
	})
	return urlBuffer.String(), err
}

func toJson(request interface{}) (io.Reader, error) {
	reqBody := bytes.Buffer{}
	if request != nil {
		err := json.NewEncoder(&reqBody).Encode(request)
		return &reqBody, err
	}
	return &reqBody, nil
}

// makeRequest makes requests of method type with customizable request and result types
func (c BaseClient) makeRequest(method string, request interface{}, result interface{}) (interface{}, error) {
	url, err := c.makeUrl(method)
	if err != nil {
		return nil, err
	}
	reqBody, err := toJson(request)
	contentType := "application/json"
	ret, err := c.doPostRequest(url, contentType, reqBody, result)
	return ret, err
}

func makeFormBody(formData *map[string]io.Reader) (io.Reader, string, error) {
	formBytes := bytes.Buffer{}
	formWriter := multipart.NewWriter(&formBytes)
	for fieldName, reader := range *formData {
		var writer io.Writer
		var err error
		if reader, ok := reader.(*os.File); ok {
			writer, err = formWriter.CreateFormFile(fieldName, reader.Name())
		} else {
			writer, err = formWriter.CreateFormField(fieldName)
		}
		if err != nil {
			return nil, "", err
		}
		if _, err := io.Copy(writer, reader); err != nil {
			return nil, "", err
		}
	}
	if err := formWriter.Close(); err != nil {
		logrus.WithError(err).Warn("closing resource")
	}
	return &formBytes, formWriter.FormDataContentType(), nil
}

func (c BaseClient) doFormRequest(
	method string, formData *map[string]io.Reader, result interface{}) (ret interface{}, err error) {

	url, err := c.makeUrl(method)
	if err != nil {
		return nil, err
	}

	bodyReader, contentType, err := makeFormBody(formData)
	if err != nil {
		return nil, err
	}

	ret, err = c.doPostRequest(url, contentType, bodyReader, result)
	return ret, err
}

type ResponseWrapper struct {
	Ok          bool               `json:"ok"`
	ErrorCode   int                `json:"error_code"`
	Parameters  ResponseParameters `json:"parameters"`
	Description string             `json:"description"`
	Result      interface{}        `json:"result"`
}

func (c BaseClient) doPostRequest(
	url string, contentType string, bodyReader io.Reader, result interface{}) (interface{}, error) {

	resp, err := c.httpClient.Post(url, contentType, bodyReader)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		logrus.Infof("got response code %d", resp.StatusCode)
	}
	defer closeOrWarn(resp.Body)

	var respWrapper ResponseWrapper
	if result != nil {
		respWrapper = ResponseWrapper{
			Result: result,
		}
	}
	// checking ErrorCode for zero to allow long polling pass through with Ok==false
	if respWrapper.ErrorCode != 0 && !respWrapper.Ok {
		return nil, fmt.Errorf("status code %d, error: %s", respWrapper.ErrorCode, respWrapper.Description)
	}
	err = json.NewDecoder(resp.Body).Decode(&respWrapper)
	if err != nil {
		return nil, err
	}
	if !respWrapper.Ok {
		return nil, fmt.Errorf(
			"response status is not OK: %d - %s", respWrapper.ErrorCode, respWrapper.Description)
	}
	return respWrapper.Result, err
}

func (c BaseClient) DownloadFile(filePath string) ([]byte, error) {
	url, err := c.makeFileURL(filePath)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer closeOrReport(resp.Body, "response body")
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("DownloadFile expected http 200 got %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

type fileParams struct {
	Token string
	Path  string
}

func (c BaseClient) makeFileURL(path string) (string, error) {
	urlBuffer := bytes.Buffer{}
	err := c.tmpl.ExecuteTemplate(&urlBuffer, "file_dl_url", fileParams{
		Token: c.apiToken,
		Path:  path,
	})
	return urlBuffer.String(), err
}

func closeOrWarn(closer io.Closer) {
	if err := closer.Close(); err != nil {
		logrus.WithError(err).Warn("closing resource")
	}
}

func closeOrReport(closer io.Closer, name string) {
	if err := closer.Close(); err != nil {
		logrus.Errorf("%s closing \n", name)
	}
}
