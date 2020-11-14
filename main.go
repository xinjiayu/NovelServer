package main

import (
	_ "NovelServer/boot"
	"NovelServer/library/version"
	_ "NovelServer/router"
	"github.com/gogf/gf/frame/g"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	g.Server().Run()
}
