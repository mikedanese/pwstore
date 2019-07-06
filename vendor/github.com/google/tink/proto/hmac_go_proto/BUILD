load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["hmac.pb.go"],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/google/tink/proto/hmac_go_proto",
    importpath = "github.com/google/tink/proto/hmac_go_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/tink/proto/common_go_proto:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)
