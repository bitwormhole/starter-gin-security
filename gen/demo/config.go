package demo

import "github.com/bitwormhole/starter/application"

func ExportConfigForSGSDemo(cb application.ConfigBuilder) error {

	return autoGenConfig(cb)
}
