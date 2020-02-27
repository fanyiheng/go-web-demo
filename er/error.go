package er

import "fmt"

type ErrUsr struct {
	Source error  `json:"source"`
	Msg    string `json:"msg"`
	Log    bool   `json:"log"`
}

func Msgf(f string, msg ...interface{}) *ErrUsr {
	return &ErrUsr{Msg: fmt.Sprintf(f, msg...)}
}

func Msg(msg string) *ErrUsr {
	return &ErrUsr{Msg: msg}
}

func (err *ErrUsr) Lg() *ErrUsr {
	err.Log = true
	return err
}
func (err *ErrUsr) Src(source error) *ErrUsr {
	err.Source = source
	err.Log = true
	return err
}

func (err *ErrUsr) Error() string {
	return err.String()
}

func (err *ErrUsr) String() string {
	return fmt.Sprintf("%s -> %v", err.Msg, err.Source)
}
