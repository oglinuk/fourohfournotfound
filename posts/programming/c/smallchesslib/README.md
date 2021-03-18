# SmallChessLib(SCL)

Have you ever wondered what the limits of a terminal is in terms of usage?
*What can be done inside a terminal*?

How about a game of [chess](https://gitlab.com/drummyfish/smallchesslib)? How?

`#include "smallchesslib.h"`

```C
// main.c
#include <stdio.h>
#include "smallchesslib.h"

void observe(char state)
{
        putchar(state);
}

int main()
{
   SCL_BoardState set_state = SCL_BOARD_START_STATE;
   SCL_SquareSet unset_state = SCL_SQUARE_SET_EMPTY;

   SCL_printBoard(set_state, observe, unset_state, SCL_PRINT_FORMAT_UTF8, 4, 1, 0);
   return 0;
}
```

```
   A B C D E F G H
8 ▒♜░♞▒♝░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒♟░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒▒░░▒▒░░▒▒
4 ▒▒░░▒▒░░▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░♙▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░♘▒♖
```

Mindblowing.

## Making The "First" Move

```C
// main.c
#include <stdio.h>

#include "smallchesslib.h"

void observe(char state)
{
        putchar(state);
}
        
int main()
{
   SCL_BoardState set_state = SCL_BOARD_START_STATE;
   SCL_SquareSet unset_state = SCL_SQUARE_SET_EMPTY;

   SCL_printBoard(set_state, observe, unset_state, SCL_PRINT_FORMAT_UTF8, 4, 1, 0);
   SCL_makeMove(set_state, SCL_stringToSquare("B1"), SCL_stringToSquare("C1"));
   SCL_printBoard(set_state, observe, unset_state, SCL_PRINT_FORMAT_UTF8, 4, 1, 0);
   return 0;
}
```

```
   A B C D E F G H
8 ▒♜░♞▒♝░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒♟░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒▒░░▒▒░░▒▒
4 ▒▒░░▒▒░░▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░♙▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░♘▒♖
   A B C D E F G H
8 ▒♜░♞▒♝░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒♟░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒▒░░▒▒░░▒▒
4 ▒▒░░▒▒░░▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░♙▒♙░♙▒♙░♙
1 ░♖▒▒░♘▒♕░♔▒♗░♘▒♖
```

Uh.

## Hello (Chess) World!


```C
// main.c
#include <stdio.h>

#include "smallchesslib.h"

void observe(char state)
{
        putchar(state);
}

int main()
{
   SCL_BoardState set_state = SCL_BOARD_START_STATE;
   SCL_SquareSet unset_state = SCL_SQUARE_SET_EMPTY;
   SCL_printBoard(set_state, observe, unset_state, SCL_PRINT_FORMAT_UTF8, 4, 1, 0);
   SCL_makeMove(set_state, SCL_stringToSquare("E2"), SCL_stringToSquare("E4"));
   SCL_makeMove(set_state, SCL_stringToSquare("C7"), SCL_stringToSquare("C5"));
   unset_state, SCL_PRINT_FORMAT_UTF8, 4, 1, 0);
   return 0;
}
```

```
   A B C D E F G H
8 ▒♜░♞▒♝░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒♟░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒▒░░▒▒░░▒▒
4 ▒▒░░▒▒░░▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░♙▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░♘▒♖
   A B C D E F G H
8 ▒♜░♞▒♝░♛▒♚░♝▒♞░♜
7 ░♟▒♟░░▒♟░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░♟▒▒░░▒▒░░▒▒
4 ▒▒░░▒▒░░▒♙░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░♙▒▒░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░♘▒♖
```


## Scholars Mate

```C
void observe(char state) {
	putchar(state);
}

void move(SCL_BoardState state, SCL_BoardState unset_state, char *from, char *to)
{
	SCL_makeMove(state, SCL_stringToSquare(from), SCL_stringToSquare(to));
	SCL_printBoard(state, observe, unset_state, SCL_PRINT_FORMAT_UTF8, 4, 1, 0);
}

int main()
{
	SCL_BoardState set_state = SCL_BOARD_START_STATE;
	SCL_SquareSet unset_state = SCL_SQUARE_SET_EMPTY;
	SCL_printBoard(set_state, observe, unset_state, SCL_PRINT_FORMAT_UTF8, 4, 1, 0);
	//SCL_makeMove(set_state, SCL_stringToSquare("E2"), SCL_stringToSquare("E4"));
	//SCL_makeMove(set_state, SCL_stringToSquare("C7"), SCL_stringToSquare("C5"));
	
	/* Scholars Mate */
	move(set_state, unset_state, "D7", "D5");
	move(set_state, unset_state, "D2", "D4");
	move(set_state, unset_state, "C8", "F5");
	move(set_state, unset_state, "G1", "F3");
	move(set_state, unset_state, "E8", "C6");
	move(set_state, unset_state, "F3", "E5");
	move(set_state, unset_state, "C6", "C2");
        /* Check Mate */
        return 0;
}
```

```
   A B C D E F G H
8 ▒♜░♞▒♝░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒♟░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒▒░░▒▒░░▒▒
4 ▒▒░░▒▒░░▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░♙▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░♘▒♖
   A B C D E F G H
8 ▒♜░♞▒♝░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒▒░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒♟░░▒▒░░▒▒
4 ▒▒░░▒▒░░▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░♙▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░♘▒♖
   A B C D E F G H
8 ▒♜░♞▒♝░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒▒░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒♟░░▒▒░░▒▒
4 ▒▒░░▒▒░♙▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░░▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░♘▒♖
   A B C D E F G H
8 ▒♜░♞▒▒░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒▒░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒♟░░▒♝░░▒▒
4 ▒▒░░▒▒░♙▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░░▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░♘▒♖
   A B C D E F G H
8 ▒♜░♞▒▒░♛▒♚░♝▒♞░♜
7 ░♟▒♟░♟▒▒░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒♟░░▒♝░░▒▒
4 ▒▒░░▒▒░♙▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒♘░░▒▒
2 ▒♙░♙▒♙░░▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░░▒♖
   A B C D E F G H
8 ▒♜░♞▒▒░♛▒▒░♝▒♞░♜
7 ░♟▒♟░♟▒▒░♟▒♟░♟▒♟
6 ▒▒░░▒♚░░▒▒░░▒▒░░
5 ░░▒▒░░▒♟░░▒♝░░▒▒
4 ▒▒░░▒▒░♙▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒♘░░▒▒
2 ▒♙░♙▒♙░░▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░░▒♖
   A B C D E F G H
8 ▒♜░♞▒▒░♛▒▒░♝▒♞░♜
7 ░♟▒♟░♟▒▒░♟▒♟░♟▒♟
6 ▒▒░░▒♚░░▒▒░░▒▒░░
5 ░░▒▒░░▒♟░♘▒♝░░▒▒
4 ▒▒░░▒▒░♙▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♙░░▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░░▒♖
   A B C D E F G H
8 ▒♜░♞▒▒░♛▒▒░♝▒♞░♜
7 ░♟▒♟░♟▒▒░♟▒♟░♟▒♟
6 ▒▒░░▒▒░░▒▒░░▒▒░░
5 ░░▒▒░░▒♟░♘▒♝░░▒▒
4 ▒▒░░▒▒░♙▒▒░░▒▒░░
3 ░░▒▒░░▒▒░░▒▒░░▒▒
2 ▒♙░♙▒♚░░▒♙░♙▒♙░♙
1 ░♖▒♘░♗▒♕░♔▒♗░░▒♖
```

Now that is beautiful ...