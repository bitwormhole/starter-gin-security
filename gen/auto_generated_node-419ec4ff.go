// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	gate0x423a60 "github.com/bitwormhole/starter-gin-security/gate"
	ram0xca67a3 "github.com/bitwormhole/starter-gin-security/gate/support/session/ram"
	controller0x723949 "github.com/bitwormhole/starter-gin-security/gate/web/controller"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComFilter struct {
	instance *gate0x423a60.Filter
	 markup0x23084a.Component `class:"rest-controller keeper-configurer"`
	Enabled bool `inject:"${security.gin-filter.enabled}"`
	TokenHeaderName string `inject:"${security.http-header.token.name}"`
	SetTokenHeaderName string `inject:"${security.http-header.settoken.name}"`
}


type pComSecurityInterceptorRegistry struct {
	instance *gate0x423a60.SecurityInterceptorRegistry
	 markup0x23084a.Component `class:"rest-interceptor-registry"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
	Permissions gate0x423a60.PermissionManager `inject:"#security-gate-permission-manager"`
}


type pComPermissionManagerImpl struct {
	instance *gate0x423a60.PermissionManagerImpl
	 markup0x23084a.Component `id:"security-gate-permission-manager" initMethod:"Init" `
	Context application0x67f6c5.Context `inject:"context"`
	ResName string `inject:"${security.permissions.properties.name}"`
}


type pComTheRAMSessionProvider struct {
	instance *ram0xca67a3.TheRAMSessionProvider
	 markup0x23084a.Component `class:"keeper-session-provider-registry"`
	HTTPHeaderSetTokenName string `inject:"${security.http-header.settoken.name}"`
	HTTPHeaderTokenName string `inject:"${security.http-header.token.name}"`
}


type pComAuthController struct {
	instance *controller0x723949.AuthController
	 markup0x23084a.Component `class:"rest-controller"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
	MyResponder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComSessionController struct {
	instance *controller0x723949.SessionController
	 markup0x23084a.Component `class:"rest-controller"`
	MyResponder glass0x47343f.MainResponder `inject:"#glass-main-responder"`
	Subjects keeper0x6d39ef.SubjectManager `inject:"#keeper-subject-manager"`
}

