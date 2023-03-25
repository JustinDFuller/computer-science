// You are given an array of colors and numbers
// 1 blue, 1 green, 2 red, 3 blue, 4 green
// You can assume that the array is sorted by number, but not by color
// Given an O(n) algorithm to sort by color but keep each color sorted by number
// 1 blue, 3 blue, 1 green, 4 green, 2 red
package exercises

import "log"

type Color struct {
	count int
	name  string
}

// sortByColor uses a bucket sort to sort by color
// each color is already sorted by number, so it retains that sorting
func sortByColor(input []Color) []Color {
	buckets := map[string][]Color{}

	for _, color := range input {
		buckets[color.name] = append(buckets[color.name], color)
	}

	var output []Color
	for _, colors := range buckets {
		output = append(output, colors...)
	}

	return output
}

func main() {
	input := []Color{
		{
			name:  "blue",
			count: 1,
		},
		{
			name:  "green",
			count: 1,
		},
		{
			name:  "red",
			count: 2,
		},
		{
			name:  "blue",
			count: 3,
		},
		{
			name:  "green",
			count: 4,
		},
	}

	log.Print(sortByColor(input))
}

// 2009/11/10 23:00:00 [{1 blue} {3 blue} {1 green} {4 green} {2 red}]
