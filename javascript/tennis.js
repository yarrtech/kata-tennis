export class Tennis {
    firstPlayerScoreTimes = 0;
    scoreLookup = {
        1: 'fifteen',
        0: 'love',
        2: 'thirty',
        3: 'forty',
    };
    secondPlayerScoreTimes = 0;
    firstPlayerName;
    secondPlayerName;

    constructor(firstPlayerName, secondPlayerName) {
        this.firstPlayerName = firstPlayerName;
        this.secondPlayerName = secondPlayerName;
    }

    score() {
        return this.isScoreDifferent()
            ? this.isReadyForGamePoint() ? this.advState() : this.lookupScore()
            : this.isDeuce() ? this.deuceScore() : this.sameScore();
    }

    isReadyForGamePoint() {
        return this.firstPlayerScoreTimes > 3 || this.secondPlayerScoreTimes > 3;
    }

    advState() {
        if (this.isAdv()) {
            return this.advScore();
        }
        return this.winScore();
    }

    isAdv() {
        return Math.abs(this.firstPlayerScoreTimes - this.secondPlayerScoreTimes) === 1;
    }

    advScore() {
        return `${this.advPlayer()} adv`;
    }

    winScore() {
        return `${this.advPlayer()} win`;
    }

    advPlayer() {
        return this.firstPlayerScoreTimes > this.secondPlayerScoreTimes
            ? this.firstPlayerName : this.secondPlayerName;
    }

    isDeuce() {
        return this.firstPlayerScoreTimes >= 3;
    }

    deuceScore() {
        return 'deuce';
    }

    isScoreDifferent() {
        return this.secondPlayerScoreTimes !== this.firstPlayerScoreTimes;
    }

    lookupScore() {
        return `${this.scoreLookup[this.firstPlayerScoreTimes]} ${this.scoreLookup[this.secondPlayerScoreTimes]}`;
    }

    sameScore() {
        return `${this.scoreLookup[this.firstPlayerScoreTimes]} all`;
    }

    firstPlayerScore() {
        this.firstPlayerScoreTimes++;
    }

    secondPlayerScore() {
        this.secondPlayerScoreTimes++;
    }
}