package config

import (
	"flag"
	"github.com/abbi-gaurav/prototype-docusign-kyma/docusign-event-gw/internal/logger"
)

type Opts struct {
	LogRequest       *bool
	Passwd           *string
	UserName         *string
	BasicAuthEnabled *bool
	AppName          *string
	EventPublishURL  *string
	BaseTopic        *string
}

var GlobalConfig *Opts

func ParseFlags() {
	logRequest := flag.Bool("verbose", false, "log each incoming event request")
	baEnabled := flag.Bool("basic-auth-enabled", false, "Enable basic Auth")
	userName := flag.String("username", "", "Basic Auth username")
	password := flag.String("password", "", "Basic Auth Password")
	appName := flag.String("app-name", "", "Application Name")
	eventPublishURL := flag.String("event-publish-url", "http://event-publish-service.kyma-system.svc.cluster.local:8080/v1/events", "URL to forward incoming events to Kyma Eventing")
	baseTopic := flag.String("base-topic", "docusign", "Base Topic defined in the Async API specification")
	flag.Parse()

	GlobalConfig = &Opts{
		LogRequest:       logRequest,
		BasicAuthEnabled: baEnabled,
		UserName:         userName,
		Passwd:           password,
		AppName:          appName,
		EventPublishURL:  eventPublishURL,
		BaseTopic:        baseTopic,
	}

	if *GlobalConfig.BasicAuthEnabled && (*GlobalConfig.UserName == "" || *GlobalConfig.Passwd == "") {
		logger.Logger.Panicw("invalid configuration - Missing credentials", "config", GlobalConfig)
	}

	if *GlobalConfig.AppName == "" {
		logger.Logger.Panicw("Invalid configuration - Missing APP Name", "config", GlobalConfig)
	}

	logger.Logger.Infow("App config", "config", GlobalConfig)
}
