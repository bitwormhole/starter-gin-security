// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	gate0x423a60 "github.com/bitwormhole/starter-gin-security/gate"
	bypass0x0becd7 "github.com/bitwormhole/starter-gin-security/gate/bypass"
	element0x711bd8 "github.com/bitwormhole/starter-gin-security/gate/element"
	security0xf61d7a "github.com/bitwormhole/starter-security/security"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
	time0x714eea "time"
)

type pComBypassAuthorizer struct {
	instance *bypass0x0becd7.BypassAuthorizer
	 markup0x23084a.Component `class:"security-authorizer" initMethod:"Init"`
	BypassPropertiesResName string `inject:"${security.filter.bypass-properties}"`
	Context application0x67f6c5.Context `inject:"context"`
	acceptPathPrefixTable map[string]bool ``
	acceptPathStringTable map[string]bool ``
	emptyTable map[string]bool ``
}


type pComDefaultAuthenticationManager struct {
	instance *element0x711bd8.DefaultAuthenticationManager
	 markup0x23084a.Component `id:"security-authentication-manager"`
	AuthenticatorList []security0xf61d7a.Authenticator `inject:".security-authenticator"`
}


type pComDefaultAuthorizationManager struct {
	instance *element0x711bd8.DefaultAuthorizationManager
	 markup0x23084a.Component `id:"security-authorization-manager"`
	AuthorizerList []security0xf61d7a.Authorizer `inject:".security-authorizer"`
}


type pComDefaultSecurityContext struct {
	instance *element0x711bd8.DefaultSecurityContext
	 markup0x23084a.Component `id:"security-context" class:"security-context"`
	AuthenticationManager security0xf61d7a.AuthenticationManager `inject:"#security-authentication-manager"`
	AuthorizationManager security0xf61d7a.AuthorizationManager `inject:"#security-authorization-manager"`
	SubjectManager security0xf61d7a.SubjectManager `inject:"#security-subject-manager"`
	SessionIDGenerator security0xf61d7a.SessionIDGenerator `inject:"#security-session-id-generator"`
	SessionFactory security0xf61d7a.SessionFactory `inject:"#security-session-factory"`
	SessionManager security0xf61d7a.SessionManager `inject:"#security-session-manager"`
}


type pComDefaultSessionFactory struct {
	instance *element0x711bd8.DefaultSessionFactory
	 markup0x23084a.Component `id:"security-session-factory"`
	Context security0xf61d7a.Context `inject:"#security-context"`
}


type pComDefaultSessionIDGenerator struct {
	instance *element0x711bd8.DefaultSessionIDGenerator
	 markup0x23084a.Component `id:"security-session-id-generator" initMethod:"Init"`
	MinSessionIDLength int `inject:"${security.session-id.min-length}"`
	index int64 ``
	time0 time0x714eea.Time ``
	prevPart string ``
}


type pComDefaultSessionManager struct {
	instance *element0x711bd8.DefaultSessionManager
	 markup0x23084a.Component `id:"security-session-manager"`
	sessionTable map[string]security0xf61d7a.Session ``
}


type pComDefaultSubjectManager struct {
	instance *element0x711bd8.DefaultSubjectManager
	 markup0x23084a.Component `id:"security-subject-manager"`
	Context security0xf61d7a.Context `inject:"#security-context"`
}


type pComKeeperFilter struct {
	instance *gate0x423a60.KeeperFilter
	 markup0x23084a.Component `class:"rest-controller"`
	Context security0xf61d7a.Context `inject:"#security-context"`
	Order int `inject:"${security.gin-filter.order}"`
}

