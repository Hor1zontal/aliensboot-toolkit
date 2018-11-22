/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/11/22
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core


//座位编号从0
type Seats []*Player

func NewSeats(seatNum int) Seats {
	return make(Seats, seatNum)
}

//新增玩家
func (seats Seats) Add(player *Player) bool {
	for index, seat := range seats {
		if seat == nil {
			seats[index] = player
			player.Seat = int32(index + 1)
			return false
		}
	}
	return true
}


func (seats Seats) Get(playerID int64) *Player {
	for _, seat := range seats {
		if seat != nil || seat.GetPlayerid() == playerID {
			return seat
		}
	}
	return nil
}

func (seats Seats) IsFull() bool {
	for _, seat := range seats {
		if seat == nil {
			return false
		}
	}
	return true
}


func (seats Seats) Clean() {

}

func (seats Seats) Foreach(callback func(player *Player)) {
	for _, seat := range seats {
		if seat != nil {
			callback(seat)
		}
	}
}

func (seats Seats) IsAllReady() bool {
	for _, seat := range seats {
		if seat == nil || !seat.IsReady(){
			return false
		}
	}
	return true
}





