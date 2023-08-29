#!/usr/bin/env bash
set -euo pipefail

cat <<EOF
package main

func configoVersion() string {
	return "${CONFIGO_VERSION}"
}
EOF
