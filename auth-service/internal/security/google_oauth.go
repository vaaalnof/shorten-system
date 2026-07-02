package security

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"auth-service/internal/model"
	"auth-service/internal/usecase/port"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const googleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"

var _ port.GoogleOAuthService = (*GoogleOAuth)(nil)

type GoogleOAuth struct {
	config *oauth2.Config
}

func NewGoogleOAuth(
	clientID string,
	clientSecret string,
	redirectURL string,
) *GoogleOAuth {

	return &GoogleOAuth{
		config: &oauth2.Config{
			ClientID: clientID,

			ClientSecret: clientSecret,

			RedirectURL: redirectURL,

			Scopes: []string{
				"openid",
				"email",
				"profile",
			},

			Endpoint: google.Endpoint,
		},
	}
}

func (g *GoogleOAuth) AuthURL(
	state string,
) string {

	return g.config.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
	)
}

func (g *GoogleOAuth) Exchange(
	ctx context.Context,
	code string,
) (*model.GoogleUser, error) {

	token, err := g.config.Exchange(
		ctx,
		code,
	)

	if err != nil {
		return nil, err
	}

	client := g.config.Client(
		ctx,
		token,
	)

	resp, err := client.Get(
		googleUserInfoURL,
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf(
			"google oauth returned status %d",
			resp.StatusCode,
		)
	}

	var user struct {
		ID string `json:"id"`

		Email string `json:"email"`

		VerifiedEmail bool `json:"verified_email"`

		GivenName string `json:"given_name"`

		FamilyName string `json:"family_name"`

		Picture string `json:"picture"`
	}

	if err := json.NewDecoder(
		resp.Body,
	).Decode(
		&user,
	); err != nil {

		return nil, err
	}

	return &model.GoogleUser{
		ID: user.ID,

		Email: user.Email,

		FirstName: user.GivenName,

		LastName: user.FamilyName,

		AvatarURL: user.Picture,

		EmailVerified: user.VerifiedEmail,
	}, nil
}
