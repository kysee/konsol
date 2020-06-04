package types

import (
	"fmt"
)

type OptDesc struct {
	Name    string
	Usage   string
	Default interface{}
}

type CmdHandler func(*Args) error
type CmdSpec struct {
	Name    string
	Usage   string
	Opts    []*OptDesc
	Handler CmdHandler
}

func (cs *CmdSpec) StringIntent(intent string) string {
	r := fmt.Sprintf("%s%v\n", intent, cs)
	return r
}

func (cs *CmdSpec) HelpString() string {
	h := cs.Usage + "\n"
	for _, o := range cs.Opts {
		h += fmt.Sprintf("\t%s: %s\n", o.Name, o.Usage)
	}
	return h
}

type Args struct {
	vals map[string]interface{}
}

func NewArgs(_vals map[string]interface{}) *Args {
	return &Args{
		vals: _vals,
	}
}

func (args *Args) String() string {
	r := ""
	for k, v := range args.vals {

		//fmt.Println(*v.(*int), vt.Kind(), vt.String())
		switch t := v.(type) {
		case *int:
			r += fmt.Sprintf("\t%s = %v", k, *t)
		case *string:
			r += fmt.Sprintf("\t%s = %v", k, *t)
		case *bool:
			r += fmt.Sprintf("\t%s = %v", k, *t)
		}

	}
	return r
}

func (args *Args) Map() map[string]interface{} {
	return args.vals
}

func (args *Args) Set(n string, v interface{}) {
	args.vals[n] = v
}

func (args *Args) Get(n string) interface{} {
	v, ok := args.vals[n]
	if !ok {
		return nil
	}
	return v
}

func (args *Args) Int(n string) int {
	v := args.Get(n)
	if v == nil {
		return 0
	}
	return *(v.(*int))
}

func (args *Args) Str(n string) string {
	v := args.Get(n)
	if v == nil {
		return ""
	}
	return *(v.(*string))
}

func (args *Args) Bool(n string) bool {
	v := args.Get(n)
	if v == nil {
		return false
	}
	return *(v.(*bool))
}
