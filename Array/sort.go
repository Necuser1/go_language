package main

import (
	"fmt"
	"sort"
)

type mobiles struct {
	price int
	speed int
}
type sortImplementer []mobiles

func (s sortImplementer) Len() int {
	return len(s)
}

func (s sortImplementer) Swap(i, j int) {
	if s[i].price == s[j].price {
		if s[i].speed < s[j].speed {
			s[i], s[j] = s[j], s[i]
		}
	} else {
		s[i], s[j] = s[j], s[i]
	}
}

func (s sortImplementer) Less(i, j int) bool {
	return s[i].price <= s[j].price
}

func main() {
	noOfMobiles := []mobiles{}
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var mobile mobiles
		fmt.Scan(&mobile.price)
		fmt.Scan(&mobile.speed)
		noOfMobiles = append(noOfMobiles, mobile)
	}

	sort.Sort(sortImplementer(noOfMobiles))
	fmt.Println(noOfMobiles)
}
