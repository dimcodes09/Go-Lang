| Topic                   | One-Line Note                                                           |
| ----------------------- | ----------------------------------------------------------------------- |
| `for {}`                | Creates an infinite loop until `break` or `return` is executed.         |
| `switch`                | Cleaner alternative to multiple `if-else` statements.                   |
| `case`                  | Executes code for the matching switch value.                            |
| `default`               | Executes when no case matches.                                          |
| `continue`              | Skips the current loop iteration and starts the next one.               |
| `return`                | Immediately exits the current function and optionally returns value(s). |
| `error`                 | Built-in Go type used to represent errors.                              |
| `nil`                   | Means "no value" or "no error".                                         |
| `errors.New()`          | Creates a custom error with a message.                                  |
| `panic()`               | Stops the program immediately with an error.                            |
| `err != nil`            | Checks whether an error occurred.                                       |
| Multiple Return Values  | A function can return more than one value.                              |
| `os.ReadFile()`         | File → `[]byte`                                                         |
| `string()`              | `[]byte` → `string`                                                     |
| `strconv.ParseFloat()`  | `string` → `float64`                                                    |
| `fmt.Sprint()`          | Any value → `string`                                                    |
| `[]byte()`              | `string` → `[]byte`                                                     |
| `os.WriteFile()`        | `[]byte` → File                                                         |
| `strconv` package       | Used to convert strings into numeric types.                             |
| `os` package            | Used for file and operating system operations.                          |
| `errors` package        | Used to create custom errors.                                           |
| File Handling           | Reading data from a file and writing data back to a file.               |
| Type Conversion         | Converting one data type into another.                                  |
| File Persistence        | Data remains saved even after the program closes.                       |
| `0644`                  | Standard file permission used while creating/updating a file.           |
| `if err != nil`         | Standard Go pattern for handling errors.                                |
| Named Return Values     | Return variables are declared in the function signature.                |
| Function Decomposition  | Breaking a large problem into smaller reusable functions.               |
| `getBalanceFromFile()`  | Reads balance from the file and returns balance + error.                |
| `writeBalanceToFile()`  | Saves the updated balance into the file.                                |
| `getUserInput()`        | Reusable function to take user input.                                   |
| `CalculateFinancials()` | Returns multiple calculated values from one function.                   |
| `fmt.Printf()`          | Prints formatted output directly to the console.                        |
| `fmt.Sprintf()`         | Returns formatted text as a string instead of printing it.              |
