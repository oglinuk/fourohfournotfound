---
title: Shell 4 ~ Manipulating Files and Directories
draft: true
---

This chapter covers *manipulating files and directories* with the commands; `cp`, `mv`, `mkdir`, `rm`, and `ln`. All of these commands are ones that any person using the shell will have come across and use daily (except for perhaps `ln`). I'm glad to see the book covers why command line is much better than using a GUI. The answer is simple; ***power && flexibility***. One example it shows is when you want to copy all files of a certain file type from one directory to another, but only the ones that don't exist in the destination (in the example its HTML files). To do this from the command line, simply run `cp -u *.html destination`.

This is a good opprotunity to talk about a particularly amazing feature; `*` aka *wildcards* or *globbing*. There are a number of wildcards available which are listed below along with their meaning.

---

**Wildcards**

* `*`: Matches any characters
* `?`: Matches any single character
* `[characters]`: Matches any character that is a member of the set
*characters*
* `[!characters]`: Matches any character that is not a member of the set
*characters*
* `[[:class:]]`: Matches any character that is a member of the specified
*class*

**Character classes**

* `[:alnum:]`: Matches any alphanumeric character
* `[:alpha:]`: Matches any alphabetic character
* `[:digit:]`: Matches any numeral
* `[:lower:]`: Matches any lowercase letter
* `[:upper:]`: Matches any uppercase letter

**Examples**

* `g*`: Matches any file beginning with `g`
* `b*.txt`: Matches any file beginning with `b` followed by any
characters and ending with `.txt`
* `Data???`: Matches any file beginning with `Data` followed by exactly
three characters
* `[abc]*`: Matches any file beginning with either an `a`, a `b`, or a
`c`
* `BACKUP.[0-9][0-9][0-9]`: Matches any file beginning with `BACKUP`
followed by exactly three numerals
* `[[:upper:]]*`: Matches any file beginning with an uppercase letter
* `[![:digit:]]*`: Matches any file not beginning with a numeral
* `*[[:lower:]123]`: Matches any file ending with a lowercase letter or
the numerals `1`, `2`, or `3`

---

## `mkdir`
