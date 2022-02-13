package gate

import (
	"crypto/sha1"
	"errors"
	"strings"
	"sync"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

// Permission ...
type Permission interface {
	Check(method, path string) error

	GetMethod() string

	GetPath() string

	GetExpression() string

	GetExpressionSum() string

	AcceptRoles() []string

	AcceptAnonymous() bool
}

// PermissionManager ...
// 【inject:"#security-gate-permission-manager"】
type PermissionManager interface {
	GetPermission(method string, path string, create bool) (Permission, error)
}

////////////////////////////////////////////////////////////////////////////////

type permissionImpl struct {
	manager *PermissionManagerImpl

	acceptRoles []string
	acceptAll   bool
	method      string
	pattern     string
	sum         string
	expression  string
}

func (inst *permissionImpl) _Impl() Permission {
	return inst
}

func (inst *permissionImpl) init(method, path string) {

	exp := inst.computeRequestExpression(method, path)
	sum := inst.computeSum(exp)

	inst.expression = exp
	inst.sum = sum
	inst.method = method
	inst.pattern = path
}

func (inst *permissionImpl) load() error {

	exp := inst.expression
	id := inst.manager.index[exp]

	if id == "" {
		sum := inst.sum
		return errors.New("no permission for expression [" + exp + "], exp.sum=" + sum)
	}

	getter := permissionPropertyGetter{}
	getter.init(inst.manager.props, id)

	method := getter.get("method")
	path := getter.get("path")
	roles := getter.get("roles")

	err := getter.err
	if err != nil {
		return err
	}

	err = inst.Check(method, path)
	if err != nil {
		return err
	}

	roleList := inst.parseRoles(roles)
	inst.acceptRoles = roleList
	inst.acceptAll = inst.hasAnonymous(roleList)
	return nil
}

func (inst *permissionImpl) parseRoles(s string) []string {
	src := strings.Split(s, ",")
	dst := make([]string, 0)
	for _, item := range src {
		item = strings.TrimSpace(item)
		item = strings.ToLower(item)
		if item == "" {
			continue
		}

		dst = append(dst, item)
	}
	return dst
}

func (inst *permissionImpl) hasAnonymous(roles []string) bool {
	for _, item := range roles {
		if item == "anon" || item == "anonymous" || item == "anyone" {
			return true
		}
	}
	return false
}

func (inst *permissionImpl) Check(method, path string) error {
	if method == inst.method && path == inst.pattern {
		return nil
	}
	return errors.New("method or path not match")
}

func (inst *permissionImpl) GetMethod() string {
	return inst.method
}

func (inst *permissionImpl) GetPath() string {
	return inst.pattern
}

func (inst *permissionImpl) GetExpression() string {
	return inst.expression
}

func (inst *permissionImpl) GetExpressionSum() string {
	return inst.sum
}

func (inst *permissionImpl) AcceptRoles() []string {
	return inst.acceptRoles
}

func (inst *permissionImpl) AcceptAnonymous() bool {
	return inst.acceptAll
}

// compute Request Expression 计算请求表达式
func (inst *permissionImpl) computeRequestExpression(method, path string) string {
	return method + ":" + path
}

func (inst *permissionImpl) computeSum(text string) string {
	sum := sha1.Sum([]byte(text))
	return util.StringifyBytes(sum[:])
}

////////////////////////////////////////////////////////////////////////////////

type permissionPropertyGetter struct {
	props collection.Properties
	id    string
	err   error
}

func (inst *permissionPropertyGetter) init(p collection.Properties, id string) {
	inst.props = p
	inst.id = id
	inst.err = nil
}

func (inst *permissionPropertyGetter) get(field string) string {
	name := "permission." + inst.id + "." + field
	value, err := inst.props.GetPropertyRequired(name)
	if err != nil {
		inst.err = err
		return ""
	}
	return value
}

////////////////////////////////////////////////////////////////////////////////

// PermissionManagerImpl ...
type PermissionManagerImpl struct {
	markup.Component `id:"security-gate-permission-manager" initMethod:"Init" `

	Context application.Context `inject:"context"`
	ResName string              `inject:"${security.permissions.properties.name}"`

	index map[string]string // map[ expression ] id
	props collection.Properties
	mutex sync.RWMutex
}

func (inst *PermissionManagerImpl) _Impl() PermissionManager {
	return inst
}

// Init ...
func (inst *PermissionManagerImpl) Init() error {
	return inst.loadProperties()
}

func (inst *PermissionManagerImpl) loadProperties() error {

	name := inst.ResName
	text, err := inst.Context.GetResources().GetText(name)
	if err != nil {
		return err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	inst.index = inst.makeIndex(props, "permission.", ".path")
	inst.props = props
	return nil
}

func (inst *PermissionManagerImpl) makeIndex(props collection.Properties, prefix string, suffix string) map[string]string {
	src := props.Export(nil)
	dst := make(map[string]string)
	for key := range src {
		if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
			id := key[len(prefix) : len(key)-len(suffix)]
			method := props.GetProperty(prefix+id+".method", "")
			path := props.GetProperty(prefix+id+".path", "")
			exp := method + ":" + path
			dst[exp] = id
		}
	}
	return dst
}

// GetPermission ...
func (inst *PermissionManagerImpl) GetPermission(method string, path string, create bool) (Permission, error) {

	// inst.mutex.Lock()
	// defer inst.mutex.Unlock()

	perm := &permissionImpl{manager: inst}
	perm.init(method, path)
	err := perm.load()
	if err != nil {
		return nil, err
	}
	return perm, nil
}

////////////////////////////////////////////////////////////////////////////////
