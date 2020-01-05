package impl

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"light-up-backend/common"
	commonProto "light-up-backend/common/proto"
	"light-up-backend/common/utils"
	lightSeeker "light-up-backend/light-seeker-service/proto"
	lighter "light-up-backend/lighter-service/proto"
	"strings"
	"time"
)

var (
	key = []byte("LIGHTmeUP#127")
)

type Service struct {
	LightSeekerClient lightSeeker.LightSeekerService
	LighterClient     lighter.LighterService
}

func (s Service) ValidateLighter(ctx context.Context, token string) (*Claim, error) {
	tokenType, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenType.Claims.(*Claim); ok && tokenType.Valid && claims.UserType == commonProto.UserTypes_LIGHTER {
		return claims, nil
	} else {
		return nil, errors.New("Lighter token is not valid.")
	}
}

func (s Service) ValidateLightSeeker(ctx context.Context, token string) (*Claim, error) {
	tokenType, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenType.Claims.(*Claim); ok && tokenType.Valid && claims.UserType == commonProto.UserTypes_LIGHT_SEEKER {
		return claims, nil
	} else {
		return nil, errors.New("Light seeker token is not valid.")
	}
}

func (s Service) ValidateAdmin(ctx context.Context, token string) (*Claim, error) {
	tokenType, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenType.Claims.(*Claim); ok && tokenType.Valid && claims.UserType == commonProto.UserTypes_ADMIN {
		return claims, nil
	} else {
		return nil, errors.New("Admin token is not valid.")
	}
}

func (s Service) LoginLightSeeker(ctx context.Context, password string, email string) (string, error) {
	emailLoginReq := &commonProto.EmailRequest{
		Email: strings.ToLower(email),
	}
	lightSeekerResponse, err := s.LightSeekerClient.GetLightSeekerByEmail(ctx, emailLoginReq)
	if err != nil {
		return "", errors.Wrap(err, "Could not find the user by email.")
	}
	if lightSeekerResponse.LightSeeker.User.IsValid == false {
		return "", errors.New("User account is not valid anymore.")
	}

	err = utils.CompareHashAndPassword(lightSeekerResponse.LightSeeker.User.Password, password)
	if err != nil {
		return "", errors.Wrap(err, "Wrong password.")
	}

	claim := &Claim{
		UserEmail: lightSeekerResponse.LightSeeker.User.Email,
		UserId:    lightSeekerResponse.LightSeeker.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: defaultTokenExpiry(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    common.AuthenticationServiceName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(key)
}

func (s Service) LoginLighter(ctx context.Context, password, email string) (string, error) {
	emailLoginReq := &commonProto.EmailRequest{
		Email: strings.ToLower(email),
	}
	lighterResponse, err := s.LighterClient.GetLighterByEmail(ctx, emailLoginReq)
	fmt.Print(lighterResponse.Lighter)
	if err != nil {
		return "", errors.Wrap(err, "Could not find the user by email.")
	}
	if lighterResponse.Lighter.User.IsValid == false {
		return "", errors.New("User is not valid anymore.")
	}

	err = utils.CompareHashAndPassword(lighterResponse.Lighter.User.Password, password)
	if err != nil {
		return "", errors.Wrap(err, "Wrong password.")
	}

	claim := &Claim{
		UserEmail: lighterResponse.Lighter.User.Email,
		UserId:    lighterResponse.Lighter.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: defaultTokenExpiry(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    common.AuthenticationServiceName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(key)
}

func defaultTokenExpiry() int64 {
	return time.Now().Add(time.Hour * 24 * 7).Unix()
}

type Claim struct {
	UserEmail string
	UserId    string
	UserType  commonProto.UserTypes
	jwt.StandardClaims
}
