#include <cmath>

const int multiplier = 10;

int string_to_number(const std::string& s) {
  int number{};
  int counter = 0;
  
  for (auto it = s.rbegin(); it != s.rend(); ++it) {
    if (*it == '-') {
      number = -number;
    } else {
      number += (*it - '0') * std::pow(multiplier, counter);
      counter++;
    }
  }
  
  return number;
}

// Sample Tests
//
// Describe(stringToNumber)
// {
//     It(shouldWorkForTheExamples)
//     {
//         AssertThat(string_to_number("123456"), Is().EqualTo(123456));
//         AssertThat(string_to_number("352605"), Is().EqualTo(352605));
//         AssertThat(string_to_number("-321405"), Is().EqualTo(-321405));
//         AssertThat(string_to_number("-7"), Is().EqualTo(-7));
//         AssertThat(string_to_number("0"), Is().EqualTo(0));
//         AssertThat(string_to_number("-0"), Is().EqualTo(0));
//     }
// };
