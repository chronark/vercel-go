package project_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/chronark/vercel-go/api"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/chronark/vercel-go/endpoints/project"
	"github.com/chronark/vercel-go/utils/mocks"
)

/* Run all tests on a personal account an on a team account */
type testCase struct {
	name    string
	handler *project.ProjectHandler
}

var (
	token string
	tests []testCase
)

func init() {
	token = os.Getenv("VERCEL_TOKEN")
	if token == "" {
		panic("VERCEL_TOKEN must be defined")
	}
	teamid := os.Getenv("VERCEL_TEAM_ID")
	if teamid == "" {
		panic("VERCEL_TEAM_ID must be defined")
	}

	tests = []testCase{
		{name: "personal", handler: project.New(api.New(api.NewClientConfig{Token: token}), "")},
		{name: "team", handler: project.New(api.New(api.NewClientConfig{Token: token}), teamid)},
	}
}

func TestListProjectsWithResponse400(t *testing.T) {
	client := &mocks.MockHTTPClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			res := &http.Response{}
			res.StatusCode = 400
			res.Body = io.NopCloser(strings.NewReader(("Hello World")))
			return res, nil
		},
	}
	handler := project.New(api.New(api.NewClientConfig{HTTPClient: client}), "")

	_, err := handler.List(project.ListProjectsRequest{})

	require.Error(t, err)

}
func TestListProjects(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()

			//Create at least 1 project and delete it afterwards
			defer func() {
				_, err := tt.handler.Delete(name)
				require.NoError(t, err)
			}()
			_, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: name,
				},
			)
			require.NoError(t, err)

			res, err := tt.handler.List(project.ListProjectsRequest{})

			require.NoError(t, err)
			require.True(t, len(res.Projects) > 0)
		})
	}
}

func TestGetProjectWithName(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()
			//Create at least 1 project and delete it afterwards
			defer func() {
				_, err := tt.handler.Delete(name)
				require.NoError(t, err)
			}()
			createRes, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: name,
				},
			)
			require.NoError(t, err)
			res, err := tt.handler.Get(
				createRes.Id,
			)
			require.NoError(t, err)
			require.True(t, res.Name == name)
		})
	}
}

func TestGetProjectWithid(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()

			defer func() {
				_, err := tt.handler.Delete(name)
				require.NoError(t, err)
			}()
			createRes, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: name,
				},
			)
			require.NoError(t, err)
			res, err := tt.handler.Get(createRes.Id)
			require.NoError(t, err)
			require.True(t, res.Name == name)
		})
	}
}

func TestCreateProject(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()
			var projectId string
			defer func() {
				_, err := tt.handler.Delete(projectId)
				require.NoError(t, err)
			}()
			res, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: name,
				},
			)
			projectId = res.Id
			require.NoError(t, err)
			foundRes, err := tt.handler.Get(res.Id)
			require.NoError(t, err)
			require.Equal(t, res.Id, foundRes.Id)
		})
	}
}

func TestUpdateProject(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()
			var projectId string
			defer func() {
				_, err := tt.handler.Delete(projectId)
				require.NoError(t, err)
			}()
			res, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: name,
				},
			)
			projectId = res.Id
			require.NoError(t, err)
			foundRes, err := tt.handler.Get(projectId)
			require.NoError(t, err)
			require.Equal(t, res.Id, foundRes.Id)

			updateRes, err := tt.handler.Update(projectId, project.UpdateProjectRequest{

				Framework: "blitzjs",
			})
			require.NoError(t, err)
			require.Equal(t, "blitzjs", updateRes.Framework)

		})
	}
}

func TestCreateEnvironmentVariable(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projectName := uuid.NewString()
			envKey := fmt.Sprintf("a%s", strings.ReplaceAll(uuid.NewString(), "-", ""))
			envValue := uuid.NewString()
			var projectId string
			defer func() {
				_, err := tt.handler.Delete(projectId)
				require.NoError(t, err)
			}()
			res, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: projectName,
				},
			)
			projectId = res.Id
			require.NoError(t, err)

			_, err = tt.handler.CreateEnvironmentVariable(projectId, project.CreateEnvironmentVariableRequest{

				Key:    envKey,
				Value:  envValue,
				Type:   "encrypted",
				Target: []string{"production"},
			})
			require.NoError(t, err)

			foundRes, err := tt.handler.Get(projectId)
			require.NoError(t, err)
			require.Equal(t, len(foundRes.Env), 1)
			require.Equal(t, envKey, foundRes.Env[0].Key)

		})
	}
}

