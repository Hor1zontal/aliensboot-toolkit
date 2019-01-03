/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/5
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package command

import (
	"github.com/KylinHe/aliensboot-toolkit/util"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	DefaultModuleName = "defaultmodule"
)

func init() {
	moduleCmd.AddCommand(moduleAddCmd)
}

var moduleAddCmd = &cobra.Command{
	Use:   "create Ex. create %module%",
	Short: "create initial module code in current path",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		ALIENSBOTHOME := os.Getenv("ALIENSBOT_HOME")
		if ALIENSBOTHOME == "" {
			fmt.Println("can not found env ALENSBOT_HOME")
			return
		}
		config := EnsureProjectConfig()
		addModule(ALIENSBOTHOME, "", config.Name, args[0])
	},
}

func addModule(homePath string, targetHomePath string, packagePath string, moduleName string) {

	srcModulePath := getPath(homePath, "src", "aliens", "testserver", "module", DefaultModuleName)

	targetModulePath := getPath(targetHomePath, "src", packagePath, "module", moduleName)

	srcInfo, err := os.Stat(targetModulePath)
	if err == nil && srcInfo.IsDir() {
		fmt.Errorf("module path already exists : %v", targetModulePath)
		return
	}

	srcConfigPath := getPath(homePath, "copy", "config", "modules", DefaultModuleName+".yml.bak")

	targetConfigPath := getPath(targetHomePath, "config", "modules", moduleName+".yml")

	srcPublicPath := getPath(homePath, "src", "aliens", "testserver", "public", DefaultModuleName+".go")

	targetPublicPath := getPath(targetHomePath, "src", packagePath, "public", moduleName+".go")

	replaceContent := make(map[string]string)
	replaceContent[DefaultModuleName] = moduleName
	replaceContent[DefaultPackagePath] = packagePath

	util.CopyDir(srcModulePath, targetModulePath, replaceContent)
	util.CopyFile(srcConfigPath, targetConfigPath, replaceContent)
	util.CopyFile(srcPublicPath, targetPublicPath, replaceContent)

}
