package rest

import (
	"aspect_apps/internal/shared"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtUserIdentityClaimsID string

type jwtUserIdentityClaims struct {
	ID        jwtUserIdentityClaimsID `json:"jti"`
	UserID    shared.UserID           `json:"uid"`
	ExpiredAt int64                   `json:"exp"`
	NotBefore int64                   `json:"nbf"`
	Issuer    string                  `json:"iss"`
	IssuedAt  int64                   `json:"iat"`
}

func newJwtUserIdentityClaims(idGenerator shared.IDGenerator, userID shared.UserID, lifeTimeDuration time.Duration) (jwtUserIdentityClaims, error) {
	id, err := idGenerator.GenerateID()
	if err != nil {
		if errors.Is(err, &shared.InfrastructureError{}) {
			log.Println(err)
		}
		return jwtUserIdentityClaims{}, err
	}

	return jwtUserIdentityClaims{
		ID:        jwtUserIdentityClaimsID(id),
		UserID:    userID,
		ExpiredAt: time.Now().Add(lifeTimeDuration).UTC().Unix(),
		NotBefore: time.Now().UTC().Unix(),
		Issuer:    userIdentityTokenIssuer,
		IssuedAt:  time.Now().UTC().Unix(),
	}, nil
}

func jwtUserIdentityClaimsFromJWTMapClaims(mapClaims jwt.MapClaims) (jwtUserIdentityClaims, error) {
	id, ok := mapClaims["jti"].(string)
	if !ok {
		return jwtUserIdentityClaims{}, errors.New("invalid access token")
	}

	userID, ok := mapClaims["uid"].(string)
	if !ok {
		return jwtUserIdentityClaims{}, errors.New("invalid access token")
	}

	expiredAt, ok := mapClaims["exp"].(float64)
	if !ok {
		return jwtUserIdentityClaims{}, errors.New("invalid access token")
	}

	notBefore, ok := mapClaims["nbf"].(float64)
	if !ok {
		return jwtUserIdentityClaims{}, errors.New("invalid access token")
	}

	issuer, ok := mapClaims["iss"].(string)
	if !ok {
		return jwtUserIdentityClaims{}, errors.New("invalid access token")
	}

	issuedAt, ok := mapClaims["iat"].(float64)
	if !ok {
		return jwtUserIdentityClaims{}, errors.New("invalid access token")
	}

	claims := jwtUserIdentityClaims{
		ID:        jwtUserIdentityClaimsID(id),
		UserID:    shared.UserID(userID),
		ExpiredAt: int64(expiredAt),
		NotBefore: int64(notBefore),
		Issuer:    issuer,
		IssuedAt:  int64(issuedAt),
	}

	return claims, nil
}

func (claim jwtUserIdentityClaims) Valid() error {
	if time.Now().UTC().Unix() >= claim.ExpiredAt {
		return jwt.NewValidationError("access token has been expired", jwt.ValidationErrorExpired)
	}

	if time.Now().UTC().Unix() < claim.NotBefore {
		return jwt.NewValidationError("access token is not usable yet", jwt.ValidationErrorNotValidYet)
	}

	if claim.Issuer != userIdentityTokenIssuer {
		return jwt.NewValidationError("invalid access token issuer", jwt.ValidationErrorIssuer)
	}

	return nil
}

type jwtUserIdentityRefreshClaims struct {
	Subject   jwtUserIdentityClaimsID `json:"sub"`
	ExpiredAt int64                   `json:"exp"`
	NotBefore int64                   `json:"nbf"`
	Issuer    string                  `json:"iss"`
	IssuedAt  int64                   `json:"iat"`
}

func newJwtUserIdentityRefreshClaims(userIdentityClaims jwtUserIdentityClaims, lifeTimeDuration time.Duration) jwtUserIdentityRefreshClaims {
	return jwtUserIdentityRefreshClaims{
		Subject:   userIdentityClaims.ID,
		ExpiredAt: time.Now().Add(lifeTimeDuration).UTC().Unix(),
		NotBefore: time.Unix(userIdentityClaims.ExpiredAt, 0).Add(-10 * time.Minute).UTC().Unix(),
		Issuer:    userIdentityTokenIssuer,
		IssuedAt:  time.Now().UTC().Unix(),
	}
}

func jwtUserIdentityRefreshClaimsFromJWTMapClaims(mapClaims jwt.MapClaims) (jwtUserIdentityRefreshClaims, error) {
	sub, ok := mapClaims["sub"].(string)
	if !ok {
		return jwtUserIdentityRefreshClaims{}, errors.New("invalid refresh token")
	}

	expiredAt, ok := mapClaims["exp"].(float64)
	if !ok {
		return jwtUserIdentityRefreshClaims{}, errors.New("invalid refresh token")
	}

	notBefore, ok := mapClaims["nbf"].(float64)
	if !ok {
		return jwtUserIdentityRefreshClaims{}, errors.New("invalid refresh token")
	}

	issuer, ok := mapClaims["iss"].(string)
	if !ok {
		return jwtUserIdentityRefreshClaims{}, errors.New("invalid refresh token")
	}

	issuedAt, ok := mapClaims["iat"].(float64)
	if !ok {
		return jwtUserIdentityRefreshClaims{}, errors.New("invalid refresh token")
	}

	claims := jwtUserIdentityRefreshClaims{
		Subject:   jwtUserIdentityClaimsID(sub),
		ExpiredAt: int64(expiredAt),
		NotBefore: int64(notBefore),
		Issuer:    issuer,
		IssuedAt:  int64(issuedAt),
	}

	return claims, nil
}

func (claim jwtUserIdentityRefreshClaims) Valid() error {
	if time.Now().UTC().Unix() >= claim.ExpiredAt {
		return jwt.NewValidationError("refresh token has been expired", jwt.ValidationErrorExpired)
	}

	if time.Now().UTC().Unix() < claim.NotBefore {
		return jwt.NewValidationError("refresh token is not usable yet", jwt.ValidationErrorNotValidYet)
	}

	if claim.Issuer != userIdentityTokenIssuer {
		return jwt.NewValidationError("refresh token token issuer", jwt.ValidationErrorIssuer)
	}

	return nil
}
