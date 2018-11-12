package aoi

type AOI struct {
	x     float32
	y     float32
	tower *tower

	viewRadius float32 //视野范围

	Callback AOICallback

	implData interface{}
	//implData interface{}
}

func (this *AOI) GetViewRadius() float32 {
	return this.viewRadius
}

func NewAOI(data AOICallback, viewRadius float32) *AOI {
	return &AOI{
		viewRadius: viewRadius,
		Callback:   data,
	}
}

type AOICallback interface {
	OnEnterAOI(other *AOI)
	OnLeaveAOI(other *AOI)
}
