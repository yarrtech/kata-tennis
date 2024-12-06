package _go

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TennisTestSuite struct {
	suite.Suite
	*tennis
}

func TestTennisTestSuite(t *testing.T) {
	suite.Run(t, new(TennisTestSuite))
}

func (s *TennisTestSuite) SetupTest() {
	s.tennis = newTennis("joey", "tom")
}

func (s *TennisTestSuite) Test_love_all() {
	s.scoreShouldBe("love all")
}
func (s *TennisTestSuite) Test_fifteen_love() {
	s.givenFirstPlayerScoreTimes(1)
	s.scoreShouldBe("fifteen love")
}
func (s *TennisTestSuite) Test_thirty_love() {
	s.givenFirstPlayerScoreTimes(2)
	s.scoreShouldBe("thirty love")
}
func (s *TennisTestSuite) Test_forty_love() {
	s.givenFirstPlayerScoreTimes(3)
	s.scoreShouldBe("forty love")
}
func (s *TennisTestSuite) Test_love_fifteen() {
	s.givenSecondPlayerScoreTimes(1)
	s.scoreShouldBe("love fifteen")
}
func (s *TennisTestSuite) Test_love_thirty() {
	s.givenSecondPlayerScoreTimes(2)
	s.scoreShouldBe("love thirty")
}
func (s *TennisTestSuite) Test_fifteen_all() {
	s.givenSecondPlayerScoreTimes(1)
	s.givenFirstPlayerScoreTimes(1)
	s.scoreShouldBe("fifteen all")
}
func (s *TennisTestSuite) Test_thirty_all() {
	s.givenSecondPlayerScoreTimes(2)
	s.givenFirstPlayerScoreTimes(2)
	s.scoreShouldBe("thirty all")
}
func (s *TennisTestSuite) Test_deuce() {
	s.givenDeuce()
	s.scoreShouldBe("deuce")
}
func (s *TennisTestSuite) Test_first_player_adv() {
	s.givenDeuce()
	s.givenFirstPlayerScoreTimes(1)
	s.scoreShouldBe("joey adv")
}
func (s *TennisTestSuite) Test_second_player_adv() {
	s.givenDeuce()
	s.givenSecondPlayerScoreTimes(1)
	s.scoreShouldBe("tom adv")
}
func (s *TennisTestSuite) Test_second_player_win() {
	s.givenDeuce()
	s.givenSecondPlayerScoreTimes(2)
	s.scoreShouldBe("tom win")
}

func (s *TennisTestSuite) givenDeuce() {
	s.givenSecondPlayerScoreTimes(3)
	s.givenFirstPlayerScoreTimes(3)
}

func (s *TennisTestSuite) givenSecondPlayerScoreTimes(times int) {
	for range times {
		s.tennis.secondPlayerScore()
	}
}

func (s *TennisTestSuite) givenFirstPlayerScoreTimes(times int) {
	for range times {
		s.tennis.firstPlayerScore()
	}
}

func (s *TennisTestSuite) scoreShouldBe(expected string) bool {
	return s.Equal(expected, s.tennis.score())
}
