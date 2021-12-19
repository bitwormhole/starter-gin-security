package main

import (
	startergin "github.com/bitwormhole/starter-gin"
	starterginsecurity "github.com/bitwormhole/starter-gin-security"
)

func main() {
	i := startergin.InitGin()
	i.Use(starterginsecurity.Module())
	i.Use(startergin.Module())
	i.Use(startergin.ModuleWithDevtools())
	i.Run()
}
