package sdk_test

import (
	"testing"

	project "github.com/puriice/pProject/pkg/sdk"
)

func TestGetProjectInfo(t *testing.T) {
	projectService := project.NewService("http://localhost:8081", nil)

	project, err := projectService.GetProjectInfo("b5e50b5a-9234-44e3-af27-054b88b20b3a")

	if err != nil {
		t.Error(err)
	}

	t.Log(project)
}
