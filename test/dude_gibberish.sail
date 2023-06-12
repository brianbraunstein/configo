
{{- define "default_params" }}
name: __required__
{{ end -}}

{{- define "Zoverrides" }}
name: {{ .params.name }}
lang: Gibberish
numbers:
  one: blaw
  two: floof
  three: bloorp
{{ end -}}

{{- define "sail" -}}
{{ hoist_file "./dude.sail" "Zoverrides" }}
{{- end -}}

