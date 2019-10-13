#include <unordered_map>
#include <ctype.h>

typedef std::unordered_map<char, bool> character_map;

bool is_isogram(std::string str) {
  bool result = true;
  character_map letters{};
  
  for (auto character : str) {
    char lowercase_character = tolower(character);
    
    if (letters[lowercase_character]) {
      result = false;
      break;
    }
    
    letters[lowercase_character] = true;
  }
  
  return result;
}

// Sample tests
// 
// Describe(Tests)
// {
//     It(isogram_or_not)
//     {
//         Assert::That(is_isogram("Dermatoglyphics"), Equals(true));
//         Assert::That(is_isogram("moose"), Equals(false));
//         Assert::That(is_isogram("isIsogram"), Equals(false));
//     }
// };
