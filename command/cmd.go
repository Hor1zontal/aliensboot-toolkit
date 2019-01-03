/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/10/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package command

import (
	"github.com/KylinHe/aliensboot-toolkit/model"
	"github.com/KylinHe/aliensboot-toolkit/util"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var projectConfig *model.ProjectConfig

var RootCmd = &cobra.Command{
	Use: "aliensbot",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func Execute() {
	projectConfig = readProjectConfig("")

	if err := RootCmd.Execute(); err != nil {
		fmt.Errorf("error: %v", err)
		os.Exit(1)
	}
}

func EnsureProjectConfig() *model.ProjectConfig {
	if projectConfig == nil {
		fmt.Println("invalid project description file 'project.yml'")
		os.Exit(1)
	}
	fmt.Printf("project config %+v \n", projectConfig)
	return projectConfig
}

func writeProjectConfig(targetHomePath string, projectConfig *model.ProjectConfig) {
	projectFilePath := targetHomePath + "project.yml"
	data, _ := yaml.Marshal(projectConfig)
	util.WriteFile(projectFilePath, data)
}

func readProjectConfig(targetHomePath string) *model.ProjectConfig {
	projectFilePath := targetHomePath + "project.yml"
	data := util.ReadFile(projectFilePath)
	if data == nil {
		return nil
	}
	result := &model.ProjectConfig{}
	err := yaml.Unmarshal(data, result)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return result
}

func getPath(basePath string, packages ...string) string {
	result := basePath
	for _, name := range packages {
		if result == "" {
			result = name
		} else {
			result = result + string(filepath.Separator) + name
		}
	}
	return result
}
