# Local changes

Source: github.com/MichaelMure/go-term-text v0.3.1, under the included MIT license.

The root module uses this copy to preserve word boundaries when wrapping text
containing accented letters or combining marks. In `wrap.go`, `runeType` groups
these characters with ASCII letters instead of allowing breaks between them.
Wide characters retain the upstream wrapping behavior. Upstream Go source and
tests are included, along with Unicode word wrapping regressions in
`unicode_wrap_test.go`.

Two variation-selector width expectations in `len_test.go` are updated for
go-runewidth v0.0.23, which the root module already uses. Other upstream files
are unchanged.
