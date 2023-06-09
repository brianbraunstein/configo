
load("//test:rules.bzl", "standard_test")

standard_test(name="template_include_run_test")
standard_test(name="blank_test")
standard_test(name="data_test")
standard_test(name="not_found_error_test")
standard_test(name="yaml_test")

standard_test(
  name="nested_test",
  data = [
    "shallow.cfgo",
    "deep.cfgo",
  ]
)

standard_test(
  name="workspace_test",
  data = glob(["workspace_test_dir/**"]),
)

standard_test(
  name="hoist_file_test",
  data = ["dude.sail"],
)
