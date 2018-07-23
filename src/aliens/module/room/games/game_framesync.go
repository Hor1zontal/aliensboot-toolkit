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

import (
	"aliens/protocol/framesync"
	"aliens/network"
)



func NewFrameSyncGame(framePerSecond uint) *FrameSyncGame {
	return &FrameSyncGame{
		TimerGame : &TimerGame{syncPerSecond:framePerSecond},
		currFrame:      newFrame(1),
		frames:         make(map[uint32]*framesync.Frame),
	}
}

func newFrame(seq uint32) *framesync.Frame {
	return &framesync.Frame{Seq:seq, Commands:[]*framesync.Command{}}
}

type FrameSyncGame struct {
	*TimerGame

	currFrame *framesync.Frame //游戏当前帧数

	frames map[uint32]*framesync.Frame //存储所有的帧数据
}

func (this *FrameSyncGame) HandleMessage(message GameMessage) bool {
	request := message.request
	frameMessage, ok := request.(*framesync.Request)
	if !ok {
		return false
	}
	if frameMessage.GetCommand() != nil {
		this.addCommand(frameMessage.GetCommand())
		return true

	} else if frameMessage.GetRequestLostFrame() != nil {
		this.handleLostFrame(message.agent, frameMessage.GetRequestLostFrame().Seq)
		return true
	}
	return false

}

func (this *FrameSyncGame) HandleTimer() {
	frame := this.nextFrame()
	message := &framesync.Response{
		SyncFrame:frame,
	}
	data, _ := message.Marshal()
	this.BroadcastAll(data)
}

func (this *FrameSyncGame) handleLostFrame(player network.Agent, seqs []uint32) {
	lostFrames := make([]*framesync.Frame, len(seqs))
	for index, seq := range seqs {
		lostFrames[index] = this.frames[seq]
	}
	message := &framesync.Response{
		RequestLostFrameRet:&framesync.RequestLostFrameRet{
			Frame:lostFrames,
		},
	}
	data, _ := message.Marshal()
	player.WriteMsg(data)
}

//接受命令
func (this *FrameSyncGame) addCommand(command *framesync.Command) {
	this.currFrame.Commands = append(this.currFrame.Commands, command)
}

//推进下一帧
func (this *FrameSyncGame) nextFrame() *framesync.Frame {
	lastFrame := this.currFrame
	this.frames[lastFrame.Seq] = lastFrame
	currSeq := lastFrame.Seq + 1
	this.currFrame = newFrame(currSeq)
	return lastFrame
}

