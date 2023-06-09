
load(":rules.bzl", "deb_pkg_for")

deb_pkg_for(arch="")
deb_pkg_for(arch="386")
deb_pkg_for(arch="amd64")
deb_pkg_for(arch="arm")
deb_pkg_for(arch="arm64")

sh_binary(
	name = "make_deb_file",
	srcs = ["make_deb_file.sh"],
)

