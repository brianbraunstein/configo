package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "lib",
    srcs = [
        "file.go",
        "global.go",
        "util.go",
    ],
    importpath = "github.com/brianbraunstein/configo/lib",
    deps = [
        "@com_github_imdario_mergo//:mergo",
        "@com_github_masterminds_sprig_v3//:sprig",
        "@in_gopkg_yaml_v3//:yaml_v3",
    ],
)

go_test(
    name = "lib_test",
    srcs = ["util_test.go"],
    embed = [":lib"],
)

# Need for cross architecture building.
# TODO: determine if there's a way to make use of the .a file produced by the
# bazel go_library rule in a way that's platform independent (-ldflags seems
# like it probably isn't).
#filegroup(
#  name = "files",
#  srcs = glob(
#    include = ["*.go"],
#    exclude = ["*_test.go"],
#  ),
#)

