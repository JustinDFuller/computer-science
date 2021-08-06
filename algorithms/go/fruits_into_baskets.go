/*
Problem:
    Array of integers representing types of fruit trees
        1 = tree type 1
        2 = tree type 2
    I have two baskets
        Each basket can hold ONE type of fruit
        Each basket can hold unlimited fruit
    I must pick a place to start picking fruit
        Once I start picking fruit, I must pick one fruit from every tree
        All trees in that range must go in one of my two baskets
        Once I get to a different type, I stop         
    Return the maximum number of fruit I can pick 
    
Questions:
    Will there always be more than one type of fruit? NO
    Can I go back and forth between fruits? 1 2 1 2? YES
    Is there any order to the fruits? NO

Examples:
    [ 1, 2, (3, 4)] -> 2
    [ (1, 2, 1, 2)] -> 4
    [] -> 0
    [(1)] -> 1
    [ (1, 2, 1, 2), 3, 4, 5 ,6] - > 4
    [ 1, 2, 1, 3, (4, 5, 4, 5)] -> 4
    [ 1, 2, (3, 4, 4, 4,) 6, 7] -> 4
    [ 1, 2, 2, 1, 3, 3] -> 

Strategies:
    Walk through the array
        if type is basket a, increment basket a
            if sum a + b > max, max = sum
        if type is basket b, increment basket b
            if sum a + b > max, max = sum
        if type is neither
            figure out which basket I picked last
            reset the other basket
            set other basket to new type
            increment count of new type
            if sum new type + old type > max, max = sum
        return sum                        

*/

type Basket struct {
    last int
    fruit int
    count int
    streak int
}

func totalFruit(fruits []int) int {
    var max int
    
    var b1 Basket
    var b2 Basket
    
    for i, fruit := range fruits {
        if b1.fruit == fruit {
            b1.last = i + 1
            b1.count++
            b1.streak++
            b2.streak = 0
        } else if b2.fruit == fruit {
            b2.last = i + 1
            b2.count++
            b2.streak++
            b1.streak = 0
        } else {
            if b1.last > b2.last {
                b2 = Basket{
                    last: i,
                    count: 1,
                    streak: 1,
                    fruit: fruit,
                }
                b1.count = b1.streak
                b1.streak = 0
            } else {
                b1 = Basket{
                    last: i + 1,
                    count: 1,
                    streak: 1,
                    fruit: fruit,
                }
                b2.count = b2.streak
                b2.streak = 0
            }
        }
        
        if sum := b1.count + b2.count; sum > max {
            max = sum   
        }
    }

   return max 
}
