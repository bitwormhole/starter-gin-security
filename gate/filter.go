package gate

import (
	"errors"
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
	checker.ctxSecurity = inst.securityContext

	// bind context
	err := checker.initContext(c)
	if err != nil {
		return err
	}

	// init access
	err = checker.initAccess()
	if err != nil {
		return err
	}

	// init adapter
	err = checker.initAdapter()
	if err != nil {
		return err
	}

	// do auth
	// bypass
	// 	return checker.auth()
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type securityChecker struct {
	parent      *Filter
	gc          *gin.Context
	holder      *keeper.Holder
	ctxSession  *keeper.SessionContext
	ctxSecurity keeper.SecurityContext
}

func (inst *securityChecker) initContext(c *gin.Context) error {

	err := contexts.SetupGinContext(c)
	if err != nil {
		return err
	}

	h, err := keeper.GetHolder(c)
	if err != nil {
		return err
	}

	sc := h.GetSessionContext()

	inst.gc = c
	inst.ctxSession = sc
	inst.holder = h
	return nil
}

func (inst *securityChecker) initAccess() error {

	ga := &GinAccess{}
	ga.SetTokenFieldName = inst.parent.SetTokenHeaderName
	ga.TokenFieldName = inst.parent.TokenHeaderName

	access, err := ga.Init(inst.gc)
	if err != nil {
		return err
	}
	inst.ctxSession.Access = access
	return nil
}

func (inst *securityChecker) initAdapter() error {
	af := inst.ctxSecurity.GetSessionProvider().GetAdapterFactory()
	adapter, err := af.Create(inst.gc)
	if err != nil {
		return err
	}
	inst.ctxSession.Adapter = adapter
	return nil
}

func (inst *securityChecker) auth() error {

	// p1 := access.Path()
	// p2 := access.PathPattern()
	// vlog.Debug(p1)
	// vlog.Debug(p2)

	ctx := inst.gc
	access := inst.ctxSession.Access
	auth := &myAuthorization{a: access}

	d1 := access.GetSessionData()
	if d1 == nil {
		d1 = []byte{123}
	} else if len(d1) == 0 {
		d1 = []byte{4, 56}
	}
	access.SetSessionData(d1)

	ok := inst.ctxSecurity.GetAuthorizations().Accept(ctx, auth)
	if !ok {
		return errors.New("Forbidden")
	}
	return nil
}
