package src

import (
	"io/ioutil"
	"os"
)

func buildRust(gen *GenOut) {
	var OUT_DIR = gen.Dirs.RustOutDir

	writeOutputRust("rpc2.rs", buildFromTemplate("rpc.rs", gen), OUT_DIR)

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
