package gen

import "github.com/bitwormhole/starter/application"

// ExportConfigForGinSecurity 导出配置
func ExportConfigForGinSecurity(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
