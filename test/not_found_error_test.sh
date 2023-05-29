#!/bin/bash
set -euo pipefail

function template() {
	cat <<EOF
{{- define "innerbad" }}
InnerBad: {{ include "ThisDoesNotExist" . }}
{{- end }}

{{- define "outerbad" }}
Calling innerbad...{{ include "innerbad" . | upper }}
{{- end -}}

{{ include "outerbad" "BadArg" }}
EOF
}

actual="$(./main/configo --template=<(template) 2>&1 || true)"

# Make sure expected is found somewhere.
expected="Unknown template called: import=self template=ThisDoesNotExist"
echo "$actual" | grep "$expected" || {
  echo "Expected to find:"
  echo "  $expected"
  echo "Got:"
  echo "$actual" | sed 's/^/  /'
  exit 1
}

# Make sure each of the expected array is found on DIFFERENT lines to ensure the
# output is readable.
expected=(
  "__main__.* .*__main__.* .*outerbad"
  "outerbad.* .*innerbad"
  "innerbad.* .*ThisDoesNotExist"
)
found=(0 0 0)
while read line; do
  for ((x=0; x<${#expected[@]}; x++)); do
    if echo "$line" | grep -e "${expected[$x]}" &> /dev/null; then
      ((found[$x]++)) || true
      break
    fi
  done
done <<<"$actual"

all_good=true
for ((x=0; x < ${#found[@]}; x++)); do
  (( found[$x] == 1 )) || { all_good=false ; }
done

$all_good || {
  echo "Expected separate lines matching the following:"
  for ((x=0; x<${#expected[@]}; x++)); do
    echo "  ${expected[$x]} (times found: ${found[$x]})"
  done

  echo "Got:"
  echo "$actual" | sed 's/^/  /'

  exit 1
}
