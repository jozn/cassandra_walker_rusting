package ant

//////////////////////////// Raw PB Types /////////////////////

////////// PB Service /////////
type PBService struct {
	Name      string
	PBMethods []PBMethod
	Comment   string
	PBOptions []PBOptions
}

type PBMethod struct {
	MethodName  string
	InTypeName  string
	OutTypeName string
	Comment     string
	PBOptions   []PBOptions
}

////////// PB Messages /////////
type PBMessage struct {
	Name      string
	PBFields  []PBMessageField
	Comment   string
	PBOptions []PBOptions
}

type PBMessageField struct {
	FieldName string
	TypeName  string
	Repeated  bool
	TagNumber int
	Comment   string
	PBOptions []PBOptions
}

////////// PB Enums /////////
type PBEnum struct {
	Name      string
	PBFields  []PBEnumField
	Comment   string
	PBOptions []PBOptions
}

type PBEnumField struct {
	FieldName string
	TagNumber int
	PosNumber int
	Comment   string
	PBOptions []PBOptions
}

////////// PB Others /////////
type PBOptions struct {
	OptionName  string
	OptionValue string
}

//////////////////////////// Views Types ////////////////////////

////////// Service /////////
type ServiceView struct {
	PBService
	StripedName string
	Methods     []MethodView
}

type MethodView struct {
	PBMethod
	MethodNameStriped string // removed Chat, Channel, Group, Direct from rpc name
	GoInTypeName      string
	GoOutTypeName     string
	Hash              uint32
	FullMethodName    string // RPC_Other.Echo
	ParentServiceName string // RPC_Other
	DartMethodName    string // camelCase
}

////////// Messages /////////
type MessageView struct {
	PBMessage
	Fields []MessageFieldView
}

type MessageFieldView struct {
	PBMessageField
	isPrimitive bool // is ? numbers, bool, string, bytes or refrence to other custom types
	GoType      string
	GoFlatType  string
	JavaType    string
	RustType    string
}

////////// Enums /////////
type EnumView struct {
	PBEnum
	Fields []EnumFieldView
}

type EnumFieldView struct {
	PBEnumField
}

/////////////////////////////////////////
///////////// Extractor /////////////////
type PBGenOut struct {
	PBServices []PBService
	PBMessages []PBMessage
	PBEnums    []PBEnum
}

type GenOut struct {
	// Used directly in templates
	Services []ServiceView
	Messages []MessageView
	Enums    []EnumView

	OutGoEnumsStr string
	OutGoRPCsStr  string
	OutJavaStr    string

	Dirs DirParam
}

type DirParam struct {
	ProtoDir       string
	RustOutDir     string
	RustProjectDir string
}
