package boot

import (
	_ "NovelServer/packed"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

func init() {
	err := gtime.SetTimeZone("Asia/Shanghai") //设置系统时区
	if err != nil {
		glog.Error(err)
	}
	logPath := g.Config().GetString("logger.path")
	err = glog.SetPath(logPath)
	if err != nil {
		glog.Error(err)
	}
}
