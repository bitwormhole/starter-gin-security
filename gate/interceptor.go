package gate

import (
	"errors"
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/users"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// SecurityInterceptorRegistry 是提供安全特性(授权)的 handler 拦截器的注册器
type SecurityInterceptorRegistry struct {
	markup.Component `class:"rest-interceptor-registry"`

	Subjects keeper.SubjectManager `inject:"#keeper-subject-manager"`

	Permissions keeper.PermissionManager `inject:"#keeper-permission-manager"`
}

func (inst *SecurityInterceptorRegistry) _Impl() glass.InterceptorRegistry {
	return inst
}

// GetRegistrationList ...
func (inst *SecurityInterceptorRegistry) GetRegistrationList() []*glass.InterceptorRegistration {

	ir := &glass.InterceptorRegistration{}
	ir.Name = "SecurityInterceptor"
	ir.Order = 0
	ir.Interceptor = &securityInterceptor{parent: inst}

	return []*glass.InterceptorRegistration{ir}
}

////////////////////////////////////////////////////////////////////////////////

// securityInterceptor 是提供安全特性(授权)的 handler 拦截器
type securityInterceptor struct {
	parent *SecurityInterceptorRegistry
}

func (inst *securityInterceptor) _Impl() glass.Interceptor {
	return inst
}

func (inst *securityInterceptor) Intercept(h gin.HandlerFunc) gin.HandlerFunc {
	i := &securityInterceptorInstance{
		target: h,
		parent: inst.parent,
	}
	return i.handle
}

////////////////////////////////////////////////////////////////////////////////

// securityInterceptorInstance 是提供安全特性(授权)的 handler 拦截器的实例
type securityInterceptorInstance struct {
	parent *SecurityInterceptorRegistry
	target gin.HandlerFunc
	// expression string
	permission keeper.PermissionTemplate
}

func (inst *securityInterceptorInstance) handle(c *gin.Context) {
	err := inst.check(c)
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}
	inst.target(c)
}

func (inst *securityInterceptorInstance) check(c *gin.Context) error {
	h, err := keeper.GetHolder(c)
	if err != nil {
		return err
	}
	ac := h.GetAccessContext()
	x := ac.SecurityAccess
	err = inst.authorize(x)
	if err != nil {
		return err
	}
	return inst.verify(x)
}

func (inst *securityInterceptorInstance) checkPerm(x keeper.PermissionTemplate, method string, pattern string) error {

	// todo ...
	return nil
}

// authorize 授权
func (inst *securityInterceptorInstance) authorize(x keeper.SecurityAccess) error {
	_, err := inst.loadPermission(x)
	if err != nil {
		return err
	}
	return nil
}

func (inst *securityInterceptorInstance) loadPermission(x keeper.SecurityAccess) (keeper.Permission, error) {

	pattern := x.PathPattern()
	method := x.Method()
	pt := inst.permission
	ctx := x.GetContext()

	if pt == nil {
		// load perm
		p2, err := inst.parent.Permissions.FindTemplate(ctx, x)
		if err != nil {
			return nil, err
		}
		pt = p2
		inst.permission = p2
		// inst.expression = p2.GetExpression()
	} else {
		err := inst.checkPerm(pt, method, pattern)
		if err != nil {
			return nil, err
		}
	}

	perm, err := pt.LoadPermission(x.Params())
	if err != nil {
		return nil, err
	}

	x.SetPermission(perm)
	return perm, nil
}

// verify 验证授权
func (inst *securityInterceptorInstance) verify(x keeper.SecurityAccess) error {

	perm := x.GetPermission()
	if perm == nil {
		ak := inst.getAccessKey(x)
		return errors.New("no permission to " + ak)
	}

	if perm.AcceptRole(users.RoleAnonymous) || perm.AcceptRole(users.RoleAnyone) {
		return nil
	}

	roles := x.GetRoles()
	if perm.AcceptRoles(roles) {
		return nil
	}

	// 获取身份
	subject := x.GetSubject()
	session, err := subject.GetSession(true)
	if err != nil {
		return err
	}

	roles = session.GetRoles()
	ident := session.GetIdentity()

	if perm.AcceptRoles(roles) {
		return nil
	}

	// check owner
	if perm.AcceptRole(users.RoleOwner) {
		if perm.IsOwner(ident) {
			return nil
		}
	}

	// check firend
	if perm.AcceptRole(users.RoleFriend) {
		if perm.IsFriend(ident) {
			return nil
		}
	}

	return errors.New("forbidden")
}

func (inst *securityInterceptorInstance) getAccessKey(x keeper.SecurityAccess) string {
	pattern := x.PathPattern()
	method := x.Method()
	return method + ":" + pattern
}

////////////////////////////////////////////////////////////////////////////////
