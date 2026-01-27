package errcode

type Error interface {
	error
	Code() int
}

func New(code int, err error) *Err {
	return &Err{
		code: code,
		err:  err,
	}
}

type Err struct {
	code int
	err  error
}

func (e *Err) Error() string {
	if e == nil {
		return ""
	}
	if e.err != nil {
		return e.err.Error()
	}
	return "unknown error"
}

func (e *Err) Code() int {
	if e != nil {
		return e.code
	}
	return 0
}

func (e *Err) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}
