package src

import (
	"fmt"
	"io/ioutil"
	"os"
	_ "os"
	"os/exec"
	"path"

	"github.com/emicklei/proto"
	"github.com/labstack/gommon/log"
)

const TEMPLATES_DIR = `/home/hamid/life/_active/pb_walker/templates_v2/`
const OUTPUT_DIR_DART = `/hamid/life/flip/flip_app2/lib/ui/`

func Run() {
	dirs := DirParam{
		ProtoDir:       `/home/hamid/life/_active/backbone/lib/shared/src/man/protos/proto/`,
		RustOutDir:     `/home/hamid/life/_active/backbone/lib/shared/src/gen/`,
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

	fmt.Println("===========================================")

	PrettyPrint(genOut)

	buildRust(genOut)
	//buildDart(genOut)

	err = exec.Command("javafmt").Run()
}

func getGenOut(prtos []*proto.Proto) *GenOut {
	pbGenOut := &PBExtract{
		PBMessages: ExtractAllPBMessages(prtos),
		PBServices: ExtractAllPBServices(prtos),
		PBEnums:    ExtractAllPBEnums(prtos),
	}

	genOut := &GenOut{
		Messages: processAllMessagesViews(pbGenOut.PBMessages),
		Services: processAllServicesViews(pbGenOut.PBServices),
		Enums:    processAllEnumsViews(pbGenOut.PBEnums),
	}

	return genOut
}
