# Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```go
(value, type)
```

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

## Notes

The main reason we use the interface type `I` rather than the concrete type `T` is flexibility and abstraction.

Using the interface `I` allows us to write functions that accept any type that satisfies the interface (implements the methods), rather than being restricted to just one concrete type like `T`.

This is useful because:

- We can pass in different concrete types to the same function as long as they satisfy the interface. For example, both T and F satisfy interface I in this code.
- We can more easily swap out or add new concrete types since the function only cares that they implement the interface, not their exact type.
- It reduces coupling between components of our code. The function only depends on the abstraction (interface) rather than specific concrete details.

So while `T` is a concrete type with specific methods, `I` is an interface that defines required behaviour but doesn't require a specific implementation. This gives us flexibility to use any supporting types interchangeably.

The interface value `i` contains both the concrete value and information about its underlying type. When we call `i.M()`, it knows which concrete type's `M()` method to call to satisfy the interface.

So in summary, using the interface gives us abstraction and flexibility, allowing different concrete types to be used interchangeably as long as they satisfy the interfaces our code requires.
