<img src="https://repository-images.githubusercontent.com/212989121/78532e2f-9749-4d93-84f2-6216605ebd1a" />

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
9. Reeeaaalllly long pattern ("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABBBBBBBBB")
10. Reeeealllly long absence of pattern ("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

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

3. Removals

When removing from an array, you should often opt to iterate in reverse, to prevent the changing length from causing problems.

You could also consider these other data structures that are better at removals:

* Hashmap or Dictionary O(1)
* Queue or Stack are potentially O(1)
* Linked List O(n)

## Concepts

- [ ] Space and Time Complexity
  - [ ] Memory Usage (Number of entries Data Structure)
  - [ ] Recursion Usage (Number of Recursions)
- [ ] How to recognize common complexities
  - [ ] Constant Complexity `O(1)`
  - [ ] Logarithmic Complexity `O(log n)`
  - [ ] Linear Complexity `O(n)`
  - [ ] Linearithmic Complexity `O(n log n)`
  - [ ] Quadratic Time `O(n²)`
  - [ ] Cubic Time `O(n³)`
  - [ ] Exponential Time `2⁰⁽ⁿ⁾`
  - [ ] Factorial Time `O(n!)`

## Data Structures

- [ ] Array
- [ ] Linked List
  - [ ] Singly Linked List
  - [ ] Doubly Linked List
  - [ ] Circular Linked List
- [ ] Hash Table
- [ ] Queue
  - [ ] Circular Queue
  - [ ] Priority Queue
  - [ ] Double Ended Queue
- [ ] Stack
  - [ ] Min Stack
  - [ ] Max Stack
  - [ ] Min-Max Stack
- [ ] Heap
- [ ] Tree
  - [ ] Binary Tree
  - [ ] Binary Search Tree
  - [ ] AVL Tree
  - [ ] Red Black Tree
  - [ ] Splay Tree
  - [ ] Treap
  - [ ] B-Tree
  - [ ] N-ary Tree 
- [ ] Trie
- [ ] Graph
  - [ ] Undirected Graph
  - [ ] Directed Graph
  - [ ] Cyclic Graph
  - [ ] Edge Labeled Graph
  - [ ] Directed Acyclic Graph
  - [ ] Disconnected Graph
  - [ ] Weighted Graph
  - [ ] Unweighted Graph
  - [ ] Adjacency List
  - [ ] Adjacency Matrix
- [ ] Bit (Bit manipulation)

## Algorithms

- Array
  - Unsorted Search
    - [ ] Sliding Window
    - [ ] Double Sliding Window
    - [ ] Reference hash table or map
  - Sorted Search
    - [ ] Binary Search
    - [ ] Jump Search
    - [ ] Interpolation Search
    - [ ] Exponential Search
    - [ ] Ternary Search
  - Sorting
    - [ ] Quick Sort
    - [ ] Merge Sort
    - [ ] Tim Sort
    - [ ] Bubble Sort
    - [ ] Radix Sort
    - [ ] Counting Sort
    - [ ] Heap Sort
- Linked List
- Hash Table
- Queue
- Stack
- Trees
  - Breadth-First Search
  - Depth-First Search
- Graphs
- Bit Manipulation
  - [ ] AND
  - [ ] OR
  - [ ] XOR
  - [ ] NOT
  - [ ] Bit Shifts

## Concurrent Programming

- Building Blocks
  - [ ] Goroutine
  - [ ] Mutex
  - [ ] Channel
  - [ ] Buffered Channel
- Strategies
  - [ ] Pipeline
