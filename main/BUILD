
package(default_visibility= ["//visibility:public"])

genrule(
  name = "main",
  outs = ["configo"],
  srcs = glob(["*"]) + [
    "//lib",
    "//:go_workspace",
  ],
  cmd = "GOCACHE=$$PWD/$(BINDIR)/.gocache GOPATH=$$PWD/$(BINDIR)/.gopath go build -o \"$@\" bristyle.com/configo/main",
)

