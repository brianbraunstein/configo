#!/bin/bash
set -euo pipefail

function template() {
	cat <<EOF
{{ "./test/shallow.cfgo" | import_as "shal" -}}
before
{{ run "shal" "needlefish" "- and zoom at the surface" }}
{{ run "shal" "nudibranch" "- chromodoris is found shallower though" }}
after
EOF
}

diff -u <(./cli/configo --template=<(template)) <(cat <<EOF
before
Needlefish like it very shallow - and zoom at the surface
Nudibranchs prefer it deeper  - especially variable neons
Nudibranchs prefer it deeper - chromodoris is found shallower though
after
EOF
)

