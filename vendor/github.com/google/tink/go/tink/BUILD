load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "aead.go",
        "deterministic_aead.go",
        "hybrid_decrypt.go",
        "hybrid_encrypt.go",
        "mac.go",
        "signer.go",
        "verifier.go",
    ],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/google/tink/go/tink",
    importpath = "github.com/google/tink/go/tink",
    visibility = ["//visibility:public"],
)
