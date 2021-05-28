package src

import (
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"
)

func buildFromTemplate(tplName string, gen *GenOut) string {
	tpl := template.New("go_interface" + tplName)
	tpl.Funcs(fns)
	tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
	noErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	noErr(err)

	buffer := bytes.NewBufferString("")
	err = tpl.Execute(buffer, gen)
	noErr(err)

	return buffer.String()
}

///////////// Template funcs /////////////

var fns = template.FuncMap{
	"toLower": strings.ToLower,
}
