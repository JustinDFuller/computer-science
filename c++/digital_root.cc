#include <string>

int digital_root(int n) {
  if (n < 10) {
    return n;
  }
  
  int sum = 0;
  
  const std::string split_int = std::to_string(n);
  
  for (const char number_char : split_int) {
    sum += number_char - '0';
  }
  
  return digital_root(sum);
}

// Sample tests
// 
// Describe(Fixed_tests)
// {
//     It(Digital_root)
//     {
//         Assert::That(digital_root(16) , Equals(7));
//         Assert::That(digital_root(195) , Equals(6));
//         Assert::That(digital_root(992) , Equals(2));
//         Assert::That(digital_root(167346) , Equals(9));
//         Assert::That(digital_root(0) , Equals(0));
//     }
// };
