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

	int count = countLines();
	printf("Num of lines from given input: %d\n", count);
	getchar();

	system("pause");
	return 0;
}