load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

load("//cli:rules.bzl", "cli_for")

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

go_library(
    name = "cli_lib",
    srcs = ["main.go"],
    importpath = "github.com/brianbraunstein/configo/cli",
    deps = ["//lib"],
)

go_binary(
    name = "cli",
    embed = [":cli_lib"],
)
