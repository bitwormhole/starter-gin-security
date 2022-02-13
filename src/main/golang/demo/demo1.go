package demo

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type Demo1controller struct {
	markup.Component `class:"rest-controller"`
}

func (inst *Demo1controller) _Impl() glass.Controller {
	return inst
}

func (inst *Demo1controller) Init(ec glass.EngineConnection) error {

	// 这种方式已弃用
	// g := gate.NewBuilder()
	// g.ForHandler(inst.handle1)
	// g.Reset()

	ec.Handle(http.MethodGet, "abc", inst.handle1)
	ec.Handle(http.MethodPost, "abc", inst.handle2)

	return nil
}

func (inst *Demo1controller) handle1(c *gin.Context) {
	c.JSON(http.StatusOK, "hello1")
}

func (inst *Demo1controller) handle2(c *gin.Context) {
	c.JSON(http.StatusOK, "hello2")
}
