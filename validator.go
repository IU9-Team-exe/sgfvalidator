package validator

import (
	"fmt"

	"github.com/rooklift/sgf"
)

// ValidateSGF validates the moves in an SGF file.
// It returns an error if any move is invalid, otherwise nil.
// The validation includes checks for move format, board boundaries, occupied points, suicide moves, and ko rule.
func ValidateSGF(sgfContent string) error {
	tree, err := sgf.Parse(sgfContent)
	if err != nil {
		return fmt.Errorf("parsing SGF: %w", err)
	}

	node := tree.Root
	for _, child := range node.Children {
		if move, ok := child.Get("B"); ok {
			if err := validateMove(node, move, "B"); err != nil {
				return fmt.Errorf("invalid black move %s: %w", move, err)
			}
			node, _ = node.Play(sgf.PointFromString(move))
		}
		if move, ok := child.Get("W"); ok {
			if err := validateMove(node, move, "W"); err != nil {
				return fmt.Errorf("invalid white move %s: %w", move, err)
			}
			node, _ = node.Play(sgf.PointFromString(move))
		}
	}
	return nil
}

// validateMove checks if a move is legal.
func validateMove(node *sgf.Node, move string, color string) error {
	point, err := sgf.ParsePoint(move)
	if err != nil {
		return fmt.Errorf("invalid move format: %w", err)
	}
	_, err = node.Play(point)
	if err != nil {
		return fmt.Errorf("illegal move: %w", err)
	}
	return nil
}
