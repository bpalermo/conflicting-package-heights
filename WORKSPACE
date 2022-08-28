# gazelle:repo bazel_gazelle
workspace(name = "bad_workspace")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

PROTO_GRPC_RULE_SHA = "bbe4db93499f5c9414926e46f9e35016999a4e9f6e3522482d3760dc61011070"

PROTO_GRPC_RULE_VERSION = "4.2.0"

JAVA_RULES_VERSION = "5.1.0"

JAVA_RULES_SHA = "d974a2d6e1a534856d1b60ad6a15e57f3970d8596fbb0bb17b9ee26ca209332a"

RULES_JVM_EXTERNAL_TAG = "4.2"

RULES_JVM_EXTERNAL_SHA = "cd1a77b7b02e8e008439ca76fd34f5b07aecb8c752961f9640dea15e9e5ba1ca"

http_archive(
    name = "rules_proto_grpc",
    sha256 = PROTO_GRPC_RULE_SHA,
    strip_prefix = "rules_proto_grpc-%s" % PROTO_GRPC_RULE_VERSION,
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/archive/%s.tar.gz" % PROTO_GRPC_RULE_VERSION],
)

http_archive(
    name = "rules_java",
    sha256 = JAVA_RULES_SHA,
    url = "https://github.com/bazelbuild/rules_java/releases/download/" + JAVA_RULES_VERSION + "/rules_java-" + JAVA_RULES_VERSION + ".tar.gz",
)

http_archive(
    name = "rules_jvm_external",
    sha256 = RULES_JVM_EXTERNAL_SHA,
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/%s.zip" % RULES_JVM_EXTERNAL_TAG,
)

# Go
load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

load("@rules_proto_grpc//:repositories.bzl", "bazel_gazelle", "io_bazel_rules_go")  # buildifier: disable=same-origin-load

io_bazel_rules_go()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    version = "1.19",
)

bazel_gazelle()

load("@rules_proto_grpc//grpc-gateway:repositories.bzl", rules_proto_grpc_gateway_repos = "gateway_repos")

rules_proto_grpc_gateway_repos()

load("@grpc_ecosystem_grpc_gateway//:repositories.bzl", "go_repositories")

go_repositories()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

gazelle_dependencies()
