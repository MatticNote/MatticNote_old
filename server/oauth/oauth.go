package oauth

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/MatticNote/MatticNote/config"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/storage"
	"net/http"
	"time"
)

var (
	oauthStorage = storage.NewExampleStore()
	oauthSecret  = func() []byte {
		if config.Config == nil {
			_, err := config.LoadConfiguration()
			if err != nil {
				panic(err)
			}
		}

		return []byte(config.Config.Server.OauthSecret)
	}()
	oauthConfig = &compose.Config{
		AccessTokenLifespan: 30 * time.Minute,
	}
	oauthPrivateKey = func() *rsa.PrivateKey {
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			panic(err)
		}
		return privateKey
	}()
	MNOAuthProvider = compose.ComposeAllEnabled(oauthConfig, oauthStorage, oauthSecret, oauthPrivateKey)
)

func AuthEndpoint(w http.ResponseWriter, r *http.Request) {
	authReq, err := MNOAuthProvider.NewAuthorizeRequest(r.Context(), r)
	if err != nil {
		MNOAuthProvider.WriteAuthorizeError(w, authReq, err)
		return
	}
}
