package sdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/puriice/golibs/pkg/messaging"
	"github.com/puriice/pProject/pkg/model"
)

const (
	slashAPISlashVersionOneSlashProjectLength = 16
	slashIDSlashUUIDLength                    = 40
)

type ProjectService struct {
	url      string
	broker   *messaging.RabbitMQ
	listener *messaging.RabbitListener
}

const (
	ProjectCreate = "project.create"
	ProjectUpdate = "project.update"
	ProjectDelete = "project.delete"
	ExchangeName  = "project.events"
)

var (
	ProjectEvents       = [...]string{ProjectCreate, ProjectUpdate, ProjectDelete}
	ErrUnknownEvent     = errors.New("Unknown event")
	ErrBrokerNotDefined = errors.New("Broker is nil")
)

type ProjectEvent struct {
	*model.Project
	EventType string `json:"event"`
}

func NewService(httpURL string, broker *messaging.RabbitMQ) *ProjectService {
	var sb strings.Builder

	sb.Grow(len(httpURL) + slashAPISlashVersionOneSlashProjectLength)
	sb.WriteString(httpURL)

	if !strings.HasSuffix(httpURL, "/") {
		sb.WriteRune('/')
	}

	sb.WriteString("api/v1/projects")

	return &ProjectService{
		url:    sb.String(),
		broker: broker,
	}
}

func (s *ProjectService) GetProjectInfo(id string) (*model.Project, error) {
	var url strings.Builder

	url.Grow(len(s.url) + slashIDSlashUUIDLength)

	url.WriteString(s.url)
	url.WriteString("/id/")
	url.WriteString(id)

	resp, err := http.Get(url.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var project *model.Project = new(model.Project)

	err = json.NewDecoder(resp.Body).Decode(project)

	return project, err
}
