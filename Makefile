
.PHONY: all clean run

all: genfiles/configo

clean:
	rm -rf genfiles

genfiles/configo: Makefile genfiles $(shell find main) $(shell find lib)
	go build -o $@ github.com/brianbraunstein/configo/main

genfiles:
	mkdir -p genfiles

.PHONY: test
test: genfiles/configo
	bazel test //test/...

