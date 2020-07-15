---
title: C++ 3 ~ Basics
date: 2019-07-23T08:37:58+12:00
tags: ["C++,", "Fundamentals"]
draft: false
---

In the previous blogs I have just dived straight into programming in C++. This blog is going to focus around getting a better understanding of the language itself, including the various components that make up the language. It will cover the basics so for an experienced C++ programmer this blog will not be very interesting. If you are like me and new to C++ however, this blog may be beneficial.

---

## Standard C++ File Structure

I found a great stackoverflow post on [directory structure for a C++ library](https://stackoverflow.com/a/1398594), which covers this topic. According to the post the file structure is as follows.

* `./`          ~ Makefile and configure scripts
* `./src`       ~ General sources
* `./include`   ~ Header files that expose the public interface and are to be installed
* `./lib`       ~ Library build directory
* `./bin`       ~ Tools build directory
* `./tools`     ~ Tools sources
* `./test`      ~ Test suites that should be run during *make test*

It is good to note what the person states about *choose a directory layout that makes sense for you (and your team). Do whatever is most sensible for your chosen development environment, build tools, and source control.*

---

## Fundamental Types

[Sourced from cppreference fundamental types](https://en.cppreference.com/w/cpp/language/types).

* `void` ~ empty set of values. Objects of type void are disallowed. There are no arrays of void, nor references of void. Pointers to void and functions returning type void are permitted.

* `std::nullptr_t` ~ null pointer literal, nullptr. Not itself a pointer type or a pointer to member type.

* `bool` ~ capable of holding one of two values: true or false. The value is implementation defined and might differ from 1.

* `int` ~ basic integer type. If no length modifiers are present, it's guaranteed to have a width of at least 16 bits. However, on 32 / 64 bit systems it is almost exclusively guaranteed to have width of at least 32 bits.
    * **Modifiers**
        * ***Signedness***
            * `signed` ~ target type will have signed representation (default if omitted).
            * `unsigned` ~ target type will have unsigned representation.

        * ***Size***
            * `short` ~ target type will be optimized for space and will have width of at least 16 bits.
            * `long` ~ target type will have width of at least 32 bits.

---

* `signed char` ~ type for signed character representation.

* `unsigned char` ~ type for unsigned character representation. 

* `char` ~ type for character representation which can be most efficiently processed on the target system. 

* `wchar_t` ~ type for wide character representation. Required to be large enough to represent any supported character code point. It has the same size, signedness, and alignment as one of the integer types, but is a distinct type.

* `float` ~ single precision floating point type. Usually IEEE-754 32 bit floating point type.

* `double` ~ double precision floating point type. Usually IEEE-754 64 bit floating point type.

* `long double` ~ extended precision floating point type. Does not necessarily map to types mandated by IEEE-754. Usually 80-bit x87 floating point type on x86 and x86-64 architectures.

---

## Assignment Operators

[Sourced from cppreference assignment operators](https://en.cppreference.com/w/cpp/language/operator_assignment)

* `simple assignment` ~ `a = b`
* `addition assignment` ~ `a += b`
* `subtraction assignment` ~ `a -= b`
* `multiplication assignment` ~ `a *= b`
* `division assignment` ~ `a /= b`
* `modulo assignment` ~ `a %= b`
* `bitwise AND assignment` ~ `a &= b`
* `bitwise OR assignment` ~ `a |= b`
* `bitwise XOR assignment` ~ `a ^= b`
* `bitwise left shift assignment` ~ `a <<= b`
* `bitwise right shift assignment` ~ `a >>= b`

---

## Increment/Decrement Operators

[Sourced from cppreference increment/decrement operators](https://en.cppreference.com/w/cpp/language/operator_incdec)

* `pre-increment` ~ `++a`
* `pre-decrement` ~ `--a`
* `post-increment` ~ `a++`
* `post-decrement` ~ `a--`

---

## Arithmetic Operators

[Sourced from cppreference arithmetic operators](https://en.cppreference.com/w/cpp/language/operator_arithmetic)

* `unary plus` ~ `+a`
* `unary minus` ~ `-a`
* `addition` ~ `a + b`
* `subtraction` ~ `a - b`
* `multiplication` ~ `a * b`
* `division` ~ `a / b`
* `modulo` ~ `a % b`
* `bitwise NOT` ~ `~a`
* `bitwise AND` ~ `a & b`
* `bitwise OR` ~ `a | b`
* `bitwise XOR` ~ `a ^ b`
* `bitwise left shift` ~ `a << b`
* `bitwise right shift` ~ `a >> b`

---

## Logical Operators

[Source from cppreference logical operators](https://en.cppreference.com/w/cpp/language/operator_logical)

* `negation` ~ `!a`
* `AND` ~ `a && b`
* `inclusive OR` ~ `a || b`

---

## Comparison Operators

[Sourced from cppreference comparison operators](https://en.cppreference.com/w/cpp/language/operator_comparison)

* `equal to` ~ `a == b`
* `not equal to` ~ `a != b`
* `less than` ~ `a < b`
* `greater than` ~ `a > b`
* `less than or equal to` ~ `a <=b`
* `greater than or equal to` ~ `a >= b`
* `three-way comparison` ~ `a <=> b`

---

## Member Access Operators

[Sourced from cppreference member access operators](https://en.cppreference.com/w/cpp/language/operator_member_access)

* `subscript` ~ `a[b]`
* `indirection` ~ `*a`
* `address-of` ~ `&a`
* `member of object` ~ `a.b`
* `member of pointer` ~ `a->b`
* `pointer to member of object` ~ `a.*b`
* `pointer to member of pointer` ~ `a->*b`

---

## Other Operators

[Sourced from cppreference other operators](https://en.cppreference.com/w/cpp/language/operator_other)

* `function call` ~ `a(a1, a2)`
* `comma` ~ `a, b`
* `conditional operator` ~ `a ? b: c`

---

## If Statements

```C++
if(a > b) {
    // ...
}
```

```C++
if(a > b) {
    // ...
} else {
    // ...
}
```

```C++
if(a > b) {
    // ...
} else if (a < b) {
    // ...
} else {
    // ...
}
```

```C++
int result = (a > b) ? true : false;
```

---

## Switch Statements

```C++
int op = 3;
switch(op) {
    case 0:
        cout << "+";
    case 1:
        cout << "-";
    case 2:
        cout << "*";
    case 3:
        cout << "/";
    case 4:
        cout << "%";
}
```

---

## Loop Statements

```C++
int i = 0;
while(i < 7) {
    i++;
}
```

```C++
int i = 0;
do {
    i++;
} while(i < 7);
```

```C++
for(int i = 0; i < 7; i++) {
    // ...
}
```

---

## Pointers and References

```C++
int v = 42;
int* ptr = &v;

std::cout << ptr << std::endl; // Outputs the memory address of v (0x7ffdf796773c)
std::cout << *ptr << std::endl; // Outputs the value of v
```

```C++
int v = 42;

std::cout << &x << std::endl; // Outputs the memory address of x (0x7fffe6d57bb4)
```

---

## Classes and Objects

```C++
#include <iostream>

// Base class
class Person
{
    private: // Access specifier

        std::string name;   // (string) Attribute
        int age;            // (int) Attribute

    public: // Access specifier

        /* Getters and Setters */
        void setName(std::string n)
        {
            this->name = n;
        }

        std::string getName()
        {
            return this->name;
        }

        void setAge(int a)
        {
            this->age = a;
        }

        int getAge()
        {
            return this->age;
        }

        /* Methods */
        void printInfo(); // Method declaration inside the class

        void speak(std::string word) // Method definition inside the class
        {
            std::cout << word << std::endl;
        }

        /* Constructors */
        Person(){}; // Default constructor

        // Constructor declaration inside the class
        Person(std::string name, int age); 
};

// Constructor definition outside the class
Person::Person(std::string name, int age)
{
    this->setName(name);
    this->setAge(age); //age = age; // Causes bug
}

// Method definition outside the class
void Person::printInfo()
{
    std::cout << "My name is " << getName() << std::endl;
    std::cout << "My age is " << getAge() << std::endl;
}

// Subclass that inherits from Person
class Programmer: public Person
{
    private:
        std::string language;

    public:
        /* Getters and Setters */
        void setLanguage(std::string l)
        {
            this->language = l;
        }

        std::string getLanguage()
        {
            return this->language;
        }

        /* Methods */
        void cry()
        {
            std::cout << "Why wont you work :*(\n";
        }
        
        void printInfo()
        {
            std::cout << "My name is " << getName() << std::endl;
            std::cout << "My age is " << getAge() << std::endl;
            std::cout << "My programming language is " << getLanguage();
        }

        /* Constructors */
        Programmer(){};
        Programmer(std::string name, int age, std::string language)
        {
            this->setName(name);
            this->setAge(age);
            this->setLanguage(language);
        };
};

int main()
{
    Person manager("Bob", 23); // Person object using constructor with params
    manager.speak("Hello world!");
    manager.printInfo();

    Programmer cppProgrammer; // Programmer object using default constructor
    cppProgrammer.setName("Alice");
    cppProgrammer.setAge(33);
    cppProgrammer.setLanguage("C++");
    cppProgrammer.printInfo();
    cppProgrammer.cry();

    return 0;
}
```
