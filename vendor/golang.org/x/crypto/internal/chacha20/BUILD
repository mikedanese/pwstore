load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "asm_arm64.s",
        "asm_ppc64le.s",
        "chacha_arm64.go",
        "chacha_generic.go",
        "chacha_noasm.go",
        "chacha_ppc64le.go",
        "chacha_s390x.go",
        "chacha_s390x.s",
        "xor.go",
    ],
    importmap = "github.com/mikedanese/pwstore/vendor/golang.org/x/crypto/internal/chacha20",
    importpath = "golang.org/x/crypto/internal/chacha20",
    visibility = ["//vendor/golang.org/x/crypto:__subpackages__"],
    deps = [
        "//vendor/golang.org/x/crypto/internal/subtle:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:s390x": [
            "//vendor/golang.org/x/sys/cpu:go_default_library",
        ],
        "//conditions:default": [],
    }),
)
