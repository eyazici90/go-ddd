package mediator

type publisher interface {
	Publish(msg interface{})
}

func (m *reflectBasedMediator) Publish(msg interface{}) {
	//
}
