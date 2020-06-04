package main

import (
	"fmt"
	"github.com/kysee/konsol"
	"github.com/kysee/konsol/types"
)

func main() {
	cmdspecs := []*types.CmdSpec{
		{"cmd1", "this is cmd1", []*types.OptDesc{{"o11", "option11", 12}, {"o12", "option12", 21}}, DoCmn1},
		{"cmd2", "this is cmd2", []*types.OptDesc{{"o21", "option21", 12}, {"o22", "option22", 21}}, DoCmn2},
		{"cmd3", "this is cmd3", []*types.OptDesc{{"o31", "option31", 12}, {"o32", "option32", false}}, DoCmn3},
	}

	kcon := konsol.NewKonsol(cmdspecs)
	kcon.Start("konsol # ")
}

func DoCmn1(args *types.Args) error {
	fmt.Printf("run cmd1: %v\n", args)
	return nil
}

func DoCmn2(args *types.Args) error {
	fmt.Printf("run cmd1: %v\n", args)
	return nil
}

func DoCmn3(args *types.Args) error {
	fmt.Printf("run cmd1: %v\n", args)
	return nil
}
