package ant

import (
	"io/ioutil"
	"os"
)

func buildRust(gen *GenOut) {
	os.MkdirAll(OUTPUT_DIR_RUST_X, os.ModePerm)

	writeOutputRust("rpc.rs", buildFromTemplate("rpc.rs", gen))
	writeOutputRust("rpc_fns_default.rs", buildFromTemplate("rpc_fns_default.rs", gen))
}

func writeOutputRust(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_DIR_RUST_X+fileName, []byte(output), os.ModePerm)
	noErr(err)
}
