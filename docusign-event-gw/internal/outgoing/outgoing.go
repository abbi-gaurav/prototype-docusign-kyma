package outgoing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/config"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/logger"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/model/events"
	"net/http"
	"time"
)

type EventForwarder struct {
	eventPublishURL *string
	client          *http.Client
}

func NewEventForwarder() *EventForwarder {
	client := &http.Client{
		Transport: &http.Transport{
			DisableCompression:  false,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 0,
			MaxConnsPerHost:     0,
			IdleConnTimeout:     30 * time.Second,
		},
	}

	return &EventForwarder{
		eventPublishURL: config.GlobalConfig.EventPublishURL,
		client:          client,
	}
}

func (e *EventForwarder) Forward(event *events.KymaEvent) (error) {
	eventBytes, err := json.Marshal(event.Data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, *e.eventPublishURL, bytes.NewReader(eventBytes))
	if err != nil {
		return err
	}
	req = e.enrichRequest(event, req)

	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("unexpected response when publishing event %d (%s)", resp.StatusCode, resp.Status)
		logger.Logger.Error(errMsg)
		return fmt.Errorf(errMsg)
	}

	return nil

}

func (e *EventForwarder) enrichRequest(event *events.KymaEvent, request *http.Request) *http.Request {
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("ce-specversion", "1.0")
	request.Header.Set("ce-type", event.EventType)
	request.Header.Set("ce-eventtypeversion", event.EventTypeVersion)
	request.Header.Set("ce-id", event.EventID)
	request.Header.Set("ce-source", *event.SourceID)
	return request
}
