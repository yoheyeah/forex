package starter

import "forex/library/files"

var (
	iniSettingFile  = "setting.ini"
	jsonSettingFile = "setting.json"
)

type Content struct {
	ConfigFile string
	App        App
	Server     Server
	Mysql      Mysql
	Mongo      Mongo
	Influx     Influx
	Redis      Redis
	Crawler    Crawler
}

func Default() *Content {
	content := &Content{
		ConfigFile: iniSettingFile,
	}
	files.BindFileToObj(content.ConfigFile, content)
	content.Builder(&content.App)
	content.Builder(&content.Server)

	return content
}

func (m *Content) Builder(b Builder) error {
	b.Builder()
	return nil
}
