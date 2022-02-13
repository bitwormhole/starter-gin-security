package ram

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bitwormhole/starter-security/keeper"
	"github.com/bitwormhole/starter-security/keeper/support/session"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/gin-gonic/gin"
)

// TheRAMSessionProvider ...
type TheRAMSessionProvider struct {
	markup.Component `class:"keeper-session-provider-registry"`

	HTTPHeaderSetTokenName string `inject:"${security.http-header.settoken.name}"`
	HTTPHeaderTokenName    string `inject:"${security.http-header.token.name}"`

	sessionStore   mySessionStore
	initialled     bool
	sessionFactory keeper.SessionFactory
	adapterFactory keeper.SessionAdapterFactory
}

func (inst *TheRAMSessionProvider) _Impl() (keeper.SessionProviderRegistry, keeper.SessionProvider) {
	return inst, inst
}

func (inst *TheRAMSessionProvider) init() {

	if inst.initialled {
		return
	}

	dsp := &session.DefaultSessionProvider{}
	dsp.Init()
	sf := dsp.GetSessionFactory()
	af := &mySessionAdapterFactory{provider: inst}

	inst.adapterFactory = af
	inst.sessionFactory = sf
	inst.initialled = true
}

// GetRegistrationList ...
func (inst *TheRAMSessionProvider) GetRegistrationList() []*keeper.SessionProviderRegistration {

	inst.init()

	reg := &keeper.SessionProviderRegistration{}
	reg.Name = "RAM"
	reg.Provider = inst

	return []*keeper.SessionProviderRegistration{reg}
}

// GetSessionFactory ...
func (inst *TheRAMSessionProvider) GetSessionFactory() keeper.SessionFactory {
	return inst.sessionFactory
}

// GetAdapterFactory ...
func (inst *TheRAMSessionProvider) GetAdapterFactory() keeper.SessionAdapterFactory {
	return inst.adapterFactory
}

////////////////////////////////////////////////////////////////////////////////

type mySessionAdapterFactory struct {
	provider *TheRAMSessionProvider
}

func (inst *mySessionAdapterFactory) _Impl() keeper.SessionAdapterFactory {
	return inst
}

func (inst *mySessionAdapterFactory) Create(ctx context.Context) (keeper.SessionAdapter, error) {

	gc, ok := ctx.(*gin.Context)
	if !ok {
		return nil, errors.New("not a gin.Context")
	}

	adapter := &mySessionAdapter{
		context:  ctx,
		gc:       gc,
		provider: inst.provider,
	}

	return adapter, nil
}

////////////////////////////////////////////////////////////////////////////////

type mySessionAdapter struct {
	provider *TheRAMSessionProvider
	context  context.Context
	gc       *gin.Context
	holder   *mySessionHolder
}

func (inst *mySessionAdapter) _Impl() keeper.SessionAdapter {
	return inst
}

func (inst *mySessionAdapter) GetContext() context.Context {
	return inst.context
}

func (inst *mySessionAdapter) Load(s keeper.Session) error {

	key := inst.provider.HTTPHeaderTokenName
	sessionID := inst.gc.Request.Header.Get(key)
	store := &inst.provider.sessionStore

	h, err := store.find(sessionID, true)
	if err != nil {
		newID := store.generateSessionID()
		h, err = store.find(newID, true)
		if err != nil {
			return err
		}
		sessionID = h.id
	}

	props, err := inst.parseProperties(h.data)
	if err != nil {
		return err
	}

	ptable := props.Export(nil)
	s.Properties().Import(ptable)
	h.access()
	return nil
}

func (inst *mySessionAdapter) Store(s keeper.Session) error {

	holder := inst.holder
	if holder == nil {
		return errors.New("no session")
	}

	data1 := holder.data
	data2 := inst.stringifyProperties(s.Properties())
	if data1 == data2 {
		return nil
	}

	key := inst.provider.HTTPHeaderSetTokenName
	inst.gc.Header(key, holder.id)
	holder.data = data2
	holder.update()
	return nil
}

