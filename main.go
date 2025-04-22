// duplicated_detector count the duplicated lines

package main

import (
	"fmt"
	"bufio"
	"os"
)
func countLines(file *os.File , countsMap *map[string]int){
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		(*countsMap)[scanner.Text()]++
	}
}

func main(){
	countsMap := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin,&countsMap)
	} else{
		for _ , file := range files{
			f , err := os.Open(file)
			if err!=nil{
				fmt.Fprintf(os.Stderr,"Error : %v",err)
				continue
			}
			countLines(f,&countsMap)
			f.Close()
		}
	}

	for line , occurence := range countsMap{
		if occurence > 1 {
			fmt.Printf("%s \t %d \n",line,occurence)
		}
	}
}