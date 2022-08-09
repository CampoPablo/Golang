# Go interface

Go interface tutorial shows how to work with interfaces in Golang.

An interface is a set of function signatures and it is a specific type. Another type implements an interface by implementing its functions. While languages like Java or C# explicitly implement interfaces, there is no explicit declaration of intent in Go.

The primary task of an interface is to provide only function signatures consisting of the name, input arguments and return types. It is up to a type (e.g. struct type) to implement functions.

Interfaces are also reffered to as exposed APIs or contracts. If a type implements the functions (the APIs) of a sortable interface, we know that it can be sorted; it adheres to the contract of being sortable.

Interfaces specify behaviour. An interface contains only the signatures of the functions, but not their implementation.

Using interfaces can make code clearer, shorter, and more readable.