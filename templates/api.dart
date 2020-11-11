import 'dart:async' as $async;
import 'package:http/http.dart' as http;
import 'package:protobuf/protobuf.dart';
import 'package:protobuf/protobuf.dart' as $pb;
import 'package:flip_app/pb/global.pb.dart';

{{range .Services}}
import 'package:flip_app/pb/{{toLower .Name}}.pb.dart';
{{- end}}

{{range .Services}}
class {{.Name}} {
	{{- range .Methods}}
  static $async.Future<{{.OutTypeName}}> {{.DartMethodName}}(
      {{.InTypeName}} param) async {
    var paramBuff = param.writeToBuffer();

    var invoke = Invoke();
    invoke.namespace = 0;
    invoke.isResponse = false;
    invoke.method = {{.Hash}}; // {{.MethodName}}
    invoke.rpcData = paramBuff;

    var invokeBuff = invoke.writeToBuffer();

    var res = await http.post(
      "http://192.168.43.159:3002/rpc",
      body: invokeBuff,
      // encoding: Encoding.getByName("utf-8")
    );
    // print('Response ## len : ${res.body.length}');
    // print('Response ## bts  : ${res.bodyBytes}');

    var response = {{.OutTypeName}}();
    response.mergeFromBuffer(res.bodyBytes);
    return response;
  }
  {{end}}
}
{{- end}}

