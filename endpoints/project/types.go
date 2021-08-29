package project

// Project houses all the information vercel offers about a project via their api
type Project struct {
	AccountID string `json:"accountId"`
	Alias     []struct {
		ConfiguredBy        string `json:"configuredBy"`
		ConfiguredChangedAt int64  `json:"configuredChangedAt"`
		CreatedAt           int64  `json:"createdAt"`
		Deployment          struct {
			Alias         []string      `json:"alias"`
			AliasAssigned int64         `json:"aliasAssigned"`
			Builds        []interface{} `json:"builds"`
			CreatedAt     int64         `json:"createdAt"`
			CreatedIn     string        `json:"createdIn"`
			Creator       struct {
				UID         string `json:"uid"`
				Email       string `json:"email"`
				Username    string `json:"username"`
				GithubLogin string `json:"githubLogin"`
			} `json:"creator"`
			DeploymentHostname string `json:"deploymentHostname"`
			Forced             bool   `json:"forced"`
			ID                 string `json:"id"`
			Meta               struct {
				GithubCommitRef         string `json:"githubCommitRef"`
				GithubRepo              string `json:"githubRepo"`
				GithubOrg               string `json:"githubOrg"`
				GithubCommitSha         string `json:"githubCommitSha"`
				GithubRepoID            string `json:"githubRepoId"`
				GithubCommitMessage     string `json:"githubCommitMessage"`
				GithubCommitAuthorLogin string `json:"githubCommitAuthorLogin"`
				GithubDeployment        string `json:"githubDeployment"`
				GithubCommitOrg         string `json:"githubCommitOrg"`
				GithubCommitAuthorName  string `json:"githubCommitAuthorName"`
				GithubCommitRepo        string `json:"githubCommitRepo"`
				GithubCommitRepoID      string `json:"githubCommitRepoId"`
			} `json:"meta"`
			Name       string `json:"name"`
			Plan       string `json:"plan"`
			Private    bool   `json:"private"`
			ReadyState string `json:"readyState"`
			Target     string `json:"target"`
			TeamID     string `json:"teamId"`
			Type       string `json:"type"`
			URL        string `json:"url"`
			UserID     string `json:"userId"`
			WithCache  bool   `json:"withCache"`
		} `json:"deployment"`
		Domain      string `json:"domain"`
		Environment string `json:"environment"`
		Target      string `json:"target"`
	} `json:"alias"`
	Analytics struct {
		ID         string `json:"id"`
		EnabledAt  int64  `json:"enabledAt"`
		DisabledAt int64  `json:"disabledAt"`
		CanceledAt int64  `json:"canceledAt"`
	} `json:"analytics"`
	AutoExposeSystemEnvs bool   `json:"autoExposeSystemEnvs"`
	BuildCommand         string `json:"buildCommand"`
	CreatedAt            int64  `json:"createdAt"`
	DevCommand           string `json:"devCommand"`
	DirectoryListing     bool   `json:"directoryListing"`
	Env                  []struct {
		Type            string      `json:"type"`
		ID              string      `json:"id"`
		Key             string      `json:"key"`
		Value           string      `json:"value"`
		Target          []string    `json:"target"`
		ConfigurationID interface{} `json:"configurationId"`
		UpdatedAt       int64       `json:"updatedAt"`
		CreatedAt       int64       `json:"createdAt"`
	} `json:"env"`
	Framework                       string `json:"framework"`
	ID                              string `json:"id"`
	InstallCommand                  string `json:"installCommand"`
	Name                            string `json:"name"`
	NodeVersion                     string `json:"nodeVersion"`
	OutputDirectory                 string `json:"outputDirectory"`
	PublicSource                    bool   `json:"publicSource"`
	RootDirectory                   string `json:"rootDirectory"`
	ServerlessFunctionRegion        string `json:"serverlessFunctionRegion"`
	SourceFilesOutsideRootDirectory bool   `json:"sourceFilesOutsideRootDirectory"`
	UpdatedAt                       int64  `json:"updatedAt"`
	Link                            struct {
		Type             string        `json:"type"`
		Repo             string        `json:"repo"`
		RepoID           int64         `json:"repoId"`
		Org              string        `json:"org"`
		GitCredentialID  string        `json:"gitCredentialId"`
		CreatedAt        int64         `json:"createdAt"`
		UpdatedAt        int64         `json:"updatedAt"`
		Sourceless       bool          `json:"sourceless"`
		ProductionBranch string        `json:"productionBranch"`
		DeployHooks      []interface{} `json:"deployHooks"`
		ProjectName      string        `json:"projectName"`
		ProjectNamespace string        `json:"projectNamespace"`
		Owner            string        `json:"owner"`
		Slug             string        `json:"slug"`
	} `json:"link"`
	LatestDeployments []struct {
		Alias         []string      `json:"alias"`
		AliasAssigned int64         `json:"aliasAssigned"`
		Builds        []interface{} `json:"builds"`
		CreatedAt     int64         `json:"createdAt"`
		CreatedIn     string        `json:"createdIn"`
		Creator       struct {
			UID         string `json:"uid"`
			Email       string `json:"email"`
			Username    string `json:"username"`
			GithubLogin string `json:"githubLogin"`
		} `json:"creator"`
		DeploymentHostname string `json:"deploymentHostname"`
		Forced             bool   `json:"forced"`
		ID                 string `json:"id"`
		Meta               struct {
			GithubCommitRef         string `json:"githubCommitRef"`
			GithubRepo              string `json:"githubRepo"`
			GithubOrg               string `json:"githubOrg"`
			GithubCommitSha         string `json:"githubCommitSha"`
			GithubCommitAuthorLogin string `json:"githubCommitAuthorLogin"`
			GithubCommitMessage     string `json:"githubCommitMessage"`
			GithubRepoID            string `json:"githubRepoId"`
			GithubDeployment        string `json:"githubDeployment"`
			GithubCommitOrg         string `json:"githubCommitOrg"`
			GithubCommitAuthorName  string `json:"githubCommitAuthorName"`
			GithubCommitRepo        string `json:"githubCommitRepo"`
			GithubCommitRepoID      string `json:"githubCommitRepoId"`
		} `json:"meta"`
		Name       string      `json:"name"`
		Plan       string      `json:"plan"`
		Private    bool        `json:"private"`
		ReadyState string      `json:"readyState"`
		Target     interface{} `json:"target"`
		TeamID     string      `json:"teamId"`
		Type       string      `json:"type"`
		URL        string      `json:"url"`
		UserID     string      `json:"userId"`
		WithCache  bool        `json:"withCache"`
	} `json:"latestDeployments"`
	Targets struct {
		Production struct {
			Alias         []string      `json:"alias"`
			AliasAssigned int64         `json:"aliasAssigned"`
			Builds        []interface{} `json:"builds"`
			CreatedAt     int64         `json:"createdAt"`
			CreatedIn     string        `json:"createdIn"`
			Creator       struct {
				UID         string `json:"uid"`
				Email       string `json:"email"`
				Username    string `json:"username"`
				GithubLogin string `json:"githubLogin"`
			} `json:"creator"`
			DeploymentHostname string `json:"deploymentHostname"`
			Forced             bool   `json:"forced"`
			ID                 string `json:"id"`
			Meta               struct {
				GithubCommitRef         string `json:"githubCommitRef"`
				GithubRepo              string `json:"githubRepo"`
				GithubOrg               string `json:"githubOrg"`
				GithubCommitSha         string `json:"githubCommitSha"`
				GithubRepoID            string `json:"githubRepoId"`
				GithubCommitMessage     string `json:"githubCommitMessage"`
				GithubCommitAuthorLogin string `json:"githubCommitAuthorLogin"`
				GithubDeployment        string `json:"githubDeployment"`
				GithubCommitOrg         string `json:"githubCommitOrg"`
				GithubCommitAuthorName  string `json:"githubCommitAuthorName"`
				GithubCommitRepo        string `json:"githubCommitRepo"`
				GithubCommitRepoID      string `json:"githubCommitRepoId"`
			} `json:"meta"`
			Name       string `json:"name"`
			Plan       string `json:"plan"`
			Private    bool   `json:"private"`
			ReadyState string `json:"readyState"`
			Target     string `json:"target"`
			TeamID     string `json:"teamId"`
			Type       string `json:"type"`
			URL        string `json:"url"`
			UserID     string `json:"userId"`
			WithCache  bool   `json:"withCache"`
		} `json:"production"`
	} `json:"targets"`
}

