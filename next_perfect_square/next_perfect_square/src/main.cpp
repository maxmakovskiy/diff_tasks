#include <iostream>
#include <math.h>


bool isPerfectSquare(long int sq)
{
	double squareRoot = sqrt(sq);
	double floatingPart = squareRoot - floor(squareRoot);

	if (floatingPart == 0.0) {
		return true;
	}

	return false;
}

// findNextSquare(121)-- > returns 144
// findNextSquare(625)-- > returns 676
// findNextSquare(114)-- > returns - 1 since 114 is not a perfect
// A perfect square is a number that can be expressed as the product of two equal integers.
long int findNextSquare(long int sq) 
{
	return -1;
}

int main()
{
	std::cout << isPerfectSquare(9) << std::endl;
	std::cin.get();

	return 0;
}