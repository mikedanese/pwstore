load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["struct.pb.go"],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/golang/protobuf/ptypes/struct",
    importpath = "github.com/golang/protobuf/ptypes/struct",
    visibility = ["//visibility:public"],
    deps = ["@com_github_golang_protobuf//proto:go_default_library"],
)
