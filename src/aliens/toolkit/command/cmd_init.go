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
	"aliens/toolkit/model"
	"aliens/toolkit/util"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
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
		Name: packagePath,
	}
	writeProjectConfig(targetHomePath, projectConfig)

	srcSrcPath := getPath(homePath, "src", "aliens", "testserver")

	targetSrcPath := getPath(targetHomePath, "src", packagePath)

	srcCopyPath := getPath(homePath, "copy")

	targetCopyPath := getPath(targetHomePath, getCurrentPath())

	//srcConfigPath := getPath(homePath, "data", "config")
	//
	//targetConfigPath := getPath(targetHomePath, "config")
	//
	//
	//srcTemplatesPath := getPath(homePath, "data", "templates")
	//
	//targetTemplatesPath := getPath(targetHomePath,"templates")
	//
	//
	//srcToolPath := getPath(homePath, "data", "tool")
	//
	//targetToolPath := getPath(targetHomePath,"tool")

	replaceContent := make(map[string]string)
	replaceContent[DefaultPackagePath] = packagePath

	util.CopyDir(srcSrcPath, targetSrcPath, replaceContent)

	util.CopyDir(srcCopyPath, targetCopyPath, replaceContent)

	//util.CopyDir(srcConfigPath, targetConfigPath, replaceContent)
	//util.CopyDir(srcTemplatesPath, targetTemplatesPath, replaceContent)
	//util.CopyDir(srcToolPath, targetToolPath, replaceContent)

}

func getCurrentPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
	}
	return path
}
