# Headers

In [C++ 1](/post/cpp-1), my examples had both a system defined header (`<iostream>`) and a user defined header (`hello.h`). In this blog we are going to cover headers in more detail. Most of the information in this blog is from a book called [C++ in a nutshell](https://www.oreilly.com/library/view/c-in-a/059600298X/), any other sources will be listed following the information. 

**Standard Library Headers**

There are 51 headers in the standard library which are listed below.

* `<algorithm>`
	* Algorithms for copying, searching, sorting, and operating on iterators/containers.
* `<bitset>`
	* Class template to hold a fixed-size sequence of bits.
* `<cassert>`
	* C header containing runtime assertion-checking.
* `<cctype>`
	* C header containing character classification and case conversion.
* `<cerrno>`
	* C header containing error codes.
* `<cfloat>`
	* C header containing limits of floating-point types.
* `<ciso646>`
	* Empty header because C declarations are incorporated in the C++ langauge.
* `<climits>`
	* C header containing limits of integer.
* `<clocale>`
	* C header containing local-specific information.
* `<cmath>`
	* C header containing mathematical functions.
* `<complex>`
	* C header containing complex numbers.
* `<csetjmp>`
	* C header containing nonlocal goto.
* `<csignal>`
	* C header containing asynchronous signals.
* `<cstdarg>`
	* C header containing macros to implement functions which take variable length arguments.
* `<cstddef>`
	* C header containing miscellaneous standard definitions.
* `<cstdio>`
	* C header containing standard input and output.
* `<cstdlib>`
	* C header containing miscellaneous functions and related declarations.
* `<cstring>`
	* C header containing string-handling functions.
* `<ctime>`
	* C header containing date/time functions and types.
* `<cwchar>`
	* C header containing wide character functions, including I/O.
* `<cwctype>`
	* C header containing wide character classification and case conversion.
* `<deque>`
	* Deque (double-ended queue) standard container.
* `<exception>`
	* Base exception class and functions related to exception-handling.
* `<fstream>`
	* File-based stream I/O.
* `<functional>`
	* Function objects, typically used with standard algorithms.
* `<iomanip>`
	* I/O manipulators used with standard I/O streams.
* `<ios>`
	* Base class declarations for I/O objects.
* `<iosfwd>`
	* Forward declarations for I/O objects.
* `<iostream>`
	* Declarations of standard I/O objects.
* `<istream>`
	* Input streams and input/output streams.
* `<iterator>`
	* Additional iterators for working with standard containers and algorithms.
* `<limits>`
	* Limits of numerical types.
* `<list>`
	* Standard linked list container.
* `<locale>`
	* Locale-specific information for formatting and parsing numbers, dates, times, and currency values, plus character-related functions for classifying, converting, and comparing characters and strings.
* `<map>`
	* Associative map (aka a dictionary) standard container.
* `<memory>`
	* Allocators, algorithms for uninitialized memory, and smart pointers (auto_ptr).
* `<new>`
	* Global operator new/delete and other functions related to managing dynamic memory.
* `<numeric>`
	* Numerical algorithms.
* `<ostream>`
	* Output streams.
* `<queue>`
	* Queue and priority queue container adapters.
* `<set>`
	* Associative set container.
* `<sstream>`
	* String-based I/O streams.
* `<stack>`
	* Stack container adapter.
* `<stdexcept>`
	* Standard exception classes.
* `<streambuf>`
	* Low-level stream buffers used by high-level I/O streams.
* `<string>`
	* Strings and wide-character strings.
* `<strstream>`
	* String streams that work with character arrays.
* `<typeinfo>`
	* Runtime type information.
* `<utility>`
	* Miscellaneous templates, such as pair, most often used with standard containers and algorithms.
* `<valarray>`
	* Numerical arrays.
* `<vector>`
	* Vector (array-like) standard container.

There are more headers in the standard library if you are using C++ 11, which are listed below. There are even more included in newer versions of C++, but I use C++ 14 so they dont apply to me. [[source]](https://en.cppreference.com/w/cpp/header)

* `<array>`
	* Static contiguous array.
* `<atomic>`
	* Atomic class template and specializations for bool, integral, and pointer types.
* `<cfenv>`
	* Floating-point environment access functions.
* `<chrono>`
	* C++ time utilities.
* `<cinttypes>`
	* Formatting macros, `intmax_t` and `uintmax_t` math and conversions.
* `<codecvt> (deprecated in C++ 17)`
	* Converts between character encodings, including UTF-8, UTF-16, and UTF-32.
* `<condition_variable>`
	* Thread waiting conditions.
* `<cstdint>`
	* Fixed-size types and limits of other types.
* `<cuchar>`
	* C-style unicode character conversion functions.
* `<forward_list>`
	* Singly-linked list
* `<future>`
	* Primitives for asynchronous computations
* `<initializer_list>`
	* Lightweight proxy object that provides access to an array of objects of type const T.
* `<mutex>`
	* Mutual exclusion primitives.
* `<random>`
	* Random number generators and distributions.
* `<ratio>`
	* Compile-time rational arithmetic.
* `<regex>`
	* Classes, algorithms and iterators to support regular expression processing.
* `<scoped_allocator>`
	* Nested allocator class.
* `<system_error>`
	* Exception class used to report conditions that have an error_code.
* `<thread>`
	* Manages a seperate thread.
* `<typeindex>`
	* Wrapper around a type_info object, that can be used as index in associative and unordered associative containers.
* `<type_traits>`
	* Compile-time type information.
* `<tuple>`
	* Implements fixed size container, which holds elements of possibly different types.
* `<unordered_set>`
	* Collection of unique keys, hashed by keys.
* `<unordered_map>`
	* Collection of key-value pairs, hashed by keys, keys are unique.
