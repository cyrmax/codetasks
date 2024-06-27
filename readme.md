# Code Tasks

In this repository I keep all the programming exercises I've done on CodeWars, LeetCode and other similar sites.
It also contains some ideas, weird algorithm implementations and other stuff that is too small for a separate repository.

## Directory layout

In the root directory are the files needed for the whole repository.
For example, this readme and some very generic gitignores covering only basic OS and IDE stuff.

Within the root directory are subdirectories for programming languages, such as python, go, rust, etc.
Each language subdirectory contains at least one language specific gitignore file and one or more exercise subfolders.
If necessary, a language folder may also contain a language-specific readme.

Each exercise subdirectory represents a self-contained project, as it would be if developed outside of this repository.
For example, it is a complete go module for the go language, a rust crate with a manifest for rust, etc.
This is necessary for reproducibility and to better learn package management in different languages and build systems.

Each exercise also contains a readme file, which must at least state the requirements for the exercise and any additional information about it.

## Commit strategy

Ideally, each commit should represent exactly one completed task, no more and no less.

The commit title should be something like "language: task name".
For example, "Go: Valid braces" or "Python: snail matrix".

For commits that are not completed tasks, but some maintenance work, the language should be "Meta", and the rest should explain what happened in a very short way.
For example, "Meta: initial commit" or "Meta: updates gitignore for Go".

# Contribute

You may want to contribute your solution to one or more tasks.
You can do this by submitting a pull request to this repository.

When making such a contribution, please keep the following in mind:

1. New solutions for an existing task should not overwrite existing code.
2. New solutions should be placed in a separate directory with the task name followed by the words "version from" and your nickname or full name.
3. The solution should contain exactly what was implemented in the original task, maybe more, but not less. For example, if the original solution included a readme, solution code, cli tool for testing, and unit tests, then your variant should include at least the same.
4. The commit naming strategy for such solution variants should follow the one described above, but should also include information about the fact that this is not a new task, but a solution variant for an older task.
5. Your readme should include a unique reference to the original task for the convenience of others.
6. Your readme and pull request body should explain why your solution is better, or what the main difference is.

# Support

If you learned something new from this repository, improved your programming skills, or just enjoyed the work, you can support the author with some coins.

[Here is a link to my Boosty page](https://boosty.to/cyrmax/donate) where you can donate or subscribe on a monthly basis.
I know the website is not the best, but I have no other way to accept donations at the moment.

# Credits

- Thanks to the DeepL Write website for making my English in this readme a little better.
