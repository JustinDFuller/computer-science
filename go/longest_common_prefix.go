/*
Problem:
    an array of strings ["foo", "far"]
    find the longest common prefix "f"   
    
Questions:
    is the array ordered in any way?
    are there any guarantees about the data? (not empty, length, charset)   
    
Test Cases:
    [] -> ""
    ["foo"] -> ""
    ["foo", "bar"] -> ""
    ["florida", "fluent", "foo"] -> "fl"
    ["bar", "fetch", "feed"] -> "fe"
    ["hey", "bar", "bat", "foo"]
    
Strategies: 
    I need to track prefixes
    
    brute force:
        loop through the strings
            loop through each prefix permutation in string
            store count of each permutation in a map 
        loop through map
            if prefix was seen in EVERY string
            if prefix is > last greatest prefix
            last greatest prefix = prefix
            length = len(prefix)
         return greatest prefix
         
     analysis:
        time: length strs * width longest string (quadratic? worse?) len = 10, longest = 100 O(10 * 100)
        space: length strs * width longest string (quadratic? worse?)
        

    sliding window (kinda)
        start i at the first string
        start j at the end of the first string
        prefix is the first string
        
        loop over strings
            if the current prefix is longer than the next string, j is length of next string
                if next string is empty (j == 0), stop! no match found, no prefix
            while string[0:j] != next string[0:j] 
                j--
                if j == 0, stop! no match found, at beginning of string, no prefix 
            found a match, longest prefix is current string [0:j]
        return longest prefix
        
        analysis:
            time: linear, O(length of array + longest string), depends on if there is a longer string or longer array
        space: linear, will create a string the size of the longest string
            
            
    
*/
func longestCommonPrefix(strs []string) string {
    /*seen := map[string]int{}
    
    for i, s := range strs {
        for i := range s {
            seen[s[:i + 1]] += 1
        }
    }
    
    var length int
    var prefix string
    
    for k, v := range seen {
        if v == len(strs) && len(k) > length  {
            length = len(k) 
            prefix = k
        }  
    }*/
    
    prefix := strs[0]
    end := len(strs[0])
    for i, s := range strs {
        if i == len(strs) - 1 {
            break
        }  
        
        if end == 0 {
            return ""
        }
        
        next := strs[i + 1]
        if len(next) < end {
            end = len(next)
            if end == 0 {
                return ""
            }
        }
        
        for next[:end] != s[:end] {
            end--
            if end < 0 {
                return ""
            }
        }
        
        prefix = s[:end]
    }
    
    return prefix
}
