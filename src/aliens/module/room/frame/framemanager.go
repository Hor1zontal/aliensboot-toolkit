/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/5/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package frame

import (
	"aliens/protocol/framesync"
	"time"
)

func NewFrameManager(framePerSecond uint) *Manager {
	return &Manager{
		framePerSecond: framePerSecond,
		currFrame:      newFrame(1),
		frames:         make(map[uint32]*framesync.Frame),
	}
}

func newFrame(seq uint32) *framesync.Frame {
	return &framesync.Frame{Seq:seq, Commands:[]*framesync.Command{}}
}

type Manager struct {
	framePerSecond uint //逻辑帧数
	currFrame *framesync.Frame //游戏当前帧数
	frames map[uint32]*framesync.Frame //存储所有的帧数据

	timer *time.Timer //逻辑帧驱动定时器
}

func (this *Manager) Start() *time.Timer {
	frameInterval := time.Second / time.Duration(this.framePerSecond)
	this.timer = time.NewTimer(frameInterval)
	return this.timer
}

func (this *Manager) Stop() {
	if this.timer != nil {
		this.timer.Stop()
	}
}

func (this *Manager) GetFrames(seqs []uint32) []*framesync.Frame {
	lostFrames := make([]*framesync.Frame, len(seqs))
	for index, seq := range seqs {
		lostFrames[index] = this.frames[seq]
	}
	return lostFrames
}

//接受命令
func (this *Manager) AcceptCommand(command *framesync.Command) {
	this.currFrame.Commands = append(this.currFrame.Commands, command)
}

//推进下一帧
func (this *Manager) NextFrame() *framesync.Frame {
	lastFrame := this.currFrame
	this.frames[lastFrame.Seq] = lastFrame
	currSeq := lastFrame.Seq + 1
	this.currFrame = newFrame(currSeq)
	return lastFrame
}


