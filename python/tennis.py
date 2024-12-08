class Tennis:
    def __init__(self, first_player_name, second_player_name):
        self.second_player_name = second_player_name
        self.first_player_name = first_player_name
        self.second_player_score_times = 0
        self.score_lookup = {
            1: "fifteen",
            0: "love",
            2: "thirty",
            3: "forty",
        }
        self.first_player_score_times = 0

    def score(self):
        if self.is_score_different():
            return self.adv_state() if self.is_ready_for_game_point() else self.lookup_score()
        return self.deuce_score() if self.is_deuce() else self.same_score()

    def is_ready_for_game_point(self):
        return self.first_player_score_times > 3 or self.second_player_score_times > 3

    def adv_state(self):
        return self.adv_score() if self.is_adv() else self.win_score()

    def is_adv(self):
        return abs(self.first_player_score_times - self.second_player_score_times) == 1

    def adv_score(self):
        return f"{self.adv_player()} adv"

    def win_score(self):
        return f"{self.adv_player()} win"

    def adv_player(self):
        return self.first_player_name if self.first_player_score_times > self.second_player_score_times else self.second_player_name

    def is_deuce(self):
        return self.first_player_score_times >= 3

    def deuce_score(self):
        return "deuce"

    def is_score_different(self):
        return self.second_player_score_times != self.first_player_score_times

    def lookup_score(self):
        return f"{self.score_lookup[self.first_player_score_times]} {self.score_lookup[self.second_player_score_times]}"

    def same_score(self):
        return f"{self.score_lookup[self.first_player_score_times]} all"

    def first_player_score(self):
        self.first_player_score_times += 1

    def second_player_score(self):
        self.second_player_score_times += 1
