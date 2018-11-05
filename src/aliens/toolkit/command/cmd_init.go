/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/10/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>k
 *******************************************************************************/
package command

import (
	"github.com/spf13/cobra"
	"aliens/toolkit/util"
	"os"
	"fmt"
	"aliens/toolkit/model"
)

const (
	DefaultPackagePath = "aliens/testserver"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [package path], Ex. aliensbot init aliens/testserver",
	Short: "initial aliensbot project",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(cmd.CommandPath())
		//if len(args)
		if len(args) == 0 {
			cmd.Help()
			return
		}


		ALIENSBOTHOME := os.Getenv("ALIENSBOT_HOME")
		if ALIENSBOTHOME == "" {
			fmt.Println("can not found env ALENSBOT_HOME")
			return
		}

		initProject(ALIENSBOTHOME, "", args[0])
	},
}



func initProject(homePath string, targetHomePath string, packagePath string) {
	util.AddFilter(DefaultModuleName)

	projectConfig := &model.ProjectConfig{
		Name:packagePath,
	}
	writeProjectConfig(targetHomePath, projectConfig)


	srcSrcPath := getPath(homePath, "src","aliens","testserver")

	targetSrcPath := getPath(targetHomePath, "src", packagePath)

	srcConfigPath := getPath(homePath, "src","aliens","config")

	targetConfigPath := getPath(targetHomePath,"config")

	replaceContent := make(map[string]string)
	replaceContent[DefaultPackagePath] = packagePath

	util.CopyDir(srcSrcPath, targetSrcPath, replaceContent)
	util.CopyDir(srcConfigPath, targetConfigPath, nil)

}

