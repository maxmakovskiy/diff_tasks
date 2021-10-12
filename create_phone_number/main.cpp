#include <iostream>
#include <string>
#include <sstream>

std::string createPhoneNumber(const int arr[10])
{
	std::stringstream ss;
	ss << "(" << arr[0] << arr[1] << arr[2] << ") ";
	ss << arr[3] << arr[4] << arr[5] << "-";
	ss << arr[6] << arr[7] << arr[8] << arr[9];

	return ss.str();
}

int main()
{
	const int arr[10] = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 0 };
	std::cout << createPhoneNumber(arr) << std::endl;
	std::cin.get();
	return 0;
}