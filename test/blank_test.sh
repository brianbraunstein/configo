#!/bin/bash
set -euo pipefail

function template() {
	cat <<EOF
before
x{{ "this isn't output at all" | blank }}x
after
EOF
}

diff -u <(./cli/cli_/cli --template=<(template)) <(cat <<EOF
before
xx
after
EOF
)

