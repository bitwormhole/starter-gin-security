// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	gate0x423a60 "github.com/bitwormhole/starter-gin-security/gate"
	bypass0x0becd7 "github.com/bitwormhole/starter-gin-security/gate/bypass"
	element0x711bd8 "github.com/bitwormhole/starter-gin-security/gate/element"
	security0xf61d7a "github.com/bitwormhole/starter-security/security"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-bypass0x0becd7.BypassAuthorizer
	cominfobuilder.Next()
	cominfobuilder.ID("com0-bypass0x0becd7.BypassAuthorizer").Class("security-authorizer").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBypassAuthorizer{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: security-authentication-manager
	cominfobuilder.Next()
	cominfobuilder.ID("security-authentication-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultAuthenticationManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: security-authorization-manager
	cominfobuilder.Next()
	cominfobuilder.ID("security-authorization-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultAuthorizationManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: security-context
	cominfobuilder.Next()
	cominfobuilder.ID("security-context").Class("security-context").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultSecurityContext{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: security-session-factory
	cominfobuilder.Next()
	cominfobuilder.ID("security-session-factory").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultSessionFactory{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: security-session-id-generator
	cominfobuilder.Next()
	cominfobuilder.ID("security-session-id-generator").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultSessionIDGenerator{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: security-session-manager
	cominfobuilder.Next()
	cominfobuilder.ID("security-session-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultSessionManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: security-subject-manager
	cominfobuilder.Next()
	cominfobuilder.ID("security-subject-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultSubjectManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com8-gate0x423a60.KeeperFilter
	cominfobuilder.Next()
	cominfobuilder.ID("com8-gate0x423a60.KeeperFilter").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComKeeperFilter{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBypassAuthorizer : the factory of component: com0-bypass0x0becd7.BypassAuthorizer
type comFactory4pComBypassAuthorizer struct {

    mPrototype * bypass0x0becd7.BypassAuthorizer

	
	mBypassPropertiesResNameSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComBypassAuthorizer) init() application.ComponentFactory {

	
	inst.mBypassPropertiesResNameSelector = config.NewInjectionSelector("${security.filter.bypass-properties}",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBypassAuthorizer) newObject() * bypass0x0becd7.BypassAuthorizer {
	return & bypass0x0becd7.BypassAuthorizer {}
}

func (inst * comFactory4pComBypassAuthorizer) castObject(instance application.ComponentInstance) * bypass0x0becd7.BypassAuthorizer {
	return instance.Get().(*bypass0x0becd7.BypassAuthorizer)
}

func (inst * comFactory4pComBypassAuthorizer) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBypassAuthorizer) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBypassAuthorizer) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBypassAuthorizer) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComBypassAuthorizer) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBypassAuthorizer) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.BypassPropertiesResName = inst.getterForFieldBypassPropertiesResNameSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldBypassPropertiesResNameSelector
func (inst * comFactory4pComBypassAuthorizer) getterForFieldBypassPropertiesResNameSelector (context application.InstanceContext) string {
    return inst.mBypassPropertiesResNameSelector.GetString(context)
}

//getterForFieldContextSelector
func (inst * comFactory4pComBypassAuthorizer) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultAuthenticationManager : the factory of component: security-authentication-manager
type comFactory4pComDefaultAuthenticationManager struct {

    mPrototype * element0x711bd8.DefaultAuthenticationManager

	
	mAuthenticatorListSelector config.InjectionSelector

}

func (inst * comFactory4pComDefaultAuthenticationManager) init() application.ComponentFactory {

	
	inst.mAuthenticatorListSelector = config.NewInjectionSelector(".security-authenticator",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultAuthenticationManager) newObject() * element0x711bd8.DefaultAuthenticationManager {
	return & element0x711bd8.DefaultAuthenticationManager {}
}

func (inst * comFactory4pComDefaultAuthenticationManager) castObject(instance application.ComponentInstance) * element0x711bd8.DefaultAuthenticationManager {
	return instance.Get().(*element0x711bd8.DefaultAuthenticationManager)
}

func (inst * comFactory4pComDefaultAuthenticationManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultAuthenticationManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultAuthenticationManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultAuthenticationManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultAuthenticationManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultAuthenticationManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AuthenticatorList = inst.getterForFieldAuthenticatorListSelector(context)
	return context.LastError()
}

//getterForFieldAuthenticatorListSelector
func (inst * comFactory4pComDefaultAuthenticationManager) getterForFieldAuthenticatorListSelector (context application.InstanceContext) []security0xf61d7a.Authenticator {
	list1 := inst.mAuthenticatorListSelector.GetList(context)
	list2 := make([]security0xf61d7a.Authenticator, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(security0xf61d7a.Authenticator)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultAuthorizationManager : the factory of component: security-authorization-manager
type comFactory4pComDefaultAuthorizationManager struct {

    mPrototype * element0x711bd8.DefaultAuthorizationManager

	
	mAuthorizerListSelector config.InjectionSelector

}

func (inst * comFactory4pComDefaultAuthorizationManager) init() application.ComponentFactory {

	
	inst.mAuthorizerListSelector = config.NewInjectionSelector(".security-authorizer",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultAuthorizationManager) newObject() * element0x711bd8.DefaultAuthorizationManager {
	return & element0x711bd8.DefaultAuthorizationManager {}
}

func (inst * comFactory4pComDefaultAuthorizationManager) castObject(instance application.ComponentInstance) * element0x711bd8.DefaultAuthorizationManager {
	return instance.Get().(*element0x711bd8.DefaultAuthorizationManager)
}

func (inst * comFactory4pComDefaultAuthorizationManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultAuthorizationManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultAuthorizationManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultAuthorizationManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultAuthorizationManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultAuthorizationManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AuthorizerList = inst.getterForFieldAuthorizerListSelector(context)
	return context.LastError()
}

//getterForFieldAuthorizerListSelector
func (inst * comFactory4pComDefaultAuthorizationManager) getterForFieldAuthorizerListSelector (context application.InstanceContext) []security0xf61d7a.Authorizer {
	list1 := inst.mAuthorizerListSelector.GetList(context)
	list2 := make([]security0xf61d7a.Authorizer, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(security0xf61d7a.Authorizer)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultSecurityContext : the factory of component: security-context
type comFactory4pComDefaultSecurityContext struct {

    mPrototype * element0x711bd8.DefaultSecurityContext

	
	mAuthenticationManagerSelector config.InjectionSelector
	mAuthorizationManagerSelector config.InjectionSelector
	mSubjectManagerSelector config.InjectionSelector
	mSessionIDGeneratorSelector config.InjectionSelector
	mSessionFactorySelector config.InjectionSelector
	mSessionManagerSelector config.InjectionSelector

}

func (inst * comFactory4pComDefaultSecurityContext) init() application.ComponentFactory {

	
	inst.mAuthenticationManagerSelector = config.NewInjectionSelector("#security-authentication-manager",nil)
	inst.mAuthorizationManagerSelector = config.NewInjectionSelector("#security-authorization-manager",nil)
	inst.mSubjectManagerSelector = config.NewInjectionSelector("#security-subject-manager",nil)
	inst.mSessionIDGeneratorSelector = config.NewInjectionSelector("#security-session-id-generator",nil)
	inst.mSessionFactorySelector = config.NewInjectionSelector("#security-session-factory",nil)
	inst.mSessionManagerSelector = config.NewInjectionSelector("#security-session-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultSecurityContext) newObject() * element0x711bd8.DefaultSecurityContext {
	return & element0x711bd8.DefaultSecurityContext {}
}

func (inst * comFactory4pComDefaultSecurityContext) castObject(instance application.ComponentInstance) * element0x711bd8.DefaultSecurityContext {
	return instance.Get().(*element0x711bd8.DefaultSecurityContext)
}

func (inst * comFactory4pComDefaultSecurityContext) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultSecurityContext) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultSecurityContext) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultSecurityContext) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSecurityContext) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSecurityContext) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AuthenticationManager = inst.getterForFieldAuthenticationManagerSelector(context)
	obj.AuthorizationManager = inst.getterForFieldAuthorizationManagerSelector(context)
	obj.SubjectManager = inst.getterForFieldSubjectManagerSelector(context)
	obj.SessionIDGenerator = inst.getterForFieldSessionIDGeneratorSelector(context)
	obj.SessionFactory = inst.getterForFieldSessionFactorySelector(context)
	obj.SessionManager = inst.getterForFieldSessionManagerSelector(context)
	return context.LastError()
}

//getterForFieldAuthenticationManagerSelector
func (inst * comFactory4pComDefaultSecurityContext) getterForFieldAuthenticationManagerSelector (context application.InstanceContext) security0xf61d7a.AuthenticationManager {

	o1 := inst.mAuthenticationManagerSelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.AuthenticationManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "security-context")
		eb.Set("field", "AuthenticationManager")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.AuthenticationManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldAuthorizationManagerSelector
func (inst * comFactory4pComDefaultSecurityContext) getterForFieldAuthorizationManagerSelector (context application.InstanceContext) security0xf61d7a.AuthorizationManager {

	o1 := inst.mAuthorizationManagerSelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.AuthorizationManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "security-context")
		eb.Set("field", "AuthorizationManager")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.AuthorizationManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSubjectManagerSelector
func (inst * comFactory4pComDefaultSecurityContext) getterForFieldSubjectManagerSelector (context application.InstanceContext) security0xf61d7a.SubjectManager {

	o1 := inst.mSubjectManagerSelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.SubjectManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "security-context")
		eb.Set("field", "SubjectManager")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.SubjectManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSessionIDGeneratorSelector
func (inst * comFactory4pComDefaultSecurityContext) getterForFieldSessionIDGeneratorSelector (context application.InstanceContext) security0xf61d7a.SessionIDGenerator {

	o1 := inst.mSessionIDGeneratorSelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.SessionIDGenerator)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "security-context")
		eb.Set("field", "SessionIDGenerator")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.SessionIDGenerator")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSessionFactorySelector
func (inst * comFactory4pComDefaultSecurityContext) getterForFieldSessionFactorySelector (context application.InstanceContext) security0xf61d7a.SessionFactory {

	o1 := inst.mSessionFactorySelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.SessionFactory)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "security-context")
		eb.Set("field", "SessionFactory")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.SessionFactory")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldSessionManagerSelector
func (inst * comFactory4pComDefaultSecurityContext) getterForFieldSessionManagerSelector (context application.InstanceContext) security0xf61d7a.SessionManager {

	o1 := inst.mSessionManagerSelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.SessionManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "security-context")
		eb.Set("field", "SessionManager")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.SessionManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultSessionFactory : the factory of component: security-session-factory
type comFactory4pComDefaultSessionFactory struct {

    mPrototype * element0x711bd8.DefaultSessionFactory

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComDefaultSessionFactory) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("#security-context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultSessionFactory) newObject() * element0x711bd8.DefaultSessionFactory {
	return & element0x711bd8.DefaultSessionFactory {}
}

func (inst * comFactory4pComDefaultSessionFactory) castObject(instance application.ComponentInstance) * element0x711bd8.DefaultSessionFactory {
	return instance.Get().(*element0x711bd8.DefaultSessionFactory)
}

func (inst * comFactory4pComDefaultSessionFactory) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultSessionFactory) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultSessionFactory) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultSessionFactory) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSessionFactory) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSessionFactory) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComDefaultSessionFactory) getterForFieldContextSelector (context application.InstanceContext) security0xf61d7a.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "security-session-factory")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultSessionIDGenerator : the factory of component: security-session-id-generator
type comFactory4pComDefaultSessionIDGenerator struct {

    mPrototype * element0x711bd8.DefaultSessionIDGenerator

	
	mMinSessionIDLengthSelector config.InjectionSelector

}

