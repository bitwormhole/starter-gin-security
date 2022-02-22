package starterginsecurity

import (
	"embed"

	"github.com/bitwormhole/starter"
	startergin "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter-gin-security/gen"
	"github.com/bitwormhole/starter-gin-security/gen/demo"
	startersecurity "github.com/bitwormhole/starter-security"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	theModuleName = "github.com/bitwormhole/starter-gin-security"
	theModuleVer  = "v0.1.3"
	theModuleRev  = 7
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
	mb.Dependency(startergin.Module())
	mb.Dependency(startersecurity.Module())

	return mb.Create()
}

////////////////////////////////////////////////////////////////////////////////

// ModuleForDemo 导出模块【github.com/bitwormhole/starter-gin-security#demo】
func ModuleForDemo() application.Module {

	parent := Module()

	mb := application.ModuleBuilder{}
	mb.Name(parent.GetName() + "#demo")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())
	mb.Resources(parent.GetResources())

	mb.OnMount(demo.ExportConfigForSGSDemo)

	mb.Dependency(parent)
	mb.Dependency(startergin.ModuleWithDevtools())
	// mb.Dependency(starterrestful.Module())
	// mb.Dependency(startersecurity.Module())

	return mb.Create()
}
