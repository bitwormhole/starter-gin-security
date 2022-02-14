package gate

import (
	"github.com/bitwormhole/starter-security/keeper"
	"github.com/gin-gonic/gin"
)

// GinAccess ...
type GinAccess struct {
	gc *gin.Context

	// TokenFieldName    string
	// SetTokenFieldName string
}

func (inst *GinAccess) _Impl() keeper.Access {
	return inst
}

// Init ...
func (inst *GinAccess) Init(c *gin.Context) (keeper.Access, error) {
	inst.gc = c
	return inst, nil
}

// Path ...
func (inst *GinAccess) Path() string {
	return inst.gc.Request.URL.Path
}

// PathPattern ...
func (inst *GinAccess) PathPattern() string {
	return inst.gc.FullPath()
}

// Method ...
func (inst *GinAccess) Method() string {
	return inst.gc.Request.Method
}

// Params ...
func (inst *GinAccess) Params() map[string]string {
	src := inst.gc.Params
	dst := make(map[string]string)
	for _, item := range src {
		dst[item.Key] = item.Value
	}
	return dst
}

// func (inst *GinAccess) GetSessionData() []byte {
// 	name := inst.TokenFieldName
// 	value := inst.gc.Request.Header.Get(name)
// 	b64 := util.Base64FromString(value)
// 	return b64.Bytes()
// }

// func (inst *GinAccess) SetSessionData(data []byte) {
// 	name := inst.SetTokenFieldName
// 	b64 := util.Base64FromBytes(data)
// 	inst.gc.Header(name, b64.String())
// }
