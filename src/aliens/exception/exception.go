package exception

type GameCode int32

const (
	NONE GameCode = iota  //value
)

func GameException(this GameCode) {
	panic(this)
}

func GameException1(err string) {
	panic(err)
}