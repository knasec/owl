package main

import (
    "strconv"
    "errors"
    "strings"
    "fmt"
)

func ValidateInt(str string, encoding string) (int, error){
    number, err := strconv.Atoi(str)

    if encoding == "ASCII"{
        asciiCodes := strings.Split(str, " ")

        var result string
        for _, code := range asciiCodes {
            charCode, err := strconv.Atoi(code)
            
            if err != nil {
                return 0, errors.New("Not a number")
            }
            result += string(rune(charCode))
        }
    }

    if (err == nil){
        return number, err
    }
    
    return 0, errors.New("Not a number")
}

func main() {
    number, err := ValidateInt("abcd", "utf8");
    fmt.Println(number, err)
}