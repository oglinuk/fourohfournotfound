# Vimtutor Lesson 1

One of the amazing things about [Vim](https://github.com/vim/vim), is that it comes with an interactive tutorial called `vimtutor`. As the tutorial states, it "is designed to describe enough of the commands that you will be able to easily use Vim as an all-purpose editor". It also states "this tutor is set up to teach by use. That means that you need to execute the commands to learn them properly. If you only read the text, you will forget the commands!"

Let's get into it!

Lesson 1.1 is ***moving the cursor*** around within Vim. There are two ways of moving the cursor; `h` (left), `j` (down), `k` (up), `l` (right) and the arrow keys. The tutor notes that if you use `hjkl` you will be able to move around much faster.

Lesson 1.2 is ***exiting*** Vim, which when I first heard about Vim, it was through a joke about opening Vim and getting non-vim users to try and exit it. To exit Vim *without saving any changes* press `esc`, then type `:q`. You can force the command by adding `!`, which looks like `:q!`.

Lesson 1.3 is ***deletion***. To delete a character, hover over it and press `x`.

Lesson 1.4 is ***insertion***. To insert text press `i` to enter *insert mode*.

Lesson 1.5 is ***appending***. To append text to the end of a line, press `A`.

Lesson 1.6 is ***editing a file***. After making changes to a file, when you want to exit Vim and *keep the changes* (unlike in lesson 1.2, where it exits without saving changes), press `esc` then execute `:wq`.

That's it for the first lesson in the `vimtutor`!
