use hyper::{Body, Request, Response, Server};
use hyper::service::{make_service_fn, service_fn};
use hyper::body;
use serde::{Deserialize, Serialize};
use quick_protobuf::{BytesReader, BytesWriter};
use quick_protobuf::{MessageRead,MessageWrite,Writer};

use crate::{pb,com, pb::sys::Invoke,com::*, rpc_fns};

pub fn server_rpc(act : Invoke) -> Result<Vec<u8>,GenErr> {
    let up = UserParam{};

    match act.method {
    {{range .Services}}
    // service: {{.Name}}
    {{- range .Methods}}
        {{.Hash}} => {
            let vec: Vec<u8> = vec![];
            let rpc_param = BytesReader::from_bytes(&vec)
                .read_message::<pb::{{.InTypeName}}>(&act.rpc_data);

            if let Ok(param) = rpc_param {
                println!("param {:?}", param);
                let result = rpc_fns::{{.MethodName}}(&up, param)?;

                let mut out_bytes = Vec::new();
                let _result = Writer::new(&mut out_bytes).write_message(&result);

                Ok(out_bytes)
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
