# Not using https://github.com/bazelbuild/rules_go yet because it's in beta.

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

def cli_for(os, arch):
  suffix = "-" + os + "_" + arch
  go_binary(
    name = "cli" + suffix,
    embed = [":cli_lib"],
    goos = os,
    goarch = arch,
  )

