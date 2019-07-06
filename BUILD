load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/mikedanese/pwstore
# gazelle:build_file_name BUILD
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/mikedanese/pwstore",
    visibility = ["//visibility:private"],
    deps = [
        "//pwdb:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
    ],
)

go_binary(
    name = "pwstore",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
