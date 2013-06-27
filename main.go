package parkour

import (
	"flag"
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

var (
	// toRefreshCfg  = flag.Bool("parkour.refresh-cfg", true, "Refresh configuration file before reading properties")
	env           = flag.String("parkour.env", "dev", "Parkour environment configration")
	configDirPath = flag.String("parkour.cfg-path", "configs", "Configration directory path, defaults to `configs`")
	unlocalized   = flag.Bool("parkour.unlocal", false, "Don't apply local configuration")

	appCfgPath   string
	envCfgPath   string
	localCfgPath string

	appCfg   *toml.TomlTree
	envCfg   *toml.TomlTree
	localCfg *toml.TomlTree

	configFiles []*toml.TomlTree
)

func init() {
	flag.Parse()

	configFiles = []*toml.TomlTree{}

	if !*unlocalized {
		localCfgPath = *configDirPath + "/local.cfg"
		if exist(localCfgPath) {
			localCfg = mustReadCfgFile(localCfgPath)
			configFiles = append(configFiles, localCfg)
		}
	}

	envCfgPath = *configDirPath + "/" + *env + ".cfg"
	if exist(envCfgPath) {
		envCfg = mustReadCfgFile(envCfgPath)
		configFiles = append(configFiles, envCfg)
	}

	appCfgPath = *configDirPath + "/app.cfg"
	if !exist(appCfgPath) {
		panic("config/app.cfg Must Exist.")
	}
	appCfg = mustReadCfgFile(appCfgPath)
	configFiles = append(configFiles, appCfg)
}

func exist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if !os.IsNotExist(err) {
		panic(err)
	}

	return false
}

func GetString(key string) string {
	return getVal(key).(string)
}

func GetBool(key string) bool {
	return getVal(key).(bool)
}

func GetFloat(key string) float64 {
	return getVal(key).(float64)
}

func GetInt(key string) int {
	return int(getVal(key).(int64))
}

func getVal(key string) interface{} {
	for _, cfg := range configFiles {
		val := cfg.Get("parkour." + key)
		if val != nil {
			return val
		}
	}

	panic(fmt.Errorf("Key %s is provided", key))

	return nil
}

func mustReadCfgFile(path string) *toml.TomlTree {
	c, err := toml.LoadFile(path)
	if err != nil {
		panic(err)
	}

	return c
}
