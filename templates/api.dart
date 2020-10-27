import 'dart:async' as $async;
import 'package:http/http.dart' as http;
import 'package:protobuf/protobuf.dart';
import 'package:protobuf/protobuf.dart' as $pb;
import 'package:flip_app/pb/sys.pb.dart';

{{range .Services}}
import 'package:flip_app/pb/{{toLower .Name}}.pb.dart';
{{- end}}

var rpcNameToHashId = {
{{- range .Services}}
	{{- range .Methods}}
        '{{.MethodName}}': {{.Hash}},
    {{- end -}}
{{- end}}
};


class FlipHttpRpcClient extends RpcClient {
  @override
  Future<T> invoke<T extends GeneratedMessage>(
      ClientContext ctx,
      String serviceName,
      String methodName,
      GeneratedMessage param,
      T emptyResponse) async {

    var paramBuff = param.writeToBuffer();

    var hashId = rpcNameToHashId[methodName];

    var invoke = Invoke();
    invoke.namespace = 0;
    invoke.isResponse = false;
    invoke.method = hashId;
    invoke.rpcData = paramBuff;

    var invokeBuff = invoke.writeToBuffer();

    var res = await http.post(
      "http://192.168.43.159:3002/rpc",
      body: invokeBuff,
      // encoding: Encoding.getByName("utf-8")
    );

    print('Response : ${res}');
  }
}


class FlipRpcPbClientContext extends $pb.ClientContext {
  
}

{{range .Services}}
{{$rn := .Name }}
class {{.NameStriped}} {
	{{- range .Methods}}
  static $async.Future<{{.OutTypeName}}> {{.DartMethodName}}({{.InTypeName}} param) async {
    var rpcClient = FlipHttpRpcClient();
    var serviceRpc = {{$rn}}Api(rpcClient);

    var ctxRpc = FlipRpcPbClientContext();

    return await serviceRpc.{{.DartMethodName}}(ctxRpc, param);
  }
  {{- end -}}
}
{{- end}}

class Auth22 {
  static $async.Future<SendConfirmCodeResponse> sendConfirmCode(SendConfirmCodeParam param) async {
    var rpcClient = FlipHttpRpcClient();
    var serviceRpc = RPC_AuthApi(rpcClient);

    var ctxRpc = FlipRpcPbClientContext();

    return await serviceRpc.sendConfirmCode(ctxRpc, param);
  }
}
