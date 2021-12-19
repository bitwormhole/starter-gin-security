package element

import (
	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/markup"
)

// DefaultSecurityContext 默认的安全上下文
type DefaultSecurityContext struct {
	markup.Component `id:"security-context" class:"security-context"`

	AuthenticationManager security.AuthenticationManager `inject:"#security-authentication-manager"`
	AuthorizationManager  security.AuthorizationManager  `inject:"#security-authorization-manager"`
	SubjectManager        security.SubjectManager        `inject:"#security-subject-manager"`

	SessionIDGenerator security.SessionIDGenerator `inject:"#security-session-id-generator"`
	SessionFactory     security.SessionFactory     `inject:"#security-session-factory"`
	SessionManager     security.SessionManager     `inject:"#security-session-manager"`
}

func (inst *DefaultSecurityContext) _Impl() security.Context {
	return inst
}

func (inst *DefaultSecurityContext) GetAuthenticationManager() security.AuthenticationManager {
	return inst.AuthenticationManager
}

func (inst *DefaultSecurityContext) GetAuthorizationManager() security.AuthorizationManager {
	return inst.AuthorizationManager
}

func (inst *DefaultSecurityContext) GetSubjectManager() security.SubjectManager {
	return inst.SubjectManager
}

func (inst *DefaultSecurityContext) GetSessionIDGenerator() security.SessionIDGenerator {
	return inst.SessionIDGenerator
}

func (inst *DefaultSecurityContext) GetSessionFactory() security.SessionFactory {
	return inst.SessionFactory
}

func (inst *DefaultSecurityContext) GetSessionManager() security.SessionManager {
	return inst.SessionManager
}
