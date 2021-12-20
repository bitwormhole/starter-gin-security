package starterginsecurity

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter-gin-security/gen"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName = "github.com/bitwormhole/starter-gin-security"
	theModuleVer  = "v0.0.3"
	theModuleRev  = 3
)

//go:embed src/main/resources
var theMainRes embed.FS

// Module 导出模块【github.com/bitwormhole/starter-gin-security】
func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVer).Revision(theModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	mb.OnMount(gen.ExportConfigForGinSecurity)

	mb.Dependency(starter.Module())

	return mb.Create()
}
