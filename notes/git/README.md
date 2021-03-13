# Git Notes

## Resources

* `man git`
* `man gitutorial`
* `man giteveryday`
* `man gitcli`
* [git scm](https://git-scm.com/)
* [git source code](https://github.com/git/git)

## Learning Plan

* `git config` 1 -> 2
* `git clone` 2 -> 3
* `git add` 2 -> 3
* `git commit` 2 -> 3
* `git merge` 1 -> 2
* `git tag` 1 -> 2
* `git rebase` 1 -> 2
* `git stash` 2 -> 3
* `git submodule` 1 -> 2
* `git remote` 2 -> 3
* `git log` 1 -> 2
* `git grep` 1 -> 2
* `git bisect` 0 -> 1
* `git request-pull` 0 -> 1
* `git format-patch` 0 -> 1
* `git send-email` 0 -> 1

## Install Git From Source

[Download source code](https://github.com/git/git/archive/master.zip) and
change into the code directory.

`make prefix=/usr`

`sudo make install`

`git --version`

## Basic Git Workflow

`git init`

`git remote add origin <url>`

`echo "# Hello Git!" >> README.md`

`git add README.md`

`git commit -m 'Initial commit'`

`git push origin master`

## Config

There are two required variables to be set via `git config`; `user.email`
and `user.name`. In this case its a global setting so the command will
look like `git config --global user.email <email>` or `git config
--global user.name <username>`. The other variable I set, but is not
required, is `core.editor`. The default is `nano`, which I dont use, so I
set it to `git config --global core.editor vim`. By default the
configuration values are written to the local configuration file localed
in the repository. There are options like `--system`, `--global`, and
`--file` that can be used to override this default.

`git config <variable> <value>` sets `<variable>` to `<value>` in the
config file.

`git config --unset <variable>` removes `<variable>` from the config
file.

`git config --global sendemail.suppresscc self` prevents `git send-email`
from adding you to the cc list. Only set this configuration if you dont
wish to recieve a copy of each patch to your email.

`git config sendemail.to <email>` sets the default email address to send
emails to when using `git send-email`.

## Clone

There are multiple protocols supported in Gits transport mechanism.
According to `man git clone`, the supported protocols are: `ssh`, `git`,
`http`, and `https`. It is important to note that although Git supports
the `ftp` and `ftps` protocols, they should never be used (as stated in
the man page). I primarily use `https`, but have started to use `ssh`
more when possible.

**clone via https**
`git clone https://host.com[:port]/path/to/repo.git`

**clone via ssh**
`git clone ssh://[user@]host.com[:port]/path/to/repo.git`

on Github you can do `git clone ssh://git@github.com/path/to/repo.git`

`git clone -b dev-django https://github.com/oglinuk/sdv701-tiered` will
clone the repository and checkout the specified branch (`dev-django`) in
this case. It is important to note that this will also fetch all of the
repositorys branches as well as the specified branch. If you want only
the single branch, you need to specify the `--single-branch` option.

If the repository to be cloned is a *private* repo, then you will need to
setup an SSH key with the hosted git service, and clone using the `ssh`
protocol.

## Add

With `git add`, you can add file content a few different ways. You can
add using the filename(s) - `git add README.md`, you can add an entire
directory - `git add src/`, and you can add via globbing - `git add
*.go`. There are a few options which enchance the experience of using
`git add`, such as `--patch`.

`git add -n <file(s)>` || `git add --dry-run <file(s)>` will not add the
`README.md` file to the index, but will display what will happen when
running the command. If the `-n` or `--dry-run` option is combined with
`--ignore-missing`, the output will show if any of the specified files
will be ignored.

`git add -f <file(s)>` || `git add --force <file(s)>` will
add otherwise ignored files to the index.

`git add -p <file(s)>` || `git add --patch <file(s)>` enters an
interactive mode when you can choose what to do with the hunks of a
patch. This option only works for *modified* file content.

`git add --chmod +x <file(s)>` overrides the executable bit of
the specified file. This only changes the executable bit in the index,
the file on disk is not changed.

## Commit

Though I started out only ever using `git commit -m 'commit message'`,
Ive decided that `git commit` should be used *without* the `-m` flag,
unless there is no additional information that is necessary.

`git commit` opens an editor (`core.editor`) for the author to write a
message to associate with the commit.

`git commit -a` || `git commit --all` stages all modified and
deleted, then open the editor to write a commit message. This command
does not affect any new files.

`git commit -p` || `git commit --patch` does the same as `git add -p` ||
`git add --patch`.

`git commit --dry-run` works the same as `git add -n` || `git add
--dry-run`. The `--dry-run` option can be combined with other options
like `--short`, `--porcelain`, and/or `--long`.

`git commit -m <msg>` || `git commit --message <msg>` is a way to write
a commit message without opening an editor.

`git commit --amend` updates the last commit, rather than creating a new
one. I dont use this option as often as I should.

## Merge

`git merge <branch>` will merge the commit history of `<branch>` into to
the current branchs history.

`git merge --commit <branch>` performs a commit after a merge. This
is the default, but it can be overridden by using the `--no-commit`
option instead.

`git merge -e <branch>` || `git merge --edit <branch>` opens an
editor before committing a successful merge to allow for further changes
to the merge message. The `--no-edit` option can be specified instead to
accept the auto-generated message (though this is discouraged according
to the man page).

`git merge --ff <branch>` will not create a merge commit if the merge is
only a fast-forward. This can be overridden by using the `--no-ff` option
instead. There is another option (`--ff-only`) that will refuse to merge
unless the current HEAD is already up to date or the merge can be
resolved as a fast-forward. A fast-forward merge is when the current
branch HEAD is an ancestor of the named commit. In this case, a new
commit is not needed, and the HEAD is updated to point at the named
commit.

`git merge --stat <branch>` shows a diff after a successful merge. This
is the default, but can be overridden by the `--no-stat` || `-n` option.

`git merge --squash <branch>` squashes the commits of `<branch>` into one
commit and merge it to the current branch. If the merge is a
fast-forward, the HEAD will not be updated with the commit from
`<branch>`.

`git merge -s <strat> <branch>` || `git merge --strategy <strat> <branch>`
allows for a specific merge strategy to be used. The default strategy is
`recursive`, `resolve`, `octopus`, `ours`, `subtree`. To learn more about
these strategies, I suggest reading [Git Merge Strategy Options and
Examples](https://www.atlassian.com/git/tutorials/using-branches/merge-strategy)
by Atlassian.

`git merge -X <option> <branch>` is a way to pass an option to the merge
strategy. For the default strategy (`recursive`), the options available
are `ours`, `theirs`, `patience`, `diff-algorithm`, `renormalize`,
`no-normalize`, `no-renames`, `find-renames=<name>`, and `subtree`. To
learn more about these options, please reference the 'Recursive Git Merge
Strategy Options' section of [Git Merge Strategy Options and
Example](https://www.atlassian.com/git/tutorials/using-branches/merge-strategy).

`git merge -m <msg> <branch>` is similar to `git commit -m <msg>`, and
will create a commit message to be used for the merge commit (if one is
created).

## Tag

Git tags are located in `refs/tags/`. Unless using the `-f` option, a tag
must not exist when creating one.

`git tag <name>` will create a tag with `<name>` if it does not exist.

`git tag -a <name>` || `git tag --annotate <name>` allows for an
annotated tag object.

`git tag -f <name>` || `git tag --force <name>` replaces and existing tag
with `<name>`.

`git tag -d <name>` || `git tag --delete <name>` deletes existing tag
`<name>`.

`git tag -l` || `git tag --list` lists all tags. You can add a
`<pattern>` to list only tags that match, for example: `git tag -l
'v1.*'`.

`git tag --sort <key>` sorts based on `<key>`. The `-` prefix is to sort
in descending order. A key can be any one of the following: `refname`,
`objecttype`, `objectsize`, `objectname`, `upstream`, `push`, `HEAD`,
`color`, `align`, `if`, `symref`, `object`, `tree`, `type`, and `tag`.
For more information on each of these keys, please reference `man git
for-each-ref`.

`git tag --contains <commit>` only lists tags which contain `<commit>` or
the HEAD of no commit is specified. You can use the `--no-contains
<commit>` option to only list tags which *dont* contain `<commit>` or
HEAD if no commit is specified.

## Rebase

The `git rebase` command has been one of the hardest for me to grasp.
Fortunately I found another great [post by
Atlassian](https://www.atlassian.com/git/tutorials/rewriting-history/git-rebase)
on this topic. According to the post, "Rebase is one of two Git utilities
that specializes in integrating changes from one branch to another. The
other change integration utility is `git merge`." I think it will be more
beneficial to first understand how these two commands differ, which
fortunately they also have a post on called [Merging vs.
Rebasing](https://www.atlassian.com/git/tutorials/merging-vs-rebasing).
The post states that both commands are designed to do the same thing,
which is integrate changes from one branch into another, but they do them
in different ways. I am not going to re-write the differences, but to sum
up what I understood, merging is a "non-destructive" operation, whereas
rebasing will re-write the history of the project by creating new commits
for each original commit. I think its important to highlight the golden
rule that is stated in the comparison post, which is "never use git
rebase on public branches".

`git rebase <branch>` takes the commits in `<branch>` and applies them to
the HEAD of the current branch. According to the post on `git rebase`,
"Rebasing is a common way to integrate upstream changes into your local
repository."

`git rebasea -i` starts a rebase in interactive mode. In this mode, you
can clean up the history by altering commits.

## Stash

`git stash` saves the working directory and index state.

`git stash list` shows all currently stashed changes.

`git stash show <stash>` shows the changes of `<stash>`. If no `<stash>`
is specified, it will show the latest stash. The diffstat of this command
can be formatted with any known formats of `git diff`.

`git stash pop <stash>` removes `<stash>` from the stash list and applies
it on top of the current working tree. If there are any conflicts, the
operation will fail and the stash will not be removed from the stash
list. You will have to resolve the conflicts and run `git stash drop`
afterwards.

`git stash apply <stash>` is similar to `git stash pop`, but does not
remove `<stash>` from the stash list.

`git stash branch <branch> <stash>` creates/checks out `<branch>`
starting from the commit of `<stash>` and applies the changes of
`<stash>` to the new working tree and index.

`git stash clear` removes all stash entries. Use with caution as the
entries may be un-recoverable.

`git stash drop <stash>` removes `<stash>` from the stash list. If no
`<stash>` is specified, it will remove the latest one.

## Submodule

`git submodule add <repository>` adds `<repository>` to the current
project. The URL can be absolute or relative.

`git submodule status <submodule>` shows the current commit, name,
and tag of `<submodule>`. If no `<submodule>` is provided, it will show
the status of all submodules.

`git submodule summary <submodule>` shows the commit summary of
`<submodule>` (defaults to HEAD).

`git submodule foreach <command>` will evaluate `<command>` in each
checkout out submodule. The `<command>` has access to variables `$name`,
`$path`, `$sha1`, and `$toplevel`. For details on these variables look at
the `foreach` section of `man git submodule`. An example command may look
like `git submodule foreach 'echo $path $name $sha1 $toplevel'`.

`git submodule sync <submodule>` synchronizes `<submodule>`. If no
`<submodule>` is specified, sync all submodules.

`git submodule absorbgitdirs` moves a `.git` directory of a submodule
into the superprojects `$GIT_DIR/modules` path and connect the Git
directory and the working directory by setting `core.worktree` and adding
a `.git` file pointing to the Git directory stored in `$GIT_DIR/modules`.

## Remote

`git remote` outputs all remotes.

`git remote add <name> <url>` adds a remote called `<name>` for the
repository at `<url>`.

`git remote rename <old> <new>` renames the remote `<old>` to `<new>`.

`git remote rm <name>` || `git remote remove <name>` removes the remote
`<name>`. All remote-tracking branches and configuration settings are
removed for `<name>`.

`git remote get-url <name>` outputs the URL(s) of `<name>`. Only the
first URL is shown by default. To list all URLs use the `--all` option.
If you want to see push URLs rather than fetch URLs, use the `--push`
option.

`git remote set-url <name> <url>` sets the remote `<name>` to `<url>`,
unless `<url>` is not valid, in which case nothing is changed. You can
add `<url>` rather than changing the existing remote URL by using the
`--add` option. If you want to remove URLs rather than change them, use
the `--delete` option. If you want to change push URLs, use the `--push`
option.

`git remote show <name>` provides information about the remote `<name>`
like the fetch/push URLs and more.

`git remote prune <name>` deletes stale references in `<name>`.

## Log

`git log` outputs the commit logs.

`git log --follow <file>` only outputs the commit logs that affect
`<file>`. Important to note it only works for one file.

`git log -<n>` || `git log -n <n>` || `git log --max-count <n>` limits
the number of commit logs to `<n>`.

`git log --skip <n>` skips `<n>` commits before outputting the commit logs.

`git log --since "<date>"` || `git log --after "<date>"` shows commits
more recent than `<date>`.

`git log --until "<date>"` || `git log --before "<date>"` shows commits
older than `<date>`.

`git log --author <pattern>` || `git log --committer <pattern>` only
shows commits where the author or committer matches `<pattern>`.

`git log --grep <pattern>` only shows commits that match `<pattern>`. To
only show commits that *dont* match <pattern>, use the `--invert-grep`
option.

`git log --merges` only shows commits that are from a merge. To only show
commits that *arent* from a merge, use the `--no-merges` option.

`git log ..@{u}` shows all commits on `remotes/branch` that are not
already on the current branch. The `@{u}` is shorthand for
`HEAD@{upstream}`.

## Grep

`git grep <pattern>` shows any *tracked* files in the current working
tree that match `<pattern>`.

`git grep --untracked <pattern>` includes untracked files when searching.

`git grep --recurse-submodules <pattern>` searches recursively in each
submodule that has been initialized/checked out in the repository.

`git grep -a <pattern>` || `git grep --text <pattern>` searches binaries
in addition as if they were text.

`git grep -n <pattern>` || `git grep --line-number <pattern>` shows the
line number of the matching line in the output.

`git grep -c <pattern>` || `git grep --count <pattern>` shows the number
of lines that match `<pattern>`, rather than showing the lines that
match.

`git grep --heading <pattern>` shows the filename above the matching
lines rather than at the start of each matching line.

`git grep -<n> <pattern>` || `git grep -C <n> <pattern>` || `git grep
--context <n>` shows `<n>` leading and trailing lines of a line matching
<pattern>. If you only want trailing lines, use the `-A <n>` ||
`--after-context <n>` options. If you only want leading lines, use the
`-B <n>` || `--before-context <n>` options.

`git grep --threads <n> <pattern>` specifies the number of grep worker
threads to use.

`git grep -f <file> <pattern>` only searches the lines of `<file>` to see
if any match `<pattern>`.

`git grep -e "<pattern>"` is an option that must be used for patterns
starting with `-`.

## Bisect

This is a command that Ive never come across until I read the post about
`git rebase` by Atlassian. At first glance it looks like a must learn
command for any person who uses Git in a production environment. The `man
git bisect` page states, "This command uses a binary search algorithm to
find which commit in your projects history introduced a bug. You use it
by first telling it a "bad" commit that is know to contain the bug, and a
"good" commit that is known to be before the bug was introduced. In fact,
`git bisect` can be used to find the commit that changed *any* property
of your project; e.g., the commit that fixed a bug, or the commit that
caused a benchmarks performance to improve."

`git bisect start` starts a session.

`git bisect bad <commit>` specifies `<commit>` as bad. If no `<commit>`
is specified, then the current commit is used.

`git bisect good <commit>` specifies `<commit>` as good. If no `<commit>`
is specified, then the current commit is used.

The above three commands are all that is needed for `git bisect` to work.

`git bisect reset` cleans up the bisection state and return to the
original HEAD after a bisect session.

`git bisect old <commit>` specifies that a commit was before the desired
change.

`git bisect new <commit>` specifis that a commit was after the desired
change.

The `old` and `new` commands are for when youre looking for a commit that
caused a change between some `old` state and `new state`. According to
`git man bisect`, "In this more general usage, you provide `git bisect`
with a 'new' commit that has some property and an 'old' commit that
doesnt have that property. Each time `git bisect` checks out a commit,
you test if that commit has the property. If it does, mark the commit as
'new'; otherwise, mark it as 'old'. When the bisection is done, `git
bisect` will report which commit introduced the property."
introduced.

## Request Pull

When reading the documentation provided by the Linux kernel maintainers
on [Creating Pull
Requests](https://www.kernel.org/doc/html/latest/maintainer/pull-requests.html)
I came across `git request-pull`, which is another command I have never
heard of. I am eager to get better with this command so I can create pull
requests from my terminal, rather than having to use the graphical
interface of the hosted git services I use.

`git request-pull <start> <url>` creates a new pull request to `<url>`
starting at commit `<start>`.

`git request-pull -p <start> <url>` includes patch details in the output.

`git request-pull <start> <url> <end>` specifies which commit to `<end>`
at. This defaults to HEAD. 

**Example `git request-pull` workflow**:

Push changes to your public repository.

`git push https://git.abc.xyz/project master`

Then run the `git request-pull` command.

`git request-pull v1.0.0 https://git.abc.xyz/project master`


## Format Patch

In my opinion the `git format-patch` command is a must when working with
creating pull requests and submitting them via a mailing list. Its entire
purpose is to format patches for emailing.

`git format-patch <branch>` extracts all commits which are in the current
branch, but not in `<branch>`.

`git format-patch --root <branch` extracts all commits which lead to
<branch> since the inception of the project.

`git format-patch -<n>` extracts the `<n>` recents commits from the current
branch.

**Example `git format-patch` workflow**:

`git checkout -b feature-branch`

... changes ...

`git commit`

`git format-patch master`

## Send Email

Just like `git request-pull` and `git format-patch`, Im interested to
improve my knowledge of `git send-email` to allow me to gain the
experience necessary to contribute to projects that are maintained over
mailing lists. It is important to note that before this command can be
used, there is some configuration that needs to be done in
`~/.gitconfig`.

```
[sendemail]
	smtpencryption = tls
	smtpserver = <your email server>
	smtpuser = <your email address>
	smtpserverport = <your email servers port>
```

`git send-email -1` sends the last commit in the current branch to the
default destination address.

`git send-email -1 <commit>` sends `<commit>` to the default destination
address.

`git send-email *.patch` sends all patches to the default destination
address which is set via `git config sendemail.to <email>`.
