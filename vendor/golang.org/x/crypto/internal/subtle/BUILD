load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["aliasing.go"],
    importmap = "github.com/mikedanese/pwstore/vendor/golang.org/x/crypto/internal/subtle",
    importpath = "golang.org/x/crypto/internal/subtle",
    visibility = ["//vendor/golang.org/x/crypto:__subpackages__"],
)
