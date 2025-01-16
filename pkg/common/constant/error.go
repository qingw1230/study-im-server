package constant

type ErrInfo struct {
	ErrCode int
	ErrMsg  string
}

func (e *ErrInfo) Error() string {
	return e.ErrMsg
}

func (e *ErrInfo) Code() int {
	return e.ErrCode
}

var (
	ErrParseToken       = ErrInfo{TokenError, ParseTokenMsg.Error()}
	ErrTokenExpired     = ErrInfo{TokenExpired, TokenExpiredMsg.Error()}
	ErrTokenInvalid     = ErrInfo{TokenInvalid, TokenInvalidMsg.Error()}
	ErrTokenMalformed   = ErrInfo{TokenMalformed, TokenMalformedMsg.Error()}
	ErrTokenNotValidYet = ErrInfo{TokenNotValidYet, TokenNotValidYetMsg.Error()}
	ErrTokenUnknown     = ErrInfo{TokenUnknown, TokenUnknownMsg.Error()}
)
