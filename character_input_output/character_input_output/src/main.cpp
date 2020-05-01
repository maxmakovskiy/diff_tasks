#include <stdio.h>
#include <stdlib.h>


#define IN 1
#define OUT 0

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

void replaceMultiBlankToSingle()
{
	int curr = 0;
	int prev = 0;

	for (; (curr = getchar()) != EOF; prev = curr)
	{
		if (curr == ' ' && prev == ' ')
		{
			continue;
		}

		putchar(curr);
	}
}

int countWords()
{
	int c;
	int numOfWords = 0;
	int state = OUT;

	while ((c = getchar()) != EOF)
	{
		if (c == ' ' || c == '\n' || c == '\t')
		{
			state = OUT;
		}
		else if (state == OUT)
		{
			state = IN;
			numOfWords++;
		}

	}

	return numOfWords;
}

int main()
{
// 	copyInputToOutput();

/*	int count = countCharacters();
	printf("Number of characters from input: %d", count); */

	/*int count = countLines();
	printf("Num of lines from given input: %d\n", count);
	getchar();*/

/*	int* counterp = countLinesBlanksTabs();
	printf("Num of lines %d\nNum of tabs %d\nNum of blanks %d\n", *counterp, *(counterp + 1), *(counterp + 2));*/

//	replaceMultiBlankToSingle();

	int count = countWords();
	printf("Num of words from given input: %d\n", count);
	

	system("pause");
	return 0;
}