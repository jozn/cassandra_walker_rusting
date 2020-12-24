pub async fn server_rpc(act: RpcInvoke, reg: impl All_Rpc_Handler) -> Result<Vec<u8>, GenErr> {

    let res_v8 = match act.rpc_service {
{{- range .Services}}
    {{$service := .}}
    RpcServiceData::{{.Name}}(method) => match method {
            {{- range .Methods}}
                {{$service.Name}}_MethodData::{{.MethodName}}(param) => {
                   let reg = reg.clone();
                   let response = reg.{{.MethodName}}(param).await?;
                   let v8 = to_vev8(&response)?;
                   v8
                },
             {{ end }}
        },

{{ end }}
    };

    Ok(res_v8)
}