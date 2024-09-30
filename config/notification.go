package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/openmeterio/openmeter/openmeter/notification/webhook"
)

type WebhookConfiguration struct {
	// Timeout for registering event types in webhook provider
	EventTypeRegistrationTimeout time.Duration
	// Skip registering event types on unsuccessful attempt instead of returning with error
	SkipEventTypeRegistrationOnError bool
}

type NotificationConfiguration struct {
	Enabled  bool
	Consumer ConsumerConfiguration

	Webhook WebhookConfiguration
}

func (c NotificationConfiguration) Validate() error {
	if err := c.Consumer.Validate(); err != nil {
		return fmt.Errorf("consumer: %w", err)
	}
	return nil
}

func ConfigureNotification(v *viper.Viper) {
	ConfigureConsumer(v, "notification.consumer")
	v.SetDefault("notification.consumer.dlq.topic", "om_sys.notification_service_dlq")
	v.SetDefault("notification.consumer.consumerGroupName", "om_notification_service")
	v.SetDefault("notification.webhook.eventTypeRegistrationTimeout", webhook.DefaultRegistrationTimeout)
	v.SetDefault("notification.webhook.skipEventTypeRegistrationOnError", false)
}
