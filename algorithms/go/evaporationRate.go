package kata

/**
 * This program tests the life of an evaporator containing a gas.
 * We know the content of the evaporator (content in ml), the percentage of foam or gas lost every day (evap_per_day) and the threshold (threshold) in percentage beyond which the evaporator is no longer useful. All numbers are strictly positive.
 * The program reports the nth day (as an integer) on which the evaporator will be out of use.
 * Note : Content is in fact not necessary in the body of the function "evaporator", you can use it or not use it, as you wish. Some people might prefer to reason with content, some other with percentages only. It's up to you but you must keep it as a parameter because the tests have it as an argument.
 */
func evaporationRate(content float64, evapPerDay int, threshold int) int {
  currentContent := content
  mlThreshold := content * (float64(threshold) / 100)
  evaporationRate := (1.0 - (float64(evapPerDay) / 100))
  day := 0
  
  for currentContent > mlThreshold {
    day += 1
    currentContent = currentContent  * evaporationRate
  }
  
  return day
}

/**
 * Sample Tests
 
package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

func dotest(content float64, evapPerDay int, threshold int, exp int) {
    var ans = 
    
}

var _ = Describe("Tests Evaporator", func() {
    It("should handle basic cases", func() {
      Expect(Evaporator(10, 10, 10)).To(Equal(22))
      Expect(Evaporator(10, 10, 5)).To(Equal(29))
      Expect(Evaporator(100, 5, 5)).To(Equal(59))
    })
})

*/
