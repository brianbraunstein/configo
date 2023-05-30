
package(default_visibility= ["//visibility:public"])

genrule(
  name = "cli",
  outs = ["configo"],
  srcs = glob(["*"]) + [
    "//lib",
    "//:go_build_files",
  ],
  cmd = "GOCACHE=$$PWD/$(BINDIR)/.gocache GOPATH=$$PWD/$(BINDIR)/.gopath go build -o \"$@\" github.com/brianbraunstein/configo/cli",
)

