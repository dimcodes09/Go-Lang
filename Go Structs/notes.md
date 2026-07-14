# 1. Structs
it helps group related data together.

# time :-
          it allows us to work with dat and time values

# * is a pointer\


Absolutely. From now on I'll follow **your exact note format**.

For this file:

---

# Which Problem Do Structs Solve?

Before structs, we store related data in **multiple variables**.

Example:

```go
firstName
lastName
birthDate
```

This works, but as the application grows, managing related data becomes difficult.

# 1. Problems Without Structs
* Related data is stored in separate variables.
* Passing multiple variables to functions makes the code longer.
* There is a higher chance of passing values in the wrong order.
* Managing related data becomes difficult as the project grows.

### Example
Instead of:
```go
outputUserDetails(firstName, lastName, birthDate)
```
We have to pass every variable separately.

# 2. Why do we use Structs?
* Structs group related data into **one single value**.
* Makes code cleaner and easier to understand.
* Makes function calls shorter.
* Reduces mistakes while passing data.
* Easier to manage related data in large applications.

### Example

Without Struct:
```go
outputUserDetails(firstName, lastName, birthDate)
```
With Struct:
```go
outputUserDetails(user)
```

# Remember
* **Without Struct** = Related data is stored in multiple variables.
* **With Struct** = Related data is grouped into one value.
* Structs make code cleaner, reusable, and easier to manage.

# Topic 04. Defining A Struct Type

Now that we know why structs are needed, let's define our own struct.
A **Struct** is a custom data type that groups related data together.

# 1. Defining a Struct

* Use the `type` keyword to create your own custom type.
* Add the `struct` keyword after the type name.
* Structs are usually defined outside functions so they can be used anywhere in the file.

### Syntax
```go
type user struct {

}
```

# 2. Adding Fields to a Struct

* A struct contains fields.
* Every field has a name and a data type.
* Fields are written inside the curly braces.

### Example
```go
type user struct {
	firstName string
	lastName  string
	birthDate string
}
```

# 3. Struct Fields Can Have Different Data Types

* A struct is not limited to one data type.
* You can mix different data types inside one struct.

Example:
```go
type user struct {
	firstName string
	lastName  string
	age       int
}
```

# 4. Structs Can Contain Other Structs
A field can also be another struct.

Example:
```go
import "time"

type user struct {
	firstName string
	lastName  string
	createdAt time.Time
}
```
`time.Time` is also a struct provided by Go.

# 5. Where Should You Define a Struct?
* Usually define structs outside all functions.
* This allows the struct to be used anywhere in the file.

Example:
```go
type user struct {
	...
}

func main() {

}
```

# 6. Naming Structs
* Lowercase = Private (inside the same package).
* Uppercase = Public (Exported).

Example:
```go
type user struct {}
```
Private
```go
type User struct {}
```
Public (Exported)

# 7. Why Do We Create a Struct?

Instead of:
```go
firstName
lastName
birthDate
```
We create:
```go
user
```

Now one variable stores all related data.

# Remember
* `type` creates a custom type.
* `struct` groups related data together.
* Every field has a name and a data type.
* A struct can contain different data types.
* A struct can contain another struct.
* Structs are usually created outside functions.
* Capital letter = Exported.
* Small letter = Private.
* One struct replaces multiple related variables.


# Topic 05. Instantiating Structs & Struct Literal Notation

Now that we have created a struct, we need to create a value based on it.

Creating a value from a struct is called **Instantiating a Struct**.

# 1. Creating a Struct Variable

* Create a normal variable.
* Its data type will be your struct.

Example:
```go
var appUser user
```

Here, `user` is our custom struct type.

# 2. Creating a Struct Instance

To create an instance of a struct, use the struct name followed by curly braces.

Syntax:
```go
user{}
```

This creates a new instance of the `user` struct.

# 3. Struct Literal Notation

We can directly assign values while creating the struct.

Example:
```go
appUser := user{
	firstName: userFirstName,
	lastName:  userLastName,
	birthDate: userBirthDate,
}
```

This is called **Struct Literal Notation**.

# 4. Assigning Values to Fields

Each field is assigned using:
```go
fieldName: value
```

Example:
```go
firstName: userFirstName
```

* Left side = Struct Field
* Right side = Variable or Value

# 5. Separating Fields

* Separate every field using a comma.
* Even the last field should have a comma.

Example:

```go
user{
	firstName: userFirstName,
	lastName:  userLastName,
}
```

# 6. Using Functions Inside Structs

You can also assign values returned by functions.

Example:

```go
createdAt: time.Now(),
```

Here, `time.Now()` returns the current time.

# 7. Structs Can Store Other Structs

`createdAt` has the type:

```go
time.Time
```

`time.Time` is also a struct provided by Go.

So a struct can contain another struct as one of its fields.

# Remember

* Creating a value from a struct is called **Instantiating a Struct**.
* `user{}` creates a new struct instance.
* Use `fieldName: value` to assign values.
* Left side = Struct Field.
* Right side = Variable or Value.
* Separate fields with commas.
* Functions can also return values that are stored inside structs.
* A struct can contain another struct as one of its fields.