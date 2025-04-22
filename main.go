// duplicated_detector prints the text of each line that appaears more than
// once in the standard input , preceded by its counts

package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		counts[scanner.Text()]++
	}

	for line , n := range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}