package aoi

type AOI struct {
	x        float32
	y        float32
	tower    *tower

	dist float32 //视野范围
	Data Data
	//implData interface{}
}

func NewAOI(data Data, dist float32) *AOI {
	return &AOI{
		dist:  dist,
		Data : data,
	}
}

type Data interface {
	OnEnterAOI(other *AOI)
	OnLeaveAOI(other *AOI)
}
