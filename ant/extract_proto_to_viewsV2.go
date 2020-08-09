package ant

import (
	"github.com/emicklei/proto"
	"strings"
)

func ExtractAllServicesViews(protos []*proto.Proto) []ServiceView {
	serviceViews := make([]ServiceView, 0)

	for _, pto := range protos {
		for _, entry := range pto.Elements {

			// Each rpc server holders
			if pbService, ok := entry.(*proto.Service); ok {
				serView := ServiceView{
					Name:    pbService.Name,
					Comment: extractCommentV2(pbService.Comment),
					Hash:    StrToInt32Hash(pbService.Name),
					Options: extractElementOptions(pbService.Elements),
				}

				// Each rpc fun
				for _, element:= range pbService.Elements {
					if m, ok := element.(*proto.RPC); ok {
						mv := MethodView{
							MethodName:        m.Name,
							InTypeName:        m.RequestType,
							GoInTypeName:      strings.Replace(m.RequestType, ".", "_", -1), // For nested messages replace . with _
							OutTypeName:       m.ReturnsType,
							GoOutTypeName:     strings.Replace(m.ReturnsType, ".", "_", -1), // For nested messages replace . with _
							Hash:              StrToInt32Hash(m.Name),
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
		for _, pbElement := range pto.Elements {
			if pbMsg, ok := pbElement.(*proto.Message); ok {
				msgView := MessageView{
					Name:    pbMsg.Name,
					Comment: extractCommentV2(pbMsg.Comment),
					Options: extractElementOptions(pbMsg.Elements),
				}

				for _, pbEle := range pbMsg.Elements {
					if field, ok := pbEle.(*proto.NormalField); ok {
						fieldView := FieldView{
							FieldName:     field.Name,
							TypeName:      field.Type,
							Repeated:      field.Repeated,
							TagNumber:     field.Sequence,
							GoType:        pbTypesToGoType(field.Type),
							isPrimitive:   pbTypesIsPrimitive(field.Type),
							GoFlatType:    pbTypesToGoFlatTypes(field.Type),
							JavaType:      pbTypesToJavaType(field.Type),
							RustType:      pbTypesToRustType(field.Type),
							Options:       protoOptionsToOptionsView(field.Options),
							RealmTypeName: pbToRealmName(pbTypesToJavaType(field.Type)),
						}
						msgView.Fields = append(msgView.Fields, fieldView)
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
		for _, pbElement := range pto.Elements {
			if enum, ok := pbElement.(*proto.Enum); ok {
				enumView := EnumView{
					Name:    enum.Name,
					Comment: extractCommentV2(enum.Comment),
					Options: extractElementOptions(enum.Elements),
				}

				pos := 0
				for _, pbEle2 := range enum.Elements {
					if value, ok := pbEle2.(*proto.EnumField); ok {
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

// Extracts last comment line if exists
func extractCommentV2(com *proto.Comment) string {
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
