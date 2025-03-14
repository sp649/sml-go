package parser

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func ParseForAsArray(data, dname string) string[] {
	
	var words []string
	
	scanner := bufio.NewScanner(data)
	
	for scanner.Scan() {
		
		line := scanner.Text()
		words := strings.Fields(line)

		ifword := strings.TrimSuffix(words[0], ":")
		
		if ifword == dname {
			return words[1:]
                } else {
			return "Couldnt not be found! :C"
		}

        }

}

func ParseFor(data, dname string) string {
	
	var words []string
	
	scanner := bufio.NewScanner(data)
	
	for scanner.Scan() {
		
		line := scanner.Text()
		words := strings.Fields(line)

		ifword := strings.TrimSuffix(words[0], ":")
		
		if ifword == dname {
			return strings.Join(words[1:], " ")
                } else {
			return "Couldnt not be found! :C"
		}

        }

}

