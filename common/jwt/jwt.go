package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/wam-lab/base-web-api/common/errno"
	"time"
)

type Token struct {
	Token   string `json:"token"`
	Expire  int64  `json:"expire"`
	Refresh int64  `json:"refresh"`
}

func GenerateToken(exp int64, secret string, claims map[string]interface{}) (*Token, error) {
	c := jwt.MapClaims(claims)

	now := time.Now().Unix()
	c["iat"] = now
	c["exp"] = now + exp

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenStr, err := t.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &Token{
		Token:   tokenStr,
		Expire:  now + exp,
		Refresh: now + exp/2,
	}, nil
}

func ParseToken(tokenStr, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errno.ErrSignMethod
		}

		return []byte(secret), nil
	})
}

func ParseTokenClaims(tokenStr, secret string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := ParseToken(tokenStr, secret)
	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}

	return nil, nil, errno.ErrInvalidToken
}
