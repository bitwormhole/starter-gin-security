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
	theModuleVer  = "v0.0.1"
	theModuleRev  = 1
)

//go:embed src/main/resources
var theMainRes embed.FS

// Module 导出模块【】
func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVer).Revision(theModuleRev)
	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
	mb.OnMount(gen.ExportConfigForGinSecurity)

	mb.Dependency(starter.Module())

	return mb.Create()
}
