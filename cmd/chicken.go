package main

import (
       "bufio"
       "flag"
       "fmt"
       "os"
       "strings"
)

func main() {

     var emoji = flag.Bool("emoji", false, "🐔")

     flag.Parse()

     var chicken string

     if *emoji {
     	chicken = "🐔"
     } else {
        chicken = "chicken"
     }

     for _, path := range flag.Args() {

     	 var buf *bufio.Scanner
	 
     	 if path == "-" {
	    buf = bufio.NewScanner(os.Stdin)
	 } else {
	 
	    fh, err := os.Open(path)

		if err != nil {
	   	panic(err)
		}
		
		buf = bufio.NewScanner(fh)
	}
	
	for buf.Scan(){

	    txt := buf.Text()
    	    // fmt.Println(txt)

	    // https://golang.org/pkg/unicode/
	    
	    chickens := make([]string, 0)

	    matches := strings.Fields(txt)
	    count := len(matches)
	    
	    for i:=0; i < count; i ++{
	    	chickens = append(chickens, chicken)
	    }

	    fmt.Println(strings.Join(chickens, " "))
	}
     }
}
