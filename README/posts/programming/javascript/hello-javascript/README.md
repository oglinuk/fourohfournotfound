---
title: Hello JavaScript!
date: 2020-04-22T17:25:41+12:00
draft: false
tags: ["Hello World,", "JavaScript"]
---

JavaScript is everywhere. Along side HTML and CSS, JavaScript is what powers most websites on the [world wide web](https://en.wikipedia.org/wiki/World_Wide_Web). I've known the very basics of JavaScript for a while now, but have never really spent time to learn it more than just simple functionality. For my learnings, I will be using another resource that was suggested by [rwxrob](https://gitlab.com/rwxrob) called [Eloquent JavaScript](https://eloquentjavascript.net/), which was written by yet another amazing human being; [Marijn Haverbeke](https://marijnhaverbeke.nl/). If you ever happen to read this blog Marijn, thank you. The book is freely available under the [Creative Commons Attribute-NonCommercial 3.0 license](https://creativecommons.org/licenses/by-nc/3.0/), and can be downloaded as a [pdf](https://eloquentjavascript.net/Eloquent_JavaScript.pdf). I would also like to thank the [325 backers](https://eloquentjavascript.net/backers3.html) who helped financially fund the book, one person who I actually know! If you ever read this blog [Belma Gaukrodger](https://belma.dev/), thank you!

Chapter 1 of the book covers `values`, `types`, and `operators`. Before it goes into that it covers *bits* and how to express numbers in binary. I have to say representing numbers in binary never really clicked until now which is kind of odd cause it seems simplistic. The number they use is 13, but I am going to challenge myself and represent the famous number 1729 for practice. 

```None
  1   1   0   1  1  0  0  0 0 0 1
1024 512 256 128 64 32 16 8 4 2 1
```

Make sense? If not then think of it this way. You add the values of a base 2 number line to sum a desired number (in this case 1729). Any numbers used to sum the number are represented by a 1, any that are notare represented by a 0. So to sum 1729 we add 1024 + 512 + 128 + 64 + 1, and so the binary representation of 1729 is `11011000001`. Lets do another, this time 7. This one is rather easy since its just 4 + 2 + 1, which means the binary representation is `111`. Cool stuff!

Another thing Eloquent JavaScript has clicked for me is the max values of *integers*, or the *number* type as its referred to in the book. In programming languages there are different types of integers ranging from int8 to normally int64. The max value of the int8 data type is 2^8 - 1, 2^16 - 1 for the int16, and so on up to 2^64 - 1 for int64. This is of course only if the integer is [*unsigned*](https://en.wikipedia.org/wiki/Integer_(computer_science)#Common_integral_data_types).

Something that I found rather interesting is the three *special numbers* in JavaScript `infinity`, `-infinity`, and `NaN`. I've come across `NaN` or *not a number* before, but not `infinity`/`-infinity`. After a quick search I found that the `infinity` value is used when a "number exceeds the upper limit of the floating point number which is 1.797693134862315E+308". [[source]](https://www.w3schools.com/jsref/jsref_infinity.asp) The same goes for a number lower than -1.797693134862316E+308. The book does note not to trust it in infinity-based computation as it isn't mathematically sound and leads to `NaN`.

Strings in JavaScript are syntactically similar to Python in the ability to use either `"valid"` or `'valid'`, but differs in also having backticks being valid for strings. Not very surprising as JavaScript is also an [interpreted language](https://en.wikipedia.org/wiki/Interpreted_language). One thing I never liked about JavaScript was how it handles string concatenation. An example is as follows.

```JavaScript
console.log("this " + "looks " + "ugly");
```

Fortunately there is a cleaner way in the form of *template literals*, but this is only available when using backticks for strings. An example of this is as follows.

```JavaScript
let better = "cleaner";

console.log(`this looks ${better}`);
```

The next topic covered in this chapter is *unary operators*, which I've come across, but have never known the name until now. The example the book uses is `typeof` in JavaScript, but one that I know from C is the `++` and `--` operators. An interesting distinction the book makes is that *binary operators* are operators that use two values, where as *unary operators* only use one.

Another cool operator that I've seen before but have never learned about is the *ternary operator* `?`. Though it may look confusing, in my opinion is an elegant way of combining a conditional operator with two possible outputs. An example of the `?` operator in action is as follows.

```JavaScript
console.log(10 > 5 ? true : false); // outputs true

console.log(1 > 5 ? true : false; // outputs false
```

*Spicy*.

An oddity in JavaScript is it's *automatic type conversion*. The book's examples of this are as follows.

```JavaScript
console.log(8 * null); // outputs 0

console.log("5" - 1); // outputs 4

console.log("5" + 1); // outputs 51

console.log("five" * 2); // outputs NaN

console.log(false == 0); // true
```

Apperantly JavaScript will convert a value to a needed type using a set of rules if an operator is applied to the "wrong" type of value. This is *type coercion*. Enter the operator unique to JavaScript;`===`. The `===` operator is exactly the same as `==` except it does not apply type coercion, instead it checks to see if the *types* are equivalent. An example of this is as follows.

```JavaScript
console.log(false == 0); // outputs same as above; true

console.log(false === 0); // outputs false
```

The final section of chapter 1 covers *short-circuit evaluation*, which is something I've multiple times but have never known the name for it until now. These are the `||` and `&&` operators.

That's it for chapter 1 and this blog! To finish off I will write the classical *hello world!* in Javascript.

```JavaScript
console.log("Hello world!");
```
