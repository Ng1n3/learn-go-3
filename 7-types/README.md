```
type Person struct {
  FirstName string
  LastName  string
  Age       int
}
```

you can use any primitive type or compund type literal to define a concrete type.

```
type Score int
type Converter func(string)Score
type TeamScores map[string]Score
```

Go Doesn't allow you to add methods to types you don't control. While you can define a method in a different file within the same package as the type declaration, it is best to keep your type definition and its associated methods together so that it's easy to follow the implementation.

Functions Versus Methods
Since you can use a method as a function, you might wonder when you should
declare a function and when you should use a method.
The differentiator is whether your function depends on other data. As I’ve covered
several times, package-level state should be effectively immutable. Anytime your
logic depends on values that are configured at startup or changed while your
program is running, those values should be stored in a struct, and that logic should
be implemented as a method. If your logic depends only on the input parameters, it
should be a function.

### Accept Interfaces, Return Structs
This means that the business logic invoked by your functions should be invoked via interfaces, but the output of your function should be a concrete type. Wehen a concrete type is returned by a function, new methods and fields can be added without breaking existing code.

## Type assertions and Type Switches
A type assertion is very different from a type conversion. Conversions
change a value to a new type, while assertions reveal the type of the value
stored in the interface. Type conversions can be applied to both concrete
types and interfaces. Type assertions can be applied only to interface types.
All type assertions are checked at runtime, so they can fail at runtime with
a panic if you don’t use the comma ok idiom. Most type conversions are
checked at compile time, so if they are invalid, your code won’t compile.
(Type conversions between slices and array pointers can fail at runtime
and don’t support the comma ok idiom, so be careful when using them!)

When an Interface could be one of multiple possible types, use a type switch instead.One common use of a type assertion is to see if the concrete type behind the interface also implements another interface.

