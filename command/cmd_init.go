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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	DefaultPackagePath = "github.com/KylinHe/aliensboot-server"
	ProjectName = "aliensboot-demo"
)

var templateName string
var moduleNames []string

func init() {
	initCmd.Flags().StringVarP(&templateName, "template", "t", ProjectName, "init template")
	initCmd.Flags().StringSliceVarP(&moduleNames, "modules", "m", []string{}, "add modules")
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
		fmt.Println(moduleNames)
	},
}

func getFilterModules(projectPath string, moduleNames []string) []string{
	var filterModules []string
	srcConfigPath := getPath(projectPath, "config", "modules")
	dir, err := ioutil.ReadDir(srcConfigPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		srcModuleName := strings.Split(fi.Name(),".")[0]
		isFilter := true
		for _, moduleName := range moduleNames {
			if srcModuleName == moduleName {
				isFilter = false
			}
		}
		if isFilter {
			filterModules = append(filterModules, srcModuleName)
		}
	}
	return filterModules
}

func initProject(homePath string, targetHomePath string, packagePath string) {
	//util.AddFilter()

	projectConfig := &model.ProjectConfig{
		Name: packagePath,
	}

	writeProjectConfig(targetHomePath, projectConfig)

	srcSrcPath := getPath(homePath, templateName, "src", DefaultPackagePath)

	targetSrcPath := getPath(targetHomePath, "src", packagePath)

	srcCopyPath := getPath(homePath, templateName)

	targetCopyPath := getPath(targetHomePath, getCurrentPath())

	srcProtocolPath := getPath(homePath, templateName, "src", DefaultPackagePath, "protocol")

	targetProtocolPath := getPath(targetHomePath, "src", packagePath, "protocol")

	replaceContent := make(map[string]string)

	replaceContent[DefaultPackagePath] = packagePath


	filterModules := getFilterModules(srcCopyPath, moduleNames)

	if moduleNames == nil || len(moduleNames) == 0 {
		util.CopyDir(srcSrcPath, targetSrcPath, replaceContent, DefaultModuleName)
		util.CopyDir(srcCopyPath, targetCopyPath, replaceContent, DefaultModuleName, "src")
	} else {
		// copy src dir to target (filter modules and protocol)
		filterModules1 := append(filterModules, "protocol")
		util.CopyDir(srcProtocolPath, targetProtocolPath, replaceContent)
		util.CopyDir(srcSrcPath, targetSrcPath, replaceContent, filterModules1...)

		// copy copy dir to target (filter modules and src)
		filterModules2 := append(filterModules, "src")
		util.CopyDir(srcCopyPath, targetCopyPath, replaceContent, filterModules2...)


	}

}

func getCurrentPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
	}
	return path
}
