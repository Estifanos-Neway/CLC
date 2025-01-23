package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	httpclient "github.com/estifanos-neway/CLC/pkg/http-client"
)

func (m *Message) Send() (*Content, error) {
	url, err := url.Parse(m.Url)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	query.Set("key", m.ApiKey)
	url.RawQuery = query.Encode()

	msgJson, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewReader(msgJson)

	req, err := http.NewRequest(http.MethodPost, url.String(), reqBody)
	if err != nil {
		return nil, err
	}

	client := httpclient.Get(30 * time.Second)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error from ai model : %s", res.Status)
	}

	var content Content
	if err := json.NewDecoder(res.Body).Decode(&content); err != nil {
		return nil, err
	}

	return &content, nil
}
