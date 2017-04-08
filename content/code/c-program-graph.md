+++
tags = ["c"]
categories = ["code"]
title = "C Program Graph"
draft = true
+++

### Create a Random Graph Using Random Edge Generation		

 Code Sample 
```c
#include<stdio.h>
#include<stdlib.h>
#include <time.h>

#define MAX_VERTICES 30
#define MAX_EDGES 10

typedef unsigned char vertex;

int main(){

    /*number of nodes in a graph*/
    srand ( time(NULL) );
    int numberOfVertices = rand() % MAX_VERTICES;

    /*number of maximum edges a vertex can have*/
    srand ( time(NULL) );
    int maxNumberOfEdges = rand() % MAX_EDGES;
    /*graphs is 2 dimensional array of pointers*/
    if( numberOfVertices == 0)
        numberOfVertices++;
    vertex ***graph;
    printf("Total Vertices = %d, Max # of Edges = %d\n",numberOfVertices, maxNumberOfEdges);

    /*generate a dynamic array of random size*/
    if ((graph = (vertex ***) malloc(sizeof(vertex **) * numberOfVertices)) == NULL){
        printf("Could not allocate memory for graph\n");
        exit(1);
    }

    /*generate space for edges*/
    int vertexCounter = 0;
    /*generate space for vertices*/
    int edgeCounter = 0;

    for (vertexCounter = 0; vertexCounter < numberOfVertices; vertexCounter++){
        if ((graph[vertexCounter] = (vertex **) malloc(sizeof(vertex *) * maxNumberOfEdges)) == NULL){
            printf("Could not allocate memory for edges\n");
            exit(1);
        }
        for (edgeCounter = 0; edgeCounter < maxNumberOfEdges; edgeCounter++){
            if ((graph[vertexCounter][edgeCounter] = (vertex *) malloc(sizeof(vertex))) == NULL){
                printf("Could not allocate memory for vertex\n");
                exit(1);
            }
        }
    }

    /*start linking the graph. All vetrices need not have same number of links*/
    vertexCounter = 0;edgeCounter = 0;
    for (vertexCounter = 0; vertexCounter < numberOfVertices; vertexCounter++){
        printf("%d:\t",vertexCounter);
        for (edgeCounter=0; edgeCounter < maxNumberOfEdges; edgeCounter++){
            if (rand()%2 == 1){ /*link the vertices*/
                int linkedVertex = rand() % numberOfVertices;
                graph[vertexCounter][edgeCounter] = graph[linkedVertex];
                printf("%d, ", linkedVertex);
            }
            else{ /*make the link NULL*/
                graph[vertexCounter][edgeCounter] = NULL;
            }
        }
        printf("\n");
    }
    return 1;
}
```

 Output 
```bash

$ gcc  random_graph.c 
$ ./a
Total Vertices = 9, Max # of Edges = 9
0:      3, 7, 2, 6, 2,
1:      7, 7, 4, 3,
2:      7, 7, 6, 1, 1,
3:      8, 0, 2, 5,
4:      2, 1, 4, 4, 5, 5,
5:      1, 2, 5,
6:      7, 2, 1, 5,
7:      0, 2, 2, 3, 4, 3,
8:      0, 5, 0, 5,
```

### Represent Graph Using Adjacency Matrix		

 Code Sample 
```c
//... A Program to represent a Graph by using an Adjacency Matrix method

void read_graph ( int adj_mat[50][50], int n )
{
    int i, j;
    char reply;
    for ( i = 1 ; i <= n ; i++ )
    {
        for ( j = 1 ; j <= n ; j++ )
        {
            if ( i == j )
            {
                adj_mat[i][j] = 0;
		        continue;
            }
            
            printf("\n Vertices %d & %d are Adjacent ? (Y/N) :",i,j);
            scanf(" %c", &reply);
            if ( reply == 'y' || reply == 'Y' )
                adj_mat[i][j] = 1;
            else
                adj_mat[i][j] = 0;
	    }
    } 
    // return;
}



void dir_graph()
{
    int adj_mat[50][50];
    int n;
    int in_deg, out_deg, i, j;
    printf("\n How Many Vertices ? : ");
    scanf("%d", &n);
    read_graph(adj_mat, n);
    printf("\n Vertex \t In_Degree \t Out_Degree \t Total_Degree ");
    for (i = 1; i <= n ; i++ )
    {
        in_deg = out_deg = 0;
        for ( j = 1 ; j <= n ; j++ )
        {
                if ( adj_mat[j][i] == 1 )
                    in_deg++;
        } 
        for ( j = 1 ; j <= n ; j++ )
            if (adj_mat[i][j] == 1 )
                out_deg++;
            printf("\n\n %5d\t\t\t%d\t\t%d\t\t%d\n\n",i,in_deg,out_deg,in_deg+out_deg);
    }
    // return;
}

void undir_graph()
{
    int adj_mat[50][50];
    int deg, i, j, n;
    printf("\n How Many Vertices ? : ");
    scanf("%d", &n);
    read_graph(adj_mat, n);
    printf("\n Vertex \t Degree ");
    for ( i = 1 ; i <= n ; i++ )
    {
        deg = 0;
        for ( j = 1 ; j <= n ; j++ )
            if ( adj_mat[i][j] == 1)
                deg++;
        printf("\n\n %5d \t\t %d\n\n", i, deg);
    } 
    // return;
} 


void main()
{
   int option;
   do
   {    	
        printf("\n A Program to represent a Graph by using an ");
        printf("Adjacency Matrix method \n ");
        printf("\n 1. Directed Graph ");
        printf("\n 2. Un-Directed Graph ");
        printf("\n 3. Exit ");
        printf("\n\n Select a proper option : ");
        scanf("%d", &option);
        switch(option)
        {
                case 1 : 
                    dir_graph();
                    break;
                case 2 : 
                    undir_graph();
                    break;
                case 3 : 
                    exit(0);
        } // switch
    }while(1);
}
```

 Output 
