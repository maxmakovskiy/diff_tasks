#include <iostream>
#include <vector>


// https://www.codewars.com/kata/55466989aeecab5aac00003e/train/cpp
// sqInRect(5, 3) should return [3, 2, 1, 1]
std::vector<int> sqInRoot(int lng, int wdth)
{
	if (lng == wdth)
	{
		return std::vector<int>();
	}

	std::vector<int> result;

	while (lng != wdth)
	{
		int b = lng > wdth ? lng : wdth;
		int s = lng < wdth ? lng : wdth;

		result.push_back(s);
		lng = s;
		wdth = b - s;
	}

	result.push_back(lng);

	return result;
}

int main() {
	
	for (const auto& el : sqInRoot(5, 3))
	{
		std::cout << "sqInRoot -> " << el << std::endl;
	}

	std::cin.get();
	return 0;
}