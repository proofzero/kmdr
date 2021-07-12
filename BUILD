load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# ktrl/BAZEL.build

load("@bazel_gazelle//:def.bzl", "gazelle")

# This declares a `gazelle` binary rule that can be run like so:
#   $ bazel run //:gazelle
#
# This will generate a BUILD.bazel file with:
# - go_library()
# - go_binary()
# - go_test()
#
# targets for each package in your project. You can re-run the same
# command in the future to update existing BUILD files with new source
# files, dependencies, and options.

# gazelle:prefix github.com/proofzero/kmdr
gazelle(name = "gazelle")

go_library(
    name = "kmdr_lib",
    srcs = ["main.go"],
    importpath = "github.com/proofzero/kmdr",
    visibility = ["//visibility:private"],
    deps = ["//cmd"],
)

go_binary(
    name = "kmdr",
    embed = [":kmdr_lib"],
    visibility = ["//visibility:public"],
)

# go_binary(
#     name = "kmdr",
#     srcs = ["main.go"],
# )
