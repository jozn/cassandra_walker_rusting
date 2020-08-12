package ant

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

func buildRust(gen *GenOut) {
	os.MkdirAll(OUTPUT_DIR_RUST_X, os.ModePerm)

	writeOutputRust("rpc.rs", buildFromTemplate("rpc.rs", gen))
}

func buildGo(gen *GenOut) {
	os.MkdirAll(OUTPUT_DIR_GO_X, os.ModePerm)
	os.MkdirAll(OUTPUT_DIR_GO_X_CONST, os.ModePerm)
	os.MkdirAll(OUTPUT_ANDROID_APP_DIR, os.ModePerm)
	os.MkdirAll(OUTPUT_ANDROID_PROTO_MOUDLE_DIR, os.ModePerm)

	OutGoRPCsStr := buildFromTemplate("rpc.tgo", gen)
	writeOutputGo("rpc.go", OutGoRPCsStr)

	OutGoRPCsEmptyStr := buildFromTemplate("rpc_empty_imple.tgo", gen)
	writeOutputGo("empty.go", OutGoRPCsEmptyStr)
	writeOutputGo("pb__gen_enum.proto", buildFromTemplate("enums.proto", gen))
	writeOutputGo("RPC_HANDLERS.java", buildFromTemplate("RPC_HANDLERS.java", gen))
	writeOutputGo("PBFlatTypes.java", buildFromTemplate("PBFlatTypes.java", gen))
	writeOutputGo("flat.go", buildFromTemplate("flat.tgo", gen))
	writeOutputConstantGo("pb.go", buildFromTemplate("xconst.tgo", gen))
	writeOutputGo("rpc_client.go", buildFromTemplate("rpc_client.tgo", gen))

	build_old(gen)

}

func writeOutputGo(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_DIR_GO_X+fileName, []byte(output), os.ModePerm)
	noErr(err)
}

func writeOutputConstantGo(fileName, output string) {
	ioutil.WriteFile(OUTPUT_DIR_GO_X_CONST+fileName, []byte(output), os.ModePerm)
}

func writeOutputRust(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_DIR_RUST_X+fileName, []byte(output), os.ModePerm)
	noErr(err)
}
///////////////////// Archives /////////////////

func build_old(gen *GenOut) {
	os.MkdirAll(OUTPUT_DIR_GO_X_CONST, os.ModePerm)

	OutGoRPCsStr := buildFromTemplate("rpc.tgo", gen)
	writeOutput("pb__gen_ant.go", OutGoRPCsStr)

	OutGoRPCsEmptyStr := buildFromTemplate("rpc_empty_imple.tgo", gen)
	writeOutput("pb__gen_ant_empty.go", OutGoRPCsEmptyStr)

	writeOutput("pb__gen_enum.proto", buildFromTemplate("enums.proto", gen))
	writeOutput("RPC_HANDLERS.java", buildFromTemplate("RPC_HANDLERS.java", gen))
	writeOutput("PBFlatTypes.java", buildFromTemplate("PBFlatTypes.java", gen))
	writeOutput("flat.go", buildFromTemplate("flat.tgo", gen))
	writeOutputConstant("pb.go", buildFromTemplate("xconst.tgo", gen))

	writeOutput("rpc_client.go", buildFromTemplate("rpc_client.tgo", gen))

	//////////////// For Android /////////////
	writeOutputAndroidProto("RPC_HANDLERS.java", buildFromTemplate("RPC_HANDLERS.java", gen))
	writeOutputAndroidProto("PBFlatTypes.java", buildFromTemplate("PBFlatTypes.java", gen))
	writeOutputAndroidApp("RPC.java", buildFromTemplate("RPC.java", gen))
	writeOutputAndroidApp("RPC_API.java", buildFromTemplate("RPC_API.java", gen))
	writeOutputAndroidApp("RpcNameToResponseMapper.java", buildFromTemplate("RpcNameToResponseMapper.java", gen))
	writeOutputAndroidProto("RPC_ResponseBase.java", buildFromTemplate("RPC_ResponseBase.java", gen))
	//copy
	writeOutputAndroidProto("Log.java", buildFromTemplate("Log.java", gen))

	/////// Enums /////////////////

}

func buildFromTemplate(tplName string, gen *GenOut) string {
	tpl := template.New("go_interface" + tplName)
	tpl.Funcs(fns)
	tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
	noErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	noErr(err)

	buffer := bytes.NewBufferString("")
	err = tpl.Execute(buffer, gen)
	noErr(err)

	return buffer.String()
}

func writeOutput(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_DIR_GO_X+fileName, []byte(output), os.ModePerm)
	noErr(err)
}

func writeOutputConstant(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_DIR_GO_X_CONST+fileName, []byte(output), os.ModePerm)
	noErr(err)
}

func writeOutputAndroidProto(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_ANDROID_PROTO_MOUDLE_DIR+fileName, []byte(output), os.ModePerm)
	noErr(err)
}

func writeOutputAndroidApp(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_ANDROID_APP_DIR+fileName, []byte(output), os.ModePerm)
	noErr(err)
}
