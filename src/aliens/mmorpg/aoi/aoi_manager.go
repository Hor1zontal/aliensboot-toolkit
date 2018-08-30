package aoi

type Manager interface {
	Enter(aoi *AOI, x, y float32)
	Leave(aoi *AOI)
	Moved(aoi *AOI, x, y float32)
}