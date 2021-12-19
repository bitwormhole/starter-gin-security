package element

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/markup"
)

type DefaultSubjectManager struct {
	markup.Component `id:"security-subject-manager"`

	Context security.Context `inject:"#security-context"`
}

func (inst *DefaultSubjectManager) _Impl() security.SubjectManager {
	return inst
}

func (inst *DefaultSubjectManager) GetSubject(ctx context.Context) security.Subject {

	const key = "github.com/bitwormhole/starter-security/security/Subject#binding"

	o1 := ctx.Value(key)
	o2, ok := inst.toSubject(o1)
	if ok {
		return o2
	}

	setter, err := contexts.GetContextSetter(ctx)
	if err != nil {
		panic(err)
	}

	subject, err := inst.createSubject()
	if err != nil {
		panic(err)
	}

	setter.SetValue(key, subject)
	return subject
}

func (inst *DefaultSubjectManager) toSubject(o interface{}) (security.Subject, bool) {
	if o == nil {
		return nil, false
	}
	o2, ok := o.(security.Subject)
	return o2, ok
}

func (inst *DefaultSubjectManager) createSubject() (security.Subject, error) {

	subject := &defaultSubject{}

	return subject, nil
}

////////////////////////////////////////////////////////////////////////////////

const (
	defaultSubjectAuthOK  = "yes"
	defaultSubjectAuthKey = "authenticated"
)

type defaultSubject struct {
	_access  security.Access
	_session security.Session
	_context security.Context
}

func (inst *defaultSubject) _Impl() security.Subject {
	return inst
}

func (inst *defaultSubject) GetAccess() security.Access {
	return inst._access
}

func (inst *defaultSubject) IsAuthenticated() bool {
	const key = defaultSubjectAuthKey
	session := inst.GetSession()
	value := session.GetProperty(key)
	return value == defaultSubjectAuthOK
}

func (inst *defaultSubject) SetAuthenticated(authenticated bool) {
	const key = defaultSubjectAuthKey
	value := ""
	session := inst.GetSession()
	if authenticated {
		value = defaultSubjectAuthOK
	}
	session.SetProperty(key, value)
}

func (inst *defaultSubject) SetAccess(a security.Access) {
	inst._access = a
}

func (inst *defaultSubject) GetSession() security.Session {
	return inst._session
}

func (inst *defaultSubject) SetSession(s security.Session) {
	inst._session = s
}

func (inst *defaultSubject) Login(ctx context.Context, a security.Authentication) (security.Identity, error) {
	authMan := inst._context.GetAuthenticationManager()
	ident, err := authMan.Authenticate(ctx, a)
	if err != nil {
		return nil, err
	}
	session := inst.GetSession()
	inst._access.SetSessionID(session.GetID())
	session.SetIdentity(ident)
	inst.SetAuthenticated(true)
	return ident, err
}

func (inst *defaultSubject) Logout() error {
	session := inst.GetSession()
	ident := &DefaultIdentity{}
	session.SetIdentity(ident)
	inst.SetAuthenticated(false)
	return nil
}

func (inst *defaultSubject) Authorize(ctx context.Context, a security.Access) (bool, error) {

	authMan := inst._context.GetAuthorizationManager()
	session := inst.GetSession()

	auth := &DefaultAuthorization{}
	auth.SetIdentity(session.GetIdentity())
	auth.SetMethod(a.Method())
	auth.SetPath(a.Path())

	if !authMan.Accept(ctx, auth) {
		return false, errors.New("forbidden")
	}

	return true, nil
}
