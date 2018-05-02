package ant

import (
	"github.com/emicklei/proto"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"ms/sun/shared/helper"
	"os"
	_ "os"
	"os/exec"
	"path"
)

const OUTPUT_DIR_GO_X = `C:/Go/_gopath/src/ms/sun/shared/x/`                                   //"./play/gen_sample_out.go"
const OUTPUT_DIR_GO_X_CONST = `C:/Go/_gopath/src/ms/sun/shared/x/xconst/`                      //"./play/gen_sample_out.go"
const OUTPUT_ANDROID_PROTO_MOUDLE_DIR = `D:\ms\social\proto\src\main\java\ir\ms\pb\` //`D:/dev_working2/MS_Native/proto/src/main/java/ir/ms/pb/` //"./play/gen_sample_out.go"
const OUTPUT_ANDROID_APP_DIR = `D:\ms\social\app\src\main\java\ir\ms\pb\`            // `D:/dev_working2/MS_Native/app/src/main/java/ir/ms/pb/`            //"./play/gen_sample_out.go"
//const TEMPLATES_DIR = "./templates/"                    //relative to main func of parent directory
const TEMPLATES_DIR = `C:/Go/_gopath/src/ms/pb_walker/templates/` //relative to main func of parent directory
const DIR_PROTOS = `C:/Go/_gopath/src/ms/sun/shared/proto`

const REALM = "realm"

const OUTPUT_ANDROID_REALM_DIR_ = `D:/ms/social/app/src/main/java/com/mardomsara/social/models_realm/pb_realm/`

/*func Run() {
	xxx()
	files, err := ioutil.ReadDir(DIR_PROTOS)
	noErr(err)
	protos := make([]string, len(files))
	for i, f := range files {
		protos[i] = f.Name()
	}

	//ast, err := parser.ParseFiles([]string{"1.proto"}, []string{"./play/"})
	ast, err := parser.ParseFiles(protos, []string{DIR_PROTOS})
	noErr(err)
	gen := &GenOut{
		Messages: ExtractAllMessagesViews(ast),
		Services: ExtractAllServicesViews(ast),
		Enums:    ExtractAllEnumsViews(ast),
	}

	build(gen)

	/////////// commeant albe ///
	os.Chdir(`C:/Go/_gopath/src/ms/sun/scripts/`)
	err = exec.Command(`C:/Go/_gopath/src/ms/sun/scripts/gen_pb.exe`).Run()
	noErr(err)
	err = exec.Command("gofmt", "-w", OUTPUT_DIR_GO_X).Run()
	//noErr(err)

	err = exec.Command("javafmt").Run()
	////
}
*/
func RunV2() {
	//xxx()
	files, err := ioutil.ReadDir(DIR_PROTOS)
	noErr(err)
	protos := make([]string, len(files))
	var prtos []*proto.Proto
	for i, f := range files {
		protos[i] = f.Name()
		reader, err := os.Open(path.Join(DIR_PROTOS, f.Name()))
		helper.NoErr(err)
		defer reader.Close()
		parser := proto.NewParser(reader)
		def, err := parser.Parse()
		if err != nil {
			log.Panic("error parsing proto: ", f.Name(), " ", err, "/n")
		}
		prtos = append(prtos, def)
	}

	gen := &GenOut{
		Messages: ExtractAllMessagesViewsV2(prtos),
		Services: ExtractAllServicesViewsV2(prtos),
		Enums:    ExtractAllEnumsViewsV2(prtos),
	}
	//gen.Realms = GetAllARealmMessageViews(gen.Messages)
	gen.Realms = GetAllARealmMessageViews_FromComments(gen.Messages)

	print("===========================================")
	//helper.PertyPrint(gen.Messages)
	//helper.PertyPrint(prtos)

	build(gen)

	/////////// commeant albe ///
	/*os.Chdir(`C:/Go/_gopath/src/ms/sun_old/scripts/`)
	  err = exec.FullMethodName(`C:/Go/_gopath/src/ms/sun_old/scripts/gen_pb.exe`).Run()
	  noErr(err)
	  err = exec.FullMethodName("gofmt", "-w", OUTPUT_DIR_GO_X).Run()*/
	//noErr(err)

	err = exec.Command("javafmt").Run()
	////
}

func alaki() {
	exec.Command(``).Run()
	os.Chdir(``)
}
