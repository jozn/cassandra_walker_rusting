package ir.ms.pb;

public class PBFlatTypes {
{{range .Messages}}
	public class {{.Name}} {
	  {{- range .Fields}}
	   public {{.JavaType}} {{.FieldName}};
	  {{- end }}
	}
	/*
	folding
	PBFlatTypes.{{.Name}} t = new PBFlatTypes.{{.Name}}();
	{{- range .Fields}}
    t.set{{.FieldName}}();
    {{- end }}
	*/

	/*
	PBFlatTypes.{{.Name}} t = new PBFlatTypes.{{.Name}}();
	{{- range .Fields}}
	t.{{.FieldName}} = ;
	{{- end }}
	*/

	/*
	{{.Name}} t = new {{.Name}}();
	{{- range .Fields}}
	t.{{.FieldName}} = m.get{{.FieldName}}() ;
	{{- end }}
	*/
{{end}}
	
}

/*
{{range .Services}}
RPC_HANDLERS.{{.Name}} {{.Name}}_Handeler = null;
{{- end}}
	
*/