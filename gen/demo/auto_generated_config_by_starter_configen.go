// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package demo

import (
	demo0x5c47af "github.com/bitwormhole/starter-gin-security/src/main/golang/demo"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-demo0x5c47af.Demo1controller
	cominfobuilder.Next()
	cominfobuilder.ID("com0-demo0x5c47af.Demo1controller").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemo1controller{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemo1controller : the factory of component: com0-demo0x5c47af.Demo1controller
type comFactory4pComDemo1controller struct {

    mPrototype * demo0x5c47af.Demo1controller

	

}

func (inst * comFactory4pComDemo1controller) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDemo1controller) newObject() * demo0x5c47af.Demo1controller {
	return & demo0x5c47af.Demo1controller {}
}

func (inst * comFactory4pComDemo1controller) castObject(instance application.ComponentInstance) * demo0x5c47af.Demo1controller {
	return instance.Get().(*demo0x5c47af.Demo1controller)
}

func (inst * comFactory4pComDemo1controller) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDemo1controller) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDemo1controller) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDemo1controller) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1controller) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1controller) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




