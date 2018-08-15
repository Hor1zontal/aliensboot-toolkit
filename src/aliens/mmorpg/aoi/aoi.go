package aoi

type AOI struct {
	x        float32
	y        float32
	tower    *tower

	dist float32 //视野范围
	Data AOIData
	//implData interface{}
}

func NewAOI(data AOIData, dist float32) *AOI {
	return &AOI{
		dist:  dist,
		Data : data,
	}
}

type AOIData interface {
	OnEnterAOI(other *AOI)
	OnLeaveAOI(other *AOI)
}
