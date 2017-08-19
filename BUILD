load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library", "go_prefix")

go_prefix("github.com/aronasorman/glidelocktobazel")

gazelle(
    name = "gazelle",
    external = "vendored",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/Masterminds/glide/cfg:go_default_library"],
)

go_binary(
    name = "glidelocktobazel",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
