package gate

import (
	"context"
	"net/http"

	"github.com/bitwormhole/starter-gin-security/gate/element"
	"github.com/bitwormhole/starter-gin/contexts"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// KeeperFilter 默认的安全过滤器
type KeeperFilter struct {
	markup.Component `class:"rest-controller"`

	Context security.Context `inject:"#security-context"`
	Order   int              `inject:"${security.gin-filter.order}"`
}

func (inst *KeeperFilter) _Impl() glass.Controller {
	return inst
}

// Init 初始化过滤器
func (inst *KeeperFilter) Init(ec glass.EngineConnection) error {
	ec.Filter(inst.Order, inst.doFilter)
	return nil
}

func (inst *KeeperFilter) doFilter(c *gin.Context) {
	ok, err := inst.check(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !ok {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Next()
}

func (inst *KeeperFilter) check(c *gin.Context) (bool, error) {

	ctx, err := contexts.GetContext2(c)
	if err != nil {
		return false, err
	}

	access := inst.getAccess(c)
	subject := inst.getSubject(c)
	session := inst.getSession(access)

	subject.SetAccess(access)
	subject.SetSession(session)

	auth := &element.DefaultAuthorization{}
	auth.SetMethod(access.Method())
	auth.SetPath(access.Path())
	auth.SetIdentity(session.GetIdentity())

	ok := inst.Context.GetAuthorizationManager().Accept(ctx, auth)
	return ok, nil
}

func (inst *KeeperFilter) getAccess(c *gin.Context) security.Access {
	access := &element.GinAccess{}
	access.Init(c)
	return access
}

func (inst *KeeperFilter) getSubject(c context.Context) security.Subject {
	subject := inst.Context.GetSubjectManager().GetSubject(c)
	return subject
}

func (inst *KeeperFilter) getSession(a security.Access) security.Session {
	sid := a.SessionID()
	if len(sid) < 1 {
		return inst.createSession(false)
	}
	man := inst.Context.GetSessionManager()
	session, _ := man.FindSession(sid)
	if session == nil {
		session = inst.createSession(true)
	}
	return session
}

func (inst *KeeperFilter) createSession(regular bool) security.Session {
	if regular {
		return inst.Context.GetSessionFactory().CreateSession()
	}
	return element.CreateEmptySession()
}
