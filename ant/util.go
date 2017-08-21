package ant

import (
	"crypto/md5"
	"encoding/binary"
	"log"
	"ms/sun/helper"
	"strings"
)

func noErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

/////////////// types converters /////////

func pbTypesToGoType(tp string) string {
	s := tp
	switch tp {
	case "int64", "sint64":
		s = "int64"
	case "double":
		s = "float64"
	case "float":
		s = "float32"
	case "int32", "sint32":
		s = "int32"
	case "uint32":
		s = "uint32"
	case "uint64":
		s = "uint64"
	case "fixed32":
		s = "uint32"
	case "fixed64":
		s = "uint64"
	case "sfixed32":
		s = "int32"
	case "sfixed64":
		s = "int64"

	case "bool":
		s = "bool"
	case "string":
		s = "string"
	case "bytes":
		s = "[]byte"
	}
	return s
}

func pbTypesToGoFlatTypes(tp string) string {
	s := tp
	switch tp {
	case "int64", "sint64", "int32",
		"sint32", "uint32", "uint64", "fixed32",
		"fixed64", "sfixed32", "sfixed64":
		s = "int"
	case "double":
		s = "float64"
	case "float":
		s = "float32"

	case "bool":
		s = "bool"
	case "string":
		s = "string"
	case "bytes":
		s = "[]byte"
	}
	return s
}

func pbTypesToJavaType(tp string) string {
	s := tp
	switch tp {
	case "int32", "sint32",
		"uint32", "fixed32",
		"sfixed32":
		s = "int"
	case "int64", "sint64",
		"uint64", "fixed64",
		"sfixed64":
		s = "int"
	case "double":
		s = "float64"
	case "float":
		s = "float32"

	case "bool":
		s = "bool"
	case "string":
		s = "string"
	case "bytes":
		s = "byte[]"
	}
	return s
}

/////////////// Hashes /////////////

func Hash(str string) int {
	sh1 := md5.Sum([]byte(str))
	b := sh1[0]
	//fmt.Println(b)
	b = b >> 1
	//fmt.Println(b)
	bytes := []byte{b, sh1[1], sh1[2], sh1[3]}

	res := binary.BigEndian.Uint32(bytes)
	//fmt.Println(res, int32(res))
	return int(res)
}

func Hash32(string string) int32 {
	return int32(Hash(string) % 10000)
}

func MyHash(string string) int {
	h := 15485862
	a := 7

	for i := 0; i < len(string); i++ {
		h = ((h * a) + int(int8(string[0]))) / 3
	}
	return h
}

func MyHash2(string string) int {
	s := strings.Split(string, ".")
	rs := helper.IntToStr((Hash(s[0])%1000)+1) + "00" + helper.IntToStr(Hash(s[1])%100000)
	//fmt.Println("+++++++"+rs)
	return helper.StrToInt(rs, -1)
}

////////////////////////////////////
