package httpclient

import (
	"net/http"
	"sync"
	"time"
)

var (
	once   sync.Once
	client *http.Client
)

func Get(timeout time.Duration) *http.Client {
	once.Do(func() {
		client = &http.Client{
			Timeout: timeout,
		}
	})
	return client
}
