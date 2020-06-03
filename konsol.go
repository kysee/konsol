package konsol

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/kysee/konsol/types"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type cmdWorker struct {
	optSet *flag.FlagSet
	//valMap map[string]reflect.Value
	intVals map[string]*int
	strVals map[string]*string

	cmdSpec *types.CmdSpec
}

func (cw *cmdWorker) defaultCmdFunc() {
	for k, v := range cw.intVals {
		fmt.Printf("%s = %d (int)\n", k, *v)
	}
	for k, v := range cw.strVals {
		fmt.Printf("%s = %d (string)\n", k, *v)
	}
}

func (cw *cmdWorker) do(args []string) error {
	if err := cw.optSet.Parse(args); err != nil {
		return err
	}
	if cw.cmdSpec.CmdFunc != nil {
		cw.cmdSpec.CmdFunc(nil, nil)
	} else {
		cw.defaultCmdFunc()
	}
	cw.resetOpts()
	return nil
}

func (cw *cmdWorker) resetOpts() {
	for k, v := range cw.intVals {
		f := cw.optSet.Lookup(k)
		i, err := strconv.Atoi(f.DefValue)
		if err != nil {
			panic("fail reset params")
		}
		*v = i
	}
	for k, v := range cw.strVals {
		f := cw.optSet.Lookup(k)
		*v = f.DefValue
	}
}

type Konsol struct {
	cmdWorkers map[string]*cmdWorker
}

func NewKonsol(cmdspecs []types.CmdSpec) *Konsol {
	kcon := &Konsol{
		cmdWorkers: make(map[string]*cmdWorker),
	}

	for _, cmdspec := range cmdspecs {
		optSet := flag.NewFlagSet(cmdspec.Name, flag.ContinueOnError)

		intVals := make(map[string]*int)
		strVals := make(map[string]*string)

		for _, opt := range cmdspec.Opts {
			va := reflect.ValueOf(opt.Defv)
			switch va.Kind() {
			case reflect.String:
				strVals[opt.Name] = new(string)
				optSet.StringVar(strVals[opt.Name], opt.Name, opt.Defv.(string), opt.Desc)
			case reflect.Int:
				intVals[opt.Name] = new(int)
				optSet.IntVar(intVals[opt.Name], opt.Name, opt.Defv.(int), opt.Desc)
			default:
				panic("not supported type for option")
			}
		}

		kcon.cmdWorkers[cmdspec.Name] = &cmdWorker{
			optSet:  optSet,
			intVals: intVals,
			strVals: strVals,
			cmdSpec: &types.CmdSpec {
				Name: cmdspec.Name,
				Opts: cmdspec.Opts,
				CmdFunc: cmdspec.CmdFunc,
			},
		}
	}

	return kcon
}

func (kcon *Konsol) Start(name string) {
	reader := bufio.NewReader(os.Stdin)

	if name == "" {
		name = "konsol # "
	}
	for {
		fmt.Printf(name)

		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)
		if len(line) <= 0 {
			continue
		}

		if kcon.procInternalCmds(line) {
			continue
		}

		if err := kcon.do(strings.Split(line, " ")); err != nil {
			kcon.usage()
		}
	}
}

func (kcon *Konsol) do(args []string) error {
	if cmd, ok := kcon.cmdWorkers[args[0]]; ok {
		return cmd.do(args[1:])
	}
	return fmt.Errorf("not supported command: %v", args)
}

func (kcon *Konsol) procInternalCmds(line string) bool {
	switch line {
	case "exit", "quit", "q":
		os.Exit(0)
	case "help", "h":
		kcon.usage()
		return true
	}
	return false
}

func (kcon *Konsol) usage() {
	r := "Command Usage is ...\n"
	for _, w := range kcon.cmdWorkers {
		r += w.cmdSpec.StringIntent("  ")
	}
	fmt.Println(r)
}
