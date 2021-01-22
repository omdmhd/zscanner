package dns

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

var values = [36]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

//Execute does lookuphost
func Execute(domain string, length int, start int) {
	for j := start; j <= length; j++ {
		fmt.Println("doing domains of length" + strconv.Itoa(j))
		domains := generateSubdomains(domain, j)
		var wg sync.WaitGroup
		groupLength := 100
		for i := 0; i < (len(domains) - groupLength); i += groupLength {
			wg.Add(1)
			group := domains[i : i+groupLength]
			go func(wg *sync.WaitGroup, domains []string) {
				for _, v := range domains {
					_, err := net.LookupHost(v)
					if err == nil {
						fmt.Println(v)
					}
				}
			}(&wg, group)
		}
		wg.Wait()
	}
}

func generateSubdomains(domain string, length int) []string {
	subData := []int{}
	for i := 0; i < length; i++ {
		subData = append(subData, 0)
	}
	domains := []string{domain, toStr(subData) + "." + domain}
	for increment(subData) {
		domains = append(domains, toStr(subData)+"."+domain)
	}
	return domains
}

func increment(posible []int) bool {
	posible[0]++
	i := 0
	for posible[i] >= len(values) {
		posible[i] = 0
		if i+1 >= len(posible) {
			return false
		}
		i++
		posible[i]++

	}
	return true
}

func toStr(vals []int) string {
	s := ""
	for i := 0; i < len(vals); i++ {
		s += values[vals[i]]
	}
	return s
}
