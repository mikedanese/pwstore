load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["cui.go"],
    importpath = "github.com/mikedanese/pwstore/cui",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/jroimartin/gocui:go_default_library",
        "//vendor/github.com/nsf/termbox-go:go_default_library",
        "@com_github_mikedanese_pwstore//bazel-pwstore/external/com_github_golang_protobuf/proto:go_default_library",
    ],
)
