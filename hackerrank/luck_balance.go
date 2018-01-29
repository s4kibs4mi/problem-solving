package hackerrank

import (
	"fmt"
	"sort"
)

/**
 * := Coded with love by Sakib Sami on 1/29/18.
 * := s4kibs4mi@gmail.com
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

// Problem Statement : https://www.hackerrank.com/challenges/luck-balance/problem

type LuckLock struct {
	Luck       int
	Importance int
}

type LuckLake []LuckLock

func (ll LuckLake) Len() int {
	return len(ll)
}

func (ll LuckLake) Swap(i, j int) {
	ll[i], ll[j] = ll[j], ll[i]
}

func (ll LuckLake) Less(i, j int) bool {
	return ll[i].Luck > ll[j].Luck
}

func Run() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	var luckLake LuckLake
	for i := 0; i < n; i++ {
		luck := LuckLock{}
		fmt.Scanf("%d %d", &luck.Luck, &luck.Importance)
		luckLake = append(luckLake, luck)
	}
	sort.Sort(&luckLake)
	sum := 0
	for i := 0; i < n; i++ {
		if luckLake[i].Importance == 0 {
			sum += luckLake[i].Luck
			//fmt.Println("Luck+ : ", luckLake[i].Luck, ", ", luckLake[i].Importance)
			continue
		}
		if k > 0 {
			sum += luckLake[i].Luck
			//fmt.Println("Luck+ : ", luckLake[i].Luck, ", ", luckLake[i].Importance)
			k--
			continue
		}
		sum -= luckLake[i].Luck
		//fmt.Println("Luck- : ", luckLake[i].Luck, ", ", luckLake[i].Importance)
	}
	fmt.Println(sum)
}
