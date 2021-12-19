package element

import (
	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/markup"
)

type DefaultSessionFactory struct {
	markup.Component `id:"security-session-factory"`

	Context security.Context `inject:"#security-context"`
}

func (inst *DefaultSessionFactory) _Impl() security.SessionFactory {
	return inst
}

func (inst *DefaultSessionFactory) CreateSession() security.Session {
	sid := inst.Context.GetSessionIDGenerator().GenerateID()
	session := &defaultSession{}
	session.init(sid)
	return session
}

////////////////////////////////////////////////////////////////////////////////

func CreateEmptySession() security.Session {
	session := &defaultSession{}
	session.init("000000")
	return session
}

////////////////////////////////////////////////////////////////////////////////

const (
	defaultSessionPropSessionID = "session-id"
	defaultSessionPropAvatar    = "avatar"
	defaultSessionPropUserID    = "uid"
	defaultSessionPropUserUUID  = "uuid"
	defaultSessionPropNickname  = "nickname"
	defaultSessionPropRoles     = "roles"
)

////////////////////////////////////////////////////////////////////////////////

type defaultSession struct {
	_props map[string]string
}

func (inst *defaultSession) _Impl() security.Session {
	return inst
}

func (inst *defaultSession) getProps() map[string]string {
	props := inst._props
	if props == nil {
		props = make(map[string]string)
		inst._props = props
	}
	return props
}

func (inst *defaultSession) init(sid string) security.Session {
	inst.SetProperty(defaultSessionPropSessionID, sid)
	return inst
}

func (inst *defaultSession) GetID() string {
	return inst.GetProperty(defaultSessionPropSessionID)
}

func (inst *defaultSession) GetProperty(name string) string {
	props := inst.getProps()
	return props[name]
}

func (inst *defaultSession) SetProperty(name, value string) {
	props := inst.getProps()
	props[name] = value
}

func (inst *defaultSession) GetIdentity() security.Identity {
	props := inst.getProps()
	ident := &DefaultIdentity{}
	ident._avatar = props[defaultSessionPropAvatar]
	ident._uuid = props[defaultSessionPropUserUUID]
	ident._uid = props[defaultSessionPropUserID]
	ident._nickname = props[defaultSessionPropNickname]
	ident._roles = props[defaultSessionPropRoles]
	return ident
}

func (inst *defaultSession) SetIdentity(ident security.Identity) {
	if ident == nil {
		return
	}
	props := inst.getProps()
	props[defaultSessionPropAvatar] = ident.Avatar()
	props[defaultSessionPropNickname] = ident.Nickname()
	props[defaultSessionPropUserID] = ident.UserID()
	props[defaultSessionPropUserUUID] = ident.UserUUID()
	props[defaultSessionPropRoles] = StringifyRoles(ident.Roles())
}
