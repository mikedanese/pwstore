load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "byteorder.go",
        "cpu.go",
        "cpu_arm.go",
        "cpu_gc_s390x.go",
        "cpu_gc_x86.go",
        "cpu_linux.go",
        "cpu_linux_arm64.go",
        "cpu_linux_ppc64x.go",
        "cpu_linux_s390x.go",
        "cpu_mips64x.go",
        "cpu_mipsx.go",
        "cpu_other_arm64.go",
        "cpu_s390x.s",
        "cpu_x86.go",
        "cpu_x86.s",
    ],
    importmap = "github.com/mikedanese/pwstore/vendor/golang.org/x/sys/cpu",
    importpath = "golang.org/x/sys/cpu",
    visibility = ["//visibility:public"],
)
