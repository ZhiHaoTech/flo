/*
 * @Author: Joe<sdauwangzh@163.com>
 * @Date: 2021-09-11 10:11:27
 * @Description:编译
 * @FilePath: /flo/flo_cli/cmd/build.go
 * JoeSay: What does not kill me, makes me stronger
 */
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "编译flo",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)
