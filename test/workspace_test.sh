#!/bin/bash
set -euo pipefail

diff -u <(./cli/configo --template=./test/workspace_test_dir/subdir/file1.configo) <(cat <<EOF
File 1 contents, which includes 2:
  File 2 contents, which also contains 3:
    File 3 contents.
And here's file 4:
  Finally the file 4 contents
EOF
)

