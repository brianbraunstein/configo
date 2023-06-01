
def standard_test(name, data=[]):
  native.sh_test(
    name = name,
    srcs = [name + ".sh"],
    data = ["//cli:configo"] + data,
  )

