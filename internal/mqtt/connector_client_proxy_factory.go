package mqtt

import (
	"context"

	"github.com/RedHatInsights/cloud-connector/internal/config"
	"github.com/RedHatInsights/cloud-connector/internal/controller"
	"github.com/RedHatInsights/cloud-connector/internal/domain"
	"github.com/RedHatInsights/cloud-connector/internal/platform/logger"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

type ConnectorClientMQTTProxyFactory struct {
	mqttClient   MQTT.Client
	topicBuilder *TopicBuilder
	config       *config.Config
}

func NewConnectorClientMQTTProxyFactory(cfg *config.Config, mqttClient MQTT.Client, topicBuilder *TopicBuilder) (controller.ConnectorClientProxyFactory, error) {
	proxyFactory := ConnectorClientMQTTProxyFactory{mqttClient: mqttClient, topicBuilder: topicBuilder, config: cfg}
	return &proxyFactory, nil
}

func (ccpf *ConnectorClientMQTTProxyFactory) CreateProxy(ctx context.Context, account domain.AccountID, client_id domain.ClientID, dispatchers domain.Dispatchers) (controller.ConnectorClient, error) {

	logger := logger.Log.WithFields(logrus.Fields{"account": account, "client_id": client_id})

	proxy := ConnectorClientMQTTProxy{
		Logger:       logger,
		Config:       ccpf.config,
		AccountID:    account,
		ClientID:     client_id,
		Client:       ccpf.mqttClient,
		TopicBuilder: ccpf.topicBuilder,
		Dispatchers:  dispatchers,
	}

	return &proxy, nil
}
