# Build and install Flashdown.
GOCMD ?= go
FLASHDOWN_BIN ?= flashdown

PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin

all: $(FLASHDOWN_BIN)

$(FLASHDOWN_BIN):
	$(GOCMD) build -o $(FLASHDOWN_BIN) ./cmd/flashdown

test:
	$(GOCMD) test ./... github.com/MichaelMure/go-term-text

install: all
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 $(FLASHDOWN_BIN) $(DESTDIR)$(BINDIR)/flashdown

clean:
	rm -f $(FLASHDOWN_BIN)

.PHONY: all $(FLASHDOWN_BIN) test install clean
