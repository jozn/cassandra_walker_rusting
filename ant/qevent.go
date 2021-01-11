package ant

import (
	"fmt"
	"os"
)

func makeQvent(gen *GenOut) {
	var qventServices []QEventService

	for i := 0; i < len(gen.Services); i++ {
		ser := gen.Services[i]
		qeventSer := QEventService{
			ServiceName: ser.NameStriped,
			Events:      nil,
		}

		for j := 0; j < len(ser.Methods); j++ {
			method := ser.Methods[j]
			qevent:= QEvent{
				EventName: method.MethodNameStriped,
				TagNum:    method.Pos + 5,
				Fields:    nil,
			}

			// Add QEvents fileds for building message
			fondMsgParam := false
			for m := 0; m < len(gen.Messages); m++ {
				msg := gen.Messages[m]
				if msg.Name == method.GoInTypeName {
					fondMsgParam = true
					// Extract each pb message fields
					for g := 0; g < len(msg.Fields); g++ {
						msgFiled := msg.Fields[g]

						msgEventField := QEventPBFields{
							Name: msgFiled.FieldName    ,
							PBType:   msgFiled.TypeName,
							Repeated: msgFiled.Repeated,
						}

						qevent.Fields = append(qevent.Fields, msgEventField)

					}
					break // found the msg exit
				}
			}
			if !fondMsgParam {
				fmt.Printf("did not found pb param message: %s\n", method.GoInTypeName)
				os.Exit(1)
			}


			qeventSer.Events = append(qeventSer.Events, qevent)
		}
		qventServices = append(qventServices, qeventSer)
	}

	PrettyPrint(qventServices[2])
}
