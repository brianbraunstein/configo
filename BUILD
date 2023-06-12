package(default_visibility = ["//visibility:public"])

filegroup(
    name = "go_build_files",
    srcs = [
        "go.mod",
        "go.sum",
    ],
)

load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/brianbraunstein/configo
gazelle(name = "gazelle")

