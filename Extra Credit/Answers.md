# Extra Credit with GO
*Note:* I just learned Go over the weekend in order to do this code assessment, so I'm still learning!

But this was a fun exercise. :)

## 1. Convert an empty interface{} to a typed var

An empty interface may hold values of any type because every type implements at least zero functions.

We can use empty interfaces to handle values of unknown type. For example:

```
var unknownType interface{}
unknownType = "hello world" // Type is now a string
fmt.Printf("%v is of type %T", unknownType, unknownType) // "hello world" is of type string

unknownType = 123 // Type is now an int
fmt.Printf("%v is of type %T", unknownType, unknownType) // 123 is of type int

```

## 2. Show an example of a Go routine and how to stop it

See goroutine.go in this directory for an example.


## 3. Create a simple in-memory cache that can safely be accessed concurrently.

I didn't have time to do this.

## 4. What is a slice and how is it related to an array?

A slice is an abstraction layer that sits on top of an array. When a slice is declared, the runtime allocates the required memory and creates an array in the background while returning the slice.

All items in a slice are of the same type, and they can be resized (unlike an array in Go). This makes slices easier to work with if you don't know the size of your list of items, or if you need to do things like sort.

## 5. What is the syntax for zero allocation, swapping of two integers?

Go has tuple assigment that we can use for this:
```
integer1, integer2 := 123, 456
integer1, integer2 = integer2, integer1
```