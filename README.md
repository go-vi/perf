# perf
This repository contains go module which are helpful to perform logical operations without bunch of lines of code

##Slice
Slice package contains operations related to slice.
 - Currently, find() is supported for slice lookup.
 - Supports int and string slice type.

 *Syntax:*
 ```
    func Find([]byte, interface{}) (bool, error)
 ```
 
 *Example:*
 ```
    listInt := []int{1,2,3}
    
    b, _ := json.Marshal(listInt)
    
    found, err := Find(b, 1)
 ```

 
