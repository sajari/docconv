package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/sajari/docconv"
)

var (
	defaultParam    = "input"
	defaultEndpoint = "http://localhost:8888/convert"
)

type client struct {
	endpoint string
	param    string
	client   http.Client
}

// Returns a client that will send to the default parameters used for the docconv pkg
func NewDefaultClient() *client {
	return &client{
		endpoint: defaultEndpoint,
		param:    defaultParam,
		client:   http.Client{},
	}
}

// Returns a client using a custom endpoint to send to. The param indicates the
// name of the parameter that the file will be send as in the form
func NewClient(endpoint, param string) (*client, error) {
	if endpoint == "" {
		return nil, errors.New("You must specify an endpoint in the format domain:port")
	}
	if defaultParam == "" {
		return nil, errors.New("You must specify an input param, the default is 'input'")
	}
	return &client{
		endpoint: defaultEndpoint,
		param:    defaultParam,
		client:   http.Client{},
	}, nil
}

// Convert a file from a local path using the http client
func (c *client) Convert(path string) (*docconv.Response, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(c.param, file.Name())

	r, err := ioutil.ReadAll(file)

	_, err = part.Write(r)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonBlob, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := new(docconv.Response)
	err = json.Unmarshal(jsonBlob, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
