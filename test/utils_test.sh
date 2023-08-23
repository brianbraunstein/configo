#!/bin/bash
set -euo pipefail

function template() {
	cat <<EOF
cidrSubnet: {{ cidrSubnet "10.0.0.0/16" 8 7 }}
EOF
}

diff -u <(./cli/cli_/cli --template=<(template)) <(cat <<EOF
cidrSubnet: 10.0.7.0/24
EOF
)

