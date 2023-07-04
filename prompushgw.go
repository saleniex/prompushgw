package prompushgw

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type PushGateway struct {
	baseUrl     string
	jobName     string
	HttpTimeout time.Duration
	httpClient  *http.Client
}

func NewPushGateway(baseUrl, jobName string) *PushGateway {
	return &PushGateway{
		baseUrl:     baseUrl,
		jobName:     jobName,
		HttpTimeout: 10,
	}
}

// Push pushes key-value list to Prometheus gateway
func (g *PushGateway) Push(keyValList *KeyValueList) error {
	url := fmt.Sprintf("%s/metrics/job/%s", g.baseUrl, g.jobName)
	body := keyValList.AsText()
	request, reqErr := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if reqErr != nil {
		return reqErr
	}

	response, doErr := g.getHttpClient().Do(request)
	if doErr != nil {
		return doErr
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with non success %d", response.StatusCode)
	}

	return nil
}

func (g *PushGateway) getHttpClient() *http.Client {
	if g.httpClient == nil {
		g.httpClient = &http.Client{
			Timeout: g.HttpTimeout * time.Second,
		}
	}
	return g.httpClient
}
