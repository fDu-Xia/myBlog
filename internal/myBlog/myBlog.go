package myBlog

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "myBlog",
	Short: "myBlog is a blog system",
	Long: `myBlog is a blog system written by golang
                Complete documentation is available at https://github.com/fDu-Xia/myBlog`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
	SilenceUsage: true,
	Args: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			if len(arg) > 0 {
				return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
			}
		}

		return nil
	},
}

func run() error {
	fmt.Println("Hello myBlog!")
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
