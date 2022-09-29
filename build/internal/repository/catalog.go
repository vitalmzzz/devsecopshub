package repository

import (
	"time"

	"gorm.io/datatypes"
)

//Структура для GitLab
type GitProjects struct {
	Id                  int64  `json: "id"`
	Name                string `json: "name"`
	Description         string `json: "description"`
	Path_with_namespace string `json: "path_with_namespace"`
	Web_url             string `json: "web_url"`
	Archived            bool   `json: "archived"`
	Owner               Owner_data
	Namespace           Namespace
}

type Owner_data struct {
	Username string `json: "username"`
	Name     string `json: "name"`
}

type Namespace struct {
	Id   int64  `json: "id"`
	Name string `json: "name"`
}

//Структура для Token
type Token struct {
	Access_token string
}

//Структура для App.Farm
type InternalSystem []struct {
	Environment string `json:"environment"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Namespace   string `json:"namespace"`
	Owner       string `json:"owner"`
	Status      struct {
		State   string `json:"state"`
		Message string `json:"message"`
	} `json:"status"`
	DateCreated    time.Time `json:"dateCreated"`
	ProjectPath    string    `json:"projectPath"`
	ResourceLimits struct {
		CPU string `json:"cpu"`
		RAM string `json:"ram"`
	} `json:"resourceLimits"`
	ResourceUsage struct {
		CPU string `json:"cpu"`
		RAM string `json:"ram"`
	} `json:"resourceUsage"`
	References struct {
		Src         string      `json:"src"`
		Artifacts   string      `json:"artifacts"`
		Monitoring  interface{} `json:"monitoring"`
		Logs        string      `json:"logs"`
		Tracing     string      `json:"tracing"`
		CodeQuality string      `json:"codeQuality"`
		Secrets     string      `json:"secrets"`
		Web         string      `json:"web"`
	} `json:"references"`
	Depart struct {
		ID          string `json:"id"`
		Path        string `json:"path"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ParentID    string `json:"parentId"`
		Members     []struct {
			User struct {
				Username string `json:"username"`
				Fullname string `json:"fullname"`
				Email    string `json:"email"`
				Company  string `json:"company"`
				JobTitle string `json:"jobTitle"`
			} `json:"user"`
			Role string `json:"role"`
		} `json:"members"`
		Owners []struct {
			User struct {
				Username string `json:"username"`
				Fullname string `json:"fullname"`
				Email    string `json:"email"`
				Company  string `json:"company"`
				JobTitle string `json:"jobTitle"`
			} `json:"user"`
			Role string `json:"role"`
		} `json:"owners"`
		AvatarURL string `json:"avatarUrl"`
	} `json:"depart"`
}

type ExternalSystems []struct {
	Environment string `json:"environment"`
	ID          string `json:"id"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Status      struct {
		State   string `json:"state"`
		Message string `json:"message"`
	} `json:"status"`
	DateCreated time.Time `json:"dateCreated"`
	References  struct {
		Src         string      `json:"src"`
		Artifacts   string      `json:"artifacts"`
		Monitoring  interface{} `json:"monitoring"`
		Logs        string      `json:"logs"`
		Tracing     string      `json:"tracing"`
		CodeQuality string      `json:"codeQuality"`
		Secrets     string      `json:"secrets"`
		Web         string      `json:"web"`
	} `json:"references"`
	Depart struct {
		ID          string `json:"id"`
		Path        string `json:"path"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ParentID    string `json:"parentId"`
		Members     []struct {
			User struct {
				Username string `json:"username"`
				Fullname string `json:"fullname"`
				Email    string `json:"email"`
				Company  string `json:"company"`
				JobTitle string `json:"jobTitle"`
			} `json:"user"`
			Role string `json:"role"`
		} `json:"members"`
		Owners []struct {
			User struct {
				Username string `json:"username"`
				Fullname string `json:"fullname"`
				Email    string `json:"email"`
				Company  string `json:"company"`
				JobTitle string `json:"jobTitle"`
			} `json:"user"`
			Role string `json:"role"`
		} `json:"owners"`
		AvatarURL string `json:"avatarUrl"`
	} `json:"depart"`
}

