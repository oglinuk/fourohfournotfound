---
title: Model Representation 
draft: false
---

"Linear regression attempts to model the relationship between two variables by fitting a linear equation to obeserved data. One variable is considered to be an explanatory variable, and the other is considered to be a dependent variable." [[source](https://stat.yale.edu/Courses/1997-98/101/linreg.htm)]

In the [Machine Learning course by Andrew Ng](https://www.coursera.org/learn/machine-learning), Andrew uses a dataset of housing prices in Portland (Oregon), which consists of the size (feet<sup>2</sup>) and price (1000s).

\

### A sample of the dataset:

| Size in feet<sup>2</sup> (x) | Price (y)                |
| :--------------------------: | :----------------------: |
| 2104 (x<sup>(1)</sup>)       | 460000 (y<sup>(1)</sup>) |
| 1416 (x<sup>(2)</sup>)       | 232000 (y<sup>(2)</sup>) |
| 852  (x<sup>(3)</sup>)       | 178000 (y<sup>(3)</sup>) |
| 1534 (x<sup>(4)</sup>)       | 315000 (y<sup>(4)</sup>) |
| ...  (x<sup>(i)</sup>)       | ...    (y<sup>(i)</sup>) |

\

### Notation:

| Symbol         | Meaning                       |
| -------------- | ----------------------------: | 
| `m`            | number of training examples   |
| `x`            | "input" variable or feature   |
| `y`            | "output" or "target" variable |
| `(x, y)`       | single training example       |
| `(x(i), y(i))` | i'th training example         |

\

The explanatory variable in this case is the size, and the dependent variable is price, since the size of the house correlates to the price of the house.

Since this is only dealing with one variable this is called *univariate linear regression*.

Linear regression is expressed as hθ(x) = θ<sub>0</sub> + θ<sub>1</sub>x

![](linear-regression.png)