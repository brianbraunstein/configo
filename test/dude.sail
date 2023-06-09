
{{- define "default_params" -}}
name: Dude McDuderson
lang: English
numbers:
  one: 1
  two: 2
  three: 3
{{- end -}}

{{- define "sail" -}}
Hi, my name is {{ .params.name }}.
I speak {{ .params.lang }}.
I can count: {{ .params.numbers.one }}, {{ .params.numbers.two }}, {{ .params.numbers.three }}, ...
{{- end -}}

