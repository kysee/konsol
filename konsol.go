package konsol

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/kysee/konsol/types"
	"os"
	"reflect"
	"strings"
)

type Konsol struct {
	cmdWorkers map[string]*cmdWorker
}

func NewKonsol(cmdspecs []*types.CmdSpec) *Konsol {
	kcon := &Konsol{
		cmdWorkers: make(map[string]*cmdWorker),
	}

	for _, cmdspec := range cmdspecs {
		flagSet := flag.NewFlagSet(cmdspec.Name, flag.ContinueOnError)
		args := make(map[string]interface{})

		for _, opt := range cmdspec.Opts {
			va := reflect.ValueOf(opt.Default)
			switch va.Kind() {
			case reflect.String:
				args[opt.Name] = flagSet.String(opt.Name, opt.Default.(string), opt.Usage)
			case reflect.Int:
				args[opt.Name] = flagSet.Int(opt.Name, opt.Default.(int), opt.Usage)
			case reflect.Bool:
				args[opt.Name] = flagSet.Bool(opt.Name, opt.Default.(bool), opt.Usage)
			default:
				panic("not supported type for option")
			}
		}

		kcon.cmdWorkers[cmdspec.Name] = &cmdWorker{
			flagSet: flagSet,
			cmdSpec: cmdspec,
			argsMap: args,
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

		csvReader := csv.NewReader(strings.NewReader(line))
		csvReader.Comma = ' '
		toks, err := csvReader.Read()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err := kcon.do(toks); err != nil {
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
	for _, w := range kcon.cmdWorkers {
		w.flagSet.Usage()
	}
}
