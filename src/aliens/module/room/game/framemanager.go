/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/5/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package game

import (
	"time"

	"aliens/protocol/room"
)

func NewFrameManager(framePerSecond uint, frameHandle func(frame *room.Frame)) *FrameManager {
	return &FrameManager{
		framePerSecond: framePerSecond,
		currFrame:      newFrame(1),
		frames:         make(map[uint]*room.Frame),
		frameHandle:    frameHandle,
	}
}


func newFrame(seq uint32) *room.Frame {
	return &room.Frame{Seq:seq, Commands:[]*room.Command{}}
}

type FrameManager struct {
	framePerSecond uint //逻辑帧数

	timer *time.Timer //逻辑帧驱动定时器
	channel chan *room.Command  //接受到的游戏指令管道

	currFrame *room.Frame //游戏当前帧数

	frameHandle func(frame *room.Frame) //帧处理器
	frames map[uint]*room.Frame //存储所有的帧数据
}

func (this *FrameManager) Start() {
	frameInterval := time.Second / time.Duration(this.framePerSecond)
	this.channel = make(chan *room.Command, 5)
	this.timer = time.NewTimer(frameInterval)
	go func() {
		for {
			select {
			case command := <-this.channel:
				if command != nil {
					this.acceptCommand(command)
				}
			case <-this.timer.C:
				this.nextFrame()
			}
		}
	}()
}

func (this *FrameManager) AddCommand(command *room.Command) {
	this.channel <- command
}

func (this *FrameManager) Close() {
	close(this.channel)
	this.timer.Stop()
}

//推进下一帧
func (this *FrameManager) nextFrame() {
	lastFrame := this.currFrame
	currSeq := lastFrame.Seq + 1
	this.currFrame = newFrame(currSeq)
	this.frameHandle(lastFrame)
}

//接受命令
func (this *FrameManager) acceptCommand(command *room.Command) {
	this.currFrame.Commands = append(this.currFrame.Commands, command)
}
