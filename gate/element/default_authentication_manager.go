package element

import (
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
