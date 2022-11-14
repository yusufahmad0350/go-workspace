package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"C:\Users\yusuf\OneDrive - University of Limerick\Desktop\votes.txt"
)

func getStrings(fileName string) ([]string, error) {
	var lines []string
	//file, err := os.Open(fileName)
	file, err := ioutil.ReadFile("C:\\Users\\yusuf\\OneDrive - University of Limerick\\Desktop\votes.txt")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
func main() {
	lines, err := dataFile.getStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
