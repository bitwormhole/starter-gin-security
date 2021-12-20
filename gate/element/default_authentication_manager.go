package element

import (
	"bytes"
	"context"
	"errors"

	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/markup"
)

// DefaultAuthenticationManager 默认的身份认证管理器
type DefaultAuthenticationManager struct {
	markup.Component `id:"security-authentication-manager"`

	AuthenticatorList []security.Authenticator `inject:".security-authenticator"`
}

func (inst *DefaultAuthenticationManager) _Impl() security.AuthenticationManager {
	return inst
}

func (inst *DefaultAuthenticationManager) Authenticate(ctx context.Context, a security.Authentication) (security.Identity, error) {
	list := inst.AuthenticatorList
	for _, item := range list {
		if !item.Supports(ctx, a) {
			continue
		}
		ident, err := item.Verify(ctx, a)
		if err == nil && ident != nil {
			return ident, nil
		}
	}
	return nil, errors.New("bad authentication")
}

////////////////////////////////////////////////////////////////////////////////

// DefaultAuthentication 默认的验证凭据
type DefaultAuthentication struct {
	_mechanism string
	_user      string
	_secret    []byte
}

func (inst *DefaultAuthentication) _Impl() security.Authentication {
	return inst
}

func (inst *DefaultAuthentication) Mechanism() string {
	return inst._mechanism
}

func (inst *DefaultAuthentication) UserID() string {
	return inst._user
}

func (inst *DefaultAuthentication) UserSecret() []byte {
	return inst._secret
}

func (inst *DefaultAuthentication) SetMechanism(s string) {
	inst._mechanism = s
}

func (inst *DefaultAuthentication) SetUserID(s string) {
	inst._user = s
}

func (inst *DefaultAuthentication) SetUserSecret(b []byte) {
	buffer := bytes.Buffer{}
	if b != nil {
		buffer.Write(b)
	}
	inst._secret = buffer.Bytes()
}

////////////////////////////////////////////////////////////////////////////////
