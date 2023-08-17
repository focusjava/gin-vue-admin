package global

{{- if .HasGlobal }}

import "gin-vue-admin/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}