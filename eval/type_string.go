// Code generated by "stringer -type Type -linecomment"; DO NOT EDIT.

package eval

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[String-1]
	_ = x[Int-2]
	_ = x[Bool-3]
	_ = x[Array-4]
	_ = x[Hash-5]
	_ = x[File-6]
	_ = x[Request-7]
	_ = x[Response-8]
	_ = x[Stream-9]
	_ = x[Name-10]
	_ = x[Key-11]
	_ = x[Yield-12]
}

const _Type_name = "stringintboolarrayhashfilerequestresponsestreamnamekeyyield"

var _Type_index = [...]uint8{0, 6, 9, 13, 18, 22, 26, 33, 41, 47, 51, 54, 59}

func (i Type) String() string {
	i -= 1
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
