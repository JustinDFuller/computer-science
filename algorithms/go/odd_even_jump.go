/*
Problem:
    Given an array of integers
    From somewhere in the array (maybe multiple places) you can make a series of jumps
    jumps != indexes
        Jumps are 1 indexed, not zero indexed
    you can jump from INDEX i -> j
        i < j
    odd jumps
        -  arr[i] <= arr[j] 
        -  arr[j] is the smallest possible value
        - so if array[i] = 1 and there is arr[j] = 3 and 2, you have to jump to 2.
    even jumps
        -   arr[i] >= arr[j]
        - arr[j] is the LARGEST possible value
        - so iff arr[i] = 3 and arr[j] can be 1 or 2, you have to jump to 2.
        
    return the number of starting points where you can reach the final index.
    
Examples:
    - [] -> 0 (no jumps)
    - [1] -> 1 (no jumps, but 0 jumps counts)
    - [2, 1] -> 1 (no jumps, but 0 jumps counts)
    - [1, 2] -> 2 (odd jump only + last index)
    - [2, 3, 1] -> 3 (does both even and odd jump)
    - [3, 2, 1] -> 2 (even jump only + last index)
    
Methods:
    Brute force
        start at first integer (i)
            1. jump is 1
            2. j is i
            3. if j == len(arr) - 1 (we made it to end of array) increment good counter, i++, break
            4. if jump is odd
                    need to find smallest number higher than arr[j]        
               if jump is even
                    need to find the biggest number smaller than arr[j] 
            5. if it does not exist, move i + 1 and start over (break loop)
            6. if it exists, 
                j = the index of that number
                jump++
                back to step 3
            7. if i == len(arr) - 1 return good jumps   
*/
func oddEvenJumps(arr []int) int {
    var good int
    
    // [] doesn't iterate, returns 0
    for i := range arr {
        // [3] stops here, returns 1
        if i == len(arr) - 1 {
            good++
            break
        }
        
        // [2,3,1,1,4]
        jump := 0
        j := i
        for j < len(arr) {
            jump++

            if j == len(arr) - 1 {
                good++
                break   
            } 
            
            if jump % 2 == 0 {
                if index := findBiggestNumberSmallerThan(arr[j], j, arr); index != 0 {
                   // found a jump, go to next j
                   j = index
                } else {
                    // couldn't find a jump, go to next i
                    break    
                }
            } else {
                if index := findSmallestNumberBiggerThan(arr[j], j, arr); index != 0 {
                    // found  ajump, go to next j
                    j = index
                } else {
                    // couldn't find a jump, go to next i
                    break    
                }
            }
        }
    }
    
    return good
}

func findBiggestNumberSmallerThan(num, start int, arr []int) int {
    var max int
    var index int

    for i := start + 1; i < len(arr); i++ {
        if arr[i] <= num && arr[i] > max {
            max = arr[i]
            index = i
        }
    }
    
    return index
}

func findSmallestNumberBiggerThan(num, start int, arr []int) int {
    var min int
    var index int
    var found bool

    for i := start + 1; i < len(arr); i++ {
        if arr[i] >= num && (arr[i] < min || !found) {
            min = arr[i]
            index = i
            found = true
        }
    }
    
    return index
}
