package main

import (
	"bufio"
	"github.com/Cgboal/DomainParser"
	"fmt"
	"os"
	"strings"
)

func main(){

	parser := parser.NewDomainParser()
	sc := bufio.NewScanner(os.Stdin)
	
	for sc.Scan() {
		line := sc.Text()
		// domain := parser.GetDomain(line)
		line = strings.ToLower(line)
		fqdn := parser.GetFQDN(line)
		fmt.Println(fqdn)
	}

}
