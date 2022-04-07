package executor

import (
	"encoding/json"
	"strings"
)

type ServiceAccount struct {
	Type         string `json:"type,omitempty"`
	ProjectID    string `json:"project_id,omitempty"`
	PrivateKeyID string `json:"private_key_id,omitempty"`
	PrivateKey   string `json:"private_key,omitempty"`
	ClientEmail  string `json:"client_email,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	AuthURI      string `json:"auth_uri,omitempty"`
	TokenURI     string `json:"token_uri,omitempty"`
	AuthProvider string `json:"auth_provider_x509_cert_url,omitempty"`
	ClientCert   string `json:"client_x509_cert_url,omitempty"`
}

func (s *ServiceAccount) cleanPrivateKey() {
	s.PrivateKey = strings.ReplaceAll(s.PrivateKey, "\n", "\\n")
}

func (s *ServiceAccount) toString() string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	r := string(b)
	return r
}
