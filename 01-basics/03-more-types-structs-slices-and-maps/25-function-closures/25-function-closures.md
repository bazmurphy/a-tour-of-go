# Function closures

Go functions may be closures. **A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.**

For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable.

## Notes

A closure in programming refers to a function that retains access to its surrounding state (lexical environment). This allows the function to access variables defined outside of its scope.

A function closure specifically refers to when a nested function accesses non-local variables from its enclosing scope. The nested function "closes over" the non-local variables, allowing it to access and modify them even when invoked outside of its original scope.

In the example, the `adder()` function returns an anonymous function that closes over the `sum` variable. `sum` is declared in the outer scope of `adder()`, but the returned function accesses and modifies this variable each time it is called.

Each call to `adder()` creates a new instance of the `sum` variable. The returned function "closes over" this instance of `sum`, allowing it to maintain a running total.

So `pos` and `neg` are two separate function closures - each one closes over its own `sum` variable that is incremented or decremented when the function is called. This demonstrates how closures allow functions to maintain state in between calls.

In summary, a closure is a function that retains access to variables in its enclosing scope. A function closure demonstrates this by returning an inner function that accesses non-local variables defined in the outer scope.
