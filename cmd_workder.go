package konsol

import (
	"flag"
	"fmt"
	"github.com/kysee/konsol/types"
	"reflect"
)

type cmdWorker struct {
	flagSet *flag.FlagSet
	cmdSpec *types.CmdSpec
	argsMap map[string]interface{}
}

func (cw *cmdWorker) defaultCmdFunc(args *types.Args) {
	for k, v := range args.Map() {
		vv := reflect.ValueOf(v)
		fmt.Printf("%s = %v (int)\n", k, vv)
	}
}

func (cw *cmdWorker) do(args []string) error {
	if err := cw.flagSet.Parse(args); err != nil {
		return err
	}
	if cw.cmdSpec.Handler != nil {
		cw.cmdSpec.Handler(types.NewArgs(cw.argsMap))
	} else {
		cw.defaultCmdFunc(types.NewArgs(cw.argsMap))
	}
	cw.resetOpts()
	return nil
}

func (cw *cmdWorker) resetOpts() {
	for _, opt := range cw.cmdSpec.Opts {
		va := reflect.ValueOf(opt.Default)
		switch va.Kind() {
		case reflect.String:
			*(cw.argsMap[opt.Name].(*string)) = opt.Default.(string)
		case reflect.Int:
			*(cw.argsMap[opt.Name].(*int)) = opt.Default.(int)
		case reflect.Bool:
			*(cw.argsMap[opt.Name].(*bool)) = opt.Default.(bool)
		default:
			panic("not supported type for option")
		}
	}
}
