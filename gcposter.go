package gcposter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	webhook    string
}

type PostData struct {
	Text   string  `json:"text"`
	Thread *Thread `json:"thread"`
}

type Thread struct {
	Name string `json:"name"`
}

func NewClient(webhook string) *Client {
	c := &Client{
		httpClient: &http.Client{},
		webhook:    webhook,
	}

	return c
}

func (c *Client) Post(message string) ([]byte, error) {
	return c.post(message, "")
}

func (c *Client) PostToThread(message, thread string) ([]byte, error) {
	return c.post(message, thread)
}

func (c *Client) post(message, thread string) ([]byte, error) {
	// Create post data
	j := PostData{message, &Thread{thread}}
	jsonStr, _ := json.Marshal(j)

	req, err := http.NewRequest(
		http.MethodPost,
		c.webhook,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return nil, err
	}

	// Set Content-Type.
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
