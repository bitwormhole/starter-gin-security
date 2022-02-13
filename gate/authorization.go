package gate

import (
	"errors"

	"github.com/bitwormhole/starter-security/keeper"
)

////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////

type myAuthorization struct {
	a keeper.Access
}

func (inst *myAuthorization) _Impl() keeper.Authorization {
	return inst
}

func (inst *myAuthorization) Identity() keeper.Identity {

	//TODO ...
	err := errors.New("no impl")
	panic(err)

	// return nil
}

func (inst *myAuthorization) Method() string {
	return inst.a.Method()
}

func (inst *myAuthorization) Path() string {
	return inst.a.Path()
}

func (inst *myAuthorization) PathPattern() string {
	return inst.a.PathPattern()
}
