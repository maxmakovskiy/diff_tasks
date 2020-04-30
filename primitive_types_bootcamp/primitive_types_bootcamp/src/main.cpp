#include <iostream>


short countBits(unsigned long a)
{
	// a = 5 -> bin(a) = 1001
	short numOfBits = 0;

	while (a)
	{
		// 1001 & 0001 = 0001
		numOfBits += a & 1;
		// 1001 >> 1 = 0100
		a >>= 1;
		// and so on
	}

	return numOfBits;
}

int main()
{
	unsigned long a = 5; // 1001
	std::cout << "Num of bist in " << a << " -> " << countBits(a) << std::endl;
	
	unsigned long b = 15; // 1111
	std::cout << "Num of bist in " << b << " -> " << countBits(b) << std::endl;
	
	
	
	
	
	std::cin.get();
	return 0;
}