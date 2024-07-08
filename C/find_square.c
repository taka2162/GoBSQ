#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define TRUE 1
#define FALSE 0
#define	MAP_SIZE 10

void	__print_map(int **map)
{
	int	i;
	int	j;

	i = 0;
	while (i < MAP_SIZE + 2)
	{
		j = 0;
		while (j < MAP_SIZE + 2)
		{
			if (0 <= map[i][j])
				printf(" ");
			if (map[i][j] == -2)
				printf("\x1b[33m");
			else if (map[i][j] == -1)
				printf("\x1b[35m");
			printf("%d\x1b[m", map[i][j]);
			j++;
		}
		printf("\n");
		i++;
	}
}

int	calculate_square_size(int **map, int i, int j, int flag)
{
	int	y;
	int	x;
	int	size;

	size = 1;
	while (1)
	{
		y = 0;
		while (y < size)
		{
			x = 0;
			while (x < size)
			{
				if (map[i + y][j + x] < 0)
					return (size - 1) ;
				x++;
			}
			y++;
		}
		size++;
	}
	return (size);
}

void	print_max_size(int **map, int i, int j)
{
	int	size;
	int	y;
	int	x;

	size = map[i][j];
	y = 0;
	while (y < size)
	{
		x = 0;
		while (x < size)
		{
			map[i + y][j + x] = -3;
			x++;
		}
		y++;
	}
}

void	write_square_size(int **map)
{
	int	i;
	int	j;
	int	max;
	int	index[2];

	max = 0;
	i = 1;
	while (map[i][1] != -2)
	{
		j = 1;
		while (map[i][j] != -2)
		{
			if (map[i][j] == 0)
			{
				map[i][j] += calculate_square_size(map, i, j, FALSE);
				if (max < map[i][j])
				{
					max = map[i][j];
					index[0] = i;
					index[1] = j;
				}
			}
			j++;
		}
		i++;
	}
	__print_map(map);
	printf("index[0] = %d  index[1] = %d\n", index[0], index[1]);
	print_max_size(map, index[0], index[1]);
}


int	main(void)
{
	int	**map;
	unsigned int	seed;
	int	i;
	int	j;

	seed = (unsigned int)time(NULL);
	srand(seed);
	map = (int **)malloc(sizeof(int *) * MAP_SIZE + 2);
	i = 0;
	while (i < MAP_SIZE + 2)
	{
		map[i] = (int *)malloc(sizeof(int) * MAP_SIZE + 2);
		i++;
	}
	i = 0;
	while (i < MAP_SIZE + 2)
	{
		j = 0;
		while (j < MAP_SIZE + 2)
		{
			if (i == 0 || j == 0 || i == MAP_SIZE + 2 - 1 || j == MAP_SIZE + 2 - 1)
				map[i][j] = -2;
			else if (rand() % 7 == 1)
				map[i][j] = -1;
			else
				map[i][j] = 0;
			j++;
		}
		i++;
	}
	__print_map(map);
	write_square_size(map);
	__print_map(map);
	i = 0;
	while (i < MAP_SIZE + 2)
	{
		free(map[i]);
		i++;
	}
	free(map);
	return (0);
}