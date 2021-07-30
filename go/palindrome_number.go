// return true if x is a palindrome
//  a palindrome is a number that is the same backwards and forwards
//
//  questions:
//      are numbers always positive? NO
//      should the sign be included? YES
//      is a single digit a palindrome? YES
//
//  strategies:
//      turn it into a string?
//          loop through string
//          compare i to len(s) - 1 - i
//
//  test cases:
//      1 true
//      11 true
//      121 true
//      1221 true
//      1231 false
//      -1 false
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    if x < 10 {
        return true
    }
    
    s := strconv.Itoa(x)
    for i, c := range s {
        j := len(s) - 2 - i 
        
        if j == i {
            return true
        }
        
        if c != rune(s[j]) {
            return false
        }
    }
   
    return true
}
