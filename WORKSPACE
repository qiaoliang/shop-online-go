load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
go_version = "1.18.3"
rules_go_version = "v0.33.0"
rules_go_sha="685052b498b6ddfe562ca7a97736741d87916fe536623afb7da2824c0211c369"

bzl_gazelle_version = "v0.25.0"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = rules_go_sha,
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/{}/rules_go-{}.zip".format(rules_go_version,rules_go_version),
        "https://github.com/bazelbuild/rules_go/releases/download/{}/rules_go-{}.zip".format(rules_go_version,rules_go_version),
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "5982e5463f171da99e3bdaeff8c0f48283a7a5f396ec5282910b9e8a49c0dd7e",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/{}/bazel-gazelle-{}.tar.gz".format(bzl_gazelle_version,bzl_gazelle_version),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/{}/bazel-gazelle-{}.tar.gz".format(bzl_gazelle_version,bzl_gazelle_version),
    ],
)


load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_rules_dependencies()

go_register_toolchains(version = go_version)

gazelle_dependencies()

