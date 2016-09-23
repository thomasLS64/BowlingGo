package bowling

import (
	"fmt"
	"testing"
)

func frameChecker(input []Frame, position int, expectedScore int, expectedError error) error {
	score, err := GetFrameScore(input, position)

	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}

	return nil
}

func scoreChecker(input []Frame, expectedScore int, expectedError error) error {
	score, err := GetScore(input)

	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}

	return nil
}

func TestNullScore(t *testing.T) {
	input := []Frame{}
	expected := 0
	errorExpected := fmt.Errorf("Taille tableau différent de 10")

	if err := scoreChecker(input, expected, errorExpected); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreNegative(t *testing.T) {
	input := []Frame{{7, 2}, {8, -2}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}}
	expected := 0
	errorExpected := fmt.Errorf("Valeur négative interdite")

	if err := scoreChecker(input, expected, errorExpected); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreLessTen(t *testing.T) {
	input := []Frame{{7, 8}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}, {4, 3}}
	expected := 0
	errorExpected := fmt.Errorf("Tuple supérieur à 10")

	if err := scoreChecker(input, expected, errorExpected); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreSpare(t *testing.T) {
	input := []Frame{{7, 3}, {6, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 7+3+6+6+1

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreSpareLimit(t *testing.T) {
	input := []Frame{{7, 2}, {6, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {7, 3}}
	expected := 7+2+6+1+7+3

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreStrike(t *testing.T) {
	input := []Frame{{10, 0}, {6, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 10+6+1+6+1

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestScoreStrikeLimit(t *testing.T) {
	input := []Frame{{5, 1}, {6, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {10, 0}}
	expected := 5+1+6+1+10

	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestFrameScore(t *testing.T) {
	input := []Frame{{5, 1}, {6, 1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 7

	if err := frameChecker(input, 1, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}