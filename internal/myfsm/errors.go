package myfsm

type MachineErrorCode string

const (
	ErrImpossibleTransition MachineErrorCode = "impossible_transition"
	ErrConditionIsFalse     MachineErrorCode = "ErrConditionIsFalse"
)

func (m MachineErrorCode) String() string {
	return string(m)
}

type MachineError struct {
	Code MachineErrorCode
}

func NewMachineError(code MachineErrorCode) *MachineError {
	return &MachineError{
		Code: code,
	}
}

func (m MachineError) Error() string {
	return m.Code.String()
}
