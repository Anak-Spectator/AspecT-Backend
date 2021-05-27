package rest

import (
	"aspect_apps/internal/shared"
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
)

type JWTUserIdentifier struct {
	signKey []byte
}

func NewJwtUserIdentifier(signKey []byte) *JWTUserIdentifier {
	return &JWTUserIdentifier{
		signKey: signKey,
	}
}

func (identifier *JWTUserIdentifier) Identify(identity shared.UserIdentity) (shared.UserID, error) {
	accessToken, ok := identity.(string)
	if !ok {
		log.Printf("%T\n", identity)
		return "", errors.New("invalid access token")
	}

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid access token", jwt.ValidationErrorSignatureInvalid)
		}

		return identifier.signKey, nil
	})
	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			return "", e
		default:
			return "", shared.NewInfrastructureError(e)
		}
	}

	if !token.Valid {
		return "", errors.New("invalid access token")
	}

	switch claims := token.Claims.(type) {
	case jwt.MapClaims:
		jwtUserClaims, err := jwtUserIdentityClaimsFromJWTMapClaims(claims)
		if err != nil {
			return "", err
		}

		if err := jwtUserClaims.Valid(); err != nil {
			return "", err
		}

		return jwtUserClaims.UserID, nil
	default:
		return "", errors.New("invalid access token")
	}
}
