# Go SGF Move Validator

This project provides a Go module to validate moves in Go game records stored in SGF (Smart Game Format) files. It checks moves for compliance with Go rules, including move format, board boundaries, occupied points, suicide moves, and the ko rule.

## Prerequisites

- Go 1.24 or later
- Git (optional, for cloning the repository)

## Installation

1. Clone the repository (or use your local copy):
   ```bash
   git clone https://github.com/your-repo/validate-go-moves.git
   cd validate-go-moves
   ```

2. Initialize the Go module (if not already done):
   ```bash
   go mod init validate-go-moves
   ```

3. Install the dependency:
   ```bash
   go get github.com/rooklift/sgf
   ```

## Usage

Run the main program to validate a hardcoded SGF string:
```bash
go run main.go
```

To validate a custom SGF file, modify `main.go` to read from a file or pass the SGF content via command-line arguments.

Run tests to verify the validator:
```bash
go test -v
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Documentation

- Package documentation: Run `godoc -http=:6060` and visit `http://localhost:6060/pkg/validate-go-moves/`.
- `rooklift/sgf` documentation: [pkg.go.dev/github.com/rooklift/sgf](https://pkg.go.dev/github.com/rooklift/sgf).
- SGF format: [red-bean.com/sgf](http://www.red-bean.com/sgf/).