type CreateProjectRequest struct {
	// The desired name for the project.
	// Required: Yes
	Name string `json:"name"`

	// The Git Repository that will be connected to the project. When this
	// is defined, any pushes to the specified connected Git Repository
	// will be automatically deployed.
	// Required: No
	GitRepository struct {
		Type string `json:"type"`
		Repo string `json:"repo"`
	} `json:"gitRepository,omitempty"`

	UpdateProjectRequest
}
type CreateProjectResponse struct {
	Project
}

type ListProjectsRequest struct {
	// Limit the number of projects returned.
	// Required: No
	Limit int64 `json:"limit,omitempty"`

	// The updatedAt point64where the list should start.
	// Required: No
	Since int64 `json:"since,omitempty"`

	// The updatedAt point64where the list should end.
	// Required: No
	Until int64 `json:"until,omitempty"`

	// Search projects by the name field.
	// Required: No
	Search string `json:"string,omitempty"`
}
type ListProjectsResponse struct {
	Projects []Project `json:"projects"`
}

// Get the information for a specific project by passing either the project id
// or name in the URL.
type GetProjectRequest struct {
	// The unique Project identifier.
	// Required: No
	Id string `json:"id,omitempty"`

	// The project name.
	// Required: No
	Name string `json:"name,omitempty"`
}

