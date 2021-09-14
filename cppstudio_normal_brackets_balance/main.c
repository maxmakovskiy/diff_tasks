#include <stdio.h>

int main()
{
	int balance, character;
	
	balance = 0;
	while((character = getchar()) != EOF)
	{
		if (character == '{')
			balance++;	
		else if (character == '}')
			balance--;	
	}	
	
	if (balance)
		printf("Balance was ruined\n");
	else
		printf("Balance is OK");

	return 0;
}
