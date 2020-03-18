package kconsole

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type CmdFunc func() error

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

type CmdLine struct {
	Name string
	Opts []*CmdOpt
	CmdFunc CmdFunc
}

type cmdWorker struct {
	optSet *flag.FlagSet
	//valMap map[string]reflect.Value
	intVals map[string]*int
	strVals map[string]*string

	cmdFunc CmdFunc
}

func (cw *cmdWorker) do(args []string) {
	if err := cw.optSet.Parse(args); err != nil {
		return
	}
	if cw.cmdFunc != nil {
		cw.cmdFunc()
	} else {
		cw.defaultCmdFunc()
	}
}

func (cw *cmdWorker) defaultCmdFunc() {
	for k, v := range cw.intVals {
		fmt.Printf("%s = %d (int)\n", k, *v)
	}
	for k, v := range cw.strVals {
		fmt.Printf("%s = %d (string)\n", k, *v)
	}
}

func (cmd *cmdWorker) clear() {

}

type KConsole struct {
	cmdWorkers map[string]*cmdWorker
}

func NewKConsole(cmdsets []*CmdLine) *KConsole {
	kcon := &KConsole{
		cmdWorkers: make(map[string]*cmdWorker),
	}

	for _, cmd := range cmdsets {
		optSet := flag.NewFlagSet(cmd.Name, flag.ContinueOnError)

		intVals := make(map[string]*int)
		strVals := make(map[string]*string)

		for _, opt := range cmd.Opts {
			va := reflect.ValueOf(opt.Defv)
			switch va.Kind() {
			case reflect.String:
				strVals[opt.Name] = new(string)
				optSet.StringVar(strVals[opt.Name], opt.Name, opt.Defv.(string), opt.Desc)
			case reflect.Int:
				intVals[opt.Name] = new(int)
				optSet.IntVar(intVals[opt.Name], opt.Name, opt.Defv.(int), opt.Desc)
			default:
				panic("not supported type")
			}
		}

		kcon.cmdWorkers[cmd.Name] = &cmdWorker{
			optSet: optSet,
			intVals: intVals,
			strVals: strVals,
			cmdFunc: cmd.CmdFunc,
		}
	}

	return kcon
}

func (kcon *KConsole) Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("kconsole # ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSpace(line)

		if kcon.isExit(line) {
			break
		}

		if len(line) > 0 {
			kcon.do(strings.Split(line, " "))
		}
	}
}

func (kcon *KConsole) do(args []string) {
	if cmd, ok := kcon.cmdWorkers[args[0]]; ok {
		cmd.do(args[1:])
	}
}
func (kcon *KConsole) isExit(line string) bool {
	switch line {
	case "exit", "quit", "q":
		return true
	}
	return false
}
