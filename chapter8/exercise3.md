# Exercise 3

## Statement

What is a package alias? How do you make one?

## Answer

The package alias is shorthand for the package name. It allows the developers to not to always specify full name of
packages they want to access some internal items.

To make an alias during the import you can bind the package to name as with usual variable. Here is an example:
```go
import m "math"

```