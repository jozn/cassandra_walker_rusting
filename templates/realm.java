package com.mardomsara.social.models_realm.pb_realm;

import io.realm.RealmObject;
import io.realm.annotations.PrimaryKey;
import ir.ms.pb.*;

{{$pk := (index .Fields 0).FieldName }}
public class {{.RealmClass}} extends RealmObject {
   {{- range $i ,$el := .Fields }}
		{{if (eq $pk  $el.FieldName ) }} @PrimaryKey 
		{{ end -}}
	public {{$el.RealmTypeName}} {{$el.FieldName }};//{{$i}} 				 PB {{$el.TagNumber}}
   {{- end }}
	

	public static {{.RealmClass}} fromPB({{.Name}} pb){
		{{.RealmClass}} r = new {{.RealmClass}}();
		{{ range $i ,$el :=  .Fields}}
		{{if (fIsRealmType $el) -}}
	    //r.{{$el.FieldName}} = pb.get{{$el.FieldName}}();//{{$i}}
	    {{- else -}}
	     r.{{$el.FieldName}} = pb.get{{$el.FieldName}}();//{{$i}}
	    {{- end}}
	    {{- end }}

	    return r;
	}

	public static {{.Name}} toPB({{.RealmClass}} rV){//realmView
		{{.Name}}.Builder pbB = {{.Name}}.newBuilder();
		{{ range $i ,$el :=  .Fields}}
		{{if (fIsRealmType $el) -}}
		//r.{{$el.FieldName}} = pb.get{{$el.FieldName}}();//{{$i}}
		{{- else -}}
		 pbB.set{{$el.FieldName}}(rV.{{$el.FieldName}});//{{$i}}
		{{- end}}
		{{- end }}

		return pbB.build();
    	}

	public static {{.Name}}.Builder toPB_Builder({{.RealmClass}} rV){//realmView
		{{.Name}}.Builder pbB = {{.Name}}.newBuilder();
		{{ range $i ,$el :=  .Fields}}
		{{if (fIsRealmType $el) -}}
		//r.{{$el.FieldName}} = pb.get{{$el.FieldName}}();//{{$i}}
		{{- else -}}
		 pbB.set{{$el.FieldName}}(rV.{{$el.FieldName}});//{{$i}}
		{{- end}}
		{{- end }}

		return pbB;
    	}



}
	/*
	folding

	//sett
	{{.RealmClass}} r = new {{.RealmClass}}();
	{{ range $i ,$el :=  .Fields}}
	r.{{$el.FieldName}} = ;//{{$i}}
	{{- end }}

	//sett - no tag number 
	{{.RealmClass}} r = new {{.RealmClass}}();
	{{ range $i ,$el :=  .Fields}}
	r.{{$el.FieldName}} = ;
	{{- end }}

	//get
	{{ range $i ,$el :=  .Fields}}
	m. = r.{{$el.FieldName}} ;//{{$i}}
	{{- end }}

	//get - no tag number
	{{ range $i ,$el :=  .Fields}}
	m. = r.{{$el.FieldName}} ;
	{{- end }}


	//for PB.Builder
	{{.Name}}.Builder b = {{.Name}}.newBuilder();
	{{ range $i ,$el :=  .Fields}}
	{{if (fIsRealmType $el) -}}
	//r.{{$el.FieldName}} = pb.get{{$el.FieldName}}();//{{$i}}
	{{- else -}}
	 b.set{{$el.FieldName}}(r.{{$el.FieldName}});//{{$i}}
	{{- end}}
	{{- end }}


	//for PB.Builder - empty
	{{.Name}}.Builder b = {{.Name}}.newBuilder();
	{{ range $i ,$el :=  .Fields}}
	{{if (fIsRealmType $el) -}}
	//r.{{$el.FieldName}} = pb.get{{$el.FieldName}}();//{{$i}}
	{{- else -}}
	 b.set{{$el.FieldName}}();//{{$i}}
	{{- end}}
	{{- end }}



	
*/
