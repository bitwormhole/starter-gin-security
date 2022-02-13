package gate

import (
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter/util"
	"github.com/gin-gonic/gin"
)

type GinAccess struct {
	gc                *gin.Context
	TokenFieldName    string
	SetTokenFieldName string
}

func (inst *GinAccess) _Impl() keeper.Access {
	return inst
}

func (inst *GinAccess) Init(c *gin.Context) (keeper.Access, error) {
	inst.gc = c
	return inst, nil
}

func (inst *GinAccess) Path() string {
	return inst.gc.Request.URL.Path
}

func (inst *GinAccess) PathPattern() string {
	return inst.gc.FullPath()
}

func (inst *GinAccess) Method() string {
	return inst.gc.Request.Method
}

func (inst *GinAccess) GetSessionData() []byte {
	name := inst.TokenFieldName
	value := inst.gc.Request.Header.Get(name)
	b64 := util.Base64FromString(value)
	return b64.Bytes()
}

func (inst *GinAccess) SetSessionData(data []byte) {
	name := inst.SetTokenFieldName
	b64 := util.Base64FromBytes(data)
	inst.gc.Header(name, b64.String())
}
