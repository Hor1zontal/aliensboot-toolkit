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
	"fmt"
	"github.com/KylinHe/aliensboot-toolkit/model"
	"github.com/KylinHe/aliensboot-toolkit/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const (
	DefaultPackagePath = "github.com/KylinHe/aliensboot-server"
	ProjectName = "aliensboot-demo"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [package path], Ex. aliensbot init github.com/KylinHe/aliensboot-server",
	Short: "initial aliensbot project",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(cmd.CommandPath())
		//if len(args)
		if len(args) == 0 {
			cmd.Help()
			return
		}

		ALIENSBOTHOME := os.Getenv("ALIENSBOOT_HOME")
		if ALIENSBOTHOME == "" {
			fmt.Println("can not found env ALENSBOT_HOME")
			return
		}

		initProject(ALIENSBOTHOME, "", args[0])
	},
}

func initProject(homePath string, targetHomePath string, packagePath string) {
	//util.AddFilter()

	projectConfig := &model.ProjectConfig{
		Name: packagePath,
	}

	writeProjectConfig(targetHomePath, projectConfig)

	srcSrcPath := getPath(homePath, ProjectName, "src", DefaultPackagePath)

	targetSrcPath := getPath(targetHomePath, "src", packagePath)

	srcCopyPath := getPath(homePath, ProjectName)

	targetCopyPath := getPath(targetHomePath, getCurrentPath())

	replaceContent := make(map[string]string)

	replaceContent[DefaultPackagePath] = packagePath

	util.CopyDir(srcSrcPath, targetSrcPath, replaceContent, DefaultModuleName)

	util.CopyDir(srcCopyPath, targetCopyPath, replaceContent, DefaultModuleName, "src")

}

func getCurrentPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
	}
	return path
}
