#include <stdio.h>

int main()
{
	int c;

	c = getchar();	
	while(c != EOF)
	{
		if (c != ' ')
		{
			putchar(c-32);		
		}		

		c = getchar();
	}	

	return 0;
}
