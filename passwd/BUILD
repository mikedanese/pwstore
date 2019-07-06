load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["passwd.go"],
    importpath = "github.com/mikedanese/pwstore/passwd",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/tink/go/subtle/aead:go_default_library",
        "//vendor/github.com/google/tink/go/subtle/random:go_default_library",
        "//vendor/github.com/google/tink/go/tink:go_default_library",
        "//vendor/golang.org/x/crypto/argon2:go_default_library",
        "//vendor/golang.org/x/crypto/chacha20poly1305:go_default_library",
        "//vendor/golang.org/x/sys/unix:go_default_library",
    ],
)
