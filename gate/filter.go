package gate

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/contexts"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// Filter 是负责安全事务的 gin 过滤器
type Filter struct {
	markup.Component `class:"rest-controller keeper-configurer"`

	Enabled            bool   `inject:"${security.gin-filter.enabled}"`
	TokenHeaderName    string `inject:"${security.http-header.token.name}"`
	SetTokenHeaderName string `inject:"${security.http-header.settoken.name}"`

	securityContext keeper.SecurityContext
}

func (inst *Filter) _Impl() (glass.Controller, keeper.Configurer) {
	return inst, inst
}

// Configure ...
func (inst *Filter) Configure(c *keeper.Context) error {
	inst.securityContext = c
	return nil
}

// Init ...
func (inst *Filter) Init(ec glass.EngineConnection) error {
	ec.Filter(0, inst.handle)
	return nil
}

func (inst *Filter) handle(c *gin.Context) {
	err := inst.doFilter(c)
	if err != nil {
		msg := err.Error()
		status := http.StatusForbidden
		h := &gin.H{
			"Error": msg,
		}
		c.AbortWithStatusJSON(status, h)
		return
	}
	c.Next()
}

func (inst *Filter) doFilter(c *gin.Context) error {

	checker := securityChecker{}
	checker.parent = inst
	checker.gc = c

	err := checker.initContext()
	if err != nil {
		return err
	}

	err = checker.initAccess()
	if err != nil {
		return err
	}

	err = checker.initAdapter()
	if err != nil {
		return err
	}

	err = checker.initSubject()
	if err != nil {
		return err
	}

	err = checker.initSession()
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type securityChecker struct {
	parent *Filter
	gc     *gin.Context
	holder *keeper.Holder
	ac     *keeper.AccessContext
}

func (inst *securityChecker) initContext() error {

	ctx := inst.gc
	err := contexts.SetupGinContext(ctx)
	if err != nil {
		return err
	}

	h, err := keeper.GetHolder(ctx)
	if err != nil {
		return err
	}

	ac := h.GetAccessContext()
	ac.SecurityContext = inst.parent.securityContext
	ac.Context = ctx

	inst.holder = h
	inst.ac = ac
	return nil
}

func (inst *securityChecker) initAccess() error {

	// access
	access := &GinAccess{}
	_, err := access.Init(inst.gc)
	if err != nil {
		return err
	}

	// security-access
	sa := &keeper.DefaultSecurityAccess{}
	sa.Init(inst.ac)
	sa.Access = access

	// done
	inst.ac.SecurityAccess = sa
	inst.ac.Access = access
	return nil
}

func (inst *securityChecker) initAdapter() error {
	ctx := inst.gc
	sp := inst.parent.securityContext.GetSessionProvider()
	adapter, err := sp.GetAdapterFactory().Create(ctx)
	if err != nil {
		return err
	}
	inst.ac.Adapter = adapter
	return nil
}

func (inst *securityChecker) initSubject() error {
	ctx := inst.gc
	sm := inst.parent.securityContext.GetSubjects()
	subject, err := sm.GetSubject(ctx)
	if err != nil {
		return err
	}
	inst.ac.Subject = subject
	return nil
}

func (inst *securityChecker) initSession() error {

	return nil
}
