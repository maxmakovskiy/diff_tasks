#include <stdio.h>
#include <stdlib.h>


void copyInputToOutput()
{
	int c;

	while ((c = getchar()) != EOF)
	{
		putchar(c);
	}
}

int countCharacters()
{
	int numOfChars;

//	for (numOfChars = 0; getchar() != EOF; ++numOfChars);
	for (numOfChars = 0; getchar() != '\n'; ++numOfChars);

	return numOfChars;
}

int* countLinesBlanksTabs()
{
	static int storage[3] = { 0, 0, 0 };
	int c;

	while ((c = getchar()) != EOF)
	{
		if (c == '\n')
		{
			storage[0]++;
		}
		else if (c == '\t')
		{
			storage[1]++;
		}
		else if (c == ' ')
		{
			storage[2]++;
		}
	}
	
	return storage;
}

int countLines()
{
	int numOfLines = 0;
	int c;

//	while ((c = getchar()) != EOF)
	while ((c = getchar()) != '-')
	{
		if (c == '\n')
		{
			++numOfLines;
		}
	}

	return numOfLines;
}

int main()
{
// 	copyInputToOutput();

/*	int count = countCharacters();
	printf("Number of characters from input: %d", count); */

	/*int count = countLines();
	printf("Num of lines from given input: %d\n", count);
	getchar();*/

	int* counterp = countLinesBlanksTabs();
	printf("Num of lines %d\nNum of tabs %d\nNum of blanks %d\n", *counterp, *(counterp + 1), *(counterp + 2));



	system("pause");
	return 0;
}