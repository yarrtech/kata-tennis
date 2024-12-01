import {Tennis} from "./tennis";

describe('score', () => {
    let tennis = new Tennis('joey', 'tom');
    beforeEach(() => {
        tennis = new Tennis('joey', 'tom');
    });

    function scoreShouldBe(expected) {
        expect(tennis.score()).toBe(expected);
    }

    it('should be love all', () => {
        scoreShouldBe('love all');
    });
    it('should be fifteen love', () => {
        givenFirstPlayerScoreTimes(1);
        scoreShouldBe('fifteen love');
    });

    function givenFirstPlayerScoreTimes(times) {
        for (let i = 0; i < times; i++) {
            tennis.firstPlayerScore();
        }
    }

    it('should be thirty love', () => {
        givenFirstPlayerScoreTimes(2);
        scoreShouldBe('thirty love');
    });

    it('should be forty love', () => {
        givenFirstPlayerScoreTimes(3);
        scoreShouldBe('forty love');
    });

    it('should be love fifteen', () => {
        givenSecondPlayerScoreTimes(1);
        scoreShouldBe('love fifteen');
    });

    function givenSecondPlayerScoreTimes(times) {
        for (let i = 0; i < times; i++) {
            tennis.secondPlayerScore();
        }
    }

    it('should be love thirty', () => {
        givenSecondPlayerScoreTimes(2);
        scoreShouldBe('love thirty');
    });

    it('should be fifteen all', () => {
        givenSecondPlayerScoreTimes(1);
        givenFirstPlayerScoreTimes(1);
        scoreShouldBe('fifteen all');
    });

    it('should be thirty all', () => {
        givenSecondPlayerScoreTimes(2);
        givenFirstPlayerScoreTimes(2);
        scoreShouldBe('thirty all');
    });

    function givenDeuce() {
        givenSecondPlayerScoreTimes(3);
        givenFirstPlayerScoreTimes(3);
    }

    it('should be deuce', () => {
        givenDeuce();
        scoreShouldBe('deuce');
    });

    it('should be first player adv', () => {
        givenDeuce();
        givenFirstPlayerScoreTimes(1);
        scoreShouldBe('joey adv');
    });

    it('should be second player adv', () => {
        givenDeuce();
        givenSecondPlayerScoreTimes(1);
        scoreShouldBe('tom adv');
    });

    it('should be second player win', () => {
        givenDeuce();
        givenSecondPlayerScoreTimes(2);
        scoreShouldBe('tom win');
    });


});