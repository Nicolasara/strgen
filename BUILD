load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "strgen",
    srcs = ["strgen.go"],
    importpath = "github.com/Nicolasara/mono/tree/main/nico/packages/go/strgen",
    visibility = ["//visibility:public"],
)

go_test(
    name = "strgen_test",
    srcs = ["strgen_test.go"],
    embed = [":strgen"],
)
