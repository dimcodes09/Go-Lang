# Packages

So till here we had learnt packages 

1. Splitting code across files in same package
-> we could create a new file and write any function there and we could call that in our 1st file works for same packages only.

#  2. Why would You use more than one package?

* In Go, we can split code into **multiple files** or **multiple packages**.
* A **package** is used to keep related code together and make it reusable in other projects.
* Utility functions like reading or writing files are good candidates for a separate package.
* This makes the project cleaner, easier to manage, and easier to reuse.

### **Steps to Create a Package**

1. Create a new folder (example: `fileops`).
2. Move related functions (like `getBalanceFromFile` and `writeBalanceToFile`) into that folder.
3. Change `package main` to `package fileops` in that file.
4. Export functions by starting their names with a capital letter (e.g., `GetBalanceFromFile`, `WriteBalanceToFile`).
5. Import the package in `main.go`:

   ```go
   import "your-module-name/fileops"
   ```
6. Call the functions using the package name:

   ```go
   balance, err := fileops.GetBalanceFromFile()
   fileops.WriteBalanceToFile(balance)
   ```

# Remember:
* **Multiple files** = Organize code within the same package.
* **Multiple packages** = Organize and reuse code across different parts of a project or even different projects.


# 4. Preparing Code for Multiple packages :-

* Make the functions **generic** so they can work with **any float value**, not just the account balance.
* Rename `getBalanceFromFile()` to `getFloatFromFile()` and `writeBalanceToFile()` to `writeFloatToFile()`.
* Pass the **file name** as a parameter instead of using a fixed file name inside the function.
* Rename variables like `balance` → `value` and `balanceText` → `valueText` to make the functions reusable.

### **Steps**

1. Rename the functions:

   * `getBalanceFromFile()` → `getFloatFromFile()`
   * `writeBalanceToFile()` → `writeFloatToFile()`
2. Add a `fileName string` parameter to both functions.
3. Replace `accountBalanceFile` inside the functions with `fileName`.
4. Rename variables (`balance` → `value`, `balanceText` → `valueText`).
5. Update all function calls:

   ```go
   value, err := getFloatFromFile(accountBalanceFile)
   writeFloatToFile(accountBalanceFile, value)
   ```
6. After making the functions generic, move them into a **separate package** for better code organization and reuse.

# 5. Splitting code across multiple packages

* Create a **new file** (e.g., `fileops.go`) and move all file-related functions into it.
* A **package cannot share the same folder** with another package. Each package must have its **own folder**.
* Create a folder named `fileops`, move `fileops.go` into it, and change `package main` to `package fileops`.
* After moving the functions, `bank.go` can't use them until you **import the `fileops` package**.

### **Steps**

1. Create a new file: `fileops.go`.
2. Move `getFloatFromFile()` and `writeFloatToFile()` into it.
3. Create a new folder: `fileops`.
4. Move `fileops.go` into the `fileops` folder.
5. Change:

   ```go
   package main
   ```
   to
   ```go
   package fileops
   ```
6. Remove unused imports from `bank.go`.
7. Import the new package in `bank.go` and use its exported functions.

**Remember:**
**One folder = One package.** You cannot have `main` and `fileops` packages in the same folder.

# 6. Importing Packages

* To use functions from your own package, you must **import the package**, just like `fmt`.
* Import using the **module path + package name**, not just the package name.
* Access functions using **packageName.functionName()**.
* If a function is **undefined**, it is usually because it is **not exported** (its name starts with a lowercase letter).

### **Steps**
1. Import your package:

   ```go
   import "your-module-name/fileops"
   ```
2. Call functions like:

   ```go
   fileops.GetFloatFromFile()
   ```
3. Make sure the function name starts with a **capital letter**:

   ```go
   func GetFloatFromFile(...) { }
   func WriteFloatToFile(...) { }
   ```
4. In Go, **only names starting with a capital letter are exported** and can be used from another package.

**Remember:**
* `func getFloatFromFile()` ❌ Private (only inside the same package)
* `func GetFloatFromFile()` ✅ Public (accessible from other packages)

# 7. Exporting & Importing Identifiers

* In Go, **only names starting with a capital letter are exported** and can be used in other packages.
* If a function starts with a lowercase letter, it is **private** to its own package.
* Rename `getFloatFromFile()` → `GetFloatFromFile()` and `writeFloatToFile()` → `WriteFloatToFile()`.
* After exporting, call them using the package name, e.g., `fileops.GetFloatFromFile()`.

### **Steps**
1. Rename the functions:

   ```go
   getFloatFromFile()  →  GetFloatFromFile()
   writeFloatToFile()  →  WriteFloatToFile()
   ```
2. Import your package:

   ```go
   import "your-module-name/fileops"
   ```
3. Use the functions:

   ```go
   balance, err := fileops.GetFloatFromFile(accountBalanceFile)
   fileops.WriteFloatToFile(accountBalanceFile, balance)
   ```
4. Save and run the program. It will work normally while using your custom package.

**Remember:**
* `getFloatFromFile()` ❌ Private
* `GetFloatFromFile()` ✅ Public (Exported)
* **Capital letter = Accessible from other packages.**

# 8. Using third party packages

* Go supports **third-party packages** created by other developers when the standard library is not enough.
* Search packages on **pkg.go.dev** or GitHub, then install them using `go get`.
* Installed packages are automatically added to the **go.mod** file as project dependencies.
* Import the package and use its functions just like built-in packages.

### **Steps**
1. Search for a package on **pkg.go.dev**.
2. Install it:

   ```bash
   go get github.com/username/package
   ```
3. Import it:

   ```go
   import "github.com/username/package"
   ```
4. Call its functions:

   ```go
   packageName.FunctionName()
   ```
5. Check the package documentation to learn available functions.

**Remember:**
* `go get` installs a package.
* `go.mod` stores all project dependencies.
* Third-party packages are used the same way as Go's standard packages.
