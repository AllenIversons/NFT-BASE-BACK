// Code generated by "stringer -type=ErrCode --linecomment"; DO NOT EDIT.

package base

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Success-100]
	_ = x[OpenSqlError-1001]
	_ = x[ConnectSqlError-1002]
	_ = x[InsertError-1003]
	_ = x[QueryError-1004]
	_ = x[WrongLoginError-1005]
	_ = x[AccountExistError-1006]
	_ = x[PasswdUpdateError-1007]
	_ = x[AuthFailed-1101]
	_ = x[AuthFormatError-1102]
	_ = x[InvalidToken-1103]
	_ = x[GenTokenError-1104]
	_ = x[LackTokenError-1105]
	_ = x[UserTokenError-1105]
}

const (
	_ErrCode_name_0 = "Operation succeed"
	_ErrCode_name_1 = "Error open Mysql databaseCannot connect to mysql databaseInsert data errorQuery db errorWrong username or passwdAccount is already registeredUpdate password failed"
	_ErrCode_name_2 = "Permission denied, lack tokenThe auth format in the request header is incorrectThe token has expired or is invalid or could not parse with claimsGenerate token error: Sign Token FailedLack token in request header"
)

var (
	_ErrCode_index_1 = [...]uint8{0, 25, 57, 74, 88, 112, 141, 163}
	_ErrCode_index_2 = [...]uint8{0, 29, 79, 145, 184, 212}
)

func (i ErrCode) String() string {
	switch {
	case i == 100:
		return _ErrCode_name_0
	case 1001 <= i && i <= 1007:
		i -= 1001
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	case 1101 <= i && i <= 1105:
		i -= 1101
		return _ErrCode_name_2[_ErrCode_index_2[i]:_ErrCode_index_2[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
