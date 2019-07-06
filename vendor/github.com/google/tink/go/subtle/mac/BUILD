load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["hmac.go"],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/google/tink/go/subtle/mac",
    importpath = "github.com/google/tink/go/subtle/mac",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/tink/go/subtle:go_default_library",
        "//vendor/github.com/google/tink/go/tink:go_default_library",
    ],
)
