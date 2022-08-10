package main

import (
  "fmt"
  "strconv"
  "os"
)



func main() {

	var min, max, length_max int
	var result string
	var err, err_2 error
	file_name := "digital_passwords_list.txt"
    min_default := 0
    max_default := 99999999

	fmt.Println("\r\nWelcome to the program for generating numerical passwords.\r\n")
	fmt.Print("Enter the number from which the generation of numbers will begin: \n>>> ")

    _, err = fmt.Scanf("%d", &min)
    if err != nil{
        fmt.Println("This is not a number! Used by default 0.")
        min = min_default
    }
    if min < 0 {
    	min = min_default
    	fmt.Println("The minimum number should be larger or equal to zero!")
    	fmt.Printf("Used by default %d\n", min_default)
    }

    fmt.Printf("Enter the maximum number to which the generation will end (maximum %d): \n>>> ", max_default)
    _, err_2 = fmt.Fscan(os.Stdin, &max) 
    if err_2 != nil{
        fmt.Printf("This is not a number! Used by default %d.", max_default)
        max = max_default
    }
    if max <= 0 {
    	max = max_default
    	fmt.Println("The maximum number should be more than zero!")
    	fmt.Printf("Used by default %d\n", max_default)
    }

    if max > max_default {
    	max = max_default
    	fmt.Printf("The maximum number should not be more than %d", max_default)
    	fmt.Printf("Used by default %d", max_default)
    }

    

    if min >= max {
    	min = min_default
    }



    length_max = len(strconv.Itoa(max))

	file, err := os.Create(file_name)
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
    } else {
    	fmt.Printf("File \"%s\" was successfully created!", file_name)
    }
    defer file.Close() 
    
	for i := min; i <= max; i++ {
		str := strconv.Itoa(i)
		new_str := str
		length_of_number := len(str)

		if length_of_number < length_max {
			for len(new_str) < length_max {
				new_str = "0" + new_str
			}
			result += new_str + "\r\n"
		} else {
			if i < max {
			    result += str + "\r\n"
			} else {
				result += str
			}
		}
		
	}

	file.WriteString(result)
}