func (inst *mySessionAdapter) stringifyProperties(p collection.Properties) string {
	if p == nil {
		return ""
	}
	table := p.Export(nil)
	list := make([]string, 0)
	for k, v := range table {
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		if k == "" || v == "" {
			continue
		}
		item := k + "=" + v + "\n"
		list = append(list, item)
	}
	sort.Strings(list)
	builder := strings.Builder{}
	for _, item := range list {
		builder.WriteString(item)
	}
	return builder.String()
}

func (inst *mySessionAdapter) parseProperties(s string) (collection.Properties, error) {
	list := strings.Split(s, "\n")
	props := collection.CreateProperties()
	for _, item := range list {
		i := strings.IndexRune(item, '\n')
		k := strings.TrimSpace(item[0:i])
		v := strings.TrimSpace(item[i+1:])
		props.SetProperty(k, v)
	}
	return props, nil
}

////////////////////////////////////////////////////////////////////////////////

type mySessionStore struct {
	table         map[string]*mySessionHolder
	mutex         sync.RWMutex
	indexForIDGen int
}

func (inst *mySessionStore) getTable() map[string]*mySessionHolder {

	// 快速获取
	t := inst.table
	if t != nil {
		return t
	}

	// 加锁
	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	// 安全获取
	t = inst.table
	if t != nil {
		return t
	}

	t = make(map[string]*mySessionHolder)
	inst.table = t
	return t
}

func (inst *mySessionStore) find(id string, create bool) (*mySessionHolder, error) {

	const minSessionIDLength = 20
	if len(id) < minSessionIDLength {
		return nil, errors.New("bad session id: " + id)
	}

	h, err := inst.findWithoutCreate(id)
	if err == nil {
		return h, nil
	} else if !create {
		return nil, err
	}
	return inst.findWithCreate(id)
}

func (inst *mySessionStore) findWithoutCreate(id string) (*mySessionHolder, error) {
	table := inst.getTable()
	inst.mutex.RLock()
	defer inst.mutex.RUnlock()
	h := table[id]
	if h == nil {
		return nil, errors.New("no session with id: " + id)
	}
	return h, nil
}

func (inst *mySessionStore) findWithCreate(id string) (*mySessionHolder, error) {
	table := inst.getTable()
	inst.mutex.Lock()
	defer inst.mutex.Unlock()
	h := table[id]
	if h == nil {
		h = &mySessionHolder{}
		h.id = id
		table[id] = h
	}
	return h, nil
}

func (inst *mySessionStore) generateSessionID() string {

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	// index
	inst.indexForIDGen++
	index := inst.indexForIDGen
	// time
	now := time.Now()
	// nonce
	nonce1 := make([]byte, 20)
	rand.Reader.Read(nonce1)
	nonce2 := util.HexFromBytes(nonce1)

	builder := strings.Builder{}
	builder.WriteString(now.String())
	builder.WriteString(nonce2.String())
	builder.WriteString(strconv.Itoa(index))

	bin := []byte(builder.String())
	sum := sha256.Sum256(bin)
	return util.StringifyBytes(sum[:])
}

////////////////////////////////////////////////////////////////////////////////

type mySessionHolder struct {
	id        string
	data      string
	createdAt util.Time
	updatedAt util.Time
	accessAt  util.Time
}

func (inst *mySessionHolder) init(id string) {
	now := inst.now()
	inst.id = id
	inst.createdAt = now
	inst.updatedAt = now
	inst.accessAt = now
}

func (inst *mySessionHolder) update() {
	now := inst.now()
	inst.updatedAt = now
	inst.accessAt = now
}

func (inst *mySessionHolder) access() {
	inst.accessAt = inst.now()
}

func (inst *mySessionHolder) now() util.Time {
	return util.Now()
}

////////////////////////////////////////////////////////////////////////////////
