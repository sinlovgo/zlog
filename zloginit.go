package zlog

import (
	"github.com/spf13/viper"
	"path/filepath"
)

//	init path use type yaml
//	if path is empty
//	read default config path: conf/config.yaml
func Init(path string) error {
	return InitType(path, "yaml")
}

//	read default type is json
//	if path is empty
//	read default config path: conf/config.yaml
func InitType(path, cfgType string) error {
	z := ZapLogger{
		ConfigPath: path,
	}

	// initialize configuration file
	if err := z.initConfig(cfgType); err != nil {
		return err
	}

	// initialization log package
	if err := z.initLog(); err != nil {
		return err
	}
	return nil
}

func (z *ZapLogger) initConfig(cfgType string) error {
	if z.ConfigPath != "" {
		viper.SetConfigFile(z.ConfigPath) // If a configuration file is specified, the specified configuration file is parsed
	} else {
		viper.AddConfigPath(filepath.Join("conf")) // If no configuration file is specified, the default configuration file is conf/config.yaml
		viper.SetConfigName("config")
	}
	viper.SetConfigType(cfgType)                 // Set the configuration file format to YAML
	if err := viper.ReadInConfig(); err != nil { // viper read
		return err
	}
	return nil
}
