# Local changes

Source: github.com/MichaelMure/go-term-text v0.3.1, under the included MIT license.

The root module uses this copy to preserve word boundaries when wrapping text
containing accented letters or combining marks. In `wrap.go`, `runeType` groups
these characters with ASCII letters instead of allowing breaks between them.
Wide characters retain the upstream wrapping behavior. Upstream Go source and
tests are included, along with Unicode word wrapping regressions in
`unicode_wrap_test.go`.

Two variation-selector width expectations in `len_test.go` match the root
module's selected go-runewidth version, which treats variation selectors as
zero width. Other upstream files are unchanged.
