package api

type Error struct {
	Message    string
	StatusCode int
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrNilSummonerData = Error{
		Message:    "Nil SummonerData",
		StatusCode: 1,
	}
)
