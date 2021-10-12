#include <iostream>
#include <string>
#include <vector>

/*
Examples:
"the-stealth-warrior" gets converted to "theStealthWarrior"
"The_Stealth_Warrior" gets converted to "TheStealthWarrior"
*/

void chooseDelimeter(const std::string text, std::string& delim)
{
	if (text.find("-") != std::string::npos)
		delim = "-";
	else if (text.find("_") != std::string::npos)
		delim = "_";
}

std::string to_camel_case(std::string text)
{
	if (text.length() == 0) return text;

	std::string delim;
	chooseDelimeter(text, delim);
	if (delim.empty()) return text;

	size_t start = 0;
	size_t end = text.find(delim);
//	if (end == std::string::npos) return text;
	
	std::vector<std::string> sources;
	while (end != std::string::npos)
	{
		sources.push_back(text.substr(start, end - start));
		start = end + delim.length();
		end = text.find(delim, start);
	}
	sources.push_back(text.substr(start, end));

	for (std::vector<std::string>::iterator it = sources.begin() + 1;
		it != sources.end(); it++)
	{
		if ((*it)[0] > 96 && (*it)[0] < 123)
		{
			(*it)[0] -= 32;
		}
	}

	std::string dest;
	for (std::vector<std::string>::iterator it = sources.begin();
		it != sources.end(); it++)
	{
		dest.append(*it);
	}
	
	return dest;
}

int main()
{
	std::string temp = to_camel_case("the-stealth-warrior");
	std::string temp2 = to_camel_case("a_pippi_isOmoshiroi");
	std::string temp3 = to_camel_case("theCat_Was_Savage");
	std::string temp4 = to_camel_case("the_Cat_Is_Omoshiroi");

	return 0;
}