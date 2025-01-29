package gemini

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	httpclient "github.com/estifanos-neway/CLC/internal/pkg/http-client"
)

func (m *Message) Send() (*Response, error) {
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
		errMsg := fmt.Sprintf("error from the ai api : %s", res.Status)
		slog.Debug(errMsg, "req", string(msgJson)) // TODO Add "res"
		return nil, errors.New(errMsg)
	}

	var messageResponse Response
	if err := json.NewDecoder(res.Body).Decode(&messageResponse); err != nil {
		return nil, err
	}

	return &messageResponse, nil
}

func CreateContent(role Role, text string) *Content {
	return &Content{
		Role: role,
		Parts: []*Part{
			{
				Text: text,
			},
		},
	}
}
