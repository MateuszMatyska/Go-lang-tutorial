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

### import 
import packages
import (
    "fmt"
    "strings"
)

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

### Arrays and Slices 

#### Arrays
Arrays has got a size. To add element to array you need to know index.
```
var namesArray1 = [3]string{"Tom", "Arthur", "John"}

var namesArray2 [3]string
namesArray2[0] = "Elliot"
namesArray2[1] = "Angela"
namesArray2[2] = "Darlene"
```

#### Slices 
Dynamics array you don't need to know what size array will have
```
var namesArray3 []string
namesArray3 = append(namesArray3, "Michael")
namesArray3 = append(namesArray3, "Jim")
namesArray3 = append(namesArray3, "Dwight")
namesArray3 = append(namesArray3, "Pam")
```

### Loops
In go is only for 

_ means variable which you don't want to use.

```
for index, item := range array {
    //action
}
```

#### Infinite loop
```
for {
    action
}
```

#### number range loop
```
for i := 1; i < 5; i++ {
	action
}
```

