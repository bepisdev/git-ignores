GOX := $(shell which go)
PREFIX := /usr/local
SRC := ./cmd/git-ignores
BIN := git-ignores
OUTDIR := dist
MANPAGE := ./git-ignores.1

git-ignores:
	mkdir -p $(OUTDIR)
	$(GOX) build \
		-v \
		-x \
		-o $(OUTDIR)/$(BIN) \
		$(SRC)
	cp $(MANPAGE) $(OUTDIR)

install:
	cp $(OUTDIR)/$(BIN) $(PREFIX)/bin/$(BIN)
	cp $(OUTDIR)/$(MANPAGE) $(PREFIX)/share/man/man1/$(MANPAGE)

clean:
	@rm -rf $(OUTDIR)

test:
	$(GOX) test ./...

.PHONY: install clean test
