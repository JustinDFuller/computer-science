/*

Problem:    
    - Array of POSITIVE integers
    - The integers represent the HEIGHT (y) of the index
    - Each index always starts at 0 (x, 0), (x, y)
    - Find the two indexes with the most area
    
Solutions:

    Brute Force
        - Loop over every index
        - Loop again starting at index + 1
        - calculate the area lowest(i, j) * (j - i)
        - if area is greater than max , max = area
        - return area
        
    Sliding window?
        - Start i := 0
        - Start j := len(arr) - 1
        - calculate area
        - if area > max, max = area
        - if i <= j, move i++
        - else if i > j, move j--
        - repeat
        
    Analysis:
        Brute force is O(n^2) time because it re-loops every loop.
        Sliding window is (n) because it hits each index once.
        Both are constant time b/c they don't use additional space or depth.

*/
func maxArea(height []int) int {
    var max int
    
    /* brute force
    // 0, 1
    for i, start := range height {
        // 1, 8
        for j := i + 1; j < len(height); j++ {
            end := height[j]
            // true
            if start < end {
                // 1 * (1 - 0) == 1
                if area := start * (j - i); area > max {
                    // max = 1
                    max = area
                }
            } else {
                if area := end * (j - i); area > max {
                    max = area
                }
            }
        }
    } */

    for i, j := 0, len(height) - 1; i < j; {
        var area int

        start := height[i]
        end := height[j]
        
        if start <= end {
            area = start * (j - i)
            i++
        } else {
            area = end * (j - i)
            j--
        }
        
        if area > max {
            max = area
        }
    }

    return max
}
