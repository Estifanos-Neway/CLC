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
	resBodyDecoder := json.NewDecoder(res.Body)
	if res.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("error from the ai api : %s", res.Status)

		var errRes ErrorResponse
		if err := resBodyDecoder.Decode(&errRes); err != nil {
			return nil, err
		}

		slog.Debug(errMsg, "req", m, "res", errRes)
		return nil, errors.New(errMsg)
	}

	var messageResponse Response
	if err := resBodyDecoder.Decode(&messageResponse); err != nil {
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
