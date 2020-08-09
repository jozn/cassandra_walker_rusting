package ant

////////// Service /////////
type ServiceView struct {
	Name    string
	Methods []MethodView
	Comment string
	Hash    uint32
	Options []OptionsView
}

type MethodView struct {
	MethodName        string
	InTypeName        string
	GoInTypeName      string
	OutTypeName       string
	GoOutTypeName     string
	Hash              uint32
	Options           []OptionsView
	FullMethodName    string // RPC_Other.Echo
	ParentServiceName string // RPC_Other
}

////////// Messages /////////
type MessageView struct {
	Name       string
	Fields     []FieldView
	Comment    string
	Options    []OptionsView
	RealmClass string
}

type FieldView struct {
	FieldName     string
	TypeName      string
	Repeated      bool
	TagNumber     int
	isPrimitive   bool // is ? numbers, bool, string, bytes or refrence to other custom types
	GoType        string
	GoFlatType    string
	JavaType      string
	RustType      string
	Options       []OptionsView
	RealmTypeName string
}

////////// Enums /////////
type EnumView struct {
	Name    string
	Fields  []EnumFieldView
	Comment string
	Options []OptionsView
}

type EnumFieldView struct {
	FieldName string
	TagNumber int
	PosNumber int
	Options   []OptionsView
}

/////////// Tag /////////
type OptionsView struct {
	OptionName  string
	OptionValue string
}

/////////////////////////////////////////
///////////// Extractor /////////////////
type GenOut struct {
	Services []ServiceView
	Messages []MessageView
	Enums    []EnumView

	Realms []MessageView

	OutGoEnumsStr string
	OutGoRPCsStr  string
	OutJavaStr    string
}
