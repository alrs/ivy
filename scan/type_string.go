// generated by stringer -type Type; DO NOT EDIT

package scan

import "fmt"

const _Type_name = "EOFErrorNewlineAssignCharGreaterOrEqualIdentifierLeftBrackLeftParenNumberOperatorOpRationalRightBrackRightParenSemicolonSpaceString"

var _Type_index = [...]uint8{0, 3, 8, 15, 21, 25, 39, 49, 58, 67, 73, 81, 83, 91, 101, 111, 120, 125, 131}

func (i Type) String() string {
	if i < 0 || i+1 >= Type(len(_Type_index)) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
