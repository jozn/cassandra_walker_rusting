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

const OUTPUT_DIR_GO_X = `C:/Go/_gopath/src/ms/sun/shared/x/`                         //"./play/gen_sample_out.go"
const OUTPUT_DIR_GO_X_CONST = `C:/Go/_gopath/src/ms/sun/shared/x/xconst/`            //"./play/gen_sample_out.go"
const OUTPUT_ANDROID_PROTO_MOUDLE_DIR = `D:\ms\social\proto\src\main\java\ir\ms\pb\` //`D:/dev_working2/MS_Native/proto/src/main/java/ir/ms/pb/` //"./play/gen_sample_out.go"
const OUTPUT_ANDROID_APP_DIR = `D:\ms\social\app\src\main\java\ir\ms\pb\`            // `D:/dev_working2/MS_Native/app/src/main/java/ir/ms/pb/`            //"./play/gen_sample_out.go"
//const TEMPLATES_DIR = "./templates/"                    //relative to main func of parent directory
const TEMPLATES_DIR = `C:/Go/_gopath/src/ms/pb_walker/templates/` //relative to main func of parent directory
const DIR_PROTOS = `C:/Go/_gopath/src/ms/sun/shared/proto`

const REALM = "realm"

const OUTPUT_ANDROID_REALM_DIR_ = `D:/ms/social/app/src/main/java/com/mardomsara/social/models_realm/pb_realm/`

func Run() {
	files, err := ioutil.ReadDir(DIR_PROTOS)
	noErr(err)
	protos := make([]string, len(files))
	var prtos []*proto.Proto
	for i, f := range files {
		protos[i] = f.Name()
		reader, err := os.Open(path.Join(DIR_PROTOS, f.Name()))
		noErr(err)
		defer reader.Close()
		parser := proto.NewParser(reader)
		def, err := parser.Parse()
		if err != nil {
			log.Panic("error parsing proto: ", f.Name(), " ", err, "/n")
		}
		prtos = append(prtos, def)
	}

	gen := &GenOut{
		Messages: ExtractAllMessagesViews(prtos),
		Services: ExtractAllServicesViews(prtos),
		Enums:    ExtractAllEnumsViews(prtos),
	}
	//gen.Realms = GetAllARealmMessageViews(gen.Messages)
	gen.Realms = GetAllARealmMessageViews_FromComments(gen.Messages)

	print("===========================================")

	build(gen)

	err = exec.Command("javafmt").Run()
}
