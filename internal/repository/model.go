package repository

import (
	"context"

	"github.com/puriice/pProject/internal/types"
)

type ProjectRepository interface {
	CreateProject(context context.Context, payload *types.ProjectPayload) (*types.Project, error)
	QueryProjectByID(context context.Context, id string) (*types.Project, error)
	QueryProjectByName(context context.Context, name string) (*types.Project, error)
	UpdateProject(context context.Context, id string, payload *types.ProjectPayload) error
	DeleteProject(context context.Context, id string) error
}
