load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_library(
    name = "go_default_library",
    srcs = [
        "atomic.go",
        "db.go",
    ],
    embed = [":pwdb_go_proto"],
    importpath = "github.com/mikedanese/pwstore/pwdb",
    visibility = ["//visibility:public"],
    deps = [
        "//passwd:go_default_library",
        "//vendor/github.com/google/tink/go/aead:go_default_library",
        "//vendor/github.com/google/tink/go/keyset:go_default_library",
        "//vendor/github.com/google/tink/go/subtle/random:go_default_library",
        "//vendor/github.com/google/tink/go/tink:go_default_library",
        "//vendor/golang.org/x/sys/unix:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
    ],
)

proto_library(
    name = "pwdb_proto",
    srcs = ["pwdb.proto"],
    visibility = ["//visibility:public"],
    deps = ["@com_google_protobuf//:timestamp_proto"],
)

go_proto_library(
    name = "pwdb_go_proto",
    importpath = "github.com/mikedanese/pwstore/pwdb",
    proto = ":pwdb_proto",
    visibility = ["//visibility:public"],
)
