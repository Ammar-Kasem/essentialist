# Flashdown

A terminal application for [spaced repetition][1] using flashcards written in
[Markdown][2]. Review scheduling uses [FSRS][3]. This fork of
[Essentialist](https://github.com/essentialist-app/essentialist) focuses on Flashdown.

- Flashcards are plain text Markdown files.
- Review history is stored locally alongside each deck.
- Runs in a terminal on Linux, macOS, and Windows.

[1]: https://en.wikipedia.org/wiki/Spaced_repetition
[2]: https://en.wikipedia.org/wiki/Markdown
[3]: https://github.com/open-spaced-repetition/fsrs4anki

See [CONTRIBUTING.md](.github/CONTRIBUTING.md) for contribution guidelines.

## Build and install

Requires Go 1.26 or newer. From this checkout, build with:

```shell
make
```

Install into Go's binary directory (normally `~/go/bin`):

```shell
go install ./cmd/flashdown
```

Alternatively, install the built binary with the Makefile:

```shell
make install PREFIX="$HOME/.local"
```

## Usage

```shell
flashdown [-a] [-n <number of cards>] <file or directory> [<file> ...]
```

Run `flashdown --help` for available options. Press Space or Enter to reveal an
answer, then rate it with 1 (Again), 2 (Hard), 3 (Good), or 4 (Easy). Press `s` to
skip, `w` to save, or `q` to save and quit.

## Flash cards syntax

Each deck of cards is a plain text Markdown file with the extension `.md` (ex:
`sample.md`). You can put all your decks in the same directory.

Each card starts with a heading level 2 (line starting with `##`) defining the
question. The answer is the content following (until the next heading level 2).

Your progress is stored as JSON beside the deck: `sample.md` uses `sample.json`.
To keep progress from an older version, rename `.sample.md.db` to `sample.json`
before opening that deck.

Example of a deck with 3 cards:

```markdown
## Question: what format is used?

Questions and answers are in **Markdown**.

## Are lists supported?

Yes, here is an example:

- one
- **two**
- three

## How to include a table in the answer?

Answer with a table.

|  A  |  B  |
| --- | --- |
| 124 | 456 |
```

To force a line break in a Flashdown question, insert `<br>` (or `<br/>` or
`<br />`) while keeping the heading on one source line:

```markdown
## [Val av preposition · MIC-W09]<br>Fyll i prepositionen: ”Jag litar … dig.”

på
```

Flashdown displays the label and question on separate lines. These tags also
work in answers; tags inside code remain literal text.

Question hashes ignore these break tags, including uppercase variants, so
adding, removing, or moving them does not change the card's hash.

![Screenshot](docs/flashdown-screenshot.png)

## Development

Build and run all application and patched wrapping-library tests:

```shell
make test
go build ./...
```

Dependencies are defined in `go.mod`. Update individual dependencies as needed,
then run `go mod tidy` and `make test`. The local `go-term-text` replacement keeps
accented words together during wrapping; see its
[patch notes](third_party/github.com/MichaelMure/go-term-text/PATCHES.md).

The Go module path remains `github.com/essentialist-app/essentialist` for existing
internal imports.

## Releases

The **Manual build** workflow builds Flashdown binaries for Linux, macOS, and
Windows. To publish a release, run **Create Release** with a version such as
`1.2.3`. It creates a `v1.2.3` tag and publishes the matching binaries. Pushing a
`v*` tag also runs the release workflow.