func (inst * comFactory4pComDefaultSessionIDGenerator) init() application.ComponentFactory {

	
	inst.mMinSessionIDLengthSelector = config.NewInjectionSelector("${security.session-id.min-length}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultSessionIDGenerator) newObject() * element0x711bd8.DefaultSessionIDGenerator {
	return & element0x711bd8.DefaultSessionIDGenerator {}
}

func (inst * comFactory4pComDefaultSessionIDGenerator) castObject(instance application.ComponentInstance) * element0x711bd8.DefaultSessionIDGenerator {
	return instance.Get().(*element0x711bd8.DefaultSessionIDGenerator)
}

func (inst * comFactory4pComDefaultSessionIDGenerator) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultSessionIDGenerator) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultSessionIDGenerator) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultSessionIDGenerator) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComDefaultSessionIDGenerator) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSessionIDGenerator) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.MinSessionIDLength = inst.getterForFieldMinSessionIDLengthSelector(context)
	return context.LastError()
}

//getterForFieldMinSessionIDLengthSelector
func (inst * comFactory4pComDefaultSessionIDGenerator) getterForFieldMinSessionIDLengthSelector (context application.InstanceContext) int {
    return inst.mMinSessionIDLengthSelector.GetInt(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultSessionManager : the factory of component: security-session-manager
type comFactory4pComDefaultSessionManager struct {

    mPrototype * element0x711bd8.DefaultSessionManager

	

}

func (inst * comFactory4pComDefaultSessionManager) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultSessionManager) newObject() * element0x711bd8.DefaultSessionManager {
	return & element0x711bd8.DefaultSessionManager {}
}

func (inst * comFactory4pComDefaultSessionManager) castObject(instance application.ComponentInstance) * element0x711bd8.DefaultSessionManager {
	return instance.Get().(*element0x711bd8.DefaultSessionManager)
}

func (inst * comFactory4pComDefaultSessionManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultSessionManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultSessionManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultSessionManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSessionManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSessionManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultSubjectManager : the factory of component: security-subject-manager
type comFactory4pComDefaultSubjectManager struct {

    mPrototype * element0x711bd8.DefaultSubjectManager

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComDefaultSubjectManager) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("#security-context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultSubjectManager) newObject() * element0x711bd8.DefaultSubjectManager {
	return & element0x711bd8.DefaultSubjectManager {}
}

func (inst * comFactory4pComDefaultSubjectManager) castObject(instance application.ComponentInstance) * element0x711bd8.DefaultSubjectManager {
	return instance.Get().(*element0x711bd8.DefaultSubjectManager)
}

func (inst * comFactory4pComDefaultSubjectManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultSubjectManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultSubjectManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultSubjectManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSubjectManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultSubjectManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComDefaultSubjectManager) getterForFieldContextSelector (context application.InstanceContext) security0xf61d7a.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "security-subject-manager")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComKeeperFilter : the factory of component: com8-gate0x423a60.KeeperFilter
type comFactory4pComKeeperFilter struct {

    mPrototype * gate0x423a60.KeeperFilter

	
	mContextSelector config.InjectionSelector
	mOrderSelector config.InjectionSelector

}

func (inst * comFactory4pComKeeperFilter) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("#security-context",nil)
	inst.mOrderSelector = config.NewInjectionSelector("${security.gin-filter.order}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComKeeperFilter) newObject() * gate0x423a60.KeeperFilter {
	return & gate0x423a60.KeeperFilter {}
}

func (inst * comFactory4pComKeeperFilter) castObject(instance application.ComponentInstance) * gate0x423a60.KeeperFilter {
	return instance.Get().(*gate0x423a60.KeeperFilter)
}

func (inst * comFactory4pComKeeperFilter) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComKeeperFilter) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComKeeperFilter) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComKeeperFilter) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComKeeperFilter) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComKeeperFilter) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	obj.Order = inst.getterForFieldOrderSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComKeeperFilter) getterForFieldContextSelector (context application.InstanceContext) security0xf61d7a.Context {

	o1 := inst.mContextSelector.GetOne(context)
	o2, ok := o1.(security0xf61d7a.Context)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com8-gate0x423a60.KeeperFilter")
		eb.Set("field", "Context")
		eb.Set("type1", "?")
		eb.Set("type2", "security0xf61d7a.Context")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldOrderSelector
func (inst * comFactory4pComKeeperFilter) getterForFieldOrderSelector (context application.InstanceContext) int {
    return inst.mOrderSelector.GetInt(context)
}




