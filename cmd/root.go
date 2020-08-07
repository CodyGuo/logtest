package cmd

import (
	"github.com/CodyGuo/logtest/internal/logs"

	"github.com/CodyGuo/glog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "日志生成器",
	Long:  "日志生成器",
	Run: func(cmd *cobra.Command, args []string) {
		if err := logs.GenerateLog(file, prefix); err != nil {
			glog.Fatal(err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

var file string
var prefix string

func init() {
	rootCmd.Flags().StringVarP(&file, "file", "f", "", "set log file")
	rootCmd.Flags().StringVarP(&prefix, "prefix", "p", "[2006/01/02 15:04:05.000]", "set line PREFIX")
	rootCmd.MarkFlagRequired("file")
}