type GetProjectResponse struct {
	Project
}

type UpdateProjectRequest struct {
	// The framework that is being used for this project. When null is used no
	// framework is selected.
	// Required: No
	Framework string `json:"framework,omitempty"`
	// Specifies whether the source code and logs of the deployments for this
	// project should be public or not.
	// Required: No
	PublicSource bool `json:"publicSource,omitempty"`
	// The install command for this project. When null is used this value will
	// be automatically detected.
	// Required: No
	InstallCommand string `json:"installCommand,omitempty"`
	// The build command for this project. When null is used this value will be
	// automatically detected.
	// Required: No
	BuildCommand string `json:"buildCommand,omitempty"`
	// The dev command for this project. When null is used this value will be
	// automatically detected.
	// Required: No
	DevCommand string `json:"devCommand,omitempty"`
	// The output directory of the project. When null is used this value will
	// be automatically detected.
	// Required: No
	OutputDirectory string `json:"outputDirectory,omitempty"`
	// The region to deploy Serverless Functions in this project.
	// Required: No
	ServerlessFunctionRegion string `json:"serverlessFunctionRegion,omitempty"`
	// The name of a directory or relative path to the source code of your
	// project. When null is used it will default to the project root.
	// Required: No
	RootDirectory string `json:"rootDirectory,omitempty"`
	// A new name for this project.
	// Required: No
	Name string `json:"name,omitempty"`
	// The Node.js Version for this project.
	// Required: No
	NodeVersion string `json:"nodeVersion,omitempty"`
	// Specifies whether PRs from Git forks should require a team member's
	// authorization before it can be deployed.
	// Required: No
	GitForkProtection bool `json:"gitFormProtection,omitempty"`
}
type UpdateProjectResponse struct {
	Project
}

type RemoveProjectRequest struct {
	// The unique Project identifier.
	// Required: No
	Id string `json:"id,omitempty"`

	// The project name.
	// Required: No
	Name string `json:"name,omitempty"`
}

// If the request is successful, you will receive a 204 HTTP Status code in the
// response.
type RemoveProjectResponse struct{}

type GetEnvironmentVariablesRequest struct {

	// The unique Project identifier.
	// Required: Yes
	Id string `json:"id"`
}
type GetEnvironmentVariablesResponse struct {
	Envs []struct {
		Type            string      `json:"type"`
		Value           string      `json:"value"`
		Target          []string    `json:"target"`
		GitBranch       string      `json:"gitBranch,omitempty"`
		ConfigurationID interface{} `json:"configurationId"`
		ID              string      `json:"id"`
		Key             string      `json:"key"`
		CreatedAt       int64       `json:"createdAt"`
		UpdatedAt       int64       `json:"updatedAt"`
		CreatedBy       string      `json:"createdBy"`
		UpdatedBy       interface{} `json:"updatedBy"`
	} `json:"envs"`
}

