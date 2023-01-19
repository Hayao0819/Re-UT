package main

import (
	"os"
	"fmt"
	//"sync"
	"strings"
	common "github.com/Hayao0819/Re-UT/common"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: convert_dic <path of id.def> <skkdict path>")
		os.Exit(1)
	}

	iddefPath := args[1]
	skkdictPath := args[2]

	err := convert_dic(iddefPath, skkdictPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert dic")
		os.Exit(1)
	}
}

func convert_dic(iddefPath, skkdictPath string) (error) {
	
	iddef, iddef_err := common.TextFileToArray(iddefPath)
	skkdict, skkdict_err := common.TextFileToArray(skkdictPath)

	if iddef_err != nil || skkdict_err != nil {
		return fmt.Errorf("Cannot_open_file")
	}

	id, err := common.GetId(iddef, "名詞,一般,*,*,*,*,*")
	if err != nil {
		return fmt.Errorf("Cannot_get_id")
	}

	for _, line := range skkdict {
		yomi_and_tango := strings.Split(line, " /")
		yomi := yomi_and_tango[0]

		// 読みに英字が含まれている場合はスキップ
		if strings.ContainsAny(yomi, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			continue
		}

		for tango_index, tango := range common.RemoveDeplication(strings.Split(yomi_and_tango[1], "/")) {
			tango := strings.Split(tango, ";")[0]

			// コストの計算
			cost := 7000 + (tango_index * 100)

			// 出力
			fmt.Println(common.MakeDictLine(yomi, tango, id, cost))
		}
	}

	return nil
}
