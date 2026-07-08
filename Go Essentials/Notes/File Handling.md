📌 What is File Handling?

Reading data from a file and writing data back to a file.
With file handling:- 
Program Starts
↓
Read balance.txt
↓
Balance = 1000
↓
Deposit 500
↓
Balance = 1500
↓
Write balance.txt
↓
Program Closes
↓
balance.txt = 1500 ✅

`Reading the File`
data, err := os.ReadFile(accountBalanceFile)
os.ReadFile() reads the complete file.

`Error Checking`
