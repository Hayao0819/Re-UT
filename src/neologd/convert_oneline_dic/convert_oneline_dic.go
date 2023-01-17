package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// convert_oneline_dic <path of id.def> <csv>
func convert_oneline_dic(iddefPath string, csvString string) string {
	// Read id.def
	iddef := textFileToArray(iddefPath)

	// Get csv
	csv := strings.Split(csvString, ",")

	// Run
	var yomi string
	var tango string
	var cost int
	var id int

	yomi = kata2hira(csv[11])
	tango = strings.ReplaceAll(csv[10], "#", "")
	cost, err := strconv.Atoi(csv[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert cost to int")
		os.Exit(1)
	}
	id = getId(iddef, csvString)

	return fmt.Sprintf("%s\t%d\t%d\t%d\t%s\n", yomi, id, id, cost, tango)
}

func textFileToArray(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open file: ", path)
		//return nil
		os.Exit(1)
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

	return lines
}

func getId(iddef []string, raw_csv string) int {
	csv := strings.Split(raw_csv, ",")
	matchString := csv[4] + "," + csv[5] + "," + csv[6] + "," + csv[7]

	for _, line := range iddef {
		if check_regexp(matchString, strings.Split(line, " ")[1]) {
			id, err := strconv.Atoi(strings.Split(line, " ")[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Cannot convert id to int")
				os.Exit(1)
			}
			return id
		}
	}
	return -1
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

	fmt.Println(convert_oneline_dic(iddefPath, csvString))
}
