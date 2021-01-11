package ant

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func makeQEventStruct(gen *GenOut) (qEventServices []QEventService) {
	for i := 0; i < len(gen.Services); i++ {
		ser := gen.Services[i]
		// only those rpc services that opt in via qevent_rpc comment
		if strings.Index(ser.Comment, "qevent_rpc") == -1 {
			continue
		}

		var qEventsArr []QEvent

		var eventTagNum int64 = 5
		for j := 0; j < len(ser.Methods); j++ {
			method := ser.Methods[j]

			// only those rpc methods that opt in via qevent comment
			if strings.Index(method.Comment, "qevent") == -1 {
				continue
			}

			// Extract qevent_id_{} number
			idStr := _qeventIdReg.FindString(method.Comment)
			if len(idStr) > 0 {
				var err error
				idStr := strings.Replace(idStr, "qevent_id_", "", 1)
				eventTagNum, err = strconv.ParseInt(idStr, 10, 32)
				noErr(err)
			}

			qevent := _buildQEvent(method, gen)
			qevent.TagNum = int(eventTagNum)

			qEventsArr = append(qEventsArr, qevent)
			eventTagNum += 1
		}

		if len(qEventsArr) > 0 {
			qeventSer := QEventService{
				ServiceName: ser.NameStriped,
				Events:      qEventsArr,
			}

			qEventServices = append(qEventServices, qeventSer)
		}
	}

	PrettyPrint(qEventServices)
	return
}

var _qeventIdReg = regexp.MustCompile(`qevent_id_(\d+)`)

func _buildQEvent(method MethodView, gen *GenOut) QEvent {

	qevent := QEvent{
		EventName: method.MethodNameStriped,
		TagNum:    method.Pos, // BEING OVERWRITEN
		Fields:    _buildQEventField(method.GoInTypeName, gen),
	}

	return qevent
}

var _msgViewMap = make(map[string]*MessageView)

func _buildQEventField(pramsMsgName string, gen *GenOut) (out []QEventPBFields) {
	// Fill the map
	if len(_msgViewMap) == 0 {
		for m := 0; m < len(gen.Messages); m++ {
			msg := gen.Messages[m]
			_msgViewMap[msg.Name] = &msg
		}
	}

	msg, found := _msgViewMap[pramsMsgName]
	if !found {
		panic(fmt.Sprintf("did not found pb param message: %s\n", pramsMsgName))
	}

	for g := 0; g < len(msg.Fields); g++ {
		msgFiled := msg.Fields[g]

		msgEventField := QEventPBFields{
			Name:     msgFiled.FieldName,
			PBType:   msgFiled.TypeName,
			Repeated: msgFiled.Repeated,
			TagNum:   msgFiled.TagNumber,
		}

		out = append(out, msgEventField)
	}
	return
}
