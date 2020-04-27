#include <iostream>

std::string reverseString(std::string const& str) {
	std::string result = str;
	
	for (int i = 0, j = result.length() - 1; i <= j; i++, j--) {
		char temp = result[i];
		result[i] = result[j];
		result[j] = temp;
	}

	return result;
}

// solution('abc', 'bc') // returns true
// solution('abc', 'd') // returns false
bool solution(std::string const& str, std::string const& ending)
{
	if (str.length() < ending.length()) {
		return false;
	}

	for (int i = str.length() - 1, j = ending.length() - 1; i >= 0, j >= 0; i--, j--) {
		if (str[i] == ending[j]) {
			continue;
		}
		return false;
	}

	return true;
}

int main()
{
	const std::string& source = "abc";
	const std::string& ending = "j";

	std::cout << "SOURCE: " << source << "; ENDING: " << ending << std::endl;
	std::cout << solution(source, ending) << std::endl;

	std::cin.get();

	return 0;
}