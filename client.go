package auvik

import (
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

func New(url, username, apiKey string) (*Auvik, error) {
	auth, err := securityprovider.NewSecurityProviderBasicAuth(username, apiKey)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(url, "/v1") {
		url += "/v1"
	}
	client, err := NewClient(url, WithRequestEditorFn(auth.Intercept))
	if err != nil {
		return nil, err
	}
	return client, nil
}
