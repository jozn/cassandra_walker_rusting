use hyper::{Body, Request, Response, Server};
use hyper::service::{make_service_fn, service_fn};
use hyper::body;
use serde::{Deserialize, Serialize};
use quick_protobuf::{BytesReader, BytesWriter};
use quick_protobuf::{MessageRead,MessageWrite,Writer,deserialize_from_slice};

use crate::{pb,com,com::*, rpc_fns};

pub mod method_ids {
    {{- range .Services}}
    // Service: {{.Name}}
    {{- range .Methods}}
    pub const {{.MethodName}}: u32 = {{.Hash}};
    {{- end}}
    {{end}}
    pub const ChangePhoneNumber8 : u32 = 79874;
}

pub async fn server_rpc(act : pb::Invoke) -> Result<Vec<u8>,GenErr> {
    let up = UserParam{};

    match act.method {
    {{range .Services}}
    // service: {{.Name}}
    {{- range .Methods}}
        method_ids::{{.MethodName}} => { // {{.Hash}}
            let vec: Vec<u8> = vec![];
            let rpc_param  : Result<pb::{{.InTypeName}}, ::prost::DecodeError> = prost::Message::decode(act.rpc_data.as_slice());

            if let Ok(param) = rpc_param {
                println!("param {:?}", param);
                let response = rpc_fns::{{.MethodName}}(&up, param).await?;

                let mut buff =vec![];
                prost::Message::encode(&response, &mut buff)?;

                Ok(buff)
            } else {
                Err(GenErr::ReadingPbParam)
            }
        }
    {{- end}}
    {{end}}
        _ => {
            Err(GenErr::NoRpcMatch)
        }
    }
}

pub struct RpcClient {
    endpoint: &'static str,
}

impl RpcClient {
    pub fn new(endpoint: &'static str) -> Self {
        RpcClient{
            endpoint: endpoint,
        }
    }

    fn get_next_action_id(&self) -> u64 {
        8
    }

{{range .Services -}}
// service: {{.Name}}
    {{- range .Methods}}
    pub async fn {{.MethodName}} (&self, param: pb::{{.InTypeName}}) -> Result<pb::{{.OutTypeName}},GenErr>{

        let mut buff =vec![];
        ::prost::Message::encode(&param, &mut buff)?;

        let invoke = pb::Invoke {
            namespace: 0,
            method: method_ids::{{.MethodName}},
            action_id: self.get_next_action_id() ,
            is_response: false,
            rpc_data: buff,
        };

        let mut buff =vec![];
        let m = prost::Message::encode(&invoke, &mut buff);

        let mut buff = Vec::new();
        ::prost::Message::encode(&invoke, &mut buff)?;

        let req = reqwest::Client::new()
            .post(self.endpoint)
            .body(buff)
            .send()
            .await?;

        let res_bytes = req.bytes().await?;
        let res_bytes = res_bytes.to_vec();

        let pb_res = ::prost::Message::decode(res_bytes.as_slice())?;
        Ok(pb_res)
    }
    {{end}}
{{end -}}
}

