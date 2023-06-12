# Not using https://github.com/bazelbuild/rules_go yet because it's in beta.

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
    # TODO: consider looking for a directory specific to the bazel workspace
    # that is writable but also sticks around between builds. GENDIR looks like
    # it's destroyed between builds.  This does somewhat go against the
    # "hermetic and deterministic" builds that bazel aims for though.
    cmd = ("GOCACHE=/tmp/gobazelhack/cache" +
           " GOPATH=/tmp/gobazelhack/path" +
           " GOOS=" + os +
           " GOARCH=" + arch +
           " go build -o $@ github.com/brianbraunstein/configo/cli"),
  )

