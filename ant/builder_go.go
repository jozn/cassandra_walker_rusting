package ant

import (
	"io/ioutil"
	"os"
)

func build_old(gen *GenOut) {
	os.MkdirAll(OUTPUT_DIR_GO_X_CONST, os.ModePerm)
	//os.MkdirAll(OUTPUT_DIR_GO_X, os.ModePerm)
	//os.MkdirAll(OUTPUT_ANDROID_APP_DIR, os.ModePerm)
	//os.MkdirAll(OUTPUT_ANDROID_PROTO_MOUDLE_DIR, os.ModePerm)

	OutGoRPCsStr := buildFromTemplate("rpc.tgo", gen)
	writeOutputGo("pb__gen_ant.go", OutGoRPCsStr)

	OutGoRPCsEmptyStr := buildFromTemplate("rpc_empty_imple.tgo", gen)
	writeOutputGo("pb__gen_ant_empty.go", OutGoRPCsEmptyStr)

	writeOutputGo("pb__gen_enum.proto", buildFromTemplate("enums.proto", gen))
	writeOutputGo("RPC_HANDLERS.java", buildFromTemplate("RPC_HANDLERS.java", gen))
	writeOutputGo("PBFlatTypes.java", buildFromTemplate("PBFlatTypes.java", gen))
	writeOutputGo("flat.go", buildFromTemplate("flat.tgo", gen))
	writeOutputConstantGo("pb.go", buildFromTemplate("xconst.tgo", gen))

	writeOutputGo("rpc_client.go", buildFromTemplate("rpc_client.tgo", gen))

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

func writeOutputGo(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_DIR_GO_X+fileName, []byte(output), os.ModePerm)
	noErr(err)
}

func writeOutputConstantGo(fileName, output string) {
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
