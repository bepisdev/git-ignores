GOX := $(shell which go)
PREFIX := /usr/local
SRC := ./cmd/git-ignores
BIN := git-ignores
OUTDIR := dist

git-ignores:
	mkdir -p $(OUTDIR)
	$(GOX) build \
		-v \
		-x \
		-o $(OUTDIR)/$(BIN) \
		$(SRC)

install:
	cp $(OUTDIR)/$(BIN) $(PREFIX)/bin/$(BIN)

clean:
	@rm -rf $(OUTDIR)

test:
	$(GOX) test ./...

.PHONY: install clean test
