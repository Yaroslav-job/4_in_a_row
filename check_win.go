package main

import (
	"fmt"
	"os"
)

var game_over = false

func check_win(type_player string, marker_type string) {
	for i := 0; i < 6; i++ {
		for k := 0; k < 7; k++ {
			if i+3 < 6 && k+3 < 7 && boardArr[i][k] == marker_type && boardArr[i][k] == boardArr[i+1][k+1] && boardArr[i+1][k+1] == boardArr[i+2][k+2] && boardArr[i+2][k+2] == boardArr[i+3][k+3] {
				fmt.Printf("Победил - %s\n", type_player)
				game_over = true
				os.Exit(0)
			}

			if i-3 >= 0 && k+3 < 7 && boardArr[i][k] == marker_type && boardArr[i][k] == boardArr[i-1][k+1] && boardArr[i-1][k+1] == boardArr[i-2][k+2] && boardArr[i-2][k+2] == boardArr[i-3][k+3] {
				fmt.Printf("Победил - %s\n", type_player)
				game_over = true
				os.Exit(0)
			}

			if i+3 < 6 && boardArr[i][k] == marker_type && boardArr[i][k] == boardArr[i+1][k] && boardArr[i+1][k] == boardArr[i+2][k] && boardArr[i+2][k] == boardArr[i+3][k] {
				fmt.Printf("Победил - %s\n", type_player)
				game_over = true
				os.Exit(0)
			}

			if k+3 < 7 && boardArr[i][k] == marker_type && boardArr[i][k] == boardArr[i][k+1] && boardArr[i][k+1] == boardArr[i][k+2] && boardArr[i][k+2] == boardArr[i][k+3] {
				fmt.Printf("Победил - %s\n", type_player)
				game_over = true
				os.Exit(0)
			}
		}
	}
}
