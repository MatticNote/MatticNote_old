package oauth

import (
	"context"
	"github.com/MatticNote/MatticNote/db"
	"github.com/ory/fosite"
	"golang.org/x/tools/go/analysis/passes/stringintconv/testdata/src/a"
	"time"
)

type MNOauthStorage struct {
}

type MNOauthStorageClient struct {
	ID             string
	HashedSecret   []byte
	RedirectURIs   []string
	GrantTypes     []string
	ResponseTypes  []string
	Scopes         []string
	IsClientPublic bool
	Audience       []string
}

func (c MNOauthStorageClient) GetID() string {
	return c.ID
}

func (c MNOauthStorageClient) GetHashedSecret() []byte {
	return c.HashedSecret
}

func (c MNOauthStorageClient) GetGrantTypes() fosite.Arguments {
	return c.GrantTypes
}

func (c MNOauthStorageClient) GetScopes() fosite.Arguments {
	return c.Scopes
}

func (c MNOauthStorageClient) IsPublic() bool {
	return c.IsClientPublic
}

func (c MNOauthStorageClient) GetAudience() fosite.Arguments {
	return c.Audience
}

func (s *MNOauthStorage) GetClient(ctx context.Context, id string) (fosite.Client, error) {
	var cli MNOauthStorageClient
	err := db.DB.QueryRow(ctx, "").Scan()
}

func (s *MNOauthStorage) ClientAssertionJWTValid(ctx context.Context, jti string) error {

}

func (s *MNOauthStorage) SetClientAssertionJWT(ctx context.Context, jti string, exp time.Time) error {

}
