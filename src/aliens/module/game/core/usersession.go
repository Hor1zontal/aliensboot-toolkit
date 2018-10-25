/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/7/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"time"
	"aliens/module/game/core/manager"
)

func newUserSession(uid int64) *UserSession {
	dataManager := manager.NewUserDataManager(uid)
	session := &UserSession{UserDataManager:dataManager, lastActiveTime: time.Now()}
	return session
}

type UserSession struct {
	*manager.UserDataManager
	lastActiveTime time.Time //上次活跃时间 没有要进行释放
}

