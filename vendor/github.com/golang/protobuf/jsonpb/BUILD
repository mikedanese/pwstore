load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["jsonpb.go"],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/golang/protobuf/jsonpb",
    importpath = "github.com/golang/protobuf/jsonpb",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@io_bazel_rules_go//proto/wkt:struct_go_proto",
    ],
)
