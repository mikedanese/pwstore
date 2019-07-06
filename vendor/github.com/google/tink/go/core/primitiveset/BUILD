load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["primitiveset.go"],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/google/tink/go/core/primitiveset",
    importpath = "github.com/google/tink/go/core/primitiveset",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/tink/go/core/cryptofmt:go_default_library",
        "//vendor/github.com/google/tink/proto/tink_go_proto:go_default_library",
    ],
)
