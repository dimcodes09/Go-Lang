# Error

in js we use try catch similarly in Go we do like
value, err := someFunction()
if err != nil {
    //handle error
}

`Why Two Return Values?`
func getBalanceFromFile() (float64, error)
This means the function can return:
1. Balance
2. Error

nil = Nothing
       No value
       No error
       Empty reference

`Creating an Error`
return 1000, errors.New("Failed to read file")
creating your own error object.

`Checking Errors`
if err != nil {
	fmt.Println(err)
}

Go returns the error so you decide how to handle it.

