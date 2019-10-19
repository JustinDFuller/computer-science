#include <cmath>

int make_negative(int num)
{
  return -std::abs(num);
}

// Sample tests
//
// Describe(make_negative_algorithm)
// {
//     It(should_handle_positive_test)
//     {
//         Assert::That(make_negative(42), Equals(-42));
//         Assert::That(make_negative(-42), Equals(-42));
//         Assert::That(make_negative(0), Equals(0));
//         Assert::That(make_negative(-0), Equals(0));
//     }
// };
