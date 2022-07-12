// package main

// import "fmt"

// func main() {
// 	PrintEditiors := map[string]map[string]string{

// 		"Alex": {
// 			"Miracle Morning": "2019",
// 		},
// 		"Sam": {
// 			"Atomic Habits": "2021",
// 		},
// 		"Clover": {
// 			"Rich Dad Poor Dad": "2015",
// 		},
// 	}

// 	keys := make([]string, 0, len(PrintEditiors))
// 	// for k := range PrintEditiors {
// 	// 	keys = append(keys, k)
// 	// }

// 	for _, v := range PrintEditiors {
// 		keys = append(keys, PrintEditiors[v])
// 		fmt.Println(v)

// 	}
// 	// fmt.Println(len(PrintEditiors), PrintEditiors)
// }

package main

import (
	"fmt"
	"sort"
)

func main() {
	infoForReaders := map[string]map[string]string{
		"Piter": {
			"name":       "Python for beginners",
			"periodical": "8",
		},
		"Stiver": {
			"name":       "JavaScript on level middle",
			"periodical": "3",
		},
		"Lilya": {},
		"John": {
			"name":       "senior Go",
			"periodical": "1",
		},
		"Geogy": {
			"name":       "JavaScript on level middle",
			"periodical": "3",
		},
		"Tom": {},
		"Mark": {
			"name":       "Java not for juniors",
			"periodical": "9",
		},
	}

	fmt.Println("Number of readers:", len(infoForReaders))
	fmt.Println("All publications:")

	keys := make([]string, 0, len(infoForReaders))

	for k := range infoForReaders {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(infoForReaders[v])
	}

	// fmt.Println(len(keys), cap(keys))
}
