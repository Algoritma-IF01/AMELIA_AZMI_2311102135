# <h1 align="center">Laporan Praktikum 2</h1>
<p align="center">Amelia Azmi - 2311102135</p>


## Hello World

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var greetings string = "Selamat datang di dunia go!"
	var a, b int

	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}

```

### Output:

![Alt text](Hello.png)

## Hipotenusa

```go
package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	var hipotenusa bool

	fmt.Print("Masukkan nilai A: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai B: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan nilai C: ")
	fmt.Scanln(&c)
	hipotenusa = (c * c) == (a*a + b*b)
	fmt.Println("Sisi c adalah hipotenusa segitiga a, b, c: ", hipotenusa)
}

```

### Output:

![Alt text](Hipotenusa.png)

## Latihan 1

```go
package main

import (
	"fmt"
)

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukkan input String: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input String: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input String: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal: ", satu+" "+dua+" "+tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir: ", satu+" "+dua+" "+tiga)
}

```

### Output:

![Alt text](Latihan1.png)

## Latihan 2

```go
package main

import (
	"fmt"
)

func main() {
	var tahun int

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	if tahun%4 == 0 {
		if tahun%100 == 0 {
			if tahun%400 == 0 {
				fmt.Println("Tahun kabisat")
			} else {
				fmt.Println("Bukan tahun kabisat")
			}
		} else {
			fmt.Println("Tahun kabisat")
		}
	} else {
		fmt.Println("Bukan tahun kabisat")
	}
}
```

### Output:

![Alt text](Latihan2.png)

## Latihan 3

```go
package main

import "fmt"
import "math"

const pi = 3.1415926535

func main() {
    var radius int
    fmt.Print("Jarijari=")
    fmt.Scan(&radius)

    volume := 4.0/3.0 * pi * math.Pow(float64(radius), 3)
    surfaceArea := 4 * pi * math.Pow(float64(radius), 2)

    fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", radius, volume, surfaceArea)
}
```

### Output:

![Alt text](Latihan3.png)

## Latihan 4

```go

```

### Output:


## Latihan 5

```go

```

### Output:









