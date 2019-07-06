load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "key_manager.go",
        "kms_client.go",
        "private_key_manager.go",
        "registry.go",
    ],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/google/tink/go/core/registry",
    importpath = "github.com/google/tink/go/core/registry",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/tink/go/tink:go_default_library",
        "//vendor/github.com/google/tink/proto/tink_go_proto:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)
