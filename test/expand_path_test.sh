#!/bin/bash
set -euo pipefail

diff -u \
  <(./cli/cli_/cli --template=./test/workspace_test_dir/another_dir/expand_path.cfgo) \
  <(cat <<EOF
relative: $PWD/test/workspace_test_dir/another_dir/foo/bar
workspace relative: $PWD/test/workspace_test_dir/foo/bar
absolute: /foo/bar
EOF
)

