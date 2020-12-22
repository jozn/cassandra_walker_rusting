//#![rustfmt::skip]

use crate::pb;
use crate::pb::{EchoParam, EchoResponse};
use crate::{errors::GenErr, UserParam};
use async_trait::async_trait;

use http::Version;
use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Error as HyperError, Request, Response, Server};
use std::convert::Infallible;
use std::net::SocketAddr;

pub struct RpcInvoke {
    method_id: i64, // correct data type should be i32,
    rpc_service: RpcServiceData,
}

pub enum RpcServiceData {
{{- range .Services}}
    {{.Name}}({{.Name}}_MethodData),
{{- end }}
}

{{range .Services}}
pub enum {{.Name}}_MethodData {
    {{- range .Methods}}
    {{.MethodName}}(pb::{{.InTypeName}}),
    {{- end }}
}
{{- end }}

{{range .Services}}
#[async_trait]
trait {{.Name}}_Handler {
    {{- range .Methods}}
    async fn {{.MethodName}}(up: &UserParam, param: pb::{{.InTypeName}}) -> Result<pb::{{.OutTypeName}}, GenErr> {
        Ok(pb::{{.OutTypeName}}::default())
    }
    {{- end }}
}
{{- end }}


{{range .Services}}
#[async_trait]
trait {{.Name}}_Handler2 {
    {{- range .Methods}}
    async fn {{.MethodName}}(&self, param: pb::{{.InTypeName}}) -> Result<pb::{{.OutTypeName}}, GenErr> {
        Ok(pb::{{.OutTypeName}}::default())
    }
    {{- end }}
}
{{- end }}


pub mod method_ids {
    {{- range .Services}}
    // Service: {{.Name}}
    {{- range .Methods}}
    pub const {{.MethodName}}: u32 = {{.Hash}};
    {{- end}}
    {{end}}
    pub const ExampleChangePhoneNumber8 : u32 = 79874;
}

pub fn invoke_to_parsed(invoke: &pb::Invoke) -> Result<RpcInvoke, GenErr>{
    use RpcServiceData::*;
    let rpc = match invoke.method {
{{- range .Services}}
    // {{.Name}}
    {{- $service := . }}
    {{- range .Methods}}
        method_ids::{{.MethodName}} => {
           let rpc_param  : Result<pb::{{.InTypeName}}, ::prost::DecodeError> =
                   prost::Message::decode(invoke.rpc_data.as_slice());
           let rpc_param  = rpc_param.unwrap();
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

pub async fn server_rpc(act: RpcInvoke, reg: RPC_Registry) -> Result<Vec<u8>, GenErr> {

    let res_v8 = match act.rpc_service {
{{- range .Services}}
    {{$service := .}}
    RpcServiceData::{{.Name}}(method) => match method {
            {{- range .Methods}}
                {{$service.Name}}_MethodData::{{.MethodName}}(rr) => {
                   // reg.{{.MethodName}}();
                   let response = reg.{{.MethodName}}(rr).await?;
                   let v8 = to_vev8(&response)?;
                   v8
                   //response
/*                    let mut buff =vec![];
                   prost::Message::encode(&response, &mut buff)?;
                   buff */
                },
             {{ end }}
        },

{{ end }}
    };

    Ok(res_v8)
}


pub struct RPC_Registry {
    // RPC_Shared: RPC_Shared,
// RPC_Chat: RPC_Chat,
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
