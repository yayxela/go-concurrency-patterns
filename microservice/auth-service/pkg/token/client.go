package token

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWT interface {
	GenerateToken(email string) (signedToken string, ttl int64, err error)
	ValidateToken(signedToken string) (claims *Claims, err error)
}

type tokenClient struct {
	exp    int
	issuer string
	secret string
}

func New(ttl int, issuer, secret string) JWT {
	return &tokenClient{
		exp:    ttl,
		issuer: issuer,
		secret: secret,
	}
}

type Claims struct {
	jwt.StandardClaims
}

var (
	ParseErr = errors.New("couldn't parse claims")
)

func (t *tokenClient) GenerateToken(email string) (signedToken string, ttl int64, err error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(t.exp)).Unix(),
			Issuer:    t.issuer,
			Subject:   email,
		},
	}
	ttl = claims.ExpiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(t.secret))
	return
}

func (t *tokenClient) ValidateToken(signedToken string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(t.secret), nil
		})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ParseErr
	}
	return
}
