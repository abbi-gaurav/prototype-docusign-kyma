package incoming

import (
	"encoding/xml"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/logger"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/model/events"
)

func Process(requestBody []byte) (*events.KymaEvent, error) {
	envelope, err := to(requestBody)
	if err != nil {
		return nil, err
	}

	kymaEvent := events.Map(envelope)
	logger.Logger.Infow("kyma Event", "kyma-event", kymaEvent)

	return kymaEvent, nil
}

func to(requestBody []byte) (*events.Envelope, error) {
	env := events.Envelope{}
	err := xml.Unmarshal(requestBody, &env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}
