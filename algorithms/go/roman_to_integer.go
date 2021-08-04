/*

    problem:
        Seven different symbols (known)
            I = 1
            Iv = 4
            V = 5
            IX = 9
            X = 10
            XL = 40
            L = 50
            XC = 90
            C = 100
            CD = 400
            D = 500
            CM = 900
            M = 1000
         Make the number by combining largest possible symbols
         Largest to smallest, left to right
    
    questions:
        are we only doing this for positive numbers? YES
        can the string be invalid roman numeral? NO
        can the string be empty? NO
        we return the NUMBER from the roman numeral
        roman numeral always fits in int? YES
        can you have combinations that would equal another numeral? ex. IVI = 5? NO
        
     test cases:
        I = 1
        IV = 4
        XIV = 14
        XXIV = 24
        XXIII = 23
        
      solutions:
        loop through string, match against know symbol, add to int
        
      improvements:
        could build up an array of changes to improve debugging
            Just print the array to see how solution was arrived
        ordered array of romanNumeral structs
            where romanNumeral is the string "CM" and value 901
            loop over ordered array to find match, apply value
        don't repeat j++ and continue, see if a 2 letter match was found, do in one place
        
      analysis:
        linear time
            loops through array len(s) times
        linear space
            number of strings increases with length of strings
            could reference chars or bytes directly to improve thsi
*/
func romanToInt(s string) int {
   var i int
    
    for j := 0; j < len(s); j++ {
        if j < len(s) - 1 {
            switch fmt.Sprintf("%c%c", s[j], s[j + 1]) {
            case "CM":
                i += 900
                j++
                continue
            case "CD":
                i += 400
                j++
                continue
            case "XC":
                i += 90
                j++
                continue
            case "XL":
                i += 40
                j++
                continue
            case "IX":
                i += 9
                j++
                continue
            case "IV":
                i += 4
                j++
                continue
            }    
        }
        
        
        switch s[j] {
        case 'M':
            i += 1000
        case 'D':
            i += 500
        case 'C':
            i += 100
        case 'L':
            i += 50
        case 'X':
            i += 10
        case 'V':
            i += 5
        case 'I':
            i += 1
        } 
    }
    
    return i
}
