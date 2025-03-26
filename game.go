package main

func game(type_game string) {
	for !game_over {
		input("[X]")
		board_print()
		check_win("player1", "[X]")

		if type_game == "player" {
			input("[0]")
			board_print()
			check_win("player_2", "[0]")
		} else if type_game == "bot" {
			bot("[0]")
			board_print()
			check_win("bot", "[0]")
		}
	}
}
