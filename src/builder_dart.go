package src

import (
	"io/ioutil"
	"os"
)

func buildDart(gen *GenOut) {
	//os.MkdirAll(OUTPUT_DIR_RUST_X, os.ModePerm)

	writeOutputDart("api.dart", buildFromTemplate("api.dart", gen))
}

func writeOutputDart(fileName, output string) {
	err := ioutil.WriteFile(OUTPUT_DIR_DART+fileName, []byte(output), os.ModePerm)
	noErr(err)
}
