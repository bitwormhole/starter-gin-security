package element

import (
	"strings"

	"github.com/bitwormhole/starter-security/security"
)

// DefaultIdentity 默认的用户身份结构
type DefaultIdentity struct {
	_uuid     string
	_uid      string
	_nickname string
	_avatar   string
	_roles    string // list of words
}

func (inst *DefaultIdentity) _Impl() security.Identity {
	return inst
}

// UserID 取用户ID
func (inst *DefaultIdentity) UserID() string {
	return inst._uid
}

// UserUUID 取用户UUID
func (inst *DefaultIdentity) UserUUID() string {
	return inst._uuid
}

// Nickname 取用户昵称
func (inst *DefaultIdentity) Nickname() string {
	return inst._nickname
}

// Avatar 取用户头像
func (inst *DefaultIdentity) Avatar() string {
	return inst._avatar
}

// Roles 取用户角色
func (inst *DefaultIdentity) Roles() []string {
	return ParseRoles(inst._roles)
}

// SetUserID 设置用户ID
func (inst *DefaultIdentity) SetUserID(value string) {
	inst._uid = value
}

// SetUserUUID 设置用户UUID
func (inst *DefaultIdentity) SetUserUUID(value string) {
	inst._uuid = value
}

// SetNickname 设置用户昵称
func (inst *DefaultIdentity) SetNickname(value string) {
	inst._nickname = value
}

// SetAvatar 取用户头像
func (inst *DefaultIdentity) SetAvatar(value string) {
	inst._avatar = value
}

// SetRoles 设置用户角色
func (inst *DefaultIdentity) SetRoles(value []string) {
	inst._roles = StringifyRoles(value)
}

////////////////////////////////////////////////////////////////////////////////

// ParseRoles 解析角色
func ParseRoles(s string) []string {
	array := strings.Split(s, ",")
	size1 := len(array)
	size2 := 0
	i1 := 0
	i2 := 0
	for ; i1 < size1; i1++ {
		item := strings.TrimSpace(array[i1])
		if item == "" {
			continue
		}
		array[i2] = item
		i2++
		size2++
	}
	if size1 == size2 {
		return array
	}
	return array[0:size2]
}

// StringifyRoles 格式化角色
func StringifyRoles(roles []string) string {
	builder := strings.Builder{}
	for _, str := range roles {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		if builder.Len() > 0 {
			builder.WriteRune(',')
		}
		builder.WriteString(str)
	}
	return builder.String()
}
