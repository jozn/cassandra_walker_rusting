package ant

import (
	"io/ioutil"
	"os"
	_ "os"
	"os/exec"
	"path"

	"github.com/emicklei/proto"
	"github.com/labstack/gommon/log"
)

const OUTPUT_DIR_GO_X_CONST = `/home/hamid/life/_active/backbone/src/x/xconst/`
const OUTPUT_ANDROID_PROTO_MOUDLE_DIR = `/home/hamid/life/_active/backbone/src/x/pb/`
const OUTPUT_ANDROID_APP_DIR = `/home/hamid/life/_active/backbone/src/x/android/`
const OUTPUT_DIR_GO_X = `/home/hamid/life/_active/backbone/src/x/go/`
//const OUTPUT_DIR_RUST_X = `/home/hamid/life/_active/backbone/lib/shared/src/`
//const RUST_PROJECT = `/home/hamid/life/_active/backbone/`
const TEMPLATES_DIR = `/home/hamid/life/_active/pb_walker/templates/`
//const DIR_PROTOS = `/home/hamid/life/_active/backbone/lib/shared/src/protos/proto/`

const OUTPUT_DIR_DART = `/hamid/life/flip/flip_app2/lib/ui/`

func Run() {
	dirs := DirParam{
		ProtoDir:   `/home/hamid/life/_active/backbone/lib/shared/src/protos/proto/`,
		//ProtoDir:   `//hamid/life/_active/pb_walker/play/pb2/`, // play codes
		RustOutDir: `/home/hamid/life/_active/backbone/lib/shared/src/`,
		RustProjectDir: `/home/hamid/life/_active/backbone/`,
	}

	protoDir := dirs.ProtoDir

	files, err := ioutil.ReadDir(protoDir)
	noErr(err)
	filesName := make([]string, len(files))
	var prtos []*proto.Proto
	for i, pbFile := range files {
		filesName[i] = pbFile.Name()
		pbReader, err := os.Open(path.Join(protoDir, pbFile.Name()))
		noErr(err)
		defer pbReader.Close()
		parser := proto.NewParser(pbReader)
		pbParesed, err := parser.Parse()
		if err != nil {
			log.Panic("error parsing proto: ", pbFile.Name(), " ", err, "/n")
		}
		prtos = append(prtos, pbParesed)
	}

	genOut := getGenOut(prtos)
	genOut.Dirs = dirs

	print("===========================================")

	PrettyPrint(genOut)

	//buildGo(genOut)
	buildRust(genOut)
	buildDart(genOut)

	err = exec.Command("javafmt").Run()
}

func getGenOut(prtos []*proto.Proto) *GenOut {
	genOut := &GenOut{
		PBMessages: ExtractAllPBMessages(prtos),
		PBServices: ExtractAllPBServices(prtos),
		PBEnums:    ExtractAllPBEnums(prtos),
	}
	genOut.Messages = processAllMessagesViews(genOut.PBMessages)
	genOut.Services = processAllServicesViews(genOut.PBServices)
	genOut.Enums = processAllEnumsViews(genOut.PBEnums)

	return genOut
}
