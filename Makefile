
.PHONY: all clean run

all: genfiles/configo

clean:
	rm -rf genfiles

genfiles/configo: Makefile genfiles $(shell find cli) $(shell find lib)
	go build -o $@ github.com/brianbraunstein/configo/cli

genfiles:
	mkdir -p genfiles

.PHONY: test
test: genfiles/configo
	bazel test --test_output=errors //test/...

