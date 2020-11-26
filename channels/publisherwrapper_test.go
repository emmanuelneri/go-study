package channels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapperListeningPublishedMessage(t *testing.T) {
	publisherWrapper := newPublisherWrapper()
	publisherWrapper.setPublisher(newAsyncPublisher())

	msgs := make([]string, 0, 4)
	done := make(chan struct{})
	go func() {
		total := 0
		for m := range publisherWrapper.Results() {
			msgs = append(msgs, m)

			total++
			if total == 4 {
				done <- struct{}{}
			}
		}
	}()

	publisherWrapper.Publish("test 1")
	publisherWrapper.Publish("test 2")
	publisherWrapper.Publish("test 3")
	publisherWrapper.Publish("test 4")
	<-done

	assert.ElementsMatch(t, []string{"test 1", "test 2", "test 3", "test 4"}, msgs)
}

func TestChangePublisherInstanceNotImpactListingMessages(t *testing.T) {
	publisherWrapper := newPublisherWrapper()
	msgs := make([]string, 0, 4)
	done := make(chan struct{})
	startListenResult := make(chan struct{})
	go func() {
		total := 0
		startListenResult <- struct{}{}
		for m := range publisherWrapper.Results() {
			msgs = append(msgs, m)

			total++
			if total == 4 {
				done <- struct{}{}
			}
		}
	}()
	<-startListenResult
	publisherWrapper.setPublisher(newAsyncPublisher())
	publisherWrapper.Publish("test 1")
	publisherWrapper.Publish("test 2")
	publisherWrapper.setPublisher(newAsyncPublisher())
	publisherWrapper.Publish("test 3")
	publisherWrapper.Publish("test 4")
	<-done

	assert.ElementsMatch(t, []string{"test 1", "test 2", "test 3", "test 4"}, msgs)
}