type InternalServices []struct {
	Environment         string `json:"environment"`
	ID                  string `json:"id"`
	InformationSystemID string `json:"informationSystemId"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Status              struct {
		Phase   string `json:"phase"`
		State   string `json:"state"`
		Message string `json:"message"`
	} `json:"status"`
	DateCreated          time.Time   `json:"dateCreated"`
	ProjectPath          string      `json:"projectPath"`
	ConnectionParameters interface{} `json:"connectionParameters"`
	ServiceType          string      `json:"serviceType"`
	HasPublishedAPI      bool        `json:"hasPublishedApi"`
	Addresses            struct {
		InternalURL string `json:"internalUrl"`
		PublicURL   string `json:"publicUrl"`
	} `json:"addresses"`
	References struct {
		Src              string      `json:"src"`
		GitHTTP          string      `json:"gitHttp"`
		GitSSH           string      `json:"gitSsh"`
		Artifacts        string      `json:"artifacts"`
		Monitoring       interface{} `json:"monitoring"`
		Logs             string      `json:"logs"`
		Tracing          string      `json:"tracing"`
		CodeQuality      string      `json:"codeQuality"`
		Secrets          string      `json:"secrets"`
		AuthAdminConsole string      `json:"authAdminConsole"`
	} `json:"references"`
	System struct {
		ID                   string    `json:"id"`
		Name                 string    `json:"name"`
		Description          string    `json:"description"`
		Namespace            string    `json:"namespace"`
		IntegrationNamespace string    `json:"integrationNamespace"`
		Owner                string    `json:"owner"`
		DateCreated          time.Time `json:"dateCreated"`
		ProjectPath          string    `json:"projectPath"`
		DepartPath           string    `json:"departPath"`
		AvatarURL            string    `json:"avatarUrl"`
		SystemType           string    `json:"systemType"`
		IsSensitiveData      bool      `json:"isSensitiveData"`
		DynamicEnvironment   struct {
		} `json:"dynamicEnvironment"`
	} `json:"system"`
	Depart struct {
		ID          int    `json:"id"`
		Path        string `json:"path"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ParentID    int    `json:"parentId"`
	} `json:"depart"`
	AvatarURL      string `json:"avatarUrl"`
	ServiceProject struct {
		LastUpdate time.Time `json:"lastUpdate"`
		Version    string    `json:"version"`
		Message    string    `json:"message"`
		Editor     struct {
			UserName       string `json:"userName"`
			EditorFullName string `json:"editorFullName"`
			Email          string `json:"email"`
			AvatarURL      string `json:"avatarUrl"`
		} `json:"editor"`
		Started  time.Time `json:"started"`
		Finished time.Time `json:"finished"`
		Duration int       `json:"duration"`
		Status   string    `json:"status"`
		ID       int       `json:"id"`
		WebURL   string    `json:"webUrl"`
	} `json:"serviceProject,omitempty"`
	ResourceLimits struct {
	} `json:"resourceLimits"`
	Side            string `json:"side"`
	ServiceProject0 struct {
		LastUpdate             time.Time `json:"lastUpdate"`
		Version                string    `json:"version"`
		DeployedVersion        string    `json:"deployedVersion"`
		DeployedCommitShortSHA string    `json:"deployedCommitShortSHA"`
		Message                string    `json:"message"`
		Editor                 struct {
			UserName       string `json:"userName"`
			EditorFullName string `json:"editorFullName"`
			Email          string `json:"email"`
			AvatarURL      string `json:"avatarUrl"`
		} `json:"editor"`
		Started  time.Time `json:"started"`
		Finished time.Time `json:"finished"`
		Duration int       `json:"duration"`
		Status   string    `json:"status"`
		ID       int       `json:"id"`
		WebURL   string    `json:"webUrl"`
	} `json:"serviceProject,omitempty"`
	ServiceProject1 struct {
		LastUpdate             time.Time `json:"lastUpdate"`
		Version                string    `json:"version"`
		DeployedVersion        string    `json:"deployedVersion"`
		DeployedCommitShortSHA string    `json:"deployedCommitShortSHA"`
		Message                string    `json:"message"`
		Editor                 struct {
			UserName       string `json:"userName"`
			EditorFullName string `json:"editorFullName"`
			Email          string `json:"email"`
			AvatarURL      string `json:"avatarUrl"`
		} `json:"editor"`
		Started  time.Time `json:"started"`
		Finished time.Time `json:"finished"`
		Duration int       `json:"duration"`
		Status   string    `json:"status"`
		ID       int       `json:"id"`
		WebURL   string    `json:"webUrl"`
	} `json:"serviceProject,omitempty"`
}

