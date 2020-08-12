package ant

import (
	"github.com/emicklei/proto"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	_ "os"
	"os/exec"
	"path"
)

const OUTPUT_DIR_GO_X_CONST = `/home/hamid/life/_active/backbone/src/x/xconst/`
const OUTPUT_ANDROID_PROTO_MOUDLE_DIR = `/home/hamid/life/_active/backbone/src/x/pb/`
const OUTPUT_ANDROID_APP_DIR = `/home/hamid/life/_active/backbone/src/x/android/`
const OUTPUT_DIR_GO_X = `/home/hamid/life/_active/backbone/src/x/go/`
const OUTPUT_DIR_RUST_X = `/home/hamid/life/_active/backbone/src/`
const TEMPLATES_DIR = `/home/hamid/life/_active/pb_walker/templates/`
const DIR_PROTOS = `/home/hamid/life/_active/pb_walker/play/pb/`

const REALM = "realm"

const OUTPUT_ANDROID_REALM_DIR_ = `D:/ms/social/app/src/main/java/com/mardomsara/social/models_realm/pb_realm/`

func Run() {
	files, err := ioutil.ReadDir(DIR_PROTOS)
	noErr(err)
	filesName := make([]string, len(files))
	var prtos []*proto.Proto
	for i, pbFile := range files {
		filesName[i] = pbFile.Name()
		pbReader, err := os.Open(path.Join(DIR_PROTOS, pbFile.Name()))
		noErr(err)
		defer pbReader.Close()
		parser := proto.NewParser(pbReader)
		pbParesed, err := parser.Parse()
		if err != nil {
			log.Panic("error parsing proto: ", pbFile.Name(), " ", err, "/n")
		}
		prtos = append(prtos, pbParesed)
	}

	genOut := &GenOut{
		Messages: ExtractAllMessagesViews(prtos),
		Services: ExtractAllServicesViews(prtos),
		Enums:    ExtractAllEnumsViews(prtos),
	}
	//genOut.Realms = GetAllARealmMessageViews(genOut.Messages)
	genOut.Realms = GetAllARealmMessageViews_FromComments(genOut.Messages)

	print("===========================================")

	PrettyPrint(genOut)

	//buildGo(genOut)
	buildRust(genOut)


	err = exec.Command("javafmt").Run()
}
