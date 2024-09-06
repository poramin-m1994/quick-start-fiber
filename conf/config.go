package configs

import (
	"fmt"
	// "strings"

	"github.com/spf13/viper"
)

const (
	RunModeDevelop    = "dev"
	RunModeStaging    = "stg"
	RunModeProduction = "prod"
)

var (
	config  = viper.New()
	RunMode = ""
)

func init() {
	// config.SetConfigName("config.yaml")
	// config.AddConfigPath("./configs")
	// config.SetConfigType("yaml")
	// config.SetEnvPrefix("QuickStart")
	// config.AutomaticEnv()
	// // แปลง _ underscore ใน env เป็น . dot notation ใน viper
	// config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// err := config.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %s \n", err))
	// }
	// getRunMode()
}

func getRunMode() {
	RunMode = config.GetString("runmode")
}

func getWithModeName(name string) string {
	if RunMode == "" {
		return name
	}
	return fmt.Sprintf("%s.%s", RunMode, name)
}

func SetConfigName(name string) {
	config.SetConfigName(getWithModeName(name))
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}

func Get(name string) interface{} {
	result := config.Get(getWithModeName(name))
	if result == nil {
		result = config.Get(name)
	}
	return result
}

func GetString(name string) string {
	result := config.GetString(getWithModeName(name))
	if result == "" {
		result = config.GetString(name)
	}
	return result
}

func GetStringSlice(name string) []string {
	result := config.GetStringSlice(getWithModeName(name))
	if len(result) == 0 {
		result = config.GetStringSlice(name)
	}
	return result
}

func GetBool(name string) bool {
	result := config.GetBool(getWithModeName(name))
	//if result == "" {
	//	result = config.GetString(name)
	//}
	return result
}

func GetInt(name string) int {
	result := config.GetInt(getWithModeName(name))
	//if result == "" {
	//	result = config.GetString(name)
	//}
	return result
}

func GetInt32(name string) int32 {
	result := config.GetInt32(getWithModeName(name))
	//if result == "" {
	//	result = config.GetString(name)
	//}
	return result
}

func GetInt54(name string) int64 {
	result := config.GetInt64(getWithModeName(name))
	//if result == "" {
	//	result = config.GetString(name)
	//}
	return result
}

func GetIntSlice(name string) []int {
	result := config.GetIntSlice(getWithModeName(name))
	//if result == "" {
	//	result = config.GetString(name)
	//}
	return result
}
