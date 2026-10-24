# Agent Instructions

## Project Purpose

Flashdown is a terminal flashcard application using Markdown decks and FSRS
review scheduling. This fork of Essentialist contains the terminal application.

## File Structure

- `cmd/flashdown/`: CLI entry point, terminal UI, and Markdown rendering.
- `internal/`: Card parsing, deck storage, and review scheduling.
- `internal/samples/`: Sample decks and parser test data.
- `third_party/github.com/MichaelMure/go-term-text/`: Patched terminal wrapping
  dependency; `PATCHES.md` documents the changes.
- `docs/`: Flashdown screenshot.
- `.github/`: Contribution guidelines and build/release workflows.

The Go module path remains `github.com/essentialist-app/essentialist`.

## Development

Follow standard Go conventions and format Go changes with `gofmt`.

- Build with `make` or `go build ./...`.
- Test with `make test`, which includes the patched wrapping dependency's tests.
- Install with `go install ./cmd/flashdown` or `make install`.

Keep Markdown headings on one source line. Flashdown renders supported `<br>`
tags as display line breaks, and card hashing ignores those tags.
