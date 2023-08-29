#!/bin/bash
set -euo pipefail

function template() {
	cat <<'EOF'
cidrSubnet: {{ cidrSubnet "10.0.0.0/16" 8 7 }}
cidrSubnet64: {{ cidrSubnet64 "10.0.0.0/16" 8 (int64 9) }}
EOF
}

diff -u <(./cli/cli_/cli --template=<(template)) <(cat <<EOF
cidrSubnet: 10.0.7.0/24
cidrSubnet64: 10.0.9.0/24
EOF
)

