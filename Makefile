
.PHONY: all clean

all: test
	bazel build //...

clean:
	bazel clean

.PHONY: test
test: genfiles/configo
	bazel test --test_output=errors //test/...

