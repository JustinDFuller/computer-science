// Given an integer X
// I need to return that integer with the digits reversed.
// 
// Contraints:
// 1. If it overflows the signed 32-bit integer range, return 0
// 2. It should not panic or return an error
// 3. It should retain the positive or negative sign
func reverse(x int) int {
    s := strconv.Itoa(x)
    
    negative := s[0] == '-'
    if negative {
        s = s[1:]
    }
    
    var chars []rune
    for _, c := range s {
        temp := []rune{c}
        chars = append(temp, chars...)
    }
    
    if negative {
        chars = append([]rune{'-'}, chars...)
    }
   
    reversed, err :=  strconv.Atoi(string(chars))
    if err != nil {
        return 0
    }
    
    if  reversed <= -2147483648 {
        return 0
    } 
    
    if reversed >= 2147483648 {
        return 0
    }
 
    
    return reversed
}

func main() {
    tests := []struct{
        given int
        expected int
    }{
        {
            given: 123,
            expected: 321,
        },
        {
            given: -123,
            expected: -321,
        },
        {
            given: int(math.Pow(2, 31)),
            expected: 0,
        },
        {
            given: int(math.Pow(2, 31)),
            expected: 0,
        },
    }

    for _, t := range tests {
        if actual := reverse(t.given); actual != t.expected {
            log.Fatalf("Given %d, expected %d, got %d", t.given, t.expected, actual)
        }
    }   
}
