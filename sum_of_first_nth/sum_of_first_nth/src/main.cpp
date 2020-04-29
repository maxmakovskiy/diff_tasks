#include <iostream>
#include <vector>
#include <string>

// Truncate double to two decimal place
double truncate(double num)
{
	return static_cast<double>(static_cast<int>(num * 100.)) / 100.;
}

std::string truncString(const std::string& s)
{
	std::string result;

	for (const auto& c : s)
	{
		if (c != '0')
		{
			result += c;
		}
	}

	return result;
}

std::vector<double> getSeriesOfDenoms(int n)
{
	std::vector<double> result;
	std::vector<int> temp;
	int value = 1.0;

	for (int i = 1; i < n; i++)
	{
		value += 3.0;
		temp.push_back(value);
	}

	for (const auto& el : temp)
	{
		result.push_back(1.0 / el);
	}

	return result;
}

std::string seriesSum(int n)
{
	double result = 1.0;

	for (const auto& el : getSeriesOfDenoms(n))
	{
		result += el;
	}

	double temp = truncate(result);
	return truncString(std::to_string(temp));
}

int main()
{
	std::cout << seriesSum(5) << std::endl;


	std::cin.get();
	return 0;
}