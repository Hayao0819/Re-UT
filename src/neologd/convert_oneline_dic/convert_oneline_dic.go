package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
		//os.Exit(0)
		return "", fmt.Errorf("Invalid_csv_format")
	}

	yomi = kata2hira(csv[11])
	tango = strings.ReplaceAll(csv[10], "#", "")
	cost, err := strconv.Atoi(csv[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert cost to int")
		//os.Exit(1)
		return "", err
	}
	id, err = getId(iddef, csvString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get id")
		return "", err
		//os.Exit(1)
	}

	return fmt.Sprintf("%s\t%d\t%d\t%d\t%s", yomi, id, id, cost, tango), nil
}

func textFileToArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open file: ", path)
		//os.Exit(1)
		return nil, err

	}

	defer file.Close()

	lines := make([]string, 0, 3000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Cannot read file: ", path)
		//return nil
		os.Exit(1)
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
				//os.Exit(1)
				return -1, err
			}
			return id, nil
		}
	}
	return -1, fmt.Errorf("Cannot_find_id")
}

func check_regexp(reg, str string) bool {
	result := regexp.MustCompile(reg).Match([]byte(str))
	/*
	fmt.Println(str)
	fmt.Println(reg)
	fmt.Println(result)
	*/
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

	checkNkf()
	err := convert_dic(iddefPath, csvString)
	if err != nil {
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

	// Run
	for _, line := range csv {
		converted, err := convert_oneline_dic(iddef, line)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Cannot convert line")
			return err
		}
		fmt.Println(converted)
	}
	return nil
}