```bash

$ gcc  graph.c 
$ ./graph
 A Program to represent a Graph by using an Adjacency Matrix method 
1.  Directed Graph 
2.  Un-Directed Graph 
3.  Exit 

 Select a proper option : 
 How Many Vertices ? : 1
 
 Vertices 1 & 2 are Adjacent ? (Y/N):  N
 Vertices 1 & 3 are Adjacent ? (Y/N):  Y
 Vertices 1 & 4 are Adjacent ? (Y/N):  Y
 Vertices 2 & 1 are Adjacent ? (Y/N):  Y
 Vertices 2 & 3 are Adjacent ? (Y/N):  Y 
 Vertices 2 & 4 are Adjacent ? (Y/N):  N 
 Vertices 3 & 1 are Adjacent ? (Y/N):  Y 
 Vertices 3 & 2 are Adjacent ? (Y/N):  Y
 Vertices 3 & 4 are Adjacent ? (Y/N):  Y
 Vertices 4 & 1 are Adjacent ? (Y/N):  Y 
 Vertices 4 & 2 are Adjacent ? (Y/N):  N 
 Vertices 4 & 3 are Adjacent ? (Y/N):  Y
 Vertex 	 In_Degree 	 Out_Degree 	 Total_Degree 
    1           2           0               2
    2           1           2               3
    3           0           1               1
    4           1           1               2


 A Program to represent a Graph by using an Adjacency Matrix method 
1.  Directed Graph 
2.  Un-Directed Graph 
3.  Exit

```


### Represent Graph Using Incidence Matrix		

 Code Sample 
```c
/*
* C Program to Describe the Representation of Graph using Incidence Matrix
*/
#include<stdio.h>

struct node
{
    int from, to;
} a[5], t;

void addEdge(int am[][5], int i, int j, int in)
{
    int p, q;
    a[in].from = i;
    a[in].to = j;
    for ( p = 0; p <= in; p++)
    {
        for (q = p + 1; q <= in; q++)
        {
            if (a[p].from > a[q].from)
            {
                t = a[p];
                a[p] = a[q];
                a[q] = t;
            }
            else if (a[p].from == a[q].from)
            {
                if (a[p].to > a[q].to)
                {
                    t = a[p];
                    a[p] = a[q];
                    a[q] = t;
                }
            }
            else
            {
                continue;
            }
        }
    }
}
int main()
{
    int n, c = 0, x, y, ch, i, j;
    int am[5][5];
    printf("Enter the no of vertices\n");
    scanf("%d", &n);

    for (i = 0; i < 5; i++)
    {
        for (j = 0; j < 5; j++)
        {
            am[i][j] = 0;
        }
    }
    while (ch != -1)
    {
        printf("Enter the nodes between which you want to introduce edge\n");
        scanf("%d%d", &x, &y);
        addEdge(am, x, y, c);
        c++;
        printf("Press -1 to exit 0 to proceed \n");
        scanf("%d", &ch);
    }    
    for (j = 0; j < c; j++)
    {
        am[a[j].from][j] = 1;
        am[a[j].to][j] = 1;
    }
    for (i = 0; i < n; i++)
    {
        for (j = 0; j < c; j++)
        {
            printf("%d\t" ,am[i][j]);
        }
        printf("\n");
    } 
}
```

 Output 
```bash

$ gcc   c-program-represent-graph-incidence-matrix.c 
$ ./a
Enter the no of vertices                                                          
5                                                                                 
Enter the nodes between which you want to introduce edge                          
0                                                                                 
1                                                                                 
Press -1 to exit 0 to proceed                                                                  
0                                                                                 
Enter the nodes between which you want to introduce edge                          
0                                                                                 
2                                                                                 
Press -1 to exit 0 to proceed                                                                 
0                                                                                 
Enter the nodes between which you want to introduce edge                          
2                                                                                 
3                                                                                 
Press -1 to exit 0 to proceed                                                                  
0                                                                                 
Enter the nodes between which you want to introduce edge                          
1                                                                                 
4                                                                                 
Press -1 to exit 0 to proceed                                                                  
0                                                                                 
Enter the nodes between which you want to introduce edge                          
0                                                                                 
3                                                                                 
Press -1 to exit 0 to proceed                                                                  
-1                                                                                
1       1       1       0       0                                                 
1       0       0       1       0                                                 
0       1       0       0       1                                                 
0       0       1       0       1                                                 
0       0       0       1       0                                                 
```