func TestDeleteEnvironmentVariable(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projectName := uuid.NewString()
			envKey := fmt.Sprintf("a%s", strings.ReplaceAll(uuid.NewString(), "-", ""))
			envValue := uuid.NewString()
			var projectId string
			defer func() {
				_, err := tt.handler.Delete(projectId)
				require.NoError(t, err)
			}()
			res, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: projectName,
				},
			)
			projectId = res.Id
			require.NoError(t, err)

			envRes, err := tt.handler.CreateEnvironmentVariable(projectId, project.CreateEnvironmentVariableRequest{

				Key:    envKey,
				Value:  envValue,
				Type:   "encrypted",
				Target: []string{"production"},
			})
			require.NoError(t, err)
			_, err = tt.handler.DeleteEnvironmentVariable(project.DeleteEnvironmentVariableRequest{
				Projectid: projectId,
				Envid:     envRes.Id,
			})
			require.NoError(t, err)

			foundRes, err := tt.handler.Get(projectId)
			require.NoError(t, err)
			require.Equal(t, len(foundRes.Env), 0)

		})
	}
}

func TestUpdateEnvironmentVariable(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projectName := uuid.NewString()
			envKey := fmt.Sprintf("a%s", strings.ReplaceAll(uuid.NewString(), "-", ""))
			envUpdateKey := fmt.Sprintf("a%s", strings.ReplaceAll(uuid.NewString(), "-", ""))
			envValue := uuid.NewString()
			var projectId string
			defer func() {
				_, err := tt.handler.Delete(projectId)
				require.NoError(t, err)
			}()
			res, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: projectName,
				},
			)
			projectId = res.Id
			require.NoError(t, err)

			envRes, err := tt.handler.CreateEnvironmentVariable(projectId, project.CreateEnvironmentVariableRequest{

				Key:    envKey,
				Value:  envValue,
				Type:   "encrypted",
				Target: []string{"production"},
			})
			require.NoError(t, err)

			_, err = tt.handler.UpdateEnvironmentVariable(projectId, project.UpdateEnvironmentVariableRequest{

				Envid: envRes.Id,
				Key:   envUpdateKey,
			})
			require.NoError(t, err)
			foundRes, err := tt.handler.Get(projectId)
			require.NoError(t, err)
			require.Equal(t, len(foundRes.Env), 1)
			require.Equal(t, envUpdateKey, foundRes.Env[0].Key)

		})
	}
}

func TestAddDomain(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domain := fmt.Sprintf("www.%s.com", strings.ReplaceAll(uuid.NewString(), "-", ""))
			var projectId string
			defer func() {
				_, err := tt.handler.Delete(projectId)
				require.NoError(t, err)
			}()
			res, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: uuid.NewString(),
				},
			)
			projectId = res.Id
			require.NoError(t, err)

			_, err = tt.handler.AddDomain(projectId, project.AddDomainRequest{
				Name: domain,
			})
			require.NoError(t, err)

			domainRes, err := tt.handler.GetDomain(projectId, domain)
			require.NoError(t, err)
			require.Equal(t, domain, domainRes.Name)

		})
	}
}

func TestUpdateDomain(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domain := fmt.Sprintf("www.%s.com", strings.ReplaceAll(uuid.NewString(), "-", ""))
			redirect := fmt.Sprintf("www.%s.com", strings.ReplaceAll(uuid.NewString(), "-", ""))

			var projectId string
			defer func() {
				_, err := tt.handler.Delete(projectId)
				require.NoError(t, err)
			}()
			res, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: uuid.NewString(),
				},
			)
			projectId = res.Id
			require.NoError(t, err)

			_, err = tt.handler.AddDomain(projectId, project.AddDomainRequest{
				Name: domain,
			})
			require.NoError(t, err)

			_, err = tt.handler.UpdateDomain(projectId, project.UpdateDomainRequest{
				Name:     domain,
				Redirect: redirect,
			})
			require.NoError(t, err)
			foundRes, err := tt.handler.GetDomain(projectId, domain)
			require.NoError(t, err)
			require.Equal(t, domain, foundRes.Name)
			require.Equal(t, redirect, foundRes.Redirect)

		})
	}
}

func TestListDomains(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var projectId string
			defer func() {
				_, err := tt.handler.Delete(projectId)
				require.NoError(t, err)
			}()
			domain := fmt.Sprintf("www.%s.com", strings.ReplaceAll(uuid.NewString(), "-", ""))
			projectName := uuid.NewString()
			res, err := tt.handler.Create(
				project.CreateProjectRequest{
					Name: projectName,
				},
			)
			projectId = res.Id
			require.NoError(t, err)

			_, err = tt.handler.AddDomain(projectId, project.AddDomainRequest{
				Name: domain,
			})
			require.NoError(t, err)

			domainRes, err := tt.handler.ListDomains(projectId, project.ListDomainsRequest{})
			require.NoError(t, err)
			require.Equal(t, 2, len(domainRes.Domains))
			require.Equal(t, domain, domainRes.Domains[0].Name)
			require.Equal(t, fmt.Sprintf("%s.vercel.app", projectName), domainRes.Domains[1].Name)

		})
	}
}
