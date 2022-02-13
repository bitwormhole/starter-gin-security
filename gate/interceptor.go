package gate

import (
	"context"
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// SecurityInterceptorRegistry 是提供安全特性(授权)的 handler 拦截器的注册器
type SecurityInterceptorRegistry struct {
	markup.Component `class:"rest-interceptor-registry"`

	Subjects    keeper.SubjectManager `inject:"#keeper-subject-manager"`
	Permissions PermissionManager     `inject:"#security-gate-permission-manager"`
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
	parent     *SecurityInterceptorRegistry
	target     gin.HandlerFunc
	expression string
	permission Permission
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

	pattern := c.FullPath()
	method := c.Request.Method
	perm := inst.permission

	if perm == nil {
		p2, err := inst.parent.Permissions.GetPermission(method, pattern, true)
		if err != nil {
			return err
		}
		perm = p2
		inst.permission = p2
		inst.expression = p2.GetExpression()
	} else {
		err := perm.Check(method, pattern)
		if err != nil {
			return err
		}
	}

	if perm.AcceptAnonymous() {
		return nil
	}

	subject, err := inst.parent.Subjects.GetSubject(c)
	if err != nil {
		return err
	}

	return inst.hasPermission(c, subject, perm)
}

func (inst *securityInterceptorInstance) hasPermission(c context.Context, sub keeper.Subject, perm Permission) error {

	// roles := perm.AcceptRoles()

	return nil
}

////////////////////////////////////////////////////////////////////////////////
