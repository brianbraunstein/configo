
.PHONY: all clean run

all: genfiles/configo

clean:
	rm -rf genfiles

genfiles/configo: genfiles $(shell find main) $(shell find lib)
	go build -o $@ bristyle.com/configo/main

genfiles:
	mkdir -p genfiles

.PHONY: test
test: genfiles/configo
	@echo --------------------------------------------------------------------------------
	@cd test; ./test.sh
	@echo --------------------------------------------------------------------------------

