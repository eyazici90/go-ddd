package mediator

import "context"

type PipelineBehaviour interface {
	Process(context.Context, interface{}, Next) error
}

type pipelineBuilder interface {
	UseBehaviour(PipelineBehaviour) Mediator
	Use(call func(context.Context, interface{}, Next) error) Mediator
}

func (m *reflectBasedMediator) UseBehaviour(pipelineBehaviour PipelineBehaviour) Mediator {
	return m.Use(pipelineBehaviour.Process)
}

func (m *reflectBasedMediator) Use(call func(context.Context, interface{}, Next) error) Mediator {
	if m.behaviour == nil {
		m.behaviour = m.send
	}
	seed := m.behaviour

	m.behaviour = func(ctx context.Context, msg interface{}) error {
		return call(ctx, msg, func(context.Context) error { return seed(ctx, msg) })
	}

	return m
}
