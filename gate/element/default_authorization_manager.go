package element

import (
	"context"

	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

type DefaultAuthorization struct {
	_path   string
	_method string
	_ident  security.Identity
}

func (inst *DefaultAuthorization) _Impl() security.Authorization {
	return inst
}

func (inst *DefaultAuthorization) Identity() security.Identity {
	return inst._ident
}

func (inst *DefaultAuthorization) Method() string {
	return inst._method
}

func (inst *DefaultAuthorization) Path() string {
	return inst._path
}

func (inst *DefaultAuthorization) SetIdentity(id security.Identity) {
	inst._ident = id
}

func (inst *DefaultAuthorization) SetMethod(method string) {
	inst._method = method
}

func (inst *DefaultAuthorization) SetPath(path string) {
	inst._path = path
}

////////////////////////////////////////////////////////////////////////////////

// DefaultAuthorizationManager 默认的授权管理器
type DefaultAuthorizationManager struct {
	markup.Component `id:"security-authorization-manager"`

	AuthorizerList []security.Authorizer `inject:".security-authorizer"`
}

func (inst *DefaultAuthorizationManager) _Impl() security.AuthorizationManager {
	return inst
}

func (inst *DefaultAuthorizationManager) Accept(ctx context.Context, a security.Authorization) bool {
	list := inst.AuthorizerList
	for _, item := range list {
		if !item.Supports(ctx, a) {
			continue
		}
		if item.Accept(ctx, a) {
			return true
		}
	}
	return false
}
