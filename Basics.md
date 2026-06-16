# How to get Most out of this course

1. basic lang
2.Watch The videos
3. Practice
4.Code Attachements

# imports in Go
must be double Quoted import "fmt"

## you can have multiple packages in Go
learn different Pckages from standard library

# fmt is go's package library 

# why its package name is main why not any other thing ?
we should use main because mainb is special package name that tells go that this package will be main entry

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