package ant

import (
	"fmt"
	"github.com/emicklei/proto"
	"os"
	"strings"
)

func ExtractAllServicesViews(protos []*proto.Proto) []ServiceView {
	serviceViews := make([]ServiceView, 0)

	for _, pto := range protos {
		for _, entry := range pto.Elements {
			if ser, ok := entry.(*proto.Service); ok {
				serView := ServiceView{
					Name:    ser.Name,
					Comment: findCommentV2(ser.Comment),
					Hash:    Hash32(ser.Name),
					Options: extractElementOptions(ser.Elements),
				}

				for _, m2 := range ser.Elements {
					if m, ok := m2.(*proto.RPC); ok {
						mv := MethodView{
							MethodName:        m.Name,
							InTypeName:        m.RequestType,
							GoInTypeName:      strings.Replace(m.RequestType, ".", "_", -1),
							OutTypeName:       m.ReturnsType,
							GoOutTypeName:     strings.Replace(m.ReturnsType, ".", "_", -1),
							Hash:              Hash32(m.Name),
							FullMethodName:    serView.Name + "." + m.Name,
							ParentServiceName: serView.Name,
						}
						serView.Methods = append(serView.Methods, mv)
					}
				}
				serviceViews = append(serviceViews, serView)
			}
		}
	}

	return serviceViews
}

func ExtractAllMessagesViews(protos []*proto.Proto) []MessageView {
	messageViews := make([]MessageView, 0)

	for _, pto := range protos {
		for _, ele := range pto.Elements {
			if msg, ok := ele.(*proto.Message); ok {
				msgView := MessageView{
					Name:    msg.Name,
					Comment: findCommentV2(msg.Comment),
					Options: extractElementOptions(msg.Elements),
				}

				for _, f2 := range msg.Elements {
					if f, ok := f2.(*proto.NormalField); ok {
						mv := FieldView{
							FieldName:     f.Name,
							TypeName:      f.Type,
							Repeated:      f.Repeated,
							TagNumber:     f.Sequence,
							GoType:        pbTypesToGoType(f.Type),
							GoFlatType:    pbTypesToGoFlatTypes(f.Type),
							JavaType:      pbTypesToJavaType(f.Type),
							Options:       protoOptionsToOptionsView(f.Options),
							RealmTypeName: pbToRealmName(pbTypesToJavaType(f.Type)),
						}
						msgView.Fields = append(msgView.Fields, mv)
					}
				}
				messageViews = append(messageViews, msgView)
			}
		}
	}

	return messageViews
}

func ExtractAllEnumsViews(protos []*proto.Proto) []EnumView {
	enumViews := make([]EnumView, 0)

	for _, pto := range protos {
		for _, ele := range pto.Elements {
			if enum, ok := ele.(*proto.Enum); ok {
				enumView := EnumView{
					Name:    enum.Name,
					Comment: findCommentV2(enum.Comment),
					Options: extractElementOptions(enum.Elements),
				}

				pos := 0
				for _, ele := range enum.Elements {
					if value, ok := ele.(*proto.EnumField); ok {
						fieldView := EnumFieldView{
							FieldName: value.Name,
							TagNumber: int(value.Integer),
							PosNumber: int(pos),
						}
						pos++
						enumView.Fields = append(enumView.Fields, fieldView)
					}
				}
				enumViews = append(enumViews, enumView)
			}
		}
	}

	return enumViews
}

func findCommentV2(com *proto.Comment) string {
	if com != nil && len(com.Lines) > 0 {
		return com.Lines[len(com.Lines)-1]
	}
	return ""
}

func extractElementOptions(element []proto.Visitee) (res []OptionsView) {
	for _, el := range element {
		if option, ok := el.(*proto.Option); ok {
			v := OptionsView{
				OptionName:  option.Name,
				OptionValue: option.Constant.Source,
			}
			res = append(res, v)
		}
	}
	return
}

func protoOptionsToOptionsView(options []*proto.Option) (res []OptionsView) {
	for _, option := range options {
		v := OptionsView{
			OptionName:  option.Name,
			OptionValue: option.Constant.Source,
		}
		res = append(res, v)
	}
	return
}

//////////////////////////////////////////////////////////////////////////////
func GetAllARealmMessageViews(msgs []MessageView) (res []MessageView) {
	for _, m := range msgs {
		for _, opt := range m.Options {
			if strings.ToLower(opt.OptionName) == REALM {
				res = append(res, m)
			}
		}
	}
	return
}

// pb meassages with  {realm} - GetAllARealmMessageViews() dosn't works with proto.exe it jus fails
func GetAllARealmMessageViews_FromComments(msgs []MessageView) (res []MessageView) {
	for _, m := range msgs {
		if strings.Contains(strings.ToLower(m.Comment), "{realm}") {
			res = append(res, m)
		}
	}
	return
}

func xxx() {

	reader, _ := os.Open(`./play/1.proto`)
	defer reader.Close()
	parser := proto.NewParser(reader)
	def, err := parser.Parse()
	noErr(err)
	messageViews := make([]MessageView, 0)
	for _, v := range def.Elements {
		if msg, ok := v.(*proto.Message); ok {
			fmt.Println(msg)
			msgView := MessageView{
				Name:    msg.Name,
				Comment: "ccc",
			}

			for _, f2 := range msg.Elements {
				if f, ok := f2.(*proto.NormalField); ok {
					mv := FieldView{
						FieldName:  f.Name,
						TypeName:   f.Type,
						Repeated:   f.Repeated,
						TagNumber:  f.Sequence,
						GoType:     pbTypesToGoType(f.Type),
						GoFlatType: pbTypesToGoFlatTypes(f.Type),
						JavaType:   pbTypesToJavaType(f.Type),
					}
					msgView.Fields = append(msgView.Fields, mv)
				}

			}
			messageViews = append(messageViews, msgView)
			//if msg, ok := v.(*proto.Message); ok {
			//	fmt.Println(msg)
			//}
		}
		//fmt.Println(v)
	}
	PertyPrint(messageViews)
	PertyPrint(def)

}
