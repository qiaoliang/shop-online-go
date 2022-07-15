load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/example/project
gazelle(name = "gazelle")

go_library(
    name = "project_lib",
    srcs = [
        "initData.go",
        "main.go",
    ],
    importpath = "github.com/example/project",
    deps =[
        "//app/configs",
        "//app/routers",
        "//app/goods",
    ],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "project",
    embed = [":project_lib"],
    visibility = ["//visibility:public"],
    data = [
        "//:datas_for_prod",
    ],
)

go_test(
    name = "project_test",
    srcs = ["initData_test.go"],
    embed = [":project_lib"],
    deps = ["@com_github_stretchr_testify//suite:go_default_library"],
)

filegroup(
    name = "datas_for_prod",
    srcs =
        ["config.yaml"],
    data =[
        ":db_files",
        ":pic_files",
    ],
)

filegroup(
    name = "cfgfile_for_test",
    srcs = [
        "config-test.yaml",
    ],
    visibility =[
        "//app/configs:__pkg__",
    ],
)

filegroup(
    name = "db_files",
    srcs = glob(["dbscripts/*.sql"]),
    visibility =[
        "//app/configs:__pkg__",
    ],
)

filegroup(
    name = "pic_files",
    srcs = glob(["static/**/*.*"]),
)