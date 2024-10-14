package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var totalBerat float64
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Masukan berat belanjaan di kedua kantong: ")
        scanner.Scan()
        input := scanner.Text()

        // Cek apakah input berupa dua buah bilangan real positif
        parts := strings.Split(input, " ")
        if len(parts) != 2 {
            fmt.Println("Input harus berupa dua buah bilangan real positif")
            continue
        }

        berat1, err1 := strconv.ParseFloat(parts[0], 64)
        berat2, err2 := strconv.ParseFloat(parts[1], 64)

        if err1 != nil || err2 != nil {
            fmt.Println("Input harus berupa bilangan real")
            continue
        }

        if berat1 < 0 || berat2 < 0 {
            fmt.Println("Proses selesai.")
            break
        }

        totalBerat += berat1 + berat2
        if totalBerat > 150 {
            fmt.Println("Proses selesai.")
            break
        }

        selisih := abs(berat1 - berat2)
        fmt.Printf("Sepeda motor pak Andi akan oleng: %v\n", selisih >= 9)
    }
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}