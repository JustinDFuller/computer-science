package kata

import (
  "strings"
)

func FindShort(s string) int {
  words := strings.Split(s, " ")
  
  minLength := len(s)
  
  for _, word := range words {
    length := len(strings.TrimSpace(word))
    
    if (length < minLength) {
      minLength = length
    }
  }
  
  return minLength
}

/**
 * Sample tests
package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Test Example", func() {
  It("should test that the solution returns the correct value", func() {
    Expect(FindShort("bitcoin take over the world maybe who knows perhaps")).To(Equal(3))
    Expect(FindShort("turns out random test cases are easier than writing out basic ones")).To(Equal(3))
  })
})
*/
