# Go lang Tutorial

## Notes:

### Main file
Main file is most important and means that apps runs from that file.
Remember to add function main to main.go file.

### Mod
To cread mod file for go application run
```
go mod init name_of_mode
```
mod contains the most important informations about application like for example go version.

### Variables
const - for constans
var for variables with changing values

### Printing strings

* Println for print string with next line 
* Printf for printing with variables(%v) inside 

### types with variables

for const it is possible to use var_name := value fr example 
```
banner := "banner"
```
type of data is always after var name for example 
```
var index int 
```

### Input
to use inputs remember about pointers to memory 
```
fmt.Scan(&variable_name)
```