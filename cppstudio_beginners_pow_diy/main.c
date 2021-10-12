#include <stdio.h>

int main()
{
	int power, number, i, result;

	printf("Enter number: ");
	scanf("%d", &number); 
	printf("Enter power: ");
	scanf("%d", &power);

	result = number;
	for (i = 1; i < power; i++)
	{
		result *= number;
	}
	
	printf("Result: %d", result);

	return 0;
}
