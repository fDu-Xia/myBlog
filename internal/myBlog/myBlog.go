package myBlog

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "myBlog",
	Short: "myBlog is a blog system",
	Long: `myBlog is a blog system written by golang
                Complete documentation is available at https://github.com/fDu-Xia/myBlog`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
	SilenceUsage: true,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFile, "configs", "c", "", "The path to the blog system configuration file. Empty string for no configuration file.")
}

func run() error {
	// 打印所有的配置项及其值
	settings, _ := json.Marshal(viper.AllSettings())
	fmt.Println(string(settings))
	// 打印 db -> username 配置项的值
	fmt.Println(viper.GetString("db.username"))
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
