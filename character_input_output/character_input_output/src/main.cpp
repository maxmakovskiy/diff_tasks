#include <stdio.h>


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


int main()
{
// 	copyInputToOutput();

	int count = countCharacters();
	printf("Number of characters from input: %d", count);

	getchar();
	return 0;
}