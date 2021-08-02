# problems

[![](https://www.codewars.com/users/JustinDFuller/badges/micro)](https://www.codewars.com/users/JustinDFuller)

Solving problems in different languages. Mostly from [codewars](https://codewars.com).

## Formulas and Patterns for solving algorithms

### Pattern Variations

Always include test cases for these standard pattern variations. These can apply to any iterable like a string or array (even numbers).

1. Empty pattern ("")
2. Single pattern ("A")
3. No pattern ("ABCDE")
4. All pattern ("AAAAA")
5. Starting pattern ("AAABCD")
6. Ending pattern ("ABCDDDD")
7. Middle pattern ("ABCCCDE")
8. Repeating pattern ("ABCABC")

### Formulas

1. Brute force: create and search all permutations for a condition.

Loop over the string, number, or array to find every permutation. Search every permutation for the condition (ex. duplications). This may involve 3 loops: loop over the iterable, loop again from each starting point, potentially loop over again if you need to check the contents of each permutation (like duplicates).

_Algorithmic Complexity:_ 
  - Time: almost Cubic O(n^3) but you can optimize to not run all 3 loops for the entirety of the array.
  - Space: Linear O(m) have to create a map of seen values, limited by charset of whatever you index values by (ex. UTF chars).

2. Sliding Window

Maintain two pointers to the iterable. Maintain a dictionary of seen values. Ideally, dictionary stores last location that value was seen.

When the condition is false (ex. repeated chars not found) move the right pointer up. When the condition is true (ex. repeating chars found), move the left pointer up. Ideally, move the left pointer to a known-good location (ex. where the last repeated char was found + 1). When the right pointer reaches the end of the iterable, you are done.

_Algorithmic Complexity_:
  - Time: Linear O(n). Worst case, you have to move each pointer to each location O(2n), which reduces to O(n) which is linear.
  - Space: Linear O(m) where m is the charset by which you index values. May be UTF characters or something else.

3. Tree

TODO
