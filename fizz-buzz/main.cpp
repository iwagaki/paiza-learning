#include <iostream>
#include <string>

using namespace std;

int main() {
    int N;
    cin >> N;

    for (int i = 1; i <= N; i++) {
        string result;
        if (i % 3 == 0)
            result += "Fizz ";
        if (i % 5 == 0)
            result += "Buzz ";
        if (result == "")
            result += to_string(i);

        auto right = result.find_first_not_of(" ");
        auto left = result.find_last_not_of(" ");
        result = result.substr(right, left - right + 1);
        cout << result << endl;
    }
    
    return 0;

}
