package http

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "http",
		Short: "http",
		Run:   CmdActionRun,
	}
)

func init() {

	//调用Http
	Cmd.Flags().StringP("config", "c", "", "file path")
	Cmd.Flags().String("name", "", "case name")
	_ = Cmd.MarkFlagRequired("config")
	_ = Cmd.MarkFlagRequired("name")
}

func CmdActionRun(cmd *cobra.Command, args []string) {
	configPath := cmd.Flag("config").Value.String()
	requestName := cmd.Flag("name").Value.String()

	//处理命令
	if err := DoHttpRequest(configPath, requestName);err!=nil{
		fmt.Println(err)
	}
}
