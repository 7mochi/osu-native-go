package osunative

type ErrorCode int32

const (
	EndOfEnumeration   ErrorCode = -2
	BufferSizeQuery    ErrorCode = -1
	Success            ErrorCode = 0
	ObjectNotResolved  ErrorCode = 1
	RulesetUnavailable ErrorCode = 2
	UnexpectedRuleset  ErrorCode = 3
	Failure            ErrorCode = 127
)

func (e ErrorCode) Error() string {
	switch e {
	case EndOfEnumeration:
		return "end of enumeration"
	case BufferSizeQuery:
		return "buffer size query"
	case Success:
		return "success"
	case ObjectNotResolved:
		return "object not resolved"
	case RulesetUnavailable:
		return "ruleset unavailable"
	case UnexpectedRuleset:
		return "unexpected ruleset"
	case Failure:
		return "failure"
	default:
		return "unknown error code"
	}
}

func (e ErrorCode) IsSuccess() bool {
	return e == Success
}
