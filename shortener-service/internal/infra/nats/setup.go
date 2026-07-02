package nats

import (
	gonats "github.com/nats-io/nats.go"

	"shortener-service/internal/config"

	natsutil "shortener-service/internal/utils/nats"
)

func SetupAnalytics(
	js gonats.JetStreamContext,
	settings *config.Settings,
) error {

	_, err := js.StreamInfo(
		natsutil.AnalyticsStream,
	)

	if err != nil {

		_, err = js.AddStream(
			&gonats.StreamConfig{
				Name: natsutil.AnalyticsStream,
				Subjects: []string{
					natsutil.AnalyticsSubject,
				},
				Storage:   gonats.FileStorage,
				Retention: gonats.LimitsPolicy,
				Replicas:  settings.NATS.Analytics.Replicas,
				MaxAge:    settings.NATS.Analytics.MaxAge,
			},
		)

		if err != nil {
			return err
		}
	}

	_, err = js.ConsumerInfo(
		natsutil.AnalyticsStream,
		natsutil.AnalyticsConsumer,
	)

	if err == nil {
		return nil
	}

	_, err = js.AddConsumer(
		natsutil.AnalyticsStream,

		&gonats.ConsumerConfig{
			Durable:    natsutil.AnalyticsConsumer,
			AckPolicy:  gonats.AckExplicitPolicy,
			AckWait:    settings.NATS.Analytics.AckWait,
			MaxDeliver: settings.NATS.Analytics.MaxDeliver,
		},
	)

	return err
}
