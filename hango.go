package hango

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Client is Webhook client for Hangout chat.
type Client struct {
	httpClient *http.Client
	webhook    string
}

// PostData represents post data for Google chat.
type PostData struct {
	Text   string  `json:"text"`
	Thread *Thread `json:"thread"`
}

// Thread represents thread.
type Thread struct {
	Name string `json:"name"`
}

// NewClient create new Webhook client for Hangout chat.
func NewClient(webhook string) *Client {
	c := &Client{
		httpClient: &http.Client{},
		webhook:    webhook,
	}

	return c
}

// Post message with a new thread.
func (c *Client) Post(message string) ([]byte, error) {
	return c.post(message, "")
}

// PostToThread post massage to an existing thread.
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
