# Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` named `v`.

## Notes

A method is just a regular function with a special receiver parameter.
The receiver binds the function to a specific type.

For example:

```go
type MyType struct {
  // ...fields...
}

func (m MyType) MyMethod() {
  // can access m's fields here
}
```

This defines a method called `MyMethod` on the type `MyType`.

To call this method:

```go
var m MyType
m.MyMethod()
```

The value `m` is passed automatically as the receiver parameter in the function call.

So methods:

- Are functions with a receiver parameter
- Bind functions to types
- Allow extending types with new behaviour
- Receive the type instance as the first parameter

The key benefit is grouping related behaviour with types.

Methods are a simple but powerful concept in Go! Let me know if this helps explain them better.
