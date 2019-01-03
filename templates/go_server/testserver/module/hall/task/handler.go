/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/11/14
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package task

import (
	"github.com/KylinHe/aliensboot/log"
	"github.com/KylinHe/aliensboot/module/base"
	"github.com/KylinHe/aliensboot/task"
	"e.coding.net/aliens/aliensboot_testserver/module/hall/match"
)

func Init(skeleton *base.Skeleton) {
	//匹配检查
	cron, err := task.NewCronExpr("*/2 * * * * *")
	if err != nil {
		log.Error("init match timer error : %v", err)
	}
	skeleton.CronFunc(cron, match.Manager.TryMatch)
}
