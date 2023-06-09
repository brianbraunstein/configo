#!/bin/bash
set -euo pipefail

function template() {
	cat <<EOF
{{- define "overrides" -}}
lang: Spanish
numbers:
  one: uno
  two: dos
{{- end -}}

{{- define "other_overrides" -}}
lang: German
numbers:
  one: eins
  two: zwei
{{- end -}}

{{ hoist_file "./test/dude.sail" "overrides" }}
{{ hoist_file "./test/dude.sail" "other_overrides" }}
EOF
}

diff -u <(./cli/configo --template=<(template)) <(cat <<EOF
Hi, my name is Dude McDuderson.
I speak Spanish.
I can count: uno, dos, 3, ...
Hi, my name is Dude McDuderson.
I speak German.
I can count: eins, zwei, 3, ...
EOF
)

