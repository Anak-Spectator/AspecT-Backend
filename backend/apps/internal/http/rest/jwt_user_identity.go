package rest

import (
	"aspect_apps/internal/shared"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const userIdentityTokenIssuer = "aspect_apps-main_backend"

type jwtUserIdentity struct {
	AccessToken  string
	RefreshToken string
}

func newJWTUserIdentity(signKey []byte, idGenerator shared.IDGenerator, userID shared.UserID) (jwtUserIdentity, error) {

	/*
		? ===== TODO ===== ?
		[ ] Reduce expired time for better security
		? === END TODO === ?
	*/
	userIdentityClaims, err := newJwtUserIdentityClaims(idGenerator, userID, 24*356*time.Hour)
	if err != nil {
		if errors.Is(err, &shared.InfrastructureError{}) {
			log.Println(err)
		}
		return jwtUserIdentity{}, err
	}

	userIdentityRefreshClaims := newJwtUserIdentityRefreshClaims(userIdentityClaims, 168*time.Hour)
	if err != nil {
		if errors.Is(err, &shared.InfrastructureError{}) {
			log.Println(err)
		}
		return jwtUserIdentity{}, err
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, userIdentityClaims).SignedString(signKey)
	if err != nil {
		log.Println(err)
		return jwtUserIdentity{}, shared.NewInfrastructureError(err)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, userIdentityRefreshClaims).SignedString(signKey)
	if err != nil {
		log.Println(err)
		return jwtUserIdentity{}, shared.NewInfrastructureError(err)
	}

	return jwtUserIdentity{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
