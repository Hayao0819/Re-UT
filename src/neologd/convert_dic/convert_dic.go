package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"github.com/miiton/kanaconv"
	common "github.com/Hayao0819/Re-UT/common"
)

const max_process = 1000

// convert_oneline_dic <id.def> <csv>
func convert_oneline_dic(iddef []string, csvString string) (string, error) {
	// Get csv
	csv := strings.Split(csvString, ",")

	// Run
	var yomi string
	var tango string
	var cost int
	var id int
	var err error

	if len(csv) < 11 {
		fmt.Fprintln(os.Stderr, "Invalid csv format")
		fmt.Fprintln(os.Stderr, csvString)
		return "", fmt.Errorf("Invalid_csv_format")
	}

	yomi = kanaconv.KatakanaToHiragana(csv[11])
	tango = strings.ReplaceAll(csv[10], "#", "")
	cost, _ = strconv.Atoi(csv[2])
	id, err = common.GetId(iddef, csvString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get id")
		return "", err
	}
	//return fmt.Sprintf("%s\t%d\t%d\t%d\t%s", yomi, id, id, cost, tango), nil
	return common.MakeDictLine(yomi, tango, id, cost), nil
}


func main(){
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: convert_dic <path of id.def> <csv>")
		os.Exit(1)
	}
	iddefPath := args[1]
	csvString := args[2]

	err := convert_dic(iddefPath, csvString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert dic")
		os.Exit(1)
	}
}


func convert_dic(iddefPath string, csvPath string) (error){
	csv, csv_err := common.TextFileToArray(csvPath)
	iddef, iddef_err := common.TextFileToArray(iddefPath)

	if csv_err != nil || iddef_err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read file")
		return fmt.Errorf("Cannot_read_file")
	}

	wg := &sync.WaitGroup{}
	
	sem := make(chan struct{}, max_process)

	for _, line := range csv {
		wg.Add(1)
		sem <- struct{}{}

		go func(wg *sync.WaitGroup, line string, iddef []string)() {
			converted, err := convert_oneline_dic(iddef, line)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Cannot convert line")
			}else{
				fmt.Println(converted)
			}
			wg.Done()
			<-sem
		}(wg, line, iddef)
	}

	wg.Wait()
	return nil
}
