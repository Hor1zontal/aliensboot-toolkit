package exception

type GameCode int32

const (
	NONE GameCode = iota  //value


	DATABASE_ERROR = 1  //数据库操作异常
)

func GameException(this GameCode) {
	panic(this)
}

func GameException1(err string) {
	panic(err)
}