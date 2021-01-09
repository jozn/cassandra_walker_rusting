package ant

import (
	"strings"
)

func processAllMessagesViews(pbMsgs []PBMessage) []MessageView {
	messageViews := make([]MessageView, 0)

	for _, pbMsg := range pbMsgs {
		var msgFields []MessageFieldView

		for _, pbField := range pbMsg.PBFields {
			fieldView := MessageFieldView{
				PBMessageField: pbField,
				GoType:         pbTypesToGoType(pbField.TypeName),
				isPrimitive:    pbTypesIsPrimitive(pbField.TypeName),
				GoFlatType:     pbTypesToGoFlatTypes(pbField.TypeName),
				JavaType:       pbTypesToJavaType(pbField.TypeName),
				RustType:       pbTypesToRustType(pbField.TypeName),
			}
			msgFields = append(msgFields, fieldView)
		}

		msgView := MessageView{
			PBMessage: pbMsg,
			Fields:    msgFields,
		}

		messageViews = append(messageViews, msgView)
	}

	return messageViews
}

func processAllServicesViews(pbMsgs []PBService) []ServiceView {
	messageViews := make([]ServiceView, 0)

	for _, pbRpcService := range pbMsgs {
		var serviceRpcs []MethodView

		for _, rpc := range pbRpcService.PBMethods {
			fieldView := MethodView{
				PBMethod:          rpc,
				GoInTypeName:      strings.Replace(rpc.InTypeName, ".", "_", -1),  // For nested messages replace . with _
				GoOutTypeName:     strings.Replace(rpc.OutTypeName, ".", "_", -1), // For nested messages replace . with _
				Hash:              uniqueMethodHash(rpc.MethodName),
				FullMethodName:    pbRpcService.Name + "." + rpc.MethodName,
				ParentServiceName: rpc.MethodName,
				DartMethodName:    strings.ToLower(rpc.MethodName[0:1]) + rpc.MethodName[1:],
			}
			serviceRpcs = append(serviceRpcs, fieldView)
		}

		msgView := ServiceView{
			PBService: pbRpcService,
			Methods:   serviceRpcs,
		}

		messageViews = append(messageViews, msgView)
	}

	return messageViews
}

func processAllEnumsViews(pbEnums []PBEnum) (out []EnumView) {
	for _, pbEnum := range pbEnums {
		var enumFields []EnumFieldView

		for _, pbEnum := range pbEnum.PBFields {
			fieldView := EnumFieldView{
				PBEnumField: pbEnum,
			}
			enumFields = append(enumFields, fieldView)
		}

		enumView := EnumView{
			PBEnum: pbEnum,
			Fields: enumFields,
		}

		out = append(out, enumView)
	}

	return out
}
