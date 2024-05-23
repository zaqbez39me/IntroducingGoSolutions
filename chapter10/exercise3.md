# Exercise 3

## Statement

That is a buffered channel? How would you create one with a capacity of 20?

## Answer

The buffered channel is channel with restricted capacity which makes the flow stop and asynchronously wait till it is
not full. To create the buffered channel you should use make with channel type and size. In our case we want to create a channel with capacity of 20: `make(chan int, 20)`