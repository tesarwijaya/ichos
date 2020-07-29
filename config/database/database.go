package database

type (
	// Message type
	Message string

	// MessageColletion collection to hold messages
	MessageColletion []Message

	// Collector struct to hold message collection
	Collector struct {
		messages MessageColletion
	}
)

func (c *Collector) Post(m Message) {
	c.messages = append(c.messages, m)
}

func (c *Collector) Get() MessageColletion {
	return c.messages
}

// InitCollector method to initialize database collector
func InitCollector() *Collector {
	return &Collector{
		messages: []Message{},
	}
}
