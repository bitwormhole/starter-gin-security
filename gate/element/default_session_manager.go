package element

import (
	"errors"

	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/markup"
)

type DefaultSessionManager struct {
	markup.Component `id:"security-session-manager"`

	sessionTable map[string]security.Session
}

func (inst *DefaultSessionManager) _Impl() security.SessionManager {
	return inst
}

func (inst *DefaultSessionManager) getShortSessionID(sid string) string {
	wantLen := 28
	haveLen := len(sid)
	if haveLen < wantLen {
		return "00000000"
	}
	return sid[0:wantLen]
}

func (inst *DefaultSessionManager) getTable() map[string]security.Session {
	table := inst.sessionTable
	if table == nil {
		table = make(map[string]security.Session)
		inst.sessionTable = table
	}
	return table
}

func (inst *DefaultSessionManager) InsertSession(session security.Session) (security.Session, error) {
	if session == nil {
		return nil, errors.New("bad param, session==nil")
	}
	shortid := inst.getShortSessionID(session.GetID())
	table := inst.getTable()
	older := table[shortid]
	if older != nil {
		return nil, errors.New("the session is exists")
	}
	table[shortid] = session
	return session, nil
}

func (inst *DefaultSessionManager) RemoveSession(sessionid string) error {
	older, err := inst.FindSession(sessionid)
	if err != nil {
		return err
	}
	shortid := inst.getShortSessionID(older.GetID())
	table := inst.getTable()
	table[shortid] = nil
	return nil
}

func (inst *DefaultSessionManager) UpdateSession(sessionid string, session security.Session) (security.Session, error) {
	if session == nil {
		return nil, errors.New("bad param, session==nil")
	}
	older, err := inst.FindSession(sessionid)
	if err != nil {
		return nil, err
	}
	older.SetIdentity(session.GetIdentity())
	return older, nil
}

func (inst *DefaultSessionManager) FindSession(sessionid string) (security.Session, error) {
	shortid := inst.getShortSessionID(sessionid)
	table := inst.getTable()
	session := table[shortid]
	if session == nil {
		return nil, errors.New("no session")
	}
	if session.GetID() != sessionid {
		return nil, errors.New("no session")
	}
	return session, nil
}
