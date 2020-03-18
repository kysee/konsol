package main

import "github.com/kysee/kconsole"

func main() {
	cmdline := []*kconsole.CmdLine{
		{"cmd1", []*kconsole.CmdOpt{{"o11", "option11", 1},{"o12", "option12", 1}}, nil},
		{"cmd2", []*kconsole.CmdOpt{{"o21", "option21", 1},{"o22", "option22", 1}}, nil},
		{"cmd3", []*kconsole.CmdOpt{{"o31", "option31", 1},{"o32", "option32", 1}}, nil},
	}

	kcon := kconsole.NewKConsole(cmdline)
	kcon.Start()
}
