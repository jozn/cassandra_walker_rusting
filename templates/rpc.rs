use hyper::{Body, Request, Response, Server};
use hyper::service::{make_service_fn, service_fn};
use hyper::body;
use serde::{Deserialize, Serialize};
use quick_protobuf::{BytesReader, BytesWriter};
use quick_protobuf::{MessageRead,MessageWrite,Writer};

use crate::{pb,com, pb::sys::Invoke,com::*};

pub fn server_rpc(act : Invoke) -> Result<Vec<u8>,GenErr> {
    let up = UserParam{};

    match act.method {
    {{range .Services}}
    // service: {{.Name}}
    {{- range .Methods}}
        {{.Hash}} => {
            let vec = "funk ".as_bytes().to_owned();

            let mut reader = BytesReader::from_bytes(&act.rpc_data);
            let rpc_param= pb::{{.InTypeName}}::from_reader(&mut reader, &act.rpc_data);

            if let Ok(param) = rpc_param {
            println!("param {:?}", param);
            let result = rpc::{{.MethodName}}(&up,param)?;

            let mut out_bytes = Vec::new();
            let mut writer = Writer::new(&mut out_bytes);
            let out = writer.write_message(&result);
            return Ok(out_bytes)
            } else {
            }
            Ok(vec)
        },
    {{- end}}
    {{end}}
        _ => {
           Err(GenErr{})
        }
    }
}

pub mod rpc {
    use super::*;

    {{range .Services}}
    // service: {{.Name}}
    {{- range .Methods}}
    pub fn {{.MethodName}}(up: &UserParam, param: pb::{{.InTypeName}}) -> Result<pb::{{.OutTypeName}}, GenErr> {
        Ok(pb::{{.OutTypeName}}::default())
    }

    {{- end}}
    {{end}}

    pub fn check_username(up: &UserParam, param: pb::EchoParam) -> Result<pb::EchoResponse, GenErr> {
        Ok(pb::EchoResponse::default())
    }
}
