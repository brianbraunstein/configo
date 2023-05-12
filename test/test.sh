#!/bin/bash -eu

set -o pipefail

cd "$(dirname "$(readlink -f "$0")")"

CG=$(readlink -f ../genfiles/configo)

function template() {
	cat <<EOF
This is the template file
which uses {{ .stuff | nindent 4 }} from the data file.
{{ define "innertmpl" -}}
and it can use inner templates too! {{ . }}
{{- end }}
{{ template "innertmpl" "yay!" }}
{{ template "innertmpl" (cheese "horray!") }}
EOF
}

function data() {
	cat <<EOF
one: two
stuff: goodies
three: four
EOF
}

$CG --template=<(template) --data=<(data)

