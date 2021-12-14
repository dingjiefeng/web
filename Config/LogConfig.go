package Config

type LogConfig struct {
	Level             string   `yaml:"level"`
	OutPutDir         string   `yaml:"outputdir"`
	DirMaxSizeG       string   `yaml:"dirMaxSizeG"`
	DirMaxFileCount   int64    `yaml:"dirMaxFileCount"`
	FileDateFmt       string   `yaml:"fileDateFmt"`
	FileSuffix        string   `yaml:"fileSuffix"`
	FilePrefix        string   `yaml:"filePrefix"`
	FileRotationHour  int64    `yaml:"fileRotationHour"`
	FileRotationCount int64    `yaml:"fileRotationCount"`
	FileMaxSizeM      int64    `yaml:"fileMaxSizeM"`
	ReportCaller      bool     `yaml:"reportCaller"`
	LevelReportCaller []string `yaml:"levelReportCaller"`
	Formatter         string   `yaml:"formatter"`
	WebLogFilePrefox  string   `yaml:"webLogFilePrefox"`
}
