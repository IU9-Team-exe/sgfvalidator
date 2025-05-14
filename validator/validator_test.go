package validator

import (
	"fmt"
	"testing"
)

func TestValidateSGF(t *testing.T) {
	tests := []struct {
		name    string
		sgf     string
		wantErr bool
	}{
		{
			name:    "Valid SGF",
			sgf:     "(;GM[1]FF[4]SZ[19]AP[MyGoApp:1.0];B[dd];W[dc];B[cd];W[cc])",
			wantErr: false,
		},
		{
			name:    "Invalid move format",
			sgf:     "(;GM[1]FF[4]SZ[19]AP[MyGoApp:1.0];B[zz])",
			wantErr: true,
		},
		{
			name:    "Occupied point",
			sgf:     "(;GM[1]FF[4]SZ[19]AP[MyGoApp:1.0];B[dd];W[dd])",
			wantErr: true,
		},
		{
			name:    "Suicide move",
			sgf:     "(;GM[1]FF[4]SZ[19]AP[MyGoApp:1.0];B[aa];W[ab];B[ba];W[bb];B[ac])",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("Testing: %s, SGF: %s\n", tt.name, tt.sgf)
			err := ValidateSGF(tt.sgf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSGF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
