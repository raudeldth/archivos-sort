package main

import (
    "fmt"
    "os"
    "bufio"
    "sort"
)

func main()  {
    var n int

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Ingresa la cantidad: ")
    fmt.Scan(&n)

    var strs []string
    for i := 0; i < n; i++ {
        scanner.Scan()
        str := scanner.Text()
        strs = append(strs, str)
    }

    sort.Strings(strs)

    file, err := os.Create("ascendente.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    for _, strAux := range strs {
        file.WriteString(strAux+"\n")
    }

    //descendente
    sort.Sort(sort.Reverse(sort.StringSlice(strs)))

    file2, err2 := os.Create("descendente.txt")
    if err2 != nil {
        fmt.Println(err2)
        return
    }

    defer file2.Close()

    for _, strAux := range strs {
        file2.WriteString(strAux+"\n")
    }

}
