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
	// From PB
	Name    string
	Comment string
	// Processed Fields
	NameStriped string
	Methods     []MethodView
}

type MethodView struct {
	// From PB
	MethodName  string
	InTypeName  string
	OutTypeName string
	Comment     string
	// Processed Fields
	MethodNameStriped string // removed Chat, Channel, Group, Direct from rpc name
	GoInTypeName      string
	GoOutTypeName     string
	Hash              uint32
	FullMethodName    string // RPC_Other.Echo
	ParentServiceName string // RPC_Other
	DartMethodName    string // camelCase
	Pos               int    // Seq number of rpc
}

////////// Messages /////////
type MessageView struct {
	// From PB
	Name    string
	Comment string
	// Processed Fields
	Fields []MessageFieldView
}

type MessageFieldView struct {
	// From PB
	FieldName string
	TypeName  string
	Repeated  bool
	TagNumber int
	Comment   string
	// Processed Fields
	isPrimitive bool // is ? numbers, bool, string, bytes or refrence to other custom types
	GoType      string
	GoFlatType  string
	JavaType    string
	RustType    string
}

////////// Enums /////////
type EnumView struct {
	// From PB
	Name    string
	Comment string
	// Processed Fields
	Fields []EnumFieldView
}

type EnumFieldView struct {
	// From PB
	FieldName string
	TagNumber int
	PosNumber int
	Comment   string
	// Processed Fields
}

///////////// QEvent //////////////////
// For transforming RPC_Chat to QEvent
type QEventService struct {
	ServiceName string
	Events      []QEvent
}

// For building event commands from *Param pb messages
type QEvent struct {
	EventName string
	TagNum    int
	Fields    []QEventPBFields
}

// For building Proto buffer from rpc *Param
type QEventPBFields struct {
	Name     string
	PBType   string
	Repeated bool
	TagNum   int
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

	QEvent []QEventService

	OutGoEnumsStr string
	OutGoRPCsStr  string
	OutJavaStr    string

	Dirs DirParam
}

type DirParam struct {
	ProtoDir       string
	RustOutDir     string
	RustProjectDir string
	ProtoOutDir    string
}
