GOX := $(shell which go)
PREFIX := /usr/local
MANPATH := ${HOME}/.local/share/man
INSTALL_DIR := bin
SRC := ./cmd/git-ignores
BIN := git-ignores
OUTDIR := dist
MANPAGE := git-ignores.1

git-ignores:
	mkdir -p $(OUTDIR)
	$(GOX) build \
		-v \
		-x \
		-o $(OUTDIR)/$(BIN) \
		$(SRC)

install:
	@mkdir -p $(PREFIX)/$(INSTALL_DIR)
	@mkdir -p $(MANPATH)/man1
	install $(OUTDIR)/$(BIN) $(PREFIX)/$(INSTALL_DIR)/$(BIN)
	install $(MANPAGE) $(MANPATH)/man1/$(MANPAGE)

clean:
	@rm -rf $(OUTDIR)

test:
	$(GOX) test ./...

.PHONY: install clean test
