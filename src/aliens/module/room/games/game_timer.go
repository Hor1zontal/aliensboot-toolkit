/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package games

import "time"

//只需要处理消息转发即可
type TimerGame struct {

	*SimpleGame

	timer *time.Timer //定时器

	syncPerSecond uint //每秒运行时间

	totalTime uint  //游戏总时间 -1 没有时间限制
}

func (this *TimerGame) Start() {
	frameInterval := time.Second / time.Duration(this.syncPerSecond)
	this.timer = time.NewTimer(frameInterval)
}

func (this *TimerGame) Stop() {
	if this.timer != nil {
		this.timer.Stop()
	}
}

func (this *TimerGame) HandleTimer() {

}


func (this *TimerGame) GetTimer() *time.Timer {
	return this.timer
}
