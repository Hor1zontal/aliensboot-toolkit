/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/6
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package template


type Config struct {
	ProtoPath string  //proto文件路径
	Modules []*ModuleConfig
}

//./codegen -proto protocol.proto -template  ../templates/auth_common_handle.template -output ../../module/${MODULENAME}/service/  -prefix 'handle_${}.go'

type ModuleConfig struct {
	Name string
	Outputs []*Output
}

type Output struct {
	Template string
	Output string
	Prefix string
	Overwrite bool
}