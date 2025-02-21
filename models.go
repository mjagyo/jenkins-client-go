package jenkins

type Jobs struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	FullName    string `json:"fullName"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Buildable   bool   `json:"buildable"`
	Color       string `json:"color"`
	InQueue     bool   `json:"inQueue"`
}

type CredentialRequest struct {
	Credentials Credential `json:"credentials"`
}

type Credential struct {
	Scope       string `json:"scope,omitempty"`
	ID          string `json:"id,omitempty"`
	AppID       string `json:"appID,omitempty"`
	PrivateKey  string `json:"privateKey,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Description string `json:"description,omitempty"`
	Class       string `json:"$class"`
	StaperClass string `json:"stapler-class"`
}

type CredentialResponse struct {
	Class       string `json:"_class"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	Fingerprint string `json:"fingerprint"`
	ID          string `json:"id"`
	TypeName    string `json:"typeName"`
}
