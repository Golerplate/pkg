package token

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"

	pkgerrors "github.com/golerplate/pkg/errors"
)

type Tokenizer struct {
	signingKey string
}

func NewTokenizer(ctx context.Context, key string) *Tokenizer {
	return &Tokenizer{
		signingKey: key,
	}
}

func (t *Tokenizer) Tokenize(args map[string]interface{}) (string, error) {
	params := make(jwt.MapClaims)
	for k, v := range args {
		params[k] = v
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), params)
	tok, err := token.SignedString([]byte(t.signingKey))
	if err != nil {
		log.Error().Err(err).Msg("pkg.token.tokenizer.Tokenizer.Tokenize: error creating token")
		return "", pkgerrors.NewInternalServerError("cannot create a token")
	}

	return tok, nil
}

func (t *Tokenizer) Parse(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.signingKey), nil
	})

	if err != nil || !token.Valid {
		log.Error().Err(err).Msg("pkg.token.tokenizer.Tokenizer.Parse: error parsing token")
		return nil, pkgerrors.NewInternalServerError("error parsing token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Error().Msg("pkg.token.tokenizer.Tokenizer.Parse: error casting values of parsed token")
		return nil, pkgerrors.NewInternalServerError("error casting values of parsed token")
	}
	return claims, nil
}
