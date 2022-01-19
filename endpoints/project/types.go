package project

// Project houses all the information vercel offers about a project via their api
type Project struct {
	Accountid string `json:"accountid"`
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
				Uid         string `json:"uid"`
				Email       string `json:"email"`
				Username    string `json:"username"`
				GithubLogin string `json:"githubLogin"`
			} `json:"creator"`
			DeploymentHostname string `json:"deploymentHostname"`
			Forced             bool   `json:"forced"`
			Id                 string `json:"id"`
			Meta               struct {
				GithubCommitRef         string `json:"githubCommitRef"`
				GithubRepo              string `json:"githubRepo"`
				GithubOrg               string `json:"githubOrg"`
				GithubCommitSha         string `json:"githubCommitSha"`
				GithubRepoid            string `json:"githubRepoid"`
				GithubCommitMessage     string `json:"githubCommitMessage"`
				GithubCommitAuthorLogin string `json:"githubCommitAuthorLogin"`
				GithubDeployment        string `json:"githubDeployment"`
				GithubCommitOrg         string `json:"githubCommitOrg"`
				GithubCommitAuthorName  string `json:"githubCommitAuthorName"`
				GithubCommitRepo        string `json:"githubCommitRepo"`
				GithubCommitRepoid      string `json:"githubCommitRepoid"`
			} `json:"meta"`
			Name       string `json:"name"`
			Plan       string `json:"plan"`
			Private    bool   `json:"private"`
			ReadyState string `json:"readyState"`
			Target     string `json:"target"`
			Teamid     string `json:"teamid"`
			Type       string `json:"type"`
			URL        string `json:"url"`
			Userid     string `json:"userid"`
			WithCache  bool   `json:"withCache"`
		} `json:"deployment"`
		Domain      string `json:"domain"`
		Environment string `json:"environment"`
		Target      string `json:"target"`
	} `json:"alias"`
	Analytics struct {
		Id         string `json:"id"`
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
		Id              string      `json:"id"`
		Key             string      `json:"key"`
		Value           string      `json:"value"`
		Target          []string    `json:"target"`
		Configurationid interface{} `json:"configurationid"`
		UpdatedAt       int64       `json:"updatedAt"`
		CreatedAt       int64       `json:"createdAt"`
	} `json:"env"`
	Framework                       string `json:"framework"`
	Id                              string `json:"id"`
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
		Repoid           int64         `json:"repoid"`
		Org              string        `json:"org"`
		GitCredentialid  string        `json:"gitCredentialid"`
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
			Uid         string `json:"uid"`
			Email       string `json:"email"`
			Username    string `json:"username"`
			GithubLogin string `json:"githubLogin"`
		} `json:"creator"`
		DeploymentHostname string `json:"deploymentHostname"`
		Forced             bool   `json:"forced"`
		Id                 string `json:"id"`
		Meta               struct {
			GithubCommitRef         string `json:"githubCommitRef"`
			GithubRepo              string `json:"githubRepo"`
			GithubOrg               string `json:"githubOrg"`
			GithubCommitSha         string `json:"githubCommitSha"`
			GithubCommitAuthorLogin string `json:"githubCommitAuthorLogin"`
			GithubCommitMessage     string `json:"githubCommitMessage"`
			GithubRepoid            string `json:"githubRepoid"`
			GithubDeployment        string `json:"githubDeployment"`
			GithubCommitOrg         string `json:"githubCommitOrg"`
			GithubCommitAuthorName  string `json:"githubCommitAuthorName"`
			GithubCommitRepo        string `json:"githubCommitRepo"`
			GithubCommitRepoid      string `json:"githubCommitRepoid"`
		} `json:"meta"`
		Name       string      `json:"name"`
		Plan       string      `json:"plan"`
		Private    bool        `json:"private"`
		ReadyState string      `json:"readyState"`
		Target     interface{} `json:"target"`
		Teamid     string      `json:"teamid"`
		Type       string      `json:"type"`
		URL        string      `json:"url"`
		Userid     string      `json:"userid"`
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
				Uid         string `json:"uid"`
				Email       string `json:"email"`
				Username    string `json:"username"`
				GithubLogin string `json:"githubLogin"`
			} `json:"creator"`
			DeploymentHostname string `json:"deploymentHostname"`
			Forced             bool   `json:"forced"`
			Id                 string `json:"id"`
			Meta               struct {
				GithubCommitRef         string `json:"githubCommitRef"`
				GithubRepo              string `json:"githubRepo"`
				GithubOrg               string `json:"githubOrg"`
				GithubCommitSha         string `json:"githubCommitSha"`
				GithubRepoid            string `json:"githubRepoid"`
				GithubCommitMessage     string `json:"githubCommitMessage"`
				GithubCommitAuthorLogin string `json:"githubCommitAuthorLogin"`
				GithubDeployment        string `json:"githubDeployment"`
				GithubCommitOrg         string `json:"githubCommitOrg"`
				GithubCommitAuthorName  string `json:"githubCommitAuthorName"`
				GithubCommitRepo        string `json:"githubCommitRepo"`
				GithubCommitRepoid      string `json:"githubCommitRepoid"`
			} `json:"meta"`
			Name       string `json:"name"`
			Plan       string `json:"plan"`
			Private    bool   `json:"private"`
			ReadyState string `json:"readyState"`
			Target     string `json:"target"`
			Teamid     string `json:"teamid"`
			Type       string `json:"type"`
			URL        string `json:"url"`
			Userid     string `json:"userid"`
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
	GitRepository *GitRepository `json:"gitRepository,omitempty"`

	UpdateProjectRequest
}
type GitRepository struct {
	Type string `json:"type,omitempty"`
	Repo string `json:"repo,omitempty"`
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

// If the request is successful, you will receive a 204 HTTP Status code in the
// response.
type DeleteProjectResponse struct{}

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
		Configurationid interface{} `json:"configurationid"`
		Id              string      `json:"id"`
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
	// If the type is secret, the secret id of the secret attached to the
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
	Id              string   `json:"id"`
	Key             string   `json:"key"`
	Value           string   `json:"value"`
	Target          []string `json:"target"`
	Configurationid string   `json:"configurationid,omitempty"`
	GitBranch       string   `json:"gitBranch,omitempty"`
	UpdatedAt       int64    `json:"updatedAt"`
	CreatedAt       int64    `json:"createdAt"`
}

