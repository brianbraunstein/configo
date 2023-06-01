#!/bin/bash
set -euo pipefail

function template() {
  cat <<"EOF"
{{- define "IHaveYaml" -}}

one: two
three:
  four: five
  six: 7
eight: |
  has a lot more going
  on than just 9.

{{- end -}}

{{- $goData := include "IHaveYaml" . | fromYaml -}}
{{- $goData = set $goData "one" "22" -}}

The modified yaml output is: {{- toYaml $goData | nindent 2 }}
EOF
}

diff -u <(./cli/configo --template=<(template)) <(cat <<EOF
The modified yaml output is:
  eight: |-
      has a lot more going
      on than just 9.
  one: "22"
  three:
      four: five
      six: 7
EOF
)

