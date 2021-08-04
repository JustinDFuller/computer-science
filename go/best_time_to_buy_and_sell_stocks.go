/**

Problem:
    Array of integer
    Is integer is the proce of a given stock
    Each index is the day
    Maximize profit by 
        choosing a single day to buy one stock
        choosing a different day in the future to sell that stock
        Return the maximum profit
        If you cannot achieve profit, return 0
        
Constraints:
    There's always at least one price
    Price is always >= 0
    You cannot sell before you buy
    
    
Examples:
    [0]
    [1]
    [1, 2]
    [2, 1]
    [5, 6, 4, 8]
    [4, 8, 3, 6, 9]
    [4, 8, 3, 9, 6]
    
Options:
    Brute force:
        Iterate over every index
            For each index (day) iterate over all days after it
            multiply day 1 * day 2. If greater than max, set new max.
            
     Analysis:
        Time: Outer loop N times, inner loop N - i times. O(n * n/2) times. 
        Space: Constant, uses no extra storage other than max int variable
        
     Track lowest and highest
        walk through prices
            if < lowest
                price = lowest
                price = highest
            if > highest
                price = highest
                diff = highest - lowest
                if diff > max
                    max = diff
     analysis:
         time: O(n) b/c it only walks through loop once.
         space: Constant b/c it only tracks 3 ints, not matter the input size.

**/
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }
    
   /*  brute force
    var max int
    
    for i, price := range prices {
        if i == len(prices) - 1 {
            break
        }
        
        for _, price2 := range prices[i + 1:] {
            if diff := price2 - price; diff > max {
                max = diff
            }
        }
    } */
    
    var max int
    lowest := prices[0]
    highest := prices[0]
    
    for _, price := range prices {
        if price < lowest {
            lowest = price
            highest = price
        } 
        
        if price > highest {
            highest = price
        } 
        
        if diff := highest - lowest; diff > max {
            max = diff
        }
    }
    
    return max
}
