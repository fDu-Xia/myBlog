package myBlog

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	//默认配置文件名.
	defaultConfigName = "config.yaml"
	projectDir        = "myBlog"
)

// initConfig 设置需要读取的配置文件名、环境变量，并读取配置文件内容到 viper 中.
func initConfig() {
	if configFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(configFile)
	} else {
		// 查找用户主目录
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)
		// 将用 `$HOME` 目录加入到配置文件的搜索路径中
		viper.AddConfigPath(filepath.Join(home, projectDir))
		// 设置配置文件格式为 YAML(YAML格式清晰易读，并且支持复杂的配置结构)
		viper.SetConfigType("yaml")

		// 配置文件名称（没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
		fmt.Println("....")
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取环境变量的前缀为 MYBLOG
	viper.SetEnvPrefix("MYBLOG")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println("Config file not found")
		} else {
			fmt.Printf("Config file is found,but something wrong: %s\n", err.Error())
		}
	}
	// 打印 viper 当前使用的配置文件，方便 Debug.
	_, _ = fmt.Fprintln(os.Stdout, "Using configs file:", viper.ConfigFileUsed())
}
