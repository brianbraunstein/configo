
load("//:configo_version.bzl", "configo_version")

def deb_pkg_for(arch):
  suffix = "-linux_" + arch
  if arch == "":
    suffix = ""
    arch = "local"

  native.genrule(
    name = "deb" + suffix,
    outs = ["configo_" + configo_version + "_" + arch + ".deb"],
    tools = [":make_deb_file"],
    srcs = [
      "deb_control.envsubst",
      "//cli:cli" + suffix,
    ],
    cmd = "cli_path=$(execpath //cli:cli" + suffix + ")" +
          " RULEDIR=$(RULEDIR)" +
          " OUT_FILE=$@" +
          " VERSION=" + configo_version +
          " ARCH=" + arch +
          " $(execpath :make_deb_file)",
  )

