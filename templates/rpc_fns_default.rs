use crate::{com, com::*, pb};

{{range .Services}}
// Service: {{.Name}}
pub mod {{.Name}} {
    use super::*;
    {{range .Methods}}
    pub fn {{.MethodName}}(up: &UserParam, param: pb::{{.InTypeName}}) -> Result<pb::{{.OutTypeName}}, GenErr> {
        Ok(pb::{{.OutTypeName}}::default())
    }
    {{- end}}
}
{{end}}

/*
{{range .Services -}}
pub use def::{{.Name}}::*;
{{end}}
 */
