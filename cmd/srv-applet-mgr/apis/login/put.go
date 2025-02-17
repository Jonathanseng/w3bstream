package login

import (
	"context"
	"time"

	base "github.com/machinefi/w3bstream/pkg/depends/base/types"
	"github.com/machinefi/w3bstream/pkg/depends/conf/jwt"
	"github.com/machinefi/w3bstream/pkg/depends/kit/httptransport/httpx"
	"github.com/machinefi/w3bstream/pkg/errors/status"
	"github.com/machinefi/w3bstream/pkg/modules/account"
	"github.com/machinefi/w3bstream/pkg/types"
)

type Login struct {
	httpx.MethodPut
	InBody `in:"body"`
}

type InBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	AccountID types.SFID     `json:"accountID"`
	Token     string         `json:"token"`
	ExpireAt  base.Timestamp `json:"expireAt"`
	Issuer    string         `json:"issuer"`
}

func (r *Login) Path() string { return "/login" }

func (r *Login) Output(ctx context.Context) (interface{}, error) {
	ac, err := account.ValidateAccountByLogin(ctx, r.Username, r.Password)
	if err != nil {
		return nil, err
	}
	j := jwt.MustConfFromContext(ctx)

	token, err := j.GenerateTokenByPayload(ac.AccountID)
	if err != nil {
		return nil, status.InternalServerError.StatusErr().WithDesc(err.Error())
	}

	return &Response{
		AccountID: ac.AccountID,
		Token:     token,
		ExpireAt:  base.Timestamp{Time: time.Now().Add(j.ExpIn.Duration())},
		Issuer:    j.Issuer,
	}, nil
}
