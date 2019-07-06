load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "binary_io.go",
        "handle.go",
        "json_io.go",
        "keyset.go",
        "manager.go",
        "mem_io.go",
        "reader.go",
        "validation.go",
        "writer.go",
    ],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/google/tink/go/keyset",
    importpath = "github.com/google/tink/go/keyset",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/tink/go/core/primitiveset:go_default_library",
        "//vendor/github.com/google/tink/go/core/registry:go_default_library",
        "//vendor/github.com/google/tink/go/internal:go_default_library",
        "//vendor/github.com/google/tink/go/subtle/random:go_default_library",
        "//vendor/github.com/google/tink/go/tink:go_default_library",
        "//vendor/github.com/google/tink/proto/tink_go_proto:go_default_library",
        "@com_github_golang_protobuf//jsonpb:go_default_library_gen",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)
