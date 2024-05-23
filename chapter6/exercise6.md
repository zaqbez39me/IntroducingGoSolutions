# Exercise 6

## Statement

What are `defer`, `panic`, and `recover`? How do you recover from a runtime panic?

## Answer

The `defer` is used to defer execution of some function call to run right before function return. It is often used to
close the access to some resource for example file or database.

The `panic` is used to exit the function with some error. Usually we want our program into panic when we meet with some
unexpected behaviour or unexpected input.

The `recover` is used to recover from `panic`. For example, we called some function that felt into `panic` and we want
to handle it in proper way.
