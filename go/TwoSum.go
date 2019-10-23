/**
 * This is actually quite bad perf. n ^ 2
 * but the randomly generated tests did not produce sorted values,
 * unlike the sample tests below, which are sorted.
 * I could order them and use a sliding window or maybe even a
 * binary search. Maybe another time.
 */
func TwoSum(numbers []int, target int) [2]int {
  for index, number := range numbers {
    for innerIndex, innerNumber := range numbers[index + 1:] {
      if (number + innerNumber == target) {
        return [2]int{index, innerIndex + index + 1}
      }
    }
  }
  
  return [2]int{}
}

/*
// Sample Tests

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Tests", func() {
    It("Basic tests", func() {
        Expect(TwoSum([]int{1, 2, 3}, 4)).To(Equal([2]int{0, 2}))
        Expect(TwoSum([]int{1235, 5678, 9012}, 14690)).To(Equal([2]int{1, 2}))
        Expect(TwoSum([]int{2, 2, 3}, 4)).To(Equal([2]int{0, 1}))
    })
})
*/
