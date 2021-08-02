/* 
Problem:
    Given a string of letters and numbers
    Find the longest substring that DOES NOT have repeating characterso
    
Questions:
    Letters and numbers? -> yes
    Are there empty strings? -> 0
    Are there always repeating letters? -> no
    What if all letters are repeating? -> 0
    do we count the dupes in the count? yes, only one
    
Examples:
    "" -> 0
    "a" -> 1
    "abc" -> 3
    "aaa" -> 1
    "aaabcd" -> 4
    "abcddd" -> 4
    "abbbcd" -> 3
    
Strategies:
    loop through the string
    count seen characters
    if character is the same as last character, found duplicate
    if count > max, max = count
    reset count, continue counting at current character
    O(n) complexity
    
    that was wrong, I need to track all seen chars, not just 2 in a row
    keep of map of seen chars, reset when duplicate is found
    
    brute force:
        1. loop through string
        2. loop through all permutations starting at [i]
        3. loop through each permutation to check for duplicates
        4. if no duplicates, check if greater length than max
        5. if greater than max, set max to length of substring
        loop + loop + loop = O(n^3) cubic time complexity
        map size = unique chars in string -> linear as charset grows or constant if known charset 
        
    sliding window:
        1. 2 pointers, both start at index 0
        2. check subs for duplicates
            if no duplicates, set max to len if len > max
            if no duplicates, extend end pointer (j++)
            if duplicates, contract start pointer (i++)
        3. repeat until i and j both at len(s) - 1 
        loop + loop = O(n^2) quadratic time complexity
        still using map that is maximum the size of the charset used
        
     sliding window with 1 loop:
        1. 2 pointers, both start at index 0
        2. at index 0, 
            seen[char] = true
            lastSeen[char] = j
        3. check if lastSeen[char] >= i (we saw it before, and in the current substring)
            move i to seen[j] + 1 (end after the last duplicate)
        4. set seen[s[j]] = j
        5. check if j - i (len of subs) is greater than max
            if yes, max = j - i
        6. j++
        
*/
func lengthOfLongestSubstring(s string) int {
   var max int
    
    if l := len(s); l < 2 {
       return l 
    }
    
   /* brute force
    for i := range s {
        for j := i + 1; j <= len(s); j++ {
            subs := s[i:j] 
            if l := len(subs); l > max && !hasDupes(subs) {
                max = l
            } 
        }
    } */
    
    /* sliding window
    i := 0
    j := 0
    for i < len(s) && j < len(s) {
        subs := s[i : j + 1] 
        if hasDupes(subs) {
            i++
        } else {
            if l := len(subs); l > max {
                max = l
            }
            j++
        }
    } */
    
    /* optimized sliding window */
    start, end := 0, 0
    jumpTo := map[byte]int{}
    for end < len(s) {
        byt := s[end]
        
        if jump := jumpTo[byt]; jump > start {
            start = jump
        } else if l := end - start + 1; l  > max {
            max = l
        }
        
        jumpTo[byt] = end + 1
        end++
    }
    
   return max 
}

/*
func hasDupes(s string) bool {
    seen := map[rune]bool{}
    for _, r := range s {
        if seen[r] {
            return true
        }
        seen[r] = true
    }   
    return false
}
*/
