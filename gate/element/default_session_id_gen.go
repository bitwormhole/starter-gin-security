package element

import (
	"crypto/sha256"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

type DefaultSessionIDGenerator struct {
	markup.Component `id:"security-session-id-generator" initMethod:"Init"`

	MinSessionIDLength int `inject:"${security.session-id.min-length}"` // 256(char)

	index    int64
	time0    time.Time
	prevPart string
}

func (inst *DefaultSessionIDGenerator) _Impl() security.SessionIDGenerator {
	return inst
}

func (inst *DefaultSessionIDGenerator) Init() error {
	inst.time0 = time.Now()
	return nil
}

func (inst *DefaultSessionIDGenerator) GenerateID() string {
	wantLength := inst.MinSessionIDLength
	builder := strings.Builder{}
	for builder.Len() < wantLength {
		part := inst.nextPart()
		builder.WriteString(part)
	}
	return builder.String()
}

func (inst *DefaultSessionIDGenerator) nextPart() string {

	const nl = "\n"
	index := strconv.FormatInt(inst.index, 10)
	now := time.Now().String()
	t0 := inst.time0.String()

	builder := strings.Builder{}
	builder.WriteString(inst.prevPart)
	builder.WriteString(nl)
	builder.WriteString(index)
	builder.WriteString(nl)
	builder.WriteString(now)
	builder.WriteString(nl)
	builder.WriteString(t0)
	builder.WriteString(nl)
	part := inst.hash(builder.String())

	inst.index++
	inst.prevPart = part

	return part
}

func (inst *DefaultSessionIDGenerator) hash(s string) string {
	sum := sha256.Sum256([]byte(s))
	return util.StringifyBytes(sum[:])
}
