#include <iostream>


std::string get_middle(std::string str)
{
	std::string result = "";
	int l = str.length();
	if (l % 2 == 0) {
		result += str[(l / 2) - 1];
		result += str[l / 2];
		return result;
	}
	else
	{
		result += str[l / 2];
		return result;
	}
}

void assertionResult(std::string result, std::string expected)
{
	int temp = expected.compare(result);
	if (temp == 0)
	{
		std::cout << "SUCCESSFULL" << std::endl;
		std::cout << result << " == " << expected << std::endl;
	}
	else
	{
		std::cout << "FAILED" << std::endl;
		std::cout << result << " != " << expected << std::endl;
	}

	std::cout << std::endl;
}

int main()
{
	std::string str1 = "test";
	std::string expect1 = "es";
	assertionResult(get_middle(str1), expect1);

	std::string str2 = "testing";
	std::string expect2 = "t";
	assertionResult(get_middle(str2), expect2);

	std::string str3 = "middle";
	std::string expect3 = "dd";
	assertionResult(get_middle(str3), expect3);

	std::string str4 = "A";
	std::string expect4 = "A";
	assertionResult(get_middle(str4), expect4);



	std::cin.get();
	return 0;
}