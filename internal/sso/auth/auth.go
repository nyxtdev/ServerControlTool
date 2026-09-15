package auth

import (
	"context"
)

type Auth struct{ Token string }

func (a Auth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"authorization": "Bearer " + a.Token}, nil
}

func (a Auth) RequireTransportSecurity() bool { return true }
