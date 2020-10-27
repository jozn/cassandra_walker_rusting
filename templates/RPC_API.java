package ir.ms.pb;

public class RPC_API {
{{range .Services}}

public static class {{ .Name }} {
	{{ $SName := .Name }}
	{{- range .Methods }}
        public static final String {{.MethodName}} = "{{$SName}}.{{.MethodName}}";
    {{- end -}}
}
{{- end }}
	
}