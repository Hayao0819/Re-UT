package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const max_process = 5000

// convert_oneline_dic <id.def> <csv>
func convert_oneline_dic(iddef []string, csvString string) (string, error) {
	// Get csv
	csv := strings.Split(csvString, ",")

	// Run
	var yomi string
	var tango string
	var cost int
	var id int

	if len(csv) < 11 {
		fmt.Fprintln(os.Stderr, "Invalid csv format")
		fmt.Fprintln(os.Stderr, csvString)
		return "", fmt.Errorf("Invalid_csv_format")
	}

	yomi, err := kata2hira(csv[11])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert katakana to hiragana")
		return "", err
	}
	tango = strings.ReplaceAll(csv[10], "#", "")
	cost, err = strconv.Atoi(csv[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert cost to int")
		return "", err
	}
	id, err = getId(iddef, csvString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get id")
		return "", err
	}

	return fmt.Sprintf("%s\t%d\t%d\t%d\t%s", yomi, id, id, cost, tango), nil
}

func textFileToArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open file: ", path)
		return nil, err

	}

	defer file.Close()

	lines := make([]string, 0, 3000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read file: ", path)
		return nil, err
	}

	return lines, nil
}

func getId(iddef []string, raw_csv string) (int, error) {
	csv := strings.Split(raw_csv, ",")
	matchString := csv[4] + "," + csv[5] + "," + csv[6] + "," + csv[7]

	for _, line := range iddef {
		if check_regexp(matchString, strings.Split(line, " ")[1]) {
			id, err := strconv.Atoi(strings.Split(line, " ")[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Cannot convert id to int")
				return -1, err
			}
			return id, nil
		}
	}
	return -1, fmt.Errorf("Cannot_find_id")
}

func check_regexp(reg, str string) bool {
	result := regexp.MustCompile(reg).Match([]byte(str))
	return result
}

func main(){
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: convert_oneline_dic <path of id.def> <csv>")
		os.Exit(1)
	}
	iddefPath := args[1]
	csvString := args[2]

	_ = checkNkf()

	err := convert_dic(iddefPath, csvString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert dic")
		os.Exit(1)
	}
}


func convert_dic(iddefPath string, csvPath string) (error){

	csv, csv_err := textFileToArray(csvPath)
	iddef, iddef_err := textFileToArray(iddefPath)

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
