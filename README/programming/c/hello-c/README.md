---
title: Hello C!
date: 2020-04-20T18:33:38+12:00
draft: false
tags: ["Hello World,", "C"]
---

Well I always new this day was going to come. Originally I was going to learn C++ and not C, but [rwxrob](https://gitlab.com/rwxrob) has convinced me otherwise. His reasoning? "The C programming language is mandatory learning for anyone who truly wants to understand how all languages and digital computers work." [[source]](https://gitlab.com/rwx.gg/README/-/tree/master/lang/c) For me, the reason why I decided on learning C is because I want to someday contribute to projects I use and love like the [Linux kernel](https://github.com/torvalds/linux), [Lynx](https://lynx.invisible-island.net), [Tmux](https://github.com/tmux/tmux), and so on.

As always lets start off with the classic *hello world!*. 

```C
#include <stdio.h>

int main() {
	puts("Hello world from C!");
	return 0;
}
```

Now if you happen to get the following error like me, then you need to install `build-essential`.

```
main.c:6:10: fatal error: stdio.h: No such file or directory
 #include <stdio.h>
          ^~~~~~~~~
compilation terminated.
```

It should output `Hello world from C!`.

---

Next for my favorite first program, a number guessing game!

```C
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <ctype.h>

int do_again() {
        int choice = 0;
        char again[2];
        printf("Play again? [y/N]: ");
        scanf("%1s", again);

        if(tolower(again[0]) == 'y') {
                choice++;
        }

        return choice;
}

int main() {
        srand(time(NULL));
        int rand_num = rand() % 10;
        int guessed;
        int guesses = 0;

        do {
                do {
                        printf("Guess: ");
                        scanf("%i", &guessed);
                        if(guessed != rand_num) {
                                if(guessed > rand_num) {
                                        puts("Guess was too high!");
                                } else {
                                        puts("Guess was too low!");
                                }
                                guesses++;
                        } else {
                                guesses++;
                                printf("You win! It took %i guesses.\n", guesses);
                        }
                } while(guessed != rand_num);
        } while(do_again() == 1);
        return 0;
}
```

The above code works, but I did run into the following error initially.

`You win! It took you 1118051169 guesses.`

I had a good laugh at this, and quickly found the issue which was simply changing 

`int guesses;` to `int guesses = 0;`
