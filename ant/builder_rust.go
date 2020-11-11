package ant

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func buildRust(gen *GenOut) {
	os.MkdirAll(OUTPUT_DIR_RUST_X, os.ModePerm)

	writeOutputRust("rpc.rs", buildFromTemplate("rpc.rs", gen))
	writeOutputRust("rpc_fns_default.rs", buildFromTemplate("rpc_fns_default.rs", gen))

	// Run cargo fmt
	currDir,err:= os.Getwd()
	noErr(err)
	os.Chdir(RUST_PROJECT)
	err = exec.Command("cargo","fmt").Run()
	noErr(err)
	os.Chdir(currDir)
}

func writeOutputRust(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_DIR_RUST_X+fileName, []byte(output), os.ModePerm)
	noErr(err)
}
