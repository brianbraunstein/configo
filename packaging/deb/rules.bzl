
version = "0.0.2"

def deb_pkg_for(arch):
  suffix = "-linux_" + arch
  if arch == "":
    suffix = ""
    arch = "local"

  native.genrule(
    name = "deb" + suffix,
    outs = ["configo_" + version + "_" + arch + ".deb"],
    tools = [":make_deb_file"],
    srcs = [
      "//cli:cli" + suffix,
      ":archive_template",
    ],
    cmd = "cli_path=$(execpath //cli:cli" + suffix + ")" +
          " RULEDIR=$(RULEDIR)" +
          " OUT_FILE=$@" +
          " $(execpath :make_deb_file)",
  )

