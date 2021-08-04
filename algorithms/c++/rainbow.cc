#include <iostream>
#include <string>
#include <vector>

// Example usage/output
// rainbow("hello", ["red", "green"]) "rrrgg"
// rainbow("hello", ["red", "green", "blue"]) "rrggb"
// rainbow("orange", ["red", "green", "blue"]) "rrggbb"

std::string rainbow(std::string message, std::vector<std::string> colors) {
    const int messageLength = message.size();
    const int colorLength = colors.size();
    
    const int dividend = messageLength / colorLength;
    int remainder = messageLength % colorLength;
    
    std::string output = "";
    
    for (int i = 0; i < colors.size(); ++i) {
        if (remainder != 0) {
            --remainder;
            output.append(std::string(dividend + 1, colors[i][0]));
        } else {
            output.append(std::string(dividend, colors[i][0]));
        }
    }
    
    return output;
}
 
int main() {
    std::cout << rainbow("hello", std::vector<std::string>{"red", "green"});
    std::cout << "\nrrrgg \n\n";
    std::cout << rainbow("hello", std::vector<std::string>{"red", "green", "blue"});
    std::cout << "\nrrggb \n\n";
    std::cout << rainbow("orange", std::vector<std::string>{"red", "green", "blue"});
    std::cout << "\nrrggbb \n\n";
    return 0;
}
