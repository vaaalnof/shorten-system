package worker

import (
	"context"

	"shortener-service/internal/entity"

	natsinfra "shortener-service/internal/infra/nats"
)

type AnalyticsProcessor interface {
	Process(
		context.Context,
		*entity.AnalyticsEvent,
	) error
}

type AnalyticsWorker struct {
	consumer  *natsinfra.Consumer
	processor AnalyticsProcessor
}

func NewAnalyticsWorker(
	consumer *natsinfra.Consumer,
	processor AnalyticsProcessor,
) *AnalyticsWorker {

	return &AnalyticsWorker{
		consumer:  consumer,
		processor: processor,
	}
}

func (w *AnalyticsWorker) Run(
	ctx context.Context,
) error {

	return w.consumer.Consume(
		ctx,
		w.processor.Process,
	)
}
