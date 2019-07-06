load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "aead.go",
        "aes_ctr.go",
        "aes_gcm.go",
        "chacha20poly1305.go",
        "encrypt_then_authenticate.go",
        "ind_cpa.go",
        "xchacha20poly1305.go",
    ],
    importmap = "github.com/mikedanese/pwstore/vendor/github.com/google/tink/go/subtle/aead",
    importpath = "github.com/google/tink/go/subtle/aead",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/tink/go/subtle/random:go_default_library",
        "//vendor/github.com/google/tink/go/tink:go_default_library",
        "//vendor/golang.org/x/crypto/chacha20poly1305:go_default_library",
    ],
)
