#include <iostream>

#define ALLOCSIZE 10000

static char allocbuf[ALLOCSIZE];
static char* allocp = allocbuf;

char* alloc(int n)
{
	if (allocbuf + ALLOCSIZE - allocp >= n)
	{
		allocp += n;
		return allocp - n;
	}
	else
	{
		return NULL;
	}
}

void afree(char* p)
{
	if ((p >= allocbuf) && (p < allocbuf + ALLOCSIZE))
	{
		allocp = p;
	}
}

int main() 
{
	char* addr;
	for (int i = 0; i < 5; i++)
	{
		addr = alloc(2000);
	}
	
	afree(addr);

	addr = alloc(2000);

	
	return 0;
}