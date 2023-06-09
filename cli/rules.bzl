
def cli_for(os, arch):
  suffix = ""
  if os + arch != "":
    suffix = "-" + os + "_" + arch
  native.genrule(
    name = "cli" + suffix,
    outs = ["configo" + suffix],
    srcs = native.glob(["**/*.go"]) + [
      "//lib",
      "//:go_build_files"
    ],
    cmd = ("GOCACHE=$$(realpath -m $(GENDIR))/.gohack/cache" +
           " GOPATH=$$(realpath -m $(GENDIR))/.gohack/path" +
           " GOOS=" + os +
           " GOARCH=" + arch +
           " go build -o $@ github.com/brianbraunstein/configo/cli"),
  )

