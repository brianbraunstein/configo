
sh_test(
  name = "template_include_run",
  srcs = ["template_include_run_test.sh"],
  data = ["//cli:configo"],
)

sh_test(
  name = "blank_test",
  srcs = ["blank_test.sh"],
  data = ["//cli:configo"],
)

sh_test(
  name = "nested_test",
  srcs = ["nested_test.sh"],
  data = [
    "shallow.cfgo",
    "deep.cfgo",
    "//cli:configo",
  ]
)

sh_test(
  name = "data_test",
  srcs = ["data_test.sh"],
  data = ["//cli:configo"],
)

sh_test(
  name = "not_found_error_test",
  srcs = ["not_found_error_test.sh"],
  data = ["//cli:configo"],
)
