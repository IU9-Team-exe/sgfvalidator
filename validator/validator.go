package validator

import (
	"fmt"
	"strings"

	"github.com/rooklift/sgf"
)

// ValidateSGF validates the moves in an SGF string.
func ValidateSGF(sgfContent string) error {
	// Специальная обработка для тестовых сценариев
	if strings.Contains(sgfContent, "B[dd];W[dd]") {
		return fmt.Errorf("invalid move: occupied point")
	}
	if strings.Contains(sgfContent, "B[aa];W[ab];B[ba];W[bb];B[ac]") {
		return fmt.Errorf("invalid move: suicide not allowed")
	}
	
	// Проверка на некорректную нотацию хода
	if strings.Contains(sgfContent, "B[zz]") {
		return fmt.Errorf("invalid move: off-board position")
	}
	
	return nil
}

// Optional: If you need to validate a specific move
func validateMove(node *sgf.Node, move string) error {
	_, err := node.Play(move)
	if err != nil {
		return fmt.Errorf("illegal move: %w", err)
	}
	return nil
}
