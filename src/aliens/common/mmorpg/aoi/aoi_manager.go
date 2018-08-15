package aoi

type AOIManager interface {
	Enter(aoi *AOI, x, y float32)
	Leave(aoi *AOI)
	Moved(aoi *AOI, x, y float32)
}