// Find the number of unique email addresses
// Unique email addresses:
//  ignore dots in the local name
//  ignore anything after a + in the local name
//  casing is ignored
//
//  test cases:
//        justin.fuller@gmail.com
//        justinfuller+foobar@gmail.com
//        justin.fuller+foobar@gmail.com
//        JUSTINFULLER@GMAIL.COM      
//        @gmail.com
//        justinfuller@    
//        justinfullergmail.com
//        justinfuller@gmail.com+foobar    
//
// data structures
//  map to store unique emails
//  can loop over emails, apply rules, add to map
//  see size of map at end
//
// edge cases:
//  No emails
//  Rules applied to domain
//  No domain
//  No local    
func numUniqueEmails(emails []string) int {
    uniques := map[string]bool{}
    
    for _, email := range emails {
        var local string
            
        email := strings.ToLower(email)
        split := strings.Split(email, "@")
        
        if len(split) != 2 {
            continue
        }
        
        for _, char := range split[0] {
            if char == '.' {
                continue
            } 
            
            if char == '+' {
                break
            }    
            
            local += string(char)
        }
        
        uniques[local + split[1]] = true
    } 
    
    return len(uniques)
}
