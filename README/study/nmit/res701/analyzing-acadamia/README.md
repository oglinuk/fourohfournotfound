---
title: Analyzing Academic Articles
date: 2019-04-05
tags: ["Res701,", "Academic Research,", "Credability,", "Academic Publishing,", "Analyzing Academic Articles"]
draft: false
---

Todays blog is about reading and understanding one of the academic articles that I chose for the [previous blog](/study/nmit/res701/academic-research). We are to answer the following questions around the article.

* Did the abstract tell you three things I said it should? If not what did it tell you?

* What seems to be the research question(s) they were trying to answer?

* What method(s) did they use to answer the question(s)?

* How credible do you think the paper is?

* Did you agree, or not, with what they wrote in their conclusion? Why?

* Briefly describle two things that you learnt from the paper.

* In no more than 250 of your own words, describe what the paper is about.

---

#### Did the abstract tell you three things I said it should? If not what did it tell you?

Yes, the abstract answered the three things; *what the research article topic is*, *what the authors/researchers did*, *what they discovered/or created/or concluded*.

**What the research article topic is**

"Class of mathematically rigorous, general, fully self-referential, self-improving, optimally efficient problem solvers." [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

**What the authors/researchers did**

"We show that such a self-rewrite is globally optimal — no local maxima! — since the code first had to prove that it is not useful to continue the proof search for alternative self-rewrites." [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

**What they discovered/or created/or concluded**

"Unlike Hutter’s previous non-self-referential methods based on hardwired proof searchers, ours not only boasts an opti-mal order of complexity but can optimally reduce any slowdowns hiddenby the O()-notation, provided the utility of such speed-ups is provable at all." [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

---

#### What seems to be the research question(s) they were trying to answer?

"Here we will eliminate the restrictive need for human effort in the most general way possible, leaving all the work including the proof search to a system that can rewrite and improve itself in arbitrary computable ways and in a most efficient fashion." [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

"Many traditional problems of computer science require just one problem-defining input at the beginning of the problem solving process." [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

"We will also consider the more general case where the problem solution requires interaction with a dynamic, initially unknown environment that produces a continual stream of inputs and feedback signals, such as in autonomous robot control tasks, where the goal may be to maximize expected cumulative future reward." [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

#### What method(s) did they use to answer the question(s)?

One method that was used in this academic article was *secondary research*, since the basis of this article was "inspired by Kurt Gödel's celebrated self-referential formulas (1931)". [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

#### How credible do you think the article is?

I believe the article is highly credible since it is published by the established professor, Jürgen Schmidhuber. Another reason I believe the article is highly credible is because it is on the arXiv.org website which is "owned and operated by Cornell University" [[source]](https://arxiv.org/).

#### Did you agree, or not, with what they wrote in their conclusion? Why?

The main point (in my opinion) of the conclusion is "it is possible to use self-referential proof systems to build optimally efficient yet conceptually very simple universal problem solvers." [[source]](https://arxiv.org/pdf/cs/0309048.pdf) ***I agree*** and the conclusion contains the reason for why that is. "The initial software p(1) of our Gödel machine runs an initial, typically sub-optimal problemsolver, e.g., one of Hutter’s approaches [16, 17] which have at least an optimal order of complexity, or some less general method [20]. Simultaneously, it runs an O()-optimal initial proof searcher using an online variant of Universal Search to test proof techniques, which are programs able to compute proofs concerning the system’s own future performance, based on an axiomatic system A encoded p(1), describing a formal utility function u, the hardware and p(1) itself. If there is no provably good, globally optimal way of rewriting p(1) at all, then humans will not find one either. But if there is one, then p(1) itself can find and exploit it.  This approach yields the first class oftheoretically sound, fully self-referential, optimally efficient, general problem solvers." [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

#### Briefly describle two things that you learnt from the article.

One thing that I learned about in this article is the limitations of Gödel machines. The limitations being "any formal system that encompasses arithmetics (or ZFC etc) is either flawed or allows for unprovable but true statements... In particular, one can construct pathological examples of environments and utility functions that make it impossible for the machine to ever prove a target theorem." [[source]](https://arxiv.org/pdf/cs/0309048.pdf) 

Another thing that I learnt about in this article was the applications of Gödel machines. These applications include; "*time-limited NP-hard optimization*, *fast theorem proving*, *maximizing expected reward with bounded resources*, and *optimize any suboptimal problem solver*." [[source]](https://arxiv.org/pdf/cs/0309048.pdf)

#### In no more than 250 of your own words, describe what the article is about.

The academic article *Gödel machines: self-referential universal problem solvers making provably optimal self-improvements* describes the possibility of creating a system which can make improvements to itself that are optimal and provable. The article also covers the various components that would make up such a system which include: proof searcher, universal search, global optimality theorem, and problem solver. Along with the previous, the article also discusses the differences between the previous works in self-learning machines and the Gödel machine.