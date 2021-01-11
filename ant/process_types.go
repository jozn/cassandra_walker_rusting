package ant

import (
	"regexp"
	"strings"
)

func processAllMessagesViews(pbMsgs []PBMessage) []MessageView {
	messageViews := make([]MessageView, 0)

	for _, pbMsg := range pbMsgs {
		var msgFields []MessageFieldView

		for _, pbField := range pbMsg.PBFields {
			fieldView := MessageFieldView{
				FieldName: pbField.FieldName,
				TypeName:  pbField.TypeName,
				Repeated:  pbField.Repeated,
				TagNumber: pbField.TagNumber,
				Comment:   pbField.Comment,
				// Processed
				isPrimitive: pbTypesIsPrimitive(pbField.TypeName),
				GoType:      pbTypesToGoType(pbField.TypeName),
				GoFlatType:  pbTypesToGoFlatTypes(pbField.TypeName),
				JavaType:    pbTypesToJavaType(pbField.TypeName),
				RustType:    pbTypesToRustType(pbField.TypeName),
			}
			msgFields = append(msgFields, fieldView)
		}

		msgView := MessageView{
			Name:    pbMsg.Name,
			Comment: pbMsg.Comment,
			Fields:  msgFields,
		}

		messageViews = append(messageViews, msgView)
	}

	return messageViews
}

func processAllServicesViews(pbServices []PBService) []ServiceView {
	messageViews := make([]ServiceView, 0)

	for _, pbRpcService := range pbServices {
		var serviceRpcs []MethodView
		var rpcServiceStripedName = strings.Replace(pbRpcService.Name, "RPC_", "", 1) // RPC_Chat > Chat

		for i, rpc := range pbRpcService.PBMethods {
			//if strings.rpc.MethodName
			fieldView := MethodView{
				MethodName:  rpc.MethodName,
				InTypeName:  rpc.InTypeName,
				OutTypeName: rpc.OutTypeName,
				Comment:     rpc.Comment,
				// Processed
				MethodNameStriped: strings.TrimPrefix(rpc.MethodName, rpcServiceStripedName), //  Every rpc prefix is the sample as rpc_service suffix
				GoInTypeName:      strings.Replace(rpc.InTypeName, ".", "_", -1),             // For nested messages replace . with _
				GoOutTypeName:     strings.Replace(rpc.OutTypeName, ".", "_", -1),            // For nested messages replace . with _
				Hash:              uniqueMethodHash(rpc.MethodName),
				FullMethodName:    pbRpcService.Name + "." + rpc.MethodName,
				ParentServiceName: rpc.MethodName,
				DartMethodName:    strings.ToLower(rpc.MethodName[0:1]) + rpc.MethodName[1:],
				Pos:               i + 1,
			}
			serviceRpcs = append(serviceRpcs, fieldView)
		}

		msgView := ServiceView{
			Name:        pbRpcService.Name,
			Comment:     pbRpcService.Comment,
			NameStriped: rpcServiceStripedName,
			Methods:     serviceRpcs,
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
				FieldName: pbEnum.FieldName,
				TagNumber: pbEnum.TagNumber,
				PosNumber: pbEnum.PosNumber,
				Comment:   pbEnum.Comment,
			}
			enumFields = append(enumFields, fieldView)
		}

		enumView := EnumView{
			Name:    pbEnum.Name,
			Comment: pbEnum.Comment,
			Fields:  enumFields,
		}

		out = append(out, enumView)
	}

	return out
}

// Old way of rpc prefix removing
var _rpcMethodPrefixRemover = regexp.MustCompile(`^(Chat|Group|Direct|Channel|Store)`)

func _stripRpcMethodName(rpcName string) string {
	out := _rpcMethodPrefixRemover.ReplaceAllString(rpcName, "")
	return out
}
