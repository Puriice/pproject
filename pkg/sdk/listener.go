package sdk

import (
	"encoding/json"

	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/pProject/pkg/model"
)

type Handler[T any] func(T)

type ProjectListener struct {
	listener *messaging.RabbitListener
	handlers map[string]Handler[*model.Project]
	onError  Handler[error]
}

func (s *ProjectService) NewListener(queueName string) (*ProjectListener, error) {
	if s.broker == nil {
		return nil, ErrBrokerNotDefined
	}

	listener, err := s.broker.NewListener(queueName, ProjectEvents[:]...)

	return &ProjectListener{
		listener: listener,
		handlers: make(map[string]Handler[*model.Project], len(ProjectEvents)),
	}, err
}

func (l *ProjectListener) OnCreate(handler Handler[*model.Project]) {
	l.handlers[ProjectCreate] = handler
}

func (l *ProjectListener) OnUpdate(handler Handler[*model.Project]) {
	l.handlers[ProjectUpdate] = handler
}

func (l *ProjectListener) OnDelete(handler Handler[string]) {
	l.handlers[ProjectDelete] = func(project *model.Project) {
		if project.ID == nil || *project.ID == "" {
			return
		}

		handler(*project.ID)
	}
}

func (l *ProjectListener) OnError(handler Handler[error]) {
	l.onError = handler
}

func (l *ProjectListener) Subscribe() error {
	return l.listener.Subscribe(func(body []byte) error {
		var event ProjectEvent
		err := json.Unmarshal(body, &event)

		if err != nil {
			l.onError(err)
			return nil
		}

		handler, ok := l.handlers[event.EventType]

		if !ok {
			l.onError(ErrUnknownEvent)
		} else {
			handler(event.Project)
		}

		return nil
	})
}
