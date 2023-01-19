package main

import (
	"os"
	"fmt"
	"strings"
	common "github.com/Hayao0819/Re-UT/common"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: convert_dic <path of id.def> <dict path>")
		os.Exit(1)
	}
	iddefPath := args[1]
	dictPath := args[2]

	err := convert_dic(iddefPath, dictPath)
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

	var cost int
	id, err := common.GetId(iddef, "名詞,一般,*,*,*,*,*")
	if err != nil {
		return fmt.Errorf("Cannot_get_id")
	}

	for _, line := range skkdict {
		var parsed []string
		var yomi string
		//if strings.Index(line, " [") == -1 {
		if ! strings.Contains(line, " ["){
			// カタカナの場合
			parsed = strings.Split(line, ";")
			for tango_index, tango := range parsed {
				yomi = common.Kata2Hira(tango)
				cost = 7000 + (tango_index * 100)
				tango = strings.Split(tango, "(")[0]
				fmt.Println(common.MakeDictLine(yomi, tango, id, cost))
			}
		}else{
			// カタカナ以外の場合
			parsed = strings.Split(line, " [")
			yomi_list := get_yomi_list(parsed[1])
			for tango_index,tango := range strings.Split(parsed[0], ";") {
				cost = 7000 + (tango_index * 100)
				tango = strings.Split(tango, "(")[0]
				for _, yomi := range yomi_list {
					fmt.Println(common.MakeDictLine(yomi, tango, id, cost))
				}
			}
		}
	}
	return nil
}


func get_yomi_list (parsed string) (yomi_list []string) {
	kagi_removed := strings.Split(parsed, "]")[0]
	hira_converted := common.Kata2Hira(kagi_removed)
	semicol_split := strings.Split(hira_converted, ";")

	var return_list []string
	for _, yomi := range semicol_split {
		kakko_removed := strings.Split(yomi, "(")[0]
		return_list = append(return_list, kakko_removed)
	}
	return common.RemoveDeplication(return_list)
}
