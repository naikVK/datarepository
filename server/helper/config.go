package helper

import (
	"github.com/BurntSushi/toml"
)

// InitConfig initConfig
func InitConfig(fpath string, config interface{}) (toml.MetaData, error) {
	return toml.DecodeFile(fpath, config)
}
