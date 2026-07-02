# How to get Most out of this course

1. basic lang
2.Watch The videos
3. Practice
4.Code Attachements

# imports in Go
must be double Quoted import "fmt"

## you can have multiple packages in Go
learn different Packages from standard library

# fmt is go's package library 

# why its package name is main why not any other thing ?
we should use main because main is special package name that tells go that this package will be main entry.

# " go: go.mod file not found in current directory or any parent directory; see 'go help modules' " whats this error??
because you did not set any folder/site

# use 
PS D:\Go> go mod init github.com/dimcodes09/Go-Lang/first-app
go: creating new go.mod: module github.com/dimcodes09/Go-Lang/first-app
go: to add module requirements and sums:
        go mod tidy
PS D:\Go> go build
PS D:\Go> ./first-app
HEllo World

A module is simply a go project 


# Mental Framework for Every Go Program
Whenever you build any program, ask these questions in order:
1. What is the problem?
2. What inputs do I need?
3. What outputs do I need?
4. Which variables should I create?
5. Which values are constants?
6. Which packages are required?
7. What calculations or logic are needed?
8. How do I display the result?

# How You Should Study Every Go Program
Whenever you see code, don't ask:
> "What is this calculator doing?"
Instead ask:
1. Which package is imported?
2. What variables are created?
3. Which variables come from user input?
4. Which new variables are created using `:=`?
5. Which Go functions are called (`Print`, `Scan`, `Println`, etc.)?
6. What is the execution order?

If you can answer those six questions, you're learning Go rather than memorizing an application's business logic.