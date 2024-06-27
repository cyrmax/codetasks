# Valid braces

From codewars.com

Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.
This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!
All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

## What is considered Valid?

A string of braces is considered valid if all braces are matched with the correct brace.

## Examples

```
"(){}[]"   =>  True
"([{}])"   =>  True
"(}"       =>  False
"[(])"     =>  False
"[({})](]" =>  False
```

## Build

To build the CLI tool for testing purposes you should use not a pure `go build` in the root directory but rather `go build ./cmd/main/main.go` to show the Go compiler that you need not a library package but an executable program.
