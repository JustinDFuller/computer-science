/*
Problem:
    We receive a string if integers
    We return a 32 bit signed integer

    Algorithm:
        Read and ignore leading whitespace
        Read + or - and add to integer
        Read chars until non-digit or end of string is reached
        Convert the digits to an integer "123" -> 123
        Return 0 if nothing was read
        If the integer is out of 32 bit signed range [-231, 231 - 1] clamp to range
        Only space is whitespace, not tab or newlines
        Don't ignore anything except leading whitespace

    Questions:
       what to do if whitespace after + or - (ignore)
       what to do with . for floats?  (stop if started, ignore if befo)

    Strategy:
        iterate through string
            if - is found, 
                set flag "negative" true or false       
            if + or - is found,
                set started flag to true     
                set starting index to  i + 1
            if non int or + or - is found
                if not started, ignore
                if started, stop here    
            switch on valid ints 0..9
                if negative, -s[i] x 10 ^ (i - started index)    
                if positive, +s[i] x 10 ^ (i - started index)
            if outside of range
                clamp to -2^31, 2^31    
                return num

*/
func myAtoi(s string) int {
    var negative bool
    var started bool
    
    var nums []int

    loop:
    for _, c := range s {
        var num int

        switch c {
            case ' ':
            if started {
                break loop
            }
            case '+':
            if started {
                break loop
            }
            started = true
            case '-':
            if started {
                break loop
            }    
            negative = true
            started = true
            case '0': 
                started = true
                num = 0
            case '1':
                num = 1
            case '2':
                num = 2
            case '3':
                num = 3
            case '4':
                num = 4
            case '5':
                num = 5
            case '6':
                num = 6
            case '7':
                num = 7
            case '8':
                num = 8
            case '9':
                num = 9
            case '.':
                if started {
                    break loop
                }
                return 0    
            default:
                if started {
                    break loop
                } else {
                    return 0
                }
        }
        
        if num > 0 {
            started = true
        }
        
        nums = append(nums, num) 
    }
    
    var sum int
    var position int
    for i := len(nums) - 1; i >= 0; i-- {
        if position > 10 && nums[i] != 0 {
            sum = int(math.Pow(2, 31))
            break
        }
        
        sum += nums[i] * int(math.Pow(10, float64(position))) 
        position++
    }
    
    if negative {
        sum = -sum
    }
    
    if sum >= int(math.Pow(2, 31)) {
        sum = int(math.Pow(2, 31)) - 1
    } else if sum <= -int(math.Pow(2, 31)) {
        sum = -int(math.Pow(2, 31)) 
    }
    
    return sum
}
