#!/bin/bash
set -euo pipefail

function template() {
	cat <<EOF
{{- define "innertmpl" -}}
Inner Template Called: {{ . }}
{{- end -}}

before
{{ template "innertmpl" "one" }}
{{ include "innertmpl" "two" | upper }}
{{ run "self" "innertmpl" "three" | lower }}
after
EOF
}

function expected() {
  cat <<EOF
before
Inner Template Called: one
INNER TEMPLATE CALLED: TWO
inner template called: three
after
EOF
}

diff -u <(./cli/cli_/cli --template=<(template)) <(expected)
