/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/13
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package games

func NewNoobFightGame() Game {
	game := &NoobFightGame{}
	game.maxPlayer = 10
	return game
}

type NoobFightGame struct {
	*SimpleGame
}