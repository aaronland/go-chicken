package main

import (
       "bufio"
       "flag"
       "fmt"
       "github.com/thisisaaronland/go-chicken"
       "os"
       "strings"
)

func main() {

     var lang = flag.String("language", "zxx", "...")
     flag.Parse()

     ch, err := chicken.GetChickenForLanguageTag(*lang)

     if err != nil {
     	panic(err)
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
	    	chickens = append(chickens, ch)
	    }

	    fmt.Println(strings.Join(chickens, " "))
	}
     }
}
