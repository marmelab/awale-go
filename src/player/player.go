package player

type Player struct {
	Number      int
	HumanPlayer bool
	MinPosition int
	MaxPosition int
	MinPick     int
	MaxPick     int
}

func New(number int, humanPlayer bool, pitCount int) Player {
	halfPit := pitCount / 2
	return Player{
		Number:      number,
		HumanPlayer: humanPlayer,
		MinPosition: number * halfPit,
		MaxPosition: (1 + number) * halfPit,
		MinPick:     (1 - number) * halfPit,
		MaxPick:     (2 - number) * halfPit,
	}
}
