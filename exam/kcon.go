package main

import (
	"github.com/kysee/konsol"
	"github.com/kysee/konsol/types"
)

func main() {
	cmdspecs := []types.CmdSpec{
		{"cmd1", []*types.CmdOpt{{"o11", "option11", 1}, {"o12", "option12", 1}}, nil},
		{"cmd2", []*types.CmdOpt{{"o21", "option21", 1}, {"o22", "option22", 1}}, nil},
		{"cmd3", []*types.CmdOpt{{"o31", "option31", 1}, {"o32", "option32", 1}}, nil},
	}

	kcon := konsol.NewKonsol(cmdspecs)
	kcon.Start("konsol # ")
}

func DoCmn1(intArgs map[string]int, strArgs map[string]string) error {
	return nil
}

func DoCmn2(intArgs map[string]int, strArgs map[string]string) error {

	return nil
}

func DoCmn3(intArgs map[string]int, strArgs map[string]string) error {

	return nil
}
