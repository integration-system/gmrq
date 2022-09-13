package grmq

import (
	"github.com/integration-system/grmq/consumer"
	"github.com/integration-system/grmq/publisher"
)

type Observer interface {
	ClientReady()
	ClientError(err error)
	ConsumerError(consumer consumer.Consumer, err error)
	PublisherError(publisher *publisher.Publisher, err error)
	PublishingFlow(publisher *publisher.Publisher, flow bool)
	ShutdownStarted()
	ShutdownDone()
}

type NoopObserver struct {
}

func (n NoopObserver) ClientReady() {

}

func (n NoopObserver) ClientError(err error) {

}

func (n NoopObserver) ConsumerError(consumer consumer.Consumer, err error) {

}

func (n NoopObserver) ShutdownStarted() {
}

func (n NoopObserver) ShutdownDone() {

}

func (n NoopObserver) PublisherError(publisher *publisher.Publisher, err error) {

}

func (n NoopObserver) PublishingFlow(publisher *publisher.Publisher, flow bool) {

}
