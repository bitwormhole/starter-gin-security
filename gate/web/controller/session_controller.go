package controller

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-restful/api/vo"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/gin-gonic/gin"
)

// SessionController ...
type SessionController struct {
	markup.Component `class:"rest-controller"`

	MyResponder glass.MainResponder   `inject:"#glass-main-responder"`
	Subjects    keeper.SubjectManager `inject:"#keeper-subject-manager"`
}

func (inst *SessionController) _Impl() glass.Controller {
	return inst
}

// Init ...
func (inst *SessionController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("session")
	ec.Handle(http.MethodGet, "", inst.doGet)
	ec.Handle(http.MethodDelete, "", inst.doDelete)
	return nil
}

// 取当前会话信息
func (inst *SessionController) doGet(c *gin.Context) {
	request := mySessionRequest{}
	request.wantRxBody = false
	err := request.open(c, inst)
	if err == nil {
		err = request.doGet()
	}
	request.send(err)
}

// 登出(logout)
func (inst *SessionController) doDelete(c *gin.Context) {
	request := mySessionRequest{}
	err := request.open(c, inst)
	if err == nil {
		err = request.doDelete()
	}
	request.send(err)
}

///////////////////////////////////////////

type mySessionRequest struct {
	parent *SessionController
	gc     *gin.Context

	wantRxBody bool

	resp   glass.Response
	rxBody vo.Session
	txBody vo.Session
}

func (inst *mySessionRequest) open(c *gin.Context, p *SessionController) error {
	inst.gc = c
	inst.parent = p
	return nil
}

func (inst *mySessionRequest) send(err error) {
	resp := &inst.resp
	resp.Context = inst.gc
	resp.Data = &inst.txBody
	inst.parent.MyResponder.Send(resp)
}

func (inst *mySessionRequest) doGet() error {

	ctx := inst.gc
	subject, err := inst.parent.Subjects.GetSubject(ctx)
	if err != nil {
		return err
	}

	session, err := subject.GetSession(false)
	if err != nil {
		return nil
	}

	src := session.Properties().Getter()
	dst := &inst.txBody.Session

	t1 := src.GetInt64(keeper.SessionFieldCreatedAt, 0)
	t2 := src.GetInt64(keeper.SessionFieldUpdatedAt, 0)

	dst.ID = src.GetString(keeper.SessionFieldUserID, "")
	dst.Name = src.GetString(keeper.SessionFieldUserName, "")
	dst.DisplayName = src.GetString(keeper.SessionFieldDisplayName, "")
	dst.Email = src.GetString(keeper.SessionFieldEmail, "")
	dst.Phone = src.GetString(keeper.SessionFieldPhone, "")
	dst.Avatar = src.GetString(keeper.SessionFieldAvatar, "")
	dst.Authenticated = src.GetBool(keeper.SessionFieldAuthenticated, false)
	dst.CreatedAt = util.Time(t1)
	dst.UpdatedAt = util.Time(t2)

	return nil
}

func (inst *mySessionRequest) doDelete() error {

	ctx := inst.gc
	subject, err := inst.parent.Subjects.GetSubject(ctx)
	if err != nil {
		return err
	}

	session, err := subject.GetSession(false)
	if err != nil {
		return err
	}

	tr := session.BeginTransaction()
	defer tr.Close()
	session.Properties().Clear()
	tr.Commit()
	return nil
}
