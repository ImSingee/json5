// Code generated by "stringer -type TokenType -output strings.go"; DO NOT EDIT.

package json5

import "strconv"

const _TokenType_name = "TypeNoneTypeArrayBeginTypeArrayEndTypeObjectBeginTypeObjectEndTypeValueSepTypePairSepTypeStringTypeIntegerTypeFloatTypeFalseTypeTrueTypeNullTypeEOF"

var _TokenType_index = [...]uint8{0, 8, 22, 34, 49, 62, 74, 85, 95, 106, 115, 124, 132, 140, 147}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
