package ant

import (
	"io/ioutil"
	"os"
)

func buildRust(gen *GenOut) {
	var OUT_DIR = gen.Dirs.RustOutDir
	//os.MkdirAll(OUTPUT_DIR_RUST_X, os.ModePerm)

	//writeOutputRust("rpc.rs", buildFromTemplate("rpc.rs", gen))
	//writeOutputRust("rpc_fns_default.rs", buildFromTemplate("rpc_fns_default.rs", gen))
	writeOutputRust("rpc.rs", buildFromTemplate("rpc.rs", gen), OUT_DIR)

	// Run cargo fmt
	currDir, err := os.Getwd()
	noErr(err)
	os.Chdir(gen.Dirs.RustProjectDir)
	// err = exec.Command("cargo", "fmt").Run()
	//noErr(err)
	os.Chdir(currDir)
}

func writeOutputRust(fileName, output, dirOut string) {
	err := ioutil.WriteFile(dirOut+fileName, []byte(output), os.ModePerm)
	noErr(err)
}