type ExternalServices []struct {
	Environment         string `json:"environment"`
	ID                  string `json:"id"`
	InformationSystemID string `json:"informationSystemId"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Status              struct {
		Phase   string `json:"phase"`
		State   string `json:"state"`
		Message string `json:"message"`
	} `json:"status"`
	DateCreated          time.Time `json:"dateCreated"`
	ProjectPath          string    `json:"projectPath"`
	ConnectionParameters []struct {
		Protocol   string `json:"protocol"`
		Parameters struct {
			Channel           string `json:"channel"`
			ConnName          string `json:"conn-name"`
			ConsumerQueueName string `json:"consumer-queue-name"`
			Password          string `json:"password"`
			ProducerQueueName string `json:"producer-queue-name"`
			QueueManager      string `json:"queue-manager"`
			User              string `json:"user"`
		} `json:"parameters"`
	} `json:"connectionParameters"`
	ServiceType     string `json:"serviceType"`
	HasPublishedAPI bool   `json:"hasPublishedApi"`
	Addresses       struct {
		InternalURL string `json:"internalUrl"`
		PublicURL   string `json:"publicUrl"`
	} `json:"addresses"`
	References struct {
		Src              string      `json:"src"`
		GitHTTP          string      `json:"gitHttp"`
		GitSSH           string      `json:"gitSsh"`
		Artifacts        string      `json:"artifacts"`
		Monitoring       interface{} `json:"monitoring"`
		Logs             string      `json:"logs"`
		Tracing          string      `json:"tracing"`
		CodeQuality      string      `json:"codeQuality"`
		Secrets          string      `json:"secrets"`
		AuthAdminConsole string      `json:"authAdminConsole"`
	} `json:"references"`
	System struct {
		ID                   string    `json:"id"`
		Name                 string    `json:"name"`
		Description          string    `json:"description"`
		Namespace            string    `json:"namespace"`
		IntegrationNamespace string    `json:"integrationNamespace"`
		Owner                string    `json:"owner"`
		DateCreated          time.Time `json:"dateCreated"`
		ProjectPath          string    `json:"projectPath"`
		DepartPath           string    `json:"departPath"`
		AvatarURL            string    `json:"avatarUrl"`
		SystemType           string    `json:"systemType"`
		Kind                 string    `json:"kind"`
		IsSensitiveData      bool      `json:"isSensitiveData"`
		DynamicEnvironment   struct {
		} `json:"dynamicEnvironment"`
	} `json:"system"`
	Depart struct {
		ID          int    `json:"id"`
		Path        string `json:"path"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ParentID    int    `json:"parentId"`
	} `json:"depart"`
	AvatarURL      string `json:"avatarUrl"`
	ServiceProject struct {
		LastUpdate             time.Time `json:"lastUpdate"`
		Version                string    `json:"version"`
		DeployedVersion        string    `json:"deployedVersion"`
		DeployedCommitShortSHA string    `json:"deployedCommitShortSHA"`
		Message                string    `json:"message"`
		Editor                 struct {
			UserName       string `json:"userName"`
			EditorFullName string `json:"editorFullName"`
			Email          string `json:"email"`
			AvatarURL      string `json:"avatarUrl"`
		} `json:"editor"`
		Started  time.Time `json:"started"`
		Finished time.Time `json:"finished"`
		Duration int       `json:"duration"`
		Status   string    `json:"status"`
		ID       int       `json:"id"`
		WebURL   string    `json:"webUrl"`
	} `json:"serviceProject"`
}

