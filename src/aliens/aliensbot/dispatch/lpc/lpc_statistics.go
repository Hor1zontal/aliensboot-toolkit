/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/5/10
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package lpc

import (
	"aliens/aliensbot/module/statistics"
	"aliens/aliensbot/module/statistics/constant"
)

var StatisticsHandler = &statisticsHandler{}

type statisticsHandler struct {
}

func (this *statisticsHandler) AddServiceStatistic(service string, no int32, interval float64) {
	statistics.ChanRPC.Go(constant.INTERNAL_STATISTICS_SERVICE_CALL, service, no, interval)
}

func (this *statisticsHandler) AddOnlineStatistic(userCount int, visitorCount int) {
	statistics.ChanRPC.Go(constant.INTERNAL_STATISTICS_ONLINE, userCount, visitorCount)
}
