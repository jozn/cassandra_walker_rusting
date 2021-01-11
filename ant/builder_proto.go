package ant

import (
	"io/ioutil"
	"os"
)

func buildProto(gen *GenOut) {
	var OUT_DIR = gen.Dirs.ProtoOutDir
	err := os.MkdirAll(OUT_DIR, os.ModePerm)
	noErr(err)

	_writeOutput("enums.proto", buildFromTemplate("pb/enums.proto", gen), OUT_DIR)

}

func _writeOutput(fileName, output, dirOut string) {
	err := ioutil.WriteFile(dirOut+fileName, []byte(output), os.ModePerm)
	noErr(err)
}
