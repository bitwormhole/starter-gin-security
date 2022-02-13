// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	gate0x423a60 "github.com/bitwormhole/starter-gin-security/gate"
	ram0xca67a3 "github.com/bitwormhole/starter-gin-security/gate/support/session/ram"
	controller0x723949 "github.com/bitwormhole/starter-gin-security/gate/web/controller"
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	keeper0x6d39ef "github.com/bitwormhole/starter-security/keeper"
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

	// component: com0-gate0x423a60.Filter
	cominfobuilder.Next()
	cominfobuilder.ID("com0-gate0x423a60.Filter").Class("rest-controller keeper-configurer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-gate0x423a60.SecurityInterceptorRegistry
	cominfobuilder.Next()
	cominfobuilder.ID("com1-gate0x423a60.SecurityInterceptorRegistry").Class("rest-interceptor-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSecurityInterceptorRegistry{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: security-gate-permission-manager
	cominfobuilder.Next()
	cominfobuilder.ID("security-gate-permission-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPermissionManagerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-ram0xca67a3.TheRAMSessionProvider
	cominfobuilder.Next()
	cominfobuilder.ID("com3-ram0xca67a3.TheRAMSessionProvider").Class("keeper-session-provider-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTheRAMSessionProvider{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-controller0x723949.AuthController
	cominfobuilder.Next()
	cominfobuilder.ID("com4-controller0x723949.AuthController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComAuthController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-controller0x723949.SessionController
	cominfobuilder.Next()
	cominfobuilder.ID("com5-controller0x723949.SessionController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComSessionController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFilter : the factory of component: com0-gate0x423a60.Filter
type comFactory4pComFilter struct {

    mPrototype * gate0x423a60.Filter

	
	mEnabledSelector config.InjectionSelector
	mTokenHeaderNameSelector config.InjectionSelector
	mSetTokenHeaderNameSelector config.InjectionSelector

}

func (inst * comFactory4pComFilter) init() application.ComponentFactory {

	
	inst.mEnabledSelector = config.NewInjectionSelector("${security.gin-filter.enabled}",nil)
	inst.mTokenHeaderNameSelector = config.NewInjectionSelector("${security.http-header.token.name}",nil)
	inst.mSetTokenHeaderNameSelector = config.NewInjectionSelector("${security.http-header.settoken.name}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFilter) newObject() * gate0x423a60.Filter {
	return & gate0x423a60.Filter {}
}

func (inst * comFactory4pComFilter) castObject(instance application.ComponentInstance) * gate0x423a60.Filter {
	return instance.Get().(*gate0x423a60.Filter)
}

func (inst * comFactory4pComFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Enabled = inst.getterForFieldEnabledSelector(context)
	obj.TokenHeaderName = inst.getterForFieldTokenHeaderNameSelector(context)
	obj.SetTokenHeaderName = inst.getterForFieldSetTokenHeaderNameSelector(context)
	return context.LastError()
}

//getterForFieldEnabledSelector
func (inst * comFactory4pComFilter) getterForFieldEnabledSelector (context application.InstanceContext) bool {
    return inst.mEnabledSelector.GetBool(context)
}

//getterForFieldTokenHeaderNameSelector
func (inst * comFactory4pComFilter) getterForFieldTokenHeaderNameSelector (context application.InstanceContext) string {
    return inst.mTokenHeaderNameSelector.GetString(context)
}

//getterForFieldSetTokenHeaderNameSelector
func (inst * comFactory4pComFilter) getterForFieldSetTokenHeaderNameSelector (context application.InstanceContext) string {
    return inst.mSetTokenHeaderNameSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSecurityInterceptorRegistry : the factory of component: com1-gate0x423a60.SecurityInterceptorRegistry
type comFactory4pComSecurityInterceptorRegistry struct {

    mPrototype * gate0x423a60.SecurityInterceptorRegistry

	
	mSubjectsSelector config.InjectionSelector
	mPermissionsSelector config.InjectionSelector

}

func (inst * comFactory4pComSecurityInterceptorRegistry) init() application.ComponentFactory {

	
	inst.mSubjectsSelector = config.NewInjectionSelector("#keeper-subject-manager",nil)
	inst.mPermissionsSelector = config.NewInjectionSelector("#security-gate-permission-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSecurityInterceptorRegistry) newObject() * gate0x423a60.SecurityInterceptorRegistry {
	return & gate0x423a60.SecurityInterceptorRegistry {}
}

func (inst * comFactory4pComSecurityInterceptorRegistry) castObject(instance application.ComponentInstance) * gate0x423a60.SecurityInterceptorRegistry {
	return instance.Get().(*gate0x423a60.SecurityInterceptorRegistry)
}

func (inst * comFactory4pComSecurityInterceptorRegistry) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSecurityInterceptorRegistry) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSecurityInterceptorRegistry) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSecurityInterceptorRegistry) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSecurityInterceptorRegistry) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSecurityInterceptorRegistry) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Subjects = inst.getterForFieldSubjectsSelector(context)
	obj.Permissions = inst.getterForFieldPermissionsSelector(context)
	return context.LastError()
}

//getterForFieldSubjectsSelector
func (inst * comFactory4pComSecurityInterceptorRegistry) getterForFieldSubjectsSelector (context application.InstanceContext) keeper0x6d39ef.SubjectManager {

	o1 := inst.mSubjectsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.SubjectManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-gate0x423a60.SecurityInterceptorRegistry")
		eb.Set("field", "Subjects")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPermissionsSelector
func (inst * comFactory4pComSecurityInterceptorRegistry) getterForFieldPermissionsSelector (context application.InstanceContext) gate0x423a60.PermissionManager {

	o1 := inst.mPermissionsSelector.GetOne(context)
	o2, ok := o1.(gate0x423a60.PermissionManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-gate0x423a60.SecurityInterceptorRegistry")
		eb.Set("field", "Permissions")
		eb.Set("type1", "?")
		eb.Set("type2", "gate0x423a60.PermissionManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPermissionManagerImpl : the factory of component: security-gate-permission-manager
type comFactory4pComPermissionManagerImpl struct {

    mPrototype * gate0x423a60.PermissionManagerImpl

	
	mContextSelector config.InjectionSelector
	mResNameSelector config.InjectionSelector

}

func (inst * comFactory4pComPermissionManagerImpl) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)
	inst.mResNameSelector = config.NewInjectionSelector("${security.permissions.properties.name}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPermissionManagerImpl) newObject() * gate0x423a60.PermissionManagerImpl {
	return & gate0x423a60.PermissionManagerImpl {}
}

func (inst * comFactory4pComPermissionManagerImpl) castObject(instance application.ComponentInstance) * gate0x423a60.PermissionManagerImpl {
	return instance.Get().(*gate0x423a60.PermissionManagerImpl)
}

func (inst * comFactory4pComPermissionManagerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPermissionManagerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPermissionManagerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPermissionManagerImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComPermissionManagerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPermissionManagerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.ResName = inst.getterForFieldResNameSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComPermissionManagerImpl) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldResNameSelector
func (inst * comFactory4pComPermissionManagerImpl) getterForFieldResNameSelector (context application.InstanceContext) string {
    return inst.mResNameSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTheRAMSessionProvider : the factory of component: com3-ram0xca67a3.TheRAMSessionProvider
type comFactory4pComTheRAMSessionProvider struct {

    mPrototype * ram0xca67a3.TheRAMSessionProvider

	
	mHTTPHeaderSetTokenNameSelector config.InjectionSelector
	mHTTPHeaderTokenNameSelector config.InjectionSelector

}

func (inst * comFactory4pComTheRAMSessionProvider) init() application.ComponentFactory {

	
	inst.mHTTPHeaderSetTokenNameSelector = config.NewInjectionSelector("${security.http-header.settoken.name}",nil)
	inst.mHTTPHeaderTokenNameSelector = config.NewInjectionSelector("${security.http-header.token.name}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTheRAMSessionProvider) newObject() * ram0xca67a3.TheRAMSessionProvider {
	return & ram0xca67a3.TheRAMSessionProvider {}
}

func (inst * comFactory4pComTheRAMSessionProvider) castObject(instance application.ComponentInstance) * ram0xca67a3.TheRAMSessionProvider {
	return instance.Get().(*ram0xca67a3.TheRAMSessionProvider)
}

func (inst * comFactory4pComTheRAMSessionProvider) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTheRAMSessionProvider) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTheRAMSessionProvider) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTheRAMSessionProvider) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheRAMSessionProvider) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTheRAMSessionProvider) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.HTTPHeaderSetTokenName = inst.getterForFieldHTTPHeaderSetTokenNameSelector(context)
	obj.HTTPHeaderTokenName = inst.getterForFieldHTTPHeaderTokenNameSelector(context)
	return context.LastError()
}

//getterForFieldHTTPHeaderSetTokenNameSelector
func (inst * comFactory4pComTheRAMSessionProvider) getterForFieldHTTPHeaderSetTokenNameSelector (context application.InstanceContext) string {
    return inst.mHTTPHeaderSetTokenNameSelector.GetString(context)
}

//getterForFieldHTTPHeaderTokenNameSelector
func (inst * comFactory4pComTheRAMSessionProvider) getterForFieldHTTPHeaderTokenNameSelector (context application.InstanceContext) string {
    return inst.mHTTPHeaderTokenNameSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComAuthController : the factory of component: com4-controller0x723949.AuthController
type comFactory4pComAuthController struct {

    mPrototype * controller0x723949.AuthController

	
	mSubjectsSelector config.InjectionSelector
	mMyResponderSelector config.InjectionSelector

}

func (inst * comFactory4pComAuthController) init() application.ComponentFactory {

	
	inst.mSubjectsSelector = config.NewInjectionSelector("#keeper-subject-manager",nil)
	inst.mMyResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComAuthController) newObject() * controller0x723949.AuthController {
	return & controller0x723949.AuthController {}
}

func (inst * comFactory4pComAuthController) castObject(instance application.ComponentInstance) * controller0x723949.AuthController {
	return instance.Get().(*controller0x723949.AuthController)
}

func (inst * comFactory4pComAuthController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComAuthController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComAuthController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComAuthController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAuthController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComAuthController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Subjects = inst.getterForFieldSubjectsSelector(context)
	obj.MyResponder = inst.getterForFieldMyResponderSelector(context)
	return context.LastError()
}

//getterForFieldSubjectsSelector
func (inst * comFactory4pComAuthController) getterForFieldSubjectsSelector (context application.InstanceContext) keeper0x6d39ef.SubjectManager {

	o1 := inst.mSubjectsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.SubjectManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com4-controller0x723949.AuthController")
		eb.Set("field", "Subjects")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldMyResponderSelector
func (inst * comFactory4pComAuthController) getterForFieldMyResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mMyResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com4-controller0x723949.AuthController")
		eb.Set("field", "MyResponder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComSessionController : the factory of component: com5-controller0x723949.SessionController
type comFactory4pComSessionController struct {

    mPrototype * controller0x723949.SessionController

	
	mMyResponderSelector config.InjectionSelector
	mSubjectsSelector config.InjectionSelector

}

func (inst * comFactory4pComSessionController) init() application.ComponentFactory {

	
	inst.mMyResponderSelector = config.NewInjectionSelector("#glass-main-responder",nil)
	inst.mSubjectsSelector = config.NewInjectionSelector("#keeper-subject-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComSessionController) newObject() * controller0x723949.SessionController {
	return & controller0x723949.SessionController {}
}

func (inst * comFactory4pComSessionController) castObject(instance application.ComponentInstance) * controller0x723949.SessionController {
	return instance.Get().(*controller0x723949.SessionController)
}

func (inst * comFactory4pComSessionController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComSessionController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComSessionController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComSessionController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSessionController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComSessionController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MyResponder = inst.getterForFieldMyResponderSelector(context)
	obj.Subjects = inst.getterForFieldSubjectsSelector(context)
	return context.LastError()
}

//getterForFieldMyResponderSelector
func (inst * comFactory4pComSessionController) getterForFieldMyResponderSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mMyResponderSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-controller0x723949.SessionController")
		eb.Set("field", "MyResponder")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSubjectsSelector
func (inst * comFactory4pComSessionController) getterForFieldSubjectsSelector (context application.InstanceContext) keeper0x6d39ef.SubjectManager {

	o1 := inst.mSubjectsSelector.GetOne(context)
	o2, ok := o1.(keeper0x6d39ef.SubjectManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-controller0x723949.SessionController")
		eb.Set("field", "Subjects")
		eb.Set("type1", "?")
		eb.Set("type2", "keeper0x6d39ef.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




