package types

import "fmt"

type CmdFunc func(map[string]int, map[string]string) error

type CmdOpt struct {
	Name string
	Desc string
	Defv interface{}
}

func NewOption(name, desc string, defval interface{}) CmdOpt {
	return CmdOpt{
		Name: name,
		Desc: desc,
		Defv: defval,
	}
}

func (co *CmdOpt) String() string {
	return fmt.Sprintf("-%s\t\t%s. (default:%v)", co.Name, co.Desc, co.Defv)
}

type CmdSpec struct {
	Name    string
	Opts    []*CmdOpt
	CmdFunc CmdFunc
}

func (cs *CmdSpec) StringIntent(intent string) string {
	r := fmt.Sprintf("%s%s\n", intent, cs.Name)
	for _, opt := range cs.Opts {
		r += fmt.Sprintf("%s\t%v\n", intent, opt)
	}
	return r
}
