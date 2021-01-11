package ant

import (
	"github.com/emicklei/proto"
)

func ExtractAllPBMessages(protos []*proto.Proto) []PBMessage {
	messageViews := make([]PBMessage, 0)

	for _, pto := range protos {
		for _, pbElement := range pto.Elements {
			if pbMsg, ok := pbElement.(*proto.Message); ok {
				msgView := PBMessage{
					Name:      pbMsg.Name,
					Comment:   extractCommentV2(pbMsg.Comment),
					PBOptions: extractElementOptions(pbMsg.Elements),
					PBFields:  nil, // setting from below code
				}

				for _, pbEle := range pbMsg.Elements {
					if field, ok := pbEle.(*proto.NormalField); ok {
						fieldView := PBMessageField{
							FieldName: field.Name,
							TypeName:  field.Type,
							Repeated:  field.Repeated,
							TagNumber: field.Sequence,
							Comment:   extractCommentV2(field.InlineComment),
							PBOptions: extractElementOptionsFromOptions(field.Options),
						}
						msgView.PBFields = append(msgView.PBFields, fieldView)
					}
				}
				messageViews = append(messageViews, msgView)
			}
		}
	}

	return messageViews
}

func ExtractAllPBServices(protos []*proto.Proto) []PBService {
	serviceViews := make([]PBService, 0)

	for _, pto := range protos {
		for _, entry := range pto.Elements {

			// Each rpc server holders
			if pbService, ok := entry.(*proto.Service); ok {
				serView := PBService{
					Name:      pbService.Name,
					Comment:   extractCommentV2(pbService.Comment),
					PBOptions: extractElementOptions(pbService.Elements),
					PBMethods: nil, // setting from below code
				}

				// Each rpc fun
				for _, element := range pbService.Elements {
					if m, ok := element.(*proto.RPC); ok {
						//PrettyPrint(element)
						mv := PBMethod{
							MethodName:  m.Name,
							InTypeName:  m.RequestType,
							OutTypeName: m.ReturnsType,
							Comment:     extractCommentV2(m.InlineComment),
							PBOptions:   extractElementOptions(m.Elements),
						}
						serView.PBMethods = append(serView.PBMethods, mv)
					}
				}
				serviceViews = append(serviceViews, serView)
			}
		}
	}

	return serviceViews
}

func ExtractAllPBEnums(protos []*proto.Proto) []PBEnum {
	enumViews := make([]PBEnum, 0)

	for _, pto := range protos {
		for _, pbElement := range pto.Elements {
			if enum, ok := pbElement.(*proto.Enum); ok {
				enumView := PBEnum{
					Name:      enum.Name,
					Comment:   extractCommentV2(enum.Comment),
					PBOptions: extractElementOptions(enum.Elements),
					PBFields:  nil, // setting from below code
				}

				pos := 0
				for _, pbEle2 := range enum.Elements {
					if value, ok := pbEle2.(*proto.EnumField); ok {
						fieldView := PBEnumField{
							FieldName: value.Name,
							TagNumber: int(value.Integer),
							PosNumber: int(pos),
							Comment:   extractCommentV2(value.InlineComment),
							PBOptions: extractElementOptions(value.Elements),
						}
						pos++
						enumView.PBFields = append(enumView.PBFields, fieldView)
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
	//PrettyPrint(com)
	if com != nil && len(com.Lines) > 0 {
		return com.Lines[len(com.Lines)-1]
	}
	return ""
}

// Extract options for message, enums, rpc
func extractElementOptions(element []proto.Visitee) (res []PBOptions) {
	//PrettyPrint(element)
	for _, el := range element {
		if option, ok := el.(*proto.Option); ok {
			//PrettyPrint(option)
			v := PBOptions{
				OptionName:  option.Name,
				OptionValue: option.Constant.Source,
			}
			res = append(res, v)
		}
	}
	return
}

// Extract options for message fields
func extractElementOptionsFromOptions(options []*proto.Option) (res []PBOptions) {
	for _, option := range options {
		v := PBOptions{
			OptionName:  option.Name,
			OptionValue: option.Constant.Source,
		}
		res = append(res, v)
	}
	return
}
