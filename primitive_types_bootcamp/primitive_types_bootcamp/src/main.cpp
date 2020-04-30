#include <iostream>

// The parity of a binary word is 1 
// if the number of 1s in the word is odd;
// otherwise, it is 0.
short parity(unsigned long a)
{ 
	// a = 0111
	short result = 0;

	while (a)
	{
		// 0111 & 0001 = 0001
		// 0001 ^ 0000 = 0001
		result ^= (a & 1);
		// 0111 >> 1 = 0011 
		a >>= 1;
		// and so on
	}

	return result;
}

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
	unsigned long b = 7; // 0111 


	std::cout << "Num of bist in " << a << " -> " << countBits(a) << std::endl;
	std::cout << "Num of bist in " << b << " -> " << countBits(b) << std::endl;

	std::cout << "\n==========================================\n" << std::endl;

	std::cout << "The parity of " << a << " is " << parity(a) << std::endl;
	std::cout << "The parity of " << b << " is " << parity(b) << std::endl;

	
	
	std::cin.get();
	return 0;
}