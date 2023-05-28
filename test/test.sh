#!/bin/bash -eu

set -o pipefail

cd "$(dirname "$(readlink -f "$0")")"

CG=$(readlink -f ../genfiles/configo)

function template() {
	cat <<EOF

{{- define "innertmpl" -}}
and it can use inner templates too! {{ . }}
{{- end }}

{{- define "innerbad" }}
InnerBad: {{ include "ZZnopeZZ" . }}
{{- end }}

{{- define "outerbad" }}
Calling innerbad...{{ include "innerbad" . | upper }}
{{- end -}}

This is the template file
which uses {{ .stuff | nindent 4 }} from the data file.

{{ template "innertmpl" "yay!" }}
{{ include "innertmpl" "uppercase me" | upper }}
{{ run "self" "innertmpl" "yay!" }}
{{ cheese "and global functions too" }}
{{ "this isn't output at all" | blank }}

{{ "./shallow.cfgo" | import_as "shal" -}}
{{ run "shal" "needlefish" "- and zoom at the surface" }}
{{ run "shal" "nudibranch" "- chromodoris is found shallower though" }}

{{ include "outerbad" "BadArg" }}
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
