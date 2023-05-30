#!/bin/bash
set -euo pipefail

function template() {
	cat <<EOF
This is the template file
which uses: {{- .stuff | nindent 4 }}
from the data file, as well as some structure {{ .nested.stuff }}
EOF
}

function data() {
	cat <<EOF
one: two
stuff: goodies
nested:
  stuff: yum
three: four
EOF
}

diff -u <(./cli/configo --template=<(template) --data=<(data)) <(cat <<EOF
This is the template file
which uses:
    goodies
from the data file, as well as some structure yum
EOF
)

