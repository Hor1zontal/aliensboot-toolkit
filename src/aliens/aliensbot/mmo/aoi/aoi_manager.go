package aoi

type Manager interface {
	ChangeViewRadius(aoi *AOI, radius float32)
	Enter(aoi *AOI, x, y float32)
	Leave(aoi *AOI)
	Moved(aoi *AOI, x, y float32)
}