type UpdateEnvironmentVariableRequest struct {
	Envid string `json:"-"`
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
type UpdateEnvironmentVariableResponse struct {
	Type            string   `json:"type"`
	Id              string   `json:"id"`
	Key             string   `json:"key"`
	Value           string   `json:"value"`
	Target          []string `json:"target"`
	Configurationid string   `json:"configurationid,omitempty"`
	GitBranch       string   `json:"gitBranch,omitempty"`
	UpdatedAt       int64    `json:"updatedAt"`
	CreatedAt       int64    `json:"createdAt"`
}

type DeleteEnvironmentVariableRequest struct {
	// Required: Yes
	// The unique Project identifier.
	Projectid string

	// Required: Yes
	// The unique Environment Variable identifier.
	Envid string
}
type DeleteEnvironmentVariableResponse struct{}

type AddDomainRequest struct {
	// The name of the production domain.
	// Required: Yes
	Name string `json:"name"`

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

type UpdateDomainRequest struct {
	// The name of the production domain.
	// Required: Yes
	Name string `json:"-"`

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
type Domain struct {
	ProjectId          string `json:"projectId"`
	Name               string `json:"name"`
	ApexName           string `json:"apexName"`
	Redirect           string `json:"redirect,omitempty"`
	RedirectStatusCode int    `json:"redirectStatusCode,omitempty"`
	GitBranch          string `json:"gitBranch,omitempty"`
	UpdatedAt          int    `json:"updatedAt,omitempty"`
	CreatedAt          int    `json:"createdAt,omitempty"`
}

type ListDomainsRequest struct {

	// Filters domains based on specific branch.
	// Required: No
	GitBranch string `json:"gitBranch,omitempty"`

	// Maximum number of domains to list from a request (max 100).
	// Required: No
	Limit int `json:"limit,omitempty"`

	// Filters only production domains when set to true.
	// Required: No
	Production bool `json:"production,omitempty"`

	// Excludes redirect project domains when "false". Includes redirect project
	// domains when "true" (default).
	// Required: No
	Redirects bool `json:"redirects,omitempty"`

	// Get domains created after this JavaScript timestamp.
	// Required: No
	Since int `json:"since,omitempty"`

	// Get domains created before this JavaScript timestamp.
	// Required: No
	Until int `json:"until,omitempty"`
}
type ListDomainsResponse struct {
	Domains    []Domain `json:"domains"`
	Pagination struct {
		Count int `json:"count"`
		Next  int `json:"next"`
		Prev  int `json:"prev"`
	} `json:"pagination"`
}

type GetDomainRequest struct{}

type SetRedirectRequest struct {
	// Target destination domain for redirect.
	// Required: No
	Redirect string `json:"redirect,omitempty"`
}
type SetRedirectResponse struct {
	Name      string `json:"name"`
	Redirect  string `json:"redirect"`
	Projectid string `json:"projectId"`
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
