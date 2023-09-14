load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

load("//cli:rules.bzl", "cli_for")
load("//:configo_version.bzl", "configo_version")

cli_for(
    arch = "386",
    os = "linux",
)

cli_for(
    arch = "amd64",
    os = "linux",
)

cli_for(
    arch = "arm",
    os = "linux",
)

cli_for(
    arch = "arm64",
    os = "linux",
)
#cli_for(os="darwin", arch="amd64")

sh_binary(
    name = "print_configoVersion.go",
    srcs = ["print_configoVersion.go.sh"],
)

genrule(
    name = "versionFuncFile",
    outs = ["configoVersion.go"],
    tools = [":print_configoVersion.go"],
    cmd = "CONFIGO_VERSION=" + configo_version +
          " $(execpath :print_configoVersion.go) > $@",
)

go_library(
    name = "versionFunc",
    srcs = [":configoVersion.go"],
    importpath = "github.com/brianbraunstein/configo/cli",
)

go_library(
    name = "cli_lib",
    srcs = ["main.go"],
    importpath = "github.com/brianbraunstein/configo/cli",
    deps = ["//lib"],
    embed = [":versionFunc"],
)

go_binary(
    name = "cli",
    embed = [
        ":cli_lib",
    ],
)
