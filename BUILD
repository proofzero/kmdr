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


# config_setting(
#  name = "dynlink",
#  flags = {
#      "dynlink": "true"
#  },
# )

go_library(
    name = "kmdr_lib",
    # srcs = select({
    #     ":dynlink": ["main.go"],
    # }),
    srcs = ["main.go"],
    importpath = "github.com/proofzero/kmdr",
    visibility = ["//visibility:private"],
    deps = ["//cmd"],
)

# Linux
# -----------------------------------------------------------------------
go_binary(
    name = "kmdr_linux_amd64",
    goos="linux",
    goarch="amd64",
    cgo=True,
    gc_goopts=[
        "-dynlink"
    ],
    linkmode="pie",
    embed = [":kmdr_lib"],
    visibility = ["//visibility:public"],
)


# Windows
# -----------------------------------------------------------------------
go_binary(
    name = "kmdr_windows_amd64",
    goos="windows",
    goarch="amd64",
    cgo=True,
    gc_goopts=[
        "-dynlink"
    ],
    linkmode="pie",
    embed = [":kmdr_lib"],
    visibility = ["//visibility:public"],
)

# OSX
# -----------------------------------------------------------------------
go_binary(
    name = "kmdr_osx_amd64",
    goos="darwin",
    goarch="amd64",
    cgo=True,
    gc_goopts=[
        "-dynlink"
    ],
    linkmode="pie",
    embed = [":kmdr_lib"],
    visibility = ["//visibility:public"],
)

# OSX