type CreateEnvironmentVariableRequest struct {
	// The type can be plain, encrypted, secret, or system.
	// Required: Yes
	Type string `json:"type"`

	// The name of the Environment Variable.
	// Required: Yes
	Key string `json:"key"`

	// If the type is plain or encrypted, a string representing the value of the
	// Environment Variable.
	// If the type is secret, the secret ID of the secret attached to the
	// Environment Variable.
	// If the type is system, the name of the System Environment Variable.
	// Required: Yes
	Value string `json:"value"`

	// The target can be a list of development, preview, or production.
	// Required: Yes
	Target []string `json:"target"`

	// The Git branch for this variable, only accepted when the target is
	// exclusively preview.
	// Required: No
	GitBranch string `json:"gitBranch,omitempty"`
}
type CreateEnvironmentVariableResponse struct {
	Type            string   `json:"type"`
	ID              string   `json:"id"`
	Key             string   `json:"key"`
	Value           string   `json:"value"`
	Target          []string `json:"target"`
	ConfigurationID string   `json:"configurationId,omitempty"`
	GitBranch       string   `json:"gitBranch,omitempty"`
	UpdatedAt       int64    `json:"updatedAt"`
	CreatedAt       int64    `json:"createdAt"`
}

type EditEnvironmentVariableRequest struct {
	// The type can be plain, encrypted, secret, or system.
	// Required: No
	Type string `json:"type,omitempty"`

	// The name of the Environment Variable.
	// Required: No
	Key string `json:"key,omitempty"`

	// The value of the Environment Variable.
	// Required: No
	Value string `json:"value,omitempty"`

	// The target can be a list of development, preview, or production.
	// Required: No
	Target []string `json:"target,omitempty"`

	// The Git branch for this variable, only accepted when the target is exclusively preview.
	// Required: No
	GitBranch string `json:"gitBranch,omitempty"`
}
type EditEnvironmentVariableResponse struct {
	Type            string   `json:"type"`
	ID              string   `json:"id"`
	Key             string   `json:"key"`
	Value           string   `json:"value"`
	Target          []string `json:"target"`
	ConfigurationID string   `json:"configurationId,omitempty"`
	GitBranch       string   `json:"gitBranch,omitempty"`
	UpdatedAt       int64    `json:"updatedAt"`
	CreatedAt       int64    `json:"createdAt"`
}

type DeleteEnvironmentVariableRequest struct {
	// Required: Yes
	// The unique Project identifier.
	ProjectId string `json:"projectId"`

	// Required: Yes
	// The unique Environment Variable identifier.
	EnvId string `json:"envId"`
}
type DeleteEnvironmentVariableResponse struct {
	Type            string   `json:"type"`
	ID              string   `json:"id"`
	Key             string   `json:"key"`
	Value           string   `json:"value"`
	Target          []string `json:"target"`
	ConfigurationID string   `json:"configurationId,omitempty"`
	GitBranch       string   `json:"gitBranch,omitempty"`
	UpdatedAt       int64    `json:"updatedAt"`
	CreatedAt       int64    `json:"createdAt"`
}

type AddDomainRequest struct {
	// The name of the production domain.
	// Required: Yes
	Name string `json:"string"`

	// Target destination domain for redirect.
	// Required: No
	Redirect string `json:"redirect,omitempty"`

	// The redirect status code (301, 302, 307, 308).
	// Required: No
	RedirectStatusCode int64 `json:"redirectStatusCode,omitempty"`

	// Git branch for the domain to be auto assigned to. The Project's production branch is the default (null).
	// Required: No
	GitBranch string `json:"gitBranch,omitempty"`
}
type AddDomainResponse struct {
	Name      string `json:"name"`
	ProjectID string `json:"projectId"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type ListDomainsRequest struct {
}
type ListDomainsResponse struct {
	Domains []struct {
		Name      string `json:"name"`
		ProjectID string `json:"projectId"`
		UpdatedAt int64  `json:"updatedAt"`
		CreatedAt int64  `json:"createdAt"`
	} `json:"domains"`
	Pagination struct {
		Count int   `json:"count"`
		Next  int64 `json:"next"`
		Prev  int64 `json:"prev"`
	} `json:"pagination"`
}

type GetDomainRequest struct{}
type GetDomainResponse struct {
	Name      string `json:"name"`
	ProjectID string `json:"projectId"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedAt int64  `json:"createdAt"`
}

type SetRedirectRequest struct {
	// Target destination domain for redirect.
	// Required: No
	Redirect string `json:"redirect,omitempty"`
}
type SetRedirectResponse struct {
	Name      string `json:"name"`
	Redirect  string `json:"redirect"`
	ProjectID string `json:"projectId"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedAt int64  `json:"createdAt"`
}

type DeleteDomainRequest struct {
	// The unique Project identifier.
	// Required: Yes
	Id string `json:"id"`

	// The Project domain name.
	// Required: Yes
	Domain string `json:"domain"`
}
type DeleteDomainResponse struct{}
