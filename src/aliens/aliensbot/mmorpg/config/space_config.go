/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/8/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package config

type SpaceConfig struct {
	Type       string  //类型名 用来分类
	MinX       float32 //x轴左边距
	MaxX       float32
	MinY       float32
	MaxY       float32
	TowerRange float32
}

type EntityConfig struct {
}
