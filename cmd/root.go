package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lc-go",
	Short: "LeetCode Go 示例运行器",
	Long: `LeetCode Go 示例运行器

用法: lc-go <命令>
使用 "lc-go --help" 查看所有可用命令。`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
