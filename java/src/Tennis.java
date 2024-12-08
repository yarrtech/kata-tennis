import java.util.HashMap;

import static java.lang.Math.abs;

public class Tennis {
    private final HashMap<Integer, String> scoreLookup = new HashMap<>() {{
        put(1, "fifteen");
        put(0, "love");
        put(2, "thirty");
        put(3, "forty");
    }};
    private final String firstPlayerName;
    private final String secondPlayerName;
    private int firstPlayerScoreTimes;
    private int secondPlayerScoreTimes;

    public Tennis(String firstPlayerName, String secondPlayerName) {
        this.firstPlayerName = firstPlayerName;
        this.secondPlayerName = secondPlayerName;
    }

    public String score() {
        return isScoreDifferent()
                ? isReadyForGamePoint() ? advState() : lookupScore()
                : isDeuce() ? deuceScore() : sameScore();
    }

    public void firstPlayerScore() {
        firstPlayerScoreTimes++;
    }

    public void secondPlayerScore() {
        secondPlayerScoreTimes++;
    }

    private String advPlayer() {
        return firstPlayerScoreTimes > secondPlayerScoreTimes ? firstPlayerName : secondPlayerName;
    }

    private String advScore() {
        return advPlayer() + " adv";
    }

    private String advState() {
        return isAdv() ? advScore() : winScore();
    }

    private String deuceScore() {
        return "deuce";
    }

    private boolean isAdv() {
        return abs(firstPlayerScoreTimes - secondPlayerScoreTimes) == 1;
    }

    private boolean isDeuce() {
        return firstPlayerScoreTimes >= 3;
    }

    private boolean isReadyForGamePoint() {
        return firstPlayerScoreTimes > 3 || secondPlayerScoreTimes > 3;
    }

    private boolean isScoreDifferent() {
        return secondPlayerScoreTimes != firstPlayerScoreTimes;
    }

    private String lookupScore() {
        return scoreLookup.get(firstPlayerScoreTimes) + " " + scoreLookup.get(secondPlayerScoreTimes);
    }

    private String sameScore() {
        return scoreLookup.get(firstPlayerScoreTimes) + " all";
    }

    private String winScore() {
        return advPlayer() + " win";
    }
}
