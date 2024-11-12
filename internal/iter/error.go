package iter

type ErrMatrix struct {
	msg string
}

func NewErrMatrix(msg string) error {
	return ErrMatrix{
		msg: msg,
	}
}

func (e ErrMatrix) Error() string {
	return e.msg
}

type ErrRHS struct {
	msg string
}

func NewErrRHS(msg string) error {
	return ErrRHS{
		msg: msg,
	}
}

func (e ErrRHS) Error() string {
	return e.msg
}

type ErrMaxIters struct{}

func NewErrMaxIters() error {
	return ErrMaxIters{}
}

func (e ErrMaxIters) Error() string {
	return "maximum number of iterations reached"
}
