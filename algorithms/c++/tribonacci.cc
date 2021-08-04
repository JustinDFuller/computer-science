std::vector<int> tribonacci(std::vector<int> signature, int n) {
    std::vector<int> result{};
    
    const int signature_size = signature.size();
    
    for (int counter = 0; counter < n; counter++) {
      if (counter < signature_size) {
        result.push_back(signature[counter]);
      } else {
        result.push_back(result[counter - 1] + result[counter - 2] + result[counter - 3]);
      }
    }
    
    return result;
}

// Sample tests
// 
// Describe(Basic_tests)
// {
//     It(Test_1)
//     {
//         std::vector<int> signature = { 1, 1, 1 };
//         std::vector<int> expected = { 1, 1, 1, 3, 5, 9, 17, 31, 57, 105 };
//         Assert::That(tribonacci(signature, 10), Equals(expected));
//     }
    
//     It(Test_2)
//     {
//         std::vector<int> signature = { 0, 0, 1 };
//         std::vector<int> expected = { 0,0,1,1,2,4,7,13,24,44 };
//         Assert::That(tribonacci(signature, 10), Equals(expected));
//     }
    
//     It(Test_3)
//     {
//         std::vector<int> signature = { 0, 1, 1 };
//         std::vector<int> expected = { 0,1,1,2,4,7,13,24,44,81 };
//         Assert::That(tribonacci(signature, 10), Equals(expected));
//     }
    
//     It(Test_4)
//     {
//         std::vector<int> signature = { 1, 0,  0 };
//         std::vector<int> expected = { 1, 0, 0, 1, 1, 2, 4, 7, 13, 24 };
//         Assert::That(tribonacci(signature, 10), Equals(expected));
//     }
    
//     It(Test_5)
//     {
//         std::vector<int> signature = { 1,2,3 };
//         std::vector<int> expected = { 1,2,3,6,11,20,37,68,125,230 };
//         Assert::That(tribonacci(signature, 10), Equals(expected));
//     }
    
//     It(Test_6)
//     {
//         std::vector<int> signature = { 1,2,3 };
//         std::vector<int> expected = { 1,2 };
//         Assert::That(tribonacci(signature, 2), Equals(expected));
//     }
    
//     It(Test_7)
//     {
//         int third = rand() % 10;
//         std::vector<int> signature = { 1,2, third };
//         std::vector<int> expected = {};
//         Assert::That(tribonacci(signature, 0), Equals(expected));
//     }
// };
