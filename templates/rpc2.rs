//#![rustfmt::skip]

use crate::{pb, common};
use crate::pb::{EchoParam, EchoResponse};
use crate::{errors::GenErr, UserParam};
use async_trait::async_trait;

use http::Version;
use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Error as HyperError, Request, Response, Server};
use std::convert::Infallible;
use std::net::SocketAddr;

#[derive(Debug)]
pub struct RpcInvoke {
    method_id: i64, // correct data type should be i32,
    rpc_service: RpcServiceData,
}

#[derive(Debug)]
pub enum RpcServiceData {
{{- range .Services}}
    {{.Name}}({{.Name}}_MethodData),
{{- end }}
}

{{range .Services}}
#[derive(Debug)]
pub enum {{.Name}}_MethodData {
    {{- range .Methods}}
    {{.MethodName}}(pb::{{.InTypeName}}),
    {{- end }}
}
{{- end }}

{{range .Services}}
#[async_trait]
pub trait {{.Name}}_Handler {
    {{- range .Methods}}
    async fn {{.MethodName}}(up: &UserParam, param: pb::{{.InTypeName}}) -> Result<pb::{{.OutTypeName}}, GenErr> {
        Ok(pb::{{.OutTypeName}}::default())
    }
    {{- end }}
}
{{- end }}


{{range .Services}}
#[async_trait]
pub trait {{.Name}}_Handler2 : Send + Sync {
    {{- range .Methods}}
    async fn {{.MethodName}}(&self, param: pb::{{.InTypeName}}) -> Result<pb::{{.OutTypeName}}, GenErr> {
        Ok(pb::{{.OutTypeName}}::default())
    }
    {{- end }}
}
{{- end }}

#[async_trait]
pub trait All_Rpc_Handler :
{{- range .Services -}}
    {{- .Name}}_Handler2 +
{{- end -}}
Clone + Send + Sync {}

pub mod method_ids {
    {{- range .Services}}
    // Service: {{.Name}}
    {{- range .Methods}}
    pub const {{.MethodName}}: u32 = {{.Hash}};
    {{- end}}
    {{end}}
    pub const ExampleChangePhoneNumber8 : u32 = 79874;
}

#[derive(Debug)]
pub enum MethodIds {
    {{- range .Services}}
    // Service: {{.Name}}
    {{- range .Methods}}
    {{.MethodName}} = {{.Hash}},
    {{- end}}
    {{end}}
}

pub fn invoke_to_parsed(invoke: &pb::Invoke) -> Result<RpcInvoke, GenErr>{
    use RpcServiceData::*;
    let rpc = match invoke.method {
{{- range .Services}}
    // {{.Name}}
    {{- $service := . }}
    {{- range .Methods}}
        method_ids::{{.MethodName}} => {
           let rpc_param: pb::{{.InTypeName}} = prost::Message::decode(invoke.rpc_data.as_slice())?;
           RpcInvoke{
                method_id: {{.Hash}} as i64,
                rpc_service: {{$service.Name}}({{$service.Name}}_MethodData::{{.MethodName}}(rpc_param)),
           }
        },
{{ end }}
{{- end }}
        _ => { panic!("sdf")}
    };
    Ok(rpc)
}

pub async fn server_rpc(act: RpcInvoke, reg: &RPC_Registry) -> Result<Vec<u8>, GenErr> {

    let res_v8 = match act.rpc_service {
{{- range .Services}}
    {{$service := .}}
    RpcServiceData::{{.Name}}(method) => match method {
            {{- range .Methods}}
                {{$service.Name}}_MethodData::{{.MethodName}}(param) => {
                   let handler = eror(&reg.{{$service.Name}})?;
                   let response = handler.{{.MethodName}}(param).await?;
                   let v8 = to_vev8(&response)?;
                   v8
                },
             {{ end }}
        },

{{ end }}
    };

    Ok(res_v8)
}

#[derive(Default)]
pub struct RPC_Registry {
{{- range .Services}}
    pub {{.Name}}: Option<Box<{{.Name}}_Handler2>>,
{{- end -}}
}

{{range .Services}}
impl {{.Name}}_Handler for RPC_Registry {}
impl {{.Name}}_Handler2 for RPC_Registry {}
{{- end }}



fn to_vev8(msg: &impl prost::Message) -> Result<Vec<u8>, GenErr> {
    let mut buff = vec![];
    prost::Message::encode(msg, &mut buff)?;
    Ok(buff)
}

fn eror<T>(input :&Option<T>) -> Result<&T, GenErr> {
    match input {
        Some(inbox) => Ok(inbox),
        None => Err(GenErr::NoRpcRegistry),
    }
}

impl common::RpcClient {

{{range .Services -}}
// service: {{.Name}}
    {{- range .Methods}}
    pub async fn {{.MethodName}} (&self, param: pb::{{.InTypeName}})
        -> Result<pb::{{.OutTypeName}},GenErr>{
        let pb_res = self.rpc_invoke(&param, method_ids::{{.MethodName}}).await?;
        Ok(pb_res)
    }
    {{end}}
{{end -}}
}

/////////////////////// Code gen for def rpc //////////////
struct _RRR_ {}
{{range .Services}}
#[async_trait]
impl {{.Name}}_Handler2 for _RRR_ {
    {{- range .Methods}}
    async fn {{.MethodName}}(&self, param: pb::{{.InTypeName}}) -> Result<pb::{{.OutTypeName}}, GenErr> {
        println!("called {{.MethodName}} in the impl code.");
        Ok(pb::{{.OutTypeName}}::default())
    }
    {{- end }}
}
{{- end }}
