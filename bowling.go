package bowling

import "fmt"

type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetFrameScore(game []Frame, position int) (int, error) {
	if game[position].firstThrow >= 0 && game[position].secondThrow >= 0 {
		totalTuple := game[position].firstThrow + game[position].secondThrow

		if totalTuple <= 10 {
			if totalTuple == 10 && position+1 < len(game) {
				if game[position].firstThrow == 10 {
					// Strike
					totalTuple += game[position+1].firstThrow + game[position+1].secondThrow
				} else {
					// Spare
					totalTuple += game[position+1].firstThrow
				}
			}
			
			return totalTuple, nil
		} else {
			return 0, fmt.Errorf("Tuple supérieur à 10")
		}
	} else {
		return 0, fmt.Errorf("Valeur négative interdite")
	}
}

func GetScore(game []Frame) (int, error) {
	if len(game) != 10 {
		return 0, fmt.Errorf("Taille tableau différent de 10")
	} else {
		score := 0

		for i := 0; i < len(game); i++ {
			scoreFrame, err := GetFrameScore(game, i)

			if (err != nil) {
				return 0, err
			} else {
				score += scoreFrame
			}
		}

		return score, nil
	}
}