//Структура для AppSecHub
type RestApplication struct {
	FilteredEntities []struct {
		ID                       int    `json:"id"`
		Name                     string `json:"name"`
		Code                     string `json:"code"`
		Supported                bool   `json:"supported"`
		IsTrackingServiceEnabled bool   `json:"isTrackingServiceEnabled"`
		Summary                  struct {
			Defects struct {
				Total    int `json:"total"`
				Blocker  int `json:"blocker"`
				Critical int `json:"critical"`
				Major    int `json:"major"`
				Minor    int `json:"minor"`
				Info     int `json:"info"`
			} `json:"defects"`
			Issues struct {
				Total    int `json:"total"`
				Critical int `json:"critical"`
				High     int `json:"high"`
				Medium   int `json:"medium"`
				Low      int `json:"low"`
			} `json:"issues"`
			SsdlTasksSummary struct {
				SsdlTotal                int `json:"ssdlTotal"`
				SsdlNew                  int `json:"ssdlNew"`
				SsdlOpen                 int `json:"ssdlOpen"`
				SsdlInProgress           int `json:"ssdlInProgress"`
				SsdlDone                 int `json:"ssdlDone"`
				SsdlIncomplete           int `json:"ssdlIncomplete"`
				SsdlNoResources          int `json:"ssdlNoResources"`
				SsdlNotApplicable        int `json:"ssdlNotApplicable"`
				SsdlClusterRunDone       int `json:"ssdlClusterRunDone"`
				SsdlClusterRunRunning    int `json:"ssdlClusterRunRunning"`
				SsdlClusterChangeDone    int `json:"ssdlClusterChangeDone"`
				SsdlClusterChangeRunning int `json:"ssdlClusterChangeRunning"`
			} `json:"ssdlTasksSummary"`
			Wri          int `json:"wri"`
			CodebaseSize int `json:"codebaseSize"`
		} `json:"summary"`
		WorkspaceID           int      `json:"workspaceId"`
		OrgStructureID        int      `json:"orgStructureId"`
		Created               int64    `json:"created"`
		StructureUnitType     string   `json:"structureUnitType"`
		IntegrityCheckEnabled bool     `json:"integrityCheckEnabled"`
		Practices             []string `json:"practices"`
		Description           string   `json:"description,omitempty"`
	} `json:"filteredEntities"`
	TotalEntitiesCount int `json:"totalEntitiesCount"`
}

type Defects struct {
	FilteredEntities []struct {
		ID         int    `json:"id"`
		AppID      int    `json:"appId"`
		Summary    string `json:"summary"`
		Status     string `json:"status"`
		PriorityID int    `json:"priorityId"`
		SeverityID int    `json:"severityId"`
		Priority   struct {
			Name string `json:"name"`
			Rank int    `json:"rank"`
		} `json:"priority"`
		Severity struct {
			Name string `json:"name"`
			Rank int    `json:"rank"`
		} `json:"severity"`
		ExternalBugs []struct {
			Product     string `json:"product"`
			DisplayName string `json:"displayName"`
			Link        string `json:"link"`
		} `json:"externalBugs"`
		IsActive   bool   `json:"isActive"`
		CreateTs   int64  `json:"createTs"`
		UpdateTs   int64  `json:"updateTs"`
		LastSync   int64  `json:"lastSync"`
		SyncStatus string `json:"syncStatus"`
	} `json:"filteredEntities"`
	TotalEntitiesCount int `json:"totalEntitiesCount"`
}

