/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NoAutoPledgeCmd represents the NoAutoPledge command
var NoAutoPledgeCmd = &cobra.Command{
	Use:   "NoAutoPledge",
	Short: "指定不分配任务机器",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NO_AUTO_PLEDGE list to empty")
		fmt.Println("args:", args)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long:  `查看NO_AUTO_PLEDGE列表`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitor NoAutoPledge list")
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add",
	Long:  `添加NO_AUTO_PLEDGE机器`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add machine ", args)
	},
}

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "del",
	Long:  `删除NO_AUTO_PLEDGE机器`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete machine ", args)
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean",
	Long:  `清空NO_AUTO_PLEDGE 列表`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("NO_AUTO_PLEDGE 列表已清空")
	},
}

func init() {
	rootCmd.AddCommand(NoAutoPledgeCmd)
	NoAutoPledgeCmd.AddCommand(listCmd, addCmd, delCmd, cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// NoAutoPledgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// NoAutoPledgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//var clean bool
	//NoAutoPledgeCmd.Flags().BoolVarP(&clean,"clean","c", true, "清空不分配任务列表")
	//NoAutoPledgeCmd.Flags().BoolP("list","l",true,"查看NO_AUTO_PLEDGE列表")
}
