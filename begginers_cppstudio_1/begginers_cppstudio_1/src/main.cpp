#include <iostream>

int main()
{
	int number;

	std::cout << "Input your number: " << std::endl;
	std::cin >>	number;

	std::cout << "1st digit " << number / 10000 << std::endl;
	std::cout << "2nd digit " <<  (number / 1000)%10 << std::endl;
	std::cout << "3rd digit " <<  (number / 100)%10 << std::endl;
	std::cout << "4th digit " <<  (number / 10)%10 << std::endl;
	std::cout << "5th digit " <<  number % 10 << std::endl;


	return 0;
}