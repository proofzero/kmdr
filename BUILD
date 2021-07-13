# ktrl/BAZEL.build

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

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


# TODO: add version numbers for genrules

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

genrule(
    name = "linux_binary",
    srcs = [":kmdr_linux_amd64"],
    outs = ["exec/kmdr_linux_amd64"],
    cmd = "cp $(SRCS) $@",
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
    out="kmdr",
)

genrule(
    name = "windows_binary",
    srcs = [":kmdr_windows_amd64"],
    outs = ["exec/kmdr_windows_amd64"],
    cmd = "cp $(SRCS) $@",
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
    out="kmdr",
)

genrule(
    name = "osx_binary",
    srcs = [":kmdr_osx_amd64"],
    outs = ["exec/kmdr_darwin_amd64"],
    cmd = "cp $(SRCS) $@",
)
