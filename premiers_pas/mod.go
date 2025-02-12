package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func estBissextile(annee int) bool {
	if annee%4 == 0 && (annee%100 != 0 || annee%400 == 0) {
		return true
	}
	return false
}

func estPremier(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func premiersNombresPremiers(n int) []int {
	var premiers []int
	i := 2
	for len(premiers) < n {
		if estPremier(i) {
			premiers = append(premiers, i)
		}
		i++
	}
	return premiers
}

func genererTableauAleatoire(n int) []int {
	rand.Seed(time.Now().UnixNano())
	tableau := make([]int, n)
	for i := range tableau {
		tableau[i] = rand.Intn(100) // Random numbers between 0 and 99
	}
	return tableau
}

func triBulles(tableau []int) {
	n := len(tableau)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if tableau[j] > tableau[j+1] {
				tableau[j], tableau[j+1] = tableau[j+1], tableau[j]
			}
		}
	}
}

func triSelection(tableau []int) {
	n := len(tableau)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if tableau[j] < tableau[minIdx] {
				minIdx = j
			}
		}
		tableau[i], tableau[minIdx] = tableau[minIdx], tableau[i]
	}
}

func rechercheDichotomique(tableau []int, x int) (bool, int) {
	idx := sort.Search(len(tableau), func(i int) bool { return tableau[i] >= x })
	if idx < len(tableau) && tableau[idx] == x {
		return true, idx
	}
	return false, -1
}

func organiserParTaille(noms []string) []string {
	sort.Slice(noms, func(i, j int) bool {
		return len(noms[i]) < len(noms[j])
	})
	return noms
}

func main() {
	// Test estBissextile
	fmt.Printf("2000 is leap year: %v\n", estBissextile(2000))
	fmt.Printf("1900 is leap year: %v\n", estBissextile(1900))
	fmt.Printf("2024 is leap year: %v\n", estBissextile(2024))
	fmt.Printf("2023 is leap year: %v\n", estBissextile(2023))

	// Test estPremier
	fmt.Printf("2 is prime: %v\n", estPremier(2))
	fmt.Printf("4 is prime: %v\n", estPremier(4))
	fmt.Printf("17 is prime: %v\n", estPremier(17))
	fmt.Printf("18 is prime: %v\n", estPremier(18))

	// Test premiersNombresPremiers
	fmt.Printf("First 10 prime numbers: %v\n", premiersNombresPremiers(10))

	// Test genererTableauAleatoire
	tableau := genererTableauAleatoire(10)
	fmt.Printf("Random array: %v\n", tableau)

	// Test triBulles
	triBulles(tableau)
	fmt.Printf("Sorted array with bubble sort: %v\n", tableau)

	// Test triSelection
	tableau = genererTableauAleatoire(10)
	triSelection(tableau)
	fmt.Printf("Sorted array with selection sort: %v\n", tableau)

	// Test rechercheDichotomique
	found, idx := rechercheDichotomique(tableau, tableau[5])
	fmt.Printf("Element found: %v at index: %d\n", found, idx)

	// Test organiserParTaille
	noms := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	groups := organiserParTaille(noms)
	fmt.Printf("Names grouped by length: %v\n", strings.Join(groups, ", "))
}
