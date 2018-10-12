package aoi

type xzaoi struct {
	aoi          *AOI
	neighbors    map[*xzaoi]struct{}
	xPrev, xNext *xzaoi
	yPrev, yNext *xzaoi
	markVal      int
}

// XZListAOIManager is an implementation of AOICalculator using XZ lists
type XZListAOIManager struct {
	//aoidist    float32
	xSweepList *xAOIList
	zSweepList *yAOIList
}

// NewXZListAOIManager creates a new XZListAOIManager
func NewXZListAOIManager(aoidist float32) Manager {
	return &XZListAOIManager{
		//aoidist:    aoidist,
		xSweepList: newXAOIList(aoidist),
		zSweepList: newYAOIList(aoidist),
	}
}

// Enter is called when Entity enters Space
func (aoiman *XZListAOIManager) Enter(aoi *AOI, x, y float32) {
	//aoi.viewRadius = aoiman.aoidist

	xzaoi := &xzaoi{
		aoi:       aoi,
		neighbors: map[*xzaoi]struct{}{},
	}
	aoi.x, aoi.y = x, y
	aoi.implData = xzaoi
	aoiman.xSweepList.Insert(xzaoi)
	aoiman.zSweepList.Insert(xzaoi)
	aoiman.adjust(xzaoi)
}

func (aoiman *XZListAOIManager) ChangeViewRadius(aoi *AOI, radius float32) {
	//TODO 后续支持
}

// Leave is called when Entity leaves Space
func (aoiman *XZListAOIManager) Leave(aoi *AOI) {
	xzaoi := aoi.implData.(*xzaoi)
	aoiman.xSweepList.Remove(xzaoi)
	aoiman.zSweepList.Remove(xzaoi)
	aoiman.adjust(xzaoi)
}

// Moved is called when Entity moves in Space
func (aoiman *XZListAOIManager) Moved(aoi *AOI, x, y float32) {
	oldX := aoi.x
	oldY := aoi.y
	aoi.x, aoi.y = x, y
	xzaoi := aoi.implData.(*xzaoi)
	if oldX != x {
		aoiman.xSweepList.Move(xzaoi, oldX)
	}
	if oldY != y {
		aoiman.zSweepList.Move(xzaoi, oldY)
	}
	aoiman.adjust(xzaoi)
}

// adjust is called by Entity to adjust neighbors
func (aoiman *XZListAOIManager) adjust(aoi *xzaoi) {
	aoiman.xSweepList.Mark(aoi)
	aoiman.zSweepList.Mark(aoi)
	// AOI marked twice are neighbors
	for neighbor := range aoi.neighbors {
		if neighbor.markVal == 2 {
			// neighbors kept
			neighbor.markVal = -2 // mark this as neighbor
		} else { // markVal < 2
			// was neighbor, but not any more
			delete(aoi.neighbors, neighbor)
			aoi.aoi.Callback.OnLeaveAOI(neighbor.aoi)
			delete(neighbor.neighbors, aoi)
			neighbor.aoi.Callback.OnLeaveAOI(aoi.aoi)
		}
	}

	// travel in X list again to find all new neighbors, whose markVal == 2
	aoiman.xSweepList.GetClearMarkedNeighbors(aoi)
	// travel in Z list again to unmark all
	aoiman.zSweepList.ClearMark(aoi)
}
