# Contraint Pada Type Parameter

Dalam Go, constraints (batasan) pada type parameters digunakan untuk mendefinisikan tipe-tipe apa saja yang dapat diterima oleh fungsi atau tipe generik. Constraints memastikan bahwa tipe yang digunakan memenuhi persyaratan tertentu sehingga kode dapat memanfaatkan fitur atau metode dari tipe tersebut.

Constraints ditentukan menggunakan interface. Misalnya, Anda bisa menentukan bahwa type parameter harus memenuhi interface tertentu, seperti berikut.

```go
package main

import "fmt"

// Definisikan interface yang digunakan sebagai constraint
type Adder interface {
    Add(a int, b int) int
}

// Struct yang memenuhi interface Adder
// Struct implementasi
type MyStruct struct{}

func (m MyStruct) Add(a int, b int) int {
    return a + b
}

// Fungsi generik dengan constraint Adder
func AddTwoValues[T Adder](adder T, a int, b int) int {
    return adder.Add(a, b)
}

func main() {
    myStruct := MyStruct{}
    result := AddTwoValues(myStruct, 3, 4)
    fmt.Println(result) // Output: 7
}
```

Pada contoh di atas:

1. **Adder** adalah interface yang memiliki metode **Add**.
2. **MyStruct** adalah struct yang mengimplementasikan interface **Adder**.
3. Fungsi **AddTwoValues** adalah fungsi generik dengan parameter tipe T yang memiliki constraint **Adder**, memastikan bahwa tipe T harus mengimplementasikan metode Add.

# Comparable

Dalam Go, comparable adalah sebuah **constraint built-in** yang digunakan untuk menunjukkan bahwa sebuah tipe dapat dibandingkan menggunakan operator perbandingan (== dan !=). Tipe-tipe yang dapat dibandingkan termasuk tipe-tipe dasar seperti **angka**, **string**, **pointer**, **array**, dan lain-lain. Tipe yang tidak dapat dibandingkan secara langsung, seperti **slice**, **map**, dan **fungsi**, tidak termasuk dalam **constraint comparable**.

Contoh penggunaan Comparable

```go
package main

import "fmt"

// Fungsi generik yang memeriksa apakah dua nilai sama
func AreEqual[T comparable](a, b T) bool {
    return a == b
}

func main() {
    fmt.Println(AreEqual(1, 1))           // Output: true
    fmt.Println(AreEqual(1, 2))           // Output: false
    fmt.Println(AreEqual("hello", "hello")) // Output: true
    fmt.Println(AreEqual("hello", "world")) // Output: false
}
```

Dalam contoh di atas:

Fungsi **AreEqual** adalah fungsi generik dengan type parameter T yang memiliki constraint **comparable**.
Fungsi ini memeriksa apakah dua nilai a dan b dari tipe T adalah sama menggunakan operator ==.
Tipe-tipe yang termasuk dalam comparable adalah:

1.  Tipe dasar:

    - bool
    - int, int8, int16, int32, int64
    - uint, uint8, uint16, uint32, uint64, uintptr
    - float32, float64
    - complex64, complex128
    - string

2.  Pointer:

    - Semua tipe pointer termasuk tipe unsafe.Pointer

3.  Array:

    - Array dengan elemen yang semuanya comparable

4.  Struct:

        - Struct dengan semua field yang comparable

Berikut adalah contoh tipe-tipe yang tidak termasuk dalam comparable:

1. Slice
2. Map
3. Function

Tipe-tipe ini tidak dapat dibandingkan langsung menggunakan operator == dan !=, sehingga tidak bisa digunakan dengan **constraint comparable**.

# Perbedaan Method dan Function

kalau function itu tanpa receiver struct `func SayHello[T any](name T) T {}`
Sedangkan kalau Method menggunakan receiver struct `func (d *Data[T]) SayHello(name T) T{}`
