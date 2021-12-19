package element

import (
	"github.com/bitwormhole/starter-security/security"
	"github.com/gin-gonic/gin"
)

// GinAccess 封装了访问者身份要素
type GinAccess struct {
	context *gin.Context

	requestSessionID string
	requestPath      string
	requestMethod    string

	setterHeaderName string
	getterHeaderName string
}

func (inst *GinAccess) _Impl() security.Access {
	return inst
}

func (inst *GinAccess) Init(ctx *gin.Context) security.Access {

	inst.context = ctx

	inst.getterHeaderName = "X-SESSION-ID"
	inst.setterHeaderName = "X-SET-SESSION-ID"

	inst.requestMethod = ctx.Request.Method
	inst.requestPath = ctx.Request.URL.Path
	inst.requestSessionID = ctx.Request.Header.Get(inst.getterHeaderName)

	return inst
}

func (inst *GinAccess) SessionID() string {
	return inst.requestSessionID
}

func (inst *GinAccess) Path() string {
	return inst.requestPath
}

func (inst *GinAccess) Method() string {
	return inst.requestMethod
}

func (inst *GinAccess) SetSessionID(sid string) {
	olderID := inst.requestSessionID
	if sid == olderID {
		return
	}
	key := inst.setterHeaderName
	inst.context.Request.Response.Header.Add(key, sid)
}
