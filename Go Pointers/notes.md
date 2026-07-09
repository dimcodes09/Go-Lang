# Pointers:- 
vars that stores value adresses instead of value.

# & (ampersand) ->
to store adressses

# Advantages:-
1. avoid unecesary value copies (go creates a copy )
2. directly mutate values  

# 1. Why do we need Pointers?

Pointers are mainly used for:
Avoid unnecessary value copies.
Modify the original value directly.
Remember
Without Pointer → Function gets a copy.
With Pointer → Function gets the memory address.

# 2. Understanding Memory

Every variable is stored somewhere in the computer's memory.
Every memory location has its own unique address.

Example:
age := 32

Think like this:
Variable
↓
Value
↓
Memory
↓
Address

A pointer stores this address.

# 3. Creating a Pointer

Use & (Ampersand) to get the memory address.

age := 32
agePointer := &age
Remember
& = Get the memory address of a variable.
4. Pointer Type

A pointer has its own data type.

Normal variable:
var age int

Pointer:
var agePointer *int

Examples:
*int
*string
*float64
*bool
Remember
*int = Pointer to an integer.
*string = Pointer to a string.

# 5. Pointer as a Value

Printing a pointer:
fmt.Println(agePointer)
Output:
0xc0000120a0
This is the memory address.

Remember:- 
Pointers store addresses, not values.

6. Dereferencing
To access the value stored at a pointer's address, use *.
fmt.Println(*agePointer)
Output:
32
Remember
& = Value → Address
* = Address → Value
7. Default Value of a Pointer

The default value of every pointer is:
nil

nil means:
No address
No value
Pointer is not pointing anywhere
8. Passing Pointers to Functions

Without Pointer
updateAge(age)
Go sends a copy of age.
With Pointer
updateAge(&age)
Go sends the memory address.

Function:
func updateAge(age *int)
Remember
*int means the function expects a pointer.

# 9. Using Pointer Inside Function

Wrong:
age - 18
Because age is a pointer.

Correct:
*age - 18
First dereference the pointer, then use the value.
Remember
Always dereference before using the value.

# 10. Using Pointer for Data Mutation
Pointers allow a function to modify the original variable.

Example:
func updateAge(age *int) {
	*age = *age - 18
}

No return value is required.
The original variable changes automatically.

# 11. fmt.Scan() Uses Pointers

Example:
var choice int
fmt.Scan(&choice)
Scan() needs the memory address so it can update the original variable.

Remember:- 

fmt.Scan() always works with pointers.

#
# Quick Revision
`Topic	Note`
Pointer	Stores the memory address of a value.
&	Returns the memory address of a variable.
*pointer	Returns the value stored at the pointer's address.
*int	Pointer to an integer.
Dereferencing	Access value using *pointer.
nil	Default value of a pointer. Means no address.
Passing Pointer	Sends the memory address instead of copying the value.
Data Mutation	Pointer lets a function modify the original variable.
fmt.Scan(&var)	Uses a pointer to update the original variable.
Remember
&
Value → Address
*
Address → Value