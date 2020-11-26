package channels

type PublisherWrapper struct {
	internalChan chan string
	publisher    *AsyncPublisher
}

func newPublisherWrapper() *PublisherWrapper {
	return &PublisherWrapper{
		internalChan: make(chan string),
	}
}

func (c *PublisherWrapper) setPublisher(publisher *AsyncPublisher) {
	c.publisher = publisher
	go func() {
		for m := range publisher.Results() {
			c.internalChan <- m
		}
	}()
}

func (c *PublisherWrapper) Publish(msg string) {
	c.publisher.Publish(msg)
}

func (c *PublisherWrapper) Results() <-chan string {
	return c.internalChan
}

type AsyncPublisher struct {
	msgs chan string
}

func newAsyncPublisher() *AsyncPublisher {
	return &AsyncPublisher{
		msgs: make(chan string),
	}
}

func (c *AsyncPublisher) Publish(msg string) {
	c.msgs <- msg
}

func (c *AsyncPublisher) Results() <-chan string {
	return c.msgs
}