type Artifacts []struct {
	ID                   int  `json:"id"`
	AppID                int  `json:"appId"`
	Active               bool `json:"active"`
	ArtifactRepositoryID int  `json:"artifactRepositoryId"`
	ArtifactRepository   struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		ToolID  int    `json:"toolId"`
		Product string `json:"product"`
	} `json:"artifactRepository"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties struct {
		BaseURL        string `json:"baseUrl"`
		Port           string `json:"port"`
		Name           string `json:"name"`
		DockerArtifact string `json:"dockerArtifact"`
		RepositoryName string `json:"repositoryName"`
	} `json:"properties"`
	PipelineID int `json:"pipelineId"`
}

type Codebase []struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Link                 string `json:"link"`
	Branch               string `json:"branch"`
	Active               bool   `json:"active"`
	AppID                int    `json:"appId"`
	VcsType              string `json:"vcsType"`
	CheckoutRelativePath string `json:"checkoutRelativePath"`
	BuildTool            string `json:"buildTool"`
	VcsRepositoryID      int    `json:"vcsRepositoryId"`
	VcsRepository        struct {
		ID                 int    `json:"id"`
		Type               string `json:"type"`
		Product            string `json:"product"`
		ProductDisplayName string `json:"productDisplayName"`
		Name               string `json:"name"`
		URL                string `json:"url"`
		ExtraData          struct {
			VcsType string `json:"vcsType"`
		} `json:"extraData"`
		WorkspaceIds    []int `json:"workspaceIds"`
		CiCdCredentials struct {
			ID                    int         `json:"id"`
			IntegrationInstanceID int         `json:"integrationInstanceId"`
			Purpose               string      `json:"purpose"`
			AuthMethod            string      `json:"authMethod"`
			Login                 string      `json:"login"`
			AccountID             string      `json:"accountId"`
			Password              interface{} `json:"password"`
			UserID                string      `json:"userId"`
			Token                 interface{} `json:"token"`
			SSHPublic             interface{} `json:"sshPublic"`
			SSHPrivate            interface{} `json:"sshPrivate"`
			SSHUsername           string      `json:"sshUsername"`
			SSHPassphrase         interface{} `json:"sshPassphrase"`
			SSHUploadedKeyName    interface{} `json:"sshUploadedKeyName"`
			Active                bool        `json:"active"`
		} `json:"ciCdCredentials"`
		Active bool `json:"active"`
	} `json:"vcsRepository"`
	PipelineID int    `json:"pipelineId"`
	Project    string `json:"project,omitempty"`
	Repo       string `json:"repo,omitempty"`
}

type Jira struct {
	Expand     string `json:"expand"`
	StartAt    int    `json:"startAt"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	Issues     []struct {
		Expand string `json:"expand"`
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Summary    string `json:"summary"`
			Components []struct {
				Self string `json:"self"`
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"components"`
			Created        string `json:"created"`
			Resolutiondate string `json:"resolutiondate"`
			Assignee       struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"assignee"`
			Aggregateprogress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
				Percent  int `json:"percent"`
			} `json:"aggregateprogress"`
			Updated string         `json:"updated"`
			Labels  datatypes.JSON `json:"labels"`
			Status  struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
		} `json:"fields"`
	} `json:"issues"`
}

type WorkTime struct {
	StartAt    int `json:"startAt"`
	MaxResults int `json:"maxResults"`
	Total      int `json:"total"`
	Worklogs   []struct {
		Self   string `json:"self"`
		Author struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"author"`
		UpdateAuthor struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"updateAuthor"`
		Comment          string `json:"comment"`
		Created          string `json:"created"`
		Updated          string `json:"updated"`
		Started          string `json:"started"`
		TimeSpent        string `json:"timeSpent"`
		TimeSpentSeconds int    `json:"timeSpentSeconds"`
		ID               string `json:"id"`
		IssueID          string `json:"issueId"`
	} `json:"worklogs"`
}
