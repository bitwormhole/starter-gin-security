package controller

import (
	"net/http"
	"time"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-restful/api/vo"
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/gin-gonic/gin"
)

// AuthController 登录控制器
type AuthController struct {
	markup.Component `class:"rest-controller"`

	Subjects    keeper.SubjectManager `inject:"#keeper-subject-manager"`
	MyResponder glass.MainResponder   `inject:"#glass-main-responder"`
}

func (inst *AuthController) _Impl() glass.Controller {
	return inst
}

// Init ...
func (inst *AuthController) Init(ec glass.EngineConnection) error {
	ec = ec.RequestMapping("auth")
	ec.Handle(http.MethodPost, "", inst.doPost)
	ec.Handle(http.MethodPut, ":id", inst.doPut)
	return nil
}

// 登录的后续步骤
func (inst *AuthController) doPut(c *gin.Context) {
	request := myAuthRequest{}
	request.wantRxBody = true
	err := request.open(c, inst)
	if err == nil {
		err = request.doPut()
	}
	request.send(err)
}

// 登录的第一步
func (inst *AuthController) doPost(c *gin.Context) {
	request := myAuthRequest{}
	request.wantRxBody = true
	err := request.open(c, inst)
	if err == nil {
		err = request.doPost()
	}
	request.send(err)
}

///////////////////////////////////////////

type myAuthRequest struct {
	parent *AuthController
	gc     *gin.Context

	wantRxBody bool

	resp   glass.Response
	rxBody vo.Auth
	txBody vo.Auth
}

func (inst *myAuthRequest) _Impl() keeper.Authentication {
	return inst
}

func (inst *myAuthRequest) Mechanism() string {
	return inst.rxBody.Auth.Mechanism
}
func (inst *myAuthRequest) User() string {
	return inst.rxBody.Auth.Account
}
func (inst *myAuthRequest) Secret() []byte {
	secret := inst.rxBody.Auth.Secret
	return secret.Bytes()
}

func (inst *myAuthRequest) open(c *gin.Context, p *AuthController) error {
	if inst.wantRxBody {
		body := &inst.rxBody
		c.BindJSON(body)
	}
	inst.gc = c
	inst.parent = p
	return nil
}

func (inst *myAuthRequest) send(err error) {
	resp := &inst.resp
	resp.Context = inst.gc
	resp.Data = &inst.txBody
	resp.Error = err
	inst.parent.MyResponder.Send(resp)
}

func (inst *myAuthRequest) doPost() error {

	ctx := inst.gc
	subject, err := inst.parent.Subjects.GetSubject(ctx)
	if err != nil {
		return err
	}

	// check auth
	identity, err := subject.Login(ctx, inst)
	if err != nil {
		time.Sleep(time.Second)
		return err
	}

	// init session
	err = inst.initSession(subject, identity)
	if err != nil {
		return err
	}

	return nil
}

func (inst *myAuthRequest) doPut() error {
	// TODO ...
	return nil
}

func (inst *myAuthRequest) initSession(subject keeper.Subject, identity keeper.Identity) error {

	session, err := subject.GetSession(true)
	if err != nil {
		return err
	}

	tr := session.BeginTransaction()
	defer tr.Close()

	now := util.Now()
	dst := session.Properties().Setter()

	session.SetIdentity(identity)
	dst.SetBool(keeper.SessionFieldAuthenticated, true)
	dst.SetInt64(keeper.SessionFieldCreatedAt, now.Int64())
	dst.SetInt64(keeper.SessionFieldUpdatedAt, now.Int64())

	tr.Commit()
	return nil
}
