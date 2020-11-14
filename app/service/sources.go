package service

import (
	"NovelServer/app/model"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

//initSourceData 初始化小说源配置文件
func initSourceData() (sourceConfigInfo map[string]*model.SourceConfig, sourceList []*model.Source) {
	//定义源配置信息文件列表
	var sourceFileList []string
	sourceConfigInfo = make(map[string]*model.SourceConfig)

	sourcePath := g.Cfg().GetString("system.sourceConfigPath")
	sourceFiles := g.Config().GetArray("system.sourceFiles")
	sourceFileList = gconv.Strings(sourceFiles)

	//如果配置文件未配置，将使用目录内全部文件
	if len(sourceFiles) == 0 {
		//获取到源文件列表
		sourceConfFileList, _ := getDataSource(sourcePath)
		sourceFileList = gconv.Strings(sourceConfFileList)
	}

	for i := 0; i < len(sourceFileList); i++ {
		configJson := getSourceConfig(sourcePath, sourceFileList[i])

		configInfo := new(model.SourceConfig)
		search := model.Search{}
		configInfo.Search = search

		if err := configJson.ToStruct(configInfo); err != nil {
			panic(err)
		}

		source := new(model.Source)
		source.SourcesCode = configInfo.SourcesCode
		source.SourcesName = configInfo.SourcesName
		sourceList = append(sourceList, source)

		sourceConfigInfo[configInfo.SourcesCode] = configInfo

	}

	return sourceConfigInfo, sourceList
}

// getDataSource 获取小说源的文件列表
func getDataSource(sourcePath string) ([]string, error) {
	files, err := gfile.DirNames(sourcePath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// getSourceConfig 获取小说源的配置文件
func getSourceConfig(sourcePath, configName string) *gjson.Json {
	sourceFile := sourcePath + "/" + configName
	sc, err := gjson.Load(sourceFile)
	if err != nil {
		glog.Error("加载配置文件出错！", err)
		return nil
	}
	return sc
}
