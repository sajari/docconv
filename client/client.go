// Package client defines types and functions for interacting with
// docconv HTTP servers.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// DefaultProtocol is the default protocol used to construct paths
// when making docconv requests.
const DefaultProtocol = "http://"

// DefaultEndpoint is the default endpoint address (host:port) for
// docconv HTTP servers.
const DefaultEndpoint = "localhost:8888"

// DefaultHTTPClient is the default HTTP client used to make requests
// to docconv HTTP servers.
var DefaultHTTPClient = http.DefaultClient

// Opt is an option used in New to create Clients.
type Opt func(*Client)

// WithEndpoint set the endpoint on a Client.
func WithEndpoint(endpoint string) Opt {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithHTTPClient sets the *http.Client used for all underlying
// calls.
func WithHTTPClient(client *http.Client) Opt {
	return func(c *Client) {
		c.httpClient = client
	}
}

// WithProtocol sets the protocol used in HTTP requests.  Currently this
// must be either http:// or https://.
func WithProtocol(protocol string) Opt {
	return func(c *Client) {
		c.protocol = protocol
	}
}

// New creates a new docconv client for interacting with a docconv HTTP
// server.
func New(opts ...Opt) *Client {
	c := &Client{
		endpoint:   DefaultEndpoint,
		protocol:   DefaultProtocol,
		httpClient: DefaultHTTPClient,
	}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Client is a docconv HTTP client.  Use New to make new Clients.
type Client struct {
	endpoint   string
	protocol   string
	httpClient *http.Client
}

// Response is from docconv.Response copied here to avoid dependency on
// the docconv package.
type Response struct {
	Body  string            `json:"body"`
	Meta  map[string]string `json:"meta"`
	MSecs uint32            `json:"msecs"`
	Error string            `json:"error"`
}

// Convert a file from a local path using the http client
func (c *Client) Convert(r io.Reader, filename string) (*Response, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	part, err := w.CreateFormFile("input", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, r); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v%v/convert", c.protocol, c.endpoint), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// ConvertPath uses the docconv Client to convert the local file
// found at path.
func ConvertPath(c *Client, path string) (*Response, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return c.Convert(f, f.Name())
}
