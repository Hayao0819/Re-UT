package reut_common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

func GetId(iddef []string, match string) (int, error) {
	for _, line := range iddef {
		if strings.Contains(strings.Split(line, " ")[1], match) {
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

func TextFileToArray(path string) ([]string, error) {
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

func RemoveDeplication(org_array []string) []string {
	mp := make(map[string]bool)
	uniq := []string{}

	for _, item := range org_array {
		if !mp[item] {
			mp[item] = true
			uniq = append(uniq, item)
		}
	}
	return uniq
}


func MakeDictLine(yomi string, tango string, id int, cost int) string {
	return fmt.Sprintf("%s\t%d\t%d\t%d\t%s", yomi, id, id, cost, tango)
}
