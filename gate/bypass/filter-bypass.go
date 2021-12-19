package bypass

import (
	"context"
	"strings"

	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
)

// BypassAuthorizer ...
type BypassAuthorizer struct {
	markup.Component `class:"security-authorizer" initMethod:"Init"`

	BypassPropertiesResName string              `inject:"${security.filter.bypass-properties}"`
	Context                 application.Context `inject:"context"`

	acceptPathPrefixTable map[string]bool
	acceptFullPathTable   map[string]bool
	emptyTable            map[string]bool
}

func (inst *BypassAuthorizer) _Impl() security.Authorizer {
	return inst
}

// Init ...
func (inst *BypassAuthorizer) Init() error {

	name := inst.BypassPropertiesResName
	text, err := inst.Context.GetResources().GetText(name)
	if err != nil {
		return err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	inst.loadPathTable(props)
	return nil
}

func (inst *BypassAuthorizer) loadPathTable(props collection.Properties) {

	table := props.Export(nil)
	prefixTable := make(map[string]bool)
	nameTable := make(map[string]bool)
	const keyPrefix = "security.bypass."

	for key, path := range table {
		if !strings.HasPrefix(key, keyPrefix) {
			continue
		}
		if strings.HasSuffix(path, "*") {
			path = path[0 : len(path)-1]
			prefixTable[path] = true
		} else {
			nameTable[path] = true
		}
	}

	inst.acceptPathPrefixTable = prefixTable
	inst.acceptFullPathTable = nameTable
}

func (inst *BypassAuthorizer) getTable(table map[string]bool) map[string]bool {
	if table == nil {
		table = inst.emptyTable
		table = make(map[string]bool)
		inst.emptyTable = table
	}
	return table
}

// Supports ...
func (inst *BypassAuthorizer) Supports(ctx context.Context, a security.Authorization) bool {
	return true
}

// Accept ...
func (inst *BypassAuthorizer) Accept(ctx context.Context, a security.Authorization) bool {

	fullpath := a.Path()

	table1 := inst.getTable(inst.acceptFullPathTable)
	ok := table1[fullpath]
	if ok {
		return true
	}

	table2 := inst.getTable(inst.acceptPathPrefixTable)
	for prefix, ok := range table2 {
		if ok && strings.HasPrefix(fullpath, prefix) {
			return true
		}
	}

	return false
}
