load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["aes_ctr.pb.go"],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/google/tink/proto/aes_ctr_go_proto",
    importpath = "github.com/google/tink/proto/aes_ctr_go_proto",
    visibility = ["//visibility:public"],
    deps = ["@com_github_golang_protobuf//proto:go_default_library"],
)
