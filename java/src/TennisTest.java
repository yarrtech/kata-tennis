import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class TennisTest {

    private final Tennis tennis = new Tennis("joey", "tom");

    @Test
    public void love_all() {
        scoreShouldBe("love all");
    }

    @Test
    public void fifteen_love() {
        givenFirstPlayerScoreTimes(1);
        scoreShouldBe("fifteen love");
    }

    @Test
    public void thirty_love() {
        givenFirstPlayerScoreTimes(2);
        scoreShouldBe("thirty love");
    }

    @Test
    public void forty_love() {
        givenFirstPlayerScoreTimes(3);
        scoreShouldBe("forty love");
    }

    @Test
    public void love_fifteen() {
        givenSecondPlayerScoreTimes(1);
        scoreShouldBe("love fifteen");
    }

    @Test
    public void love_thirty() {
        givenSecondPlayerScoreTimes(2);
        scoreShouldBe("love thirty");
    }

    @Test
    public void fifteen_all() {
        givenSecondPlayerScoreTimes(1);
        givenFirstPlayerScoreTimes(1);
        scoreShouldBe("fifteen all");
    }

    @Test
    public void thirty_all() {
        givenSecondPlayerScoreTimes(2);
        givenFirstPlayerScoreTimes(2);
        scoreShouldBe("thirty all");
    }

    @Test
    public void deuce() {
        givenDeuce();
        scoreShouldBe("deuce");
    }

    @Test
    public void first_player_adv() {
        givenDeuce();
        givenFirstPlayerScoreTimes(1);
        scoreShouldBe("joey adv");
    }

    @Test
    public void second_player_adv() {
        givenDeuce();
        givenSecondPlayerScoreTimes(1);
        scoreShouldBe("tom adv");
    }

    @Test
    public void second_player_win() {
        givenDeuce();
        givenSecondPlayerScoreTimes(2);
        scoreShouldBe("tom win");
    }

    private void givenDeuce() {
        givenSecondPlayerScoreTimes(3);
        givenFirstPlayerScoreTimes(3);
    }

    private void givenFirstPlayerScoreTimes(int times) {
        for (int i = 0; i < times; i++) {
            tennis.firstPlayerScore();
        }
    }

    private void givenSecondPlayerScoreTimes(int times) {
        for (int i = 0; i < times; i++) {
            tennis.secondPlayerScore();
        }
    }

    private void scoreShouldBe(String expected) {
        Assertions.assertEquals(expected, tennis.score());
    }
}
