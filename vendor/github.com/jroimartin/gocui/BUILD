load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "attribute.go",
        "doc.go",
        "edit.go",
        "escape.go",
        "gui.go",
        "keybinding.go",
        "view.go",
    ],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/jroimartin/gocui",
    importpath = "github.com/jroimartin/gocui",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/nsf/termbox-go:go_default_library"],
)
