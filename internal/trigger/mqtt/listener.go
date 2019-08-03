package listener

import (
	"fmt"

	"github.com/app-functions-sdk-go/pkg/transforms"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/edgexfoundry/app-functions-sdk-go/internal/common"
	"github.com/edgexfoundry/app-functions-sdk-go/internal/runtime"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/coredata"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/edgexfoundry/go-mod-messaging/pkg/types"
)

// Trigger implements Trigger to support MQTTListener
type Trigger struct {
	Configuration common.ConfigurationStruct
	Runtime       runtime.GolangRuntime
	logging       logger.LoggingClient
	topics        []types.TopicChannel
	EventClient   coredata.EventClient
}

// Initialize ...
func (trigger *Trigger) Initialize(logger logger.LoggingClient) error {
	adr := models.Addressable{}
	config := transforms.MqttConfig{}
	mqttthing := transforms.NewMQTTSender(logger, adr, "", "", config)
	if !mqttthing.client.IsConnected() {
		logger.Info("Connecting to mqtt server")
		if token := mqttthing.client.Connect(); token.Wait() && token.Error() != nil {
			return fmt.Errorf("Could not connect to mqtt server, drop event. Error: %s", token.Error().Error())
		}
		logger.Info("Connected to mqtt server")
	}
	var mc MQTT.Client
	topics := make(map[string]byte)
	topics["test"] = 0
	token := mc.SubscribeMultiple(topics, onIncomingDataReceived)
	if token.Wait() && token.Error() != nil {
		logger.Info(fmt.Sprintf("[Incoming listener] Stop incoming data listening. Cause:%v", token.Error()))
		return token.Error()
	}

	logger.Info("[Incoming listener] Start incoming data listening. ")

	return nil
}

func onIncomingDataReceived(client MQTT.Client, message MQTT.Message) {
	// edgexContext := &appcontext.Context{
	// 	Configuration: trigger.Configuration,
	// 	LoggingClient: trigger.logging,
	// 	// CorrelationID: msgs.CorrelationID,
	// 	EventClient:   trigger.EventClient,
	// }

	// trigger.Runtime.ProcessEvent(edgexContext, msgs)
	// if edgexContext.OutputData != nil {
	// outputEnvelope := types.MessageEnvelope{
	// 	CorrelationID: edgexContext.CorrelationID,
	// 	Payload:       edgexContext.OutputData,
	// 	ContentType:   clients.ContentTypeJSON,
	// }
	// err := trigger.client.Publish(outputEnvelope, trigger.Configuration.Binding.PublishTopic)
	// if err != nil {
	// 	trigger.logging.Error(fmt.Sprintf("Failed to publish Message to bus, %v", err))
	// }

	// trigger.logging.Trace("Published message to bus", "topic", trigger.Configuration.Binding.PublishTopic, clients.CorrelationHeader, msgs.CorrelationID)
	//}
}
