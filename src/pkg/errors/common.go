package errors

import "fmt"

type ErrorCode int

const (
	BaseErrorCodeUnknown ErrorCode = iota*1000 + 1000
	BaseErrorCodeGateway
	BaseErrorCodeUser
	BaseErrorCodeWallet
	BaseErrorCodeToken
	BaseErrorCodeOrder
	BaseErrorCodePnl
	BaseErrorCodeTradingPair
	BaseErrorCodeCandleStick
)

func (e ErrorCode) GetRangeName() string {
	switch {
	case e >= BaseErrorCodeUnknown && e < BaseErrorCodeGateway:
		return "Unknown"
	case e >= BaseErrorCodeGateway && e < BaseErrorCodeUser:
		return "Gateway"
	case e >= BaseErrorCodeUser && e < BaseErrorCodeWallet:
		return "User"
	case e >= BaseErrorCodeWallet && e < BaseErrorCodeToken:
		return "Wallet"
	case e >= BaseErrorCodeToken && e < BaseErrorCodeOrder:
		return "Token"
	case e >= BaseErrorCodeOrder && e < BaseErrorCodePnl:
		return "Order"
	case e >= BaseErrorCodePnl && e < BaseErrorCodeTradingPair:
		return "Pnl"
	case e >= BaseErrorCodeTradingPair && e < BaseErrorCodeCandleStick:
		return "TradingPair"
	case e >= BaseErrorCodeCandleStick:
		return "CandleStick"
	default:
		return "Unknown"
	}
}

type AppError struct {
	Code    ErrorCode
	Message string
}

func (a *AppError) Error() string {
	return fmt.Sprintf("App error with code: %d, Message: %s", a.Code, a.Message)
}

// Equal compare two error
func (a *AppError) Equal(err error) bool {
	if err == nil {
		return false
	}

	if a.Code != err.(*AppError).Code {
		return false
	}

	if a.Message != err.(*AppError).Message {
		return false
	}

	return true
}

func NewAppError(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: fmt.Sprintf("[Module %s]: %s", code.GetRangeName(), message),
	}
}
