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
	"e.coding.net/aliens/aliensboot_toolkit/model"
	"e.coding.net/aliens/aliensboot_toolkit/protobuf/template"
	"e.coding.net/aliens/aliensboot_toolkit/util"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	moduleCmd.AddCommand(codeCmd)
}

var codeCmd = &cobra.Command{
	Use:   "gen Ex. gen %module%",
	Short: "auto generate module code",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		config := EnsureProjectConfig()
		GenCode(args[0], config.Name, config.TemplatePath, "")
	},
}

func GenCode(module string, packageName string, templatePath string, rootPath string) {
	protocolPath := getPath(rootPath, "src", packageName, "protocol", "protocol.proto")
	templatePath = getPath(templatePath, "templates", "protocol")

	config := &model.CodeGenConfig{
		Package:      packageName,
		ProtoPath:    protocolPath,
		TemplatePath: templatePath,
		Modules:      []*model.ModuleConfig{},
		//TemplatePath:templatePath,
	}

	moduleName := strings.ToLower(module)

	buildModuleConfig(rootPath, config, moduleName)

	fmt.Sprintf("config data %+v", config)
	template.Convert(config)
}

func buildModuleConfig(rootPath string, config *model.CodeGenConfig, moduleName string) {
	moduleConfig := &model.ModuleConfig{
		Name:    moduleName,
		Outputs: []*model.Output{},
	}

	moduleConfig.Outputs = append(moduleConfig.Outputs, &model.Output{
		Template:  getModuleTemplatePath(config.TemplatePath, moduleName, "service.template"),
		Output:    getPath(rootPath, "src", config.Package, "module", moduleName, "service", "service.go"),
		Overwrite: true,
	})

	moduleConfig.Outputs = append(moduleConfig.Outputs, &model.Output{
		Template:  getModuleTemplatePath(config.TemplatePath, moduleName, "handle.template"),
		Output:    getPath(rootPath, "src", config.Package, "module", moduleName, "service"),
		Prefix:    "handle_${}.go",
		Overwrite: false,
	})

	moduleConfig.Outputs = append(moduleConfig.Outputs, &model.Output{
		Template:  getModuleTemplatePath(config.TemplatePath, moduleName, "rpc.template"),
		Output:    getPath(rootPath, "src", config.Package, "dispatch", "rpc", moduleName+".go"),
		Overwrite: true,
	})

	config.Modules = append(config.Modules, moduleConfig)
}

func getModuleTemplatePath(templateRoot string, module string, name string) string {
	moduleTemplatePath := getPath(templateRoot, module, name)
	exist, _ := util.PathExists(moduleTemplatePath)
	if exist {
		return moduleTemplatePath
	} else {
		return getPath(templateRoot, name)
	}
}
