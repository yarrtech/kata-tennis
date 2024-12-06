package _go

import "math"

type tennis struct {
	firstPlayerScoreTimes  int
	secondPlayerScoreTimes int
	firstPlayerName        string
	secondPlayerName       string
}

var scoreLookup = map[int]string{
	1: "fifteen",
	0: "love",
	2: "thirty",
	3: "forty",
}

func (t *tennis) score() string {
	if t.isScoreDifferent() {
		if t.isReadyForGamePoint() {
			return t.advState()
		}
		return t.lookupScore()
	}
	if t.isDeuce() {
		return t.deuceScore()
	}
	return t.sameScore()
}

func (t *tennis) isReadyForGamePoint() bool {
	return t.firstPlayerScoreTimes > 3 || t.secondPlayerScoreTimes > 3
}

func (t *tennis) advState() string {
	if t.isAdv() {
		return t.advScore()
	}
	return t.winScore()
}

func (t *tennis) isAdv() bool {
	return math.Abs(float64(t.firstPlayerScoreTimes-t.secondPlayerScoreTimes)) == 1
}

func (t *tennis) advScore() string {
	return t.advPlayer() + " adv"
}

func (t *tennis) winScore() string {
	return t.advPlayer() + " win"
}

func (t *tennis) advPlayer() string {
	if t.firstPlayerScoreTimes > t.secondPlayerScoreTimes {
		return t.firstPlayerName
	}
	return t.secondPlayerName
}

func (t *tennis) isDeuce() bool {
	return t.firstPlayerScoreTimes >= 3
}

func (t *tennis) deuceScore() string {
	return "deuce"
}

func (t *tennis) isScoreDifferent() bool {
	return t.secondPlayerScoreTimes != t.firstPlayerScoreTimes
}

func (t *tennis) lookupScore() string {
	return scoreLookup[t.firstPlayerScoreTimes] + " " + scoreLookup[t.secondPlayerScoreTimes]
}

func (t *tennis) sameScore() string {
	return scoreLookup[t.firstPlayerScoreTimes] + " all"
}

func (t *tennis) firstPlayerScore() {
	t.firstPlayerScoreTimes += 1
}

func (t *tennis) secondPlayerScore() {
	t.secondPlayerScoreTimes += 1
}

func newTennis(firstPlayerName string, secondPlayerName string) *tennis {
	return &tennis{firstPlayerName: firstPlayerName, secondPlayerName: secondPlayerName}
}
