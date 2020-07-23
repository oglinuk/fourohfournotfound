---
title: Hello C++!
date: 2019-07-21T06:17:33+12:00
tags: ["Hello World,", "C++"]
draft: false
---

For a while I've been meaning to dive into a lower language such as C++, but never found the time to do so. Now I have an excuse as my new employment involves working with C++. To start, like with anything I made a simple *hello world* program.

```C++
#include <iostream>

int main() {
    std::cout << "Hello world from C++!\n";
    return 0;
}
```

Now to compile the program via `g++ main.cpp -o main` and run it via `./main`.

The output is `Hello world from C++!`.

The next program I made was a number guessing game, similar to the one I made when first learning [Python](/python).

```C++
#include <iostream>
#include <time.h>

void play() {
    srand(time(NULL)); // Setting the seed
    int32_t rngNum = rand() % 10; // Generate random number between 0-10
    int32_t guesses = 0;
    int32_t guess = -1;

     while(guess != rngNum) {
        guess = getGuess();

        if (guess != rngNum) {
            if (guess > rngNum) {
                std::cout << "Guess was too high!\n";
            } else {
                std::cout << "Guess was too low!\n";
            }
            guesses++;
        } else {
            guesses++;
            std::cout << "You win! It took you " << guesses << " guesses.\n";
        }
    };
}

int32_t getGuess() {
    int32_t guess = -1;

    std::cout << "Guess: ";
    std::cin >> guess;

    while(std::cin.fail()) {
        std::cout << "Please enter a number ...\n";
        std::cin.clear(); // Unset failbit; see std::cin.fail()

        // https://en.cppreference.com/w/cpp/io/basic_istream/ignore
        std::cin.ignore(256,'\n');
    }

    return guess;
}

bool again() {
    char ans = 'n';
    
    std::cout << "Play again? [y/N]: ";
    std::cin >> ans;

    if(std::tolower(ans) == 'y') {
        return true;
    }
            
    system("clear");
    return false;
}

int main()
{
    bool doAgain;

    while(doAgain == 1) {
        system("clear");
        play();
        doAgain = again();
    };

    return 0;
}
```

One thing that I noticed was the return value of `Again` is either a 0 (false) or 1 (true), which confused me at first since I was originally checking `doAgain` to either be `true` or `false`.

Nothing complicated, but still interesting and fun! These programs can be found on my [ptp](https://github.com/OGLinuk/ptp) repository.
