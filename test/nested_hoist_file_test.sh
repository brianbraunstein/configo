#!/bin/bash
set -euo pipefail

function template() {
	cat <<EOF
{{- define "overrides" -}}
name: Nonsensey McBabblerson
{{- end -}}

{{- define "other_overrides" -}}
name: Gibby McGibberson
{{- end -}}

{{ hoist_file "./test/dude_gibberish.sail" "overrides" }}
{{ hoist_file "./test/dude_gibberish.sail" "other_overrides" }}
EOF
}

diff -u <(./cli/cli_/cli --template=<(template)) <(cat <<EOF
Hi, my name is Nonsensey McBabblerson.
I speak Gibberish.
I can count: blaw, floof, bloorp, ...
Hi, my name is Gibby McGibberson.
I speak Gibberish.
I can count: blaw, floof, bloorp, ...
EOF
)

