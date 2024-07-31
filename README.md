# Type Parameter

1. Untuk menandai bahwa sebuah function adalah bertipe generic, kita perlu menambahkan **Type Parameter** pada function tersebut.
2. Untuk membuat sebuah **Type Parameter**, dengan menggunakan simbol `[]` (_kurung kotak_). Nah, didalam kurung kotak tersebut, kita bisa menentukan nama **Type Parameter** nya.
3. Sama dengan bahasa pemograman lainnya seperti Java, Typescript. Type Parameter menggunakan satu huruf, contohnya T, V, K, dan lain-lain. Walaupun bisa saja lebih dari satu huruf.

# Type Constraint

1. Jika dibahasa pemograman seperti Java, C# dan lainnya, **Type Parameter** biasanya tidak perlu kita tentukan tipe datanya, berbeda dengan di Go-Lang.
2. Dari pengalaman yang dilakukan para pengembang Go-Lang, akhirnya di Go-Lang, **Type Parameter** wajib memiliki **constraint** (batasan).
3. **Type Constraint** merupakan aturan yang digunakan untuk menentukan tipe data yang diperbolehkan pada **Type Parameter**.
4. Contoh, jika kita ingin **Type Parameter** bisa digunakan untuk semua tipe data, kita bisa gunakan `interface{}` (kosong) sebagai **constraint** nya.
5. **Type Constraint** yang lebih detail akan kita bahas di materi **Type Sets**
6. Materi tambahan untuk pengetian Contraint, ada di `summary.md`

Kode: Type Constraint

```go
// function Length dengan parameter T yang memiliki
// constraint Interface kosong.
func Length[T interface{}]() {

}
```

# Tipe Data any

Di Go-Lang 1.18, diperkenalkan alias baru bernama any untuk interface{} (kosong), ini bisa mempermudah kita ketika membuat Type Parameter dengan constraint interface{}, jadi kita cukup gunakan constraint any

```go
// any adalah sebuah alias untuk interface{} kosong dan
// sama dengan interface{} di semua cara
type any = interface{}
```

Kode: Type Any

```go
func Length[T any]() {

}
```

## Menggunakan Type Parameter

- Setelah kita buat **Type Parameter** di function, selanjutnya kita bisa menggunakan **Type Parameter** tersebut sebagai tipe data di dalam function tersebut.
- Misal nya digunakan untuk return type atau function parameter
- Kita cukup gunakan nama **Type Parameter** nya saja
- **Type Parameter** hanya bisa digunakan di functionnya saja, tidak bisa digunakan di luar function.

```go

func Length[T any](param T) T {
    fmt.Println(param)
    return param
}

func TestLength(t *testing.T) {
    var result string = Length[string]("Budi")
    Length(result)

    var resultNumber int = Length[int](100)
    Length(resultNumber)

}
```

# Multiple Type Parameter

- Penggunakan **Type Parameter** bisa lebih dari satu, jika kita ingin menambahkan multiple **Type Parameter**, kita cukup gunakan tanda , (koma) sebagai pemisah antar **Type Parameter**.
- Nama **Type Parameter** harus berbeda, tidak boleh sama jika kita menambah **Type Parameter** lebih dari satu

```go

func MultipleParameter[T any, V any](param1 T, param2 V) {

}

func TestMultipleParameter(t *testing.T) {

    MultipleParameter[string, int]("Budi", 22)
    MultipleParameter[int, string](100, "Es Teh")
}
```

# Comparable

- Selain **any**, di Go-Lang versi 1.18 juga terdapat tipe data bernama **comparable**.
- comparable merupakan interface yang diimplementasikan oleh tipe data yang bisa dibandingkan (menggunakan operator != dan ==), seperti **booleans**, **numbers**, **strings**, **pointers**, **channels**, **interfaces**, **array** yang isinya ada comparable type, atau **structs** yang fields nya adalah comparable type.
- Materi tambahan untuk pengetian **Comparable**, ada di `summary.md`

# Type Parameter Inheritance

- Go-Lang sendiri sebenarnya tidak memiliki pewarisan, namun seperti kita ketahui, jika kita membuat sebuah type yang sesuai dengan kontrak interface, maka dianggap sebagai implementasi interface tersebut.
- **Type Parameter** juga mendukung hal serupa, kita bisa gunakan constraint dengan menggunakan interface, maka secara otomatis semua interface yang compatible dengan type constraint tersebut bisa kita gunakan.

# Type Sets

- Salah satu fitur yang menarik di Go-Lang Generic adalah Type Sets
- Dengan fitur ini, kita bisa menentukan lebih dari satu tipe constraint yang diperbolehkan pada type parameter.

## Membuat Type Sets

- Type Set adalah sebuah interface
- Cara membuat Type Set :

```go
type NamaTypeSet interface {
     P | Q | R
}
```

- TypeSet hanya bisa digunakan pada **type parameter**, tidak bisa digunakan sebagai **tipe data field** atau **variable**.
- Jika operator bisa digunakan di semua tipe data di dalam type set, maka operator tersebut bisa digunakan dalam kode generic.

Kode: Type Set Interface

```go
type Number interface{
    int | int8 | int16 | int32 | int64
    | float32 | float64
}

func Min[T Number](first T, second T) T {

    if first < second {

        return first
    } else {

        return second
    }

}
func TestMin(t *testing.T) {

    assert.Equal(t, int(100), Min[int](100, 200))
    assert.Equal(t, int64(100), Min[int64](100, 200))
    assert.Equal(t, float64(100.0), Min[float64](100.0, 200.0))
}
```

# Type Approximation

## Type Declaration

- Kadang, kita sering membuat **Type Declaration** di Golang untuk tipe data lain, misal kita membuat tipe data Age untuk tipe data int.
- Secara default, jika kita gunakan Age sebagai **type declaration** untuk int, lalu kita membuat **Type Set** yang berisi constraint int, maka tipe data Age dianggap tidak compatible dengan Type Set yang kita buat.

```go
Type Age int

type Number interface{
    int | int8 | int16 | int32 | int64
    | float32 | float64
}

func TestMin(t *testing.T) {

    assert.Equal(t, int(100), Min[int](100, 200))
    assert.Equal(t, int64(100), Min[int64](100, 200))
    assert.Equal(t, float64(100.0), Min[float64](100.0, 200.0))
    assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))
}

```

## Type Approximation

- Untungnya, Go-Lang memiliki feature bernama **Type Approximation**, dimana kita bisa menyebutkan bahwa semua constraint dengan tipe tersebut dan juga yang memiliki tipe dasarnya adalah tipe tersebut, maka bisa digunakan.
- Untuk menggunakan **Type Approximation**, kita bisa gunakan tanda `~` (**tilde**)

```go
// constraint Age
// memiliki tipe dasar int
Type Age int

type Number interface{
    ~int | int8 | int16 | int32 | int64
    | float32 | float64
}

```

# Type Inference

- Type Inference merupakan fitur dimana kita tidak perlu menyebutkan **Type Parameter** ketika memanggil kode generic.
- Tipe data **Type Parameter** bisa dibaca secara otomatis misal dari parameter yang kita kirim.
- Namun perlu diingat, pada beberapa kasus, jika terjadi error karena **Type Inference**, kita bisa dengan mudah memperbaikinya dengan cara menyebutkan **Type Parameter** nya saja.

```go

func TestMin(t *testing.T) {

    assert.Equal(t, int(100), Min[int](100, 200))
    assert.Equal(t, int64(100), Min[int64](100, 200))
    assert.Equal(t, float64(100.0), Min[float64](100.0, 200.0))
    assert.Equal(t, Age(100), Min[Age](Age(100), Age(200)))

}

func TestTypeInterface(t *testing.T) {

    assert.Equal(t, int(100), Min(100, 200))
    assert.Equal(t, int64(100), Min(100, 200))
    assert.Equal(t, float64(100.0), Min(100.0, 200.0))
    assert.Equal(t, Age(100), Min(Age(100), Age(200)))

}
```

# Generic Type

- Sebelumnya kita sudah bahas tentang generic di function.
- Generic juga bisa digunakan ketika membuat type.

Kode: Generic Type

```go
type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {

    for _, value := range bag {
        fmt.Println(value)
    }

}

func TestBag(t *testing.T) {

    numbers := Bag[int]{1, 2, 3, 4, 5}
    PrintBag[int](numbers)

    names := Bag[string]{"joko", "dono", "indro"}
    fmt.Println(names)
    PrintBag[string](numbers)
}

```

```go
type Bag[T any] []T

```

# Generic Struct

- Struct juga mendukung generic.
- Dengan menggunakan generic, kita bisa membuat Field dengan tipe data yang sesuai dengan **Type Parameter**.

```go
type Data[T any] struct {

    First T
    Second T

}

func TestData(t *testing.T) {

    data := Data[string] {
        First: "budi",
        Second: "Joko",
    }

    fmt.Println(data)
}

```

# Generic Method

- Selain di function, kita juga bisa tambahkan generic di method (function di struct) / **receiver**
- Namun, generic di method merupakan **generic yang terdapat di struct nya**
- Kita wajib menyebutkan semua **type parameter** yang terdapat di **struct**, walaupun tidak kita gunakan misalnya, atau jika tidak ingin kita gunakan, kita bisa gunakan `_` (**garis bawah**) sebagai pengganti **type parameter** nya
- Method tidak bisa memiliki type parameter yang mirip dengan di function.

Kode: Generic Method

```go
type Data[T any] struct {
    First T
    Second T
}

func (d * Data[_]) SayHello(name string) string {
    return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
    d.First = first
    return first
}


func TestGenericMethod(t *testing.T) {
    data := Data[string] {
        First: "budi",
        Second: "Joko",
    }

    assert.Equal(t, "Dior", data.ChangeFirst("Dior"))
    assert.Equal(t, "Hello Budi", data.SayHello("Budi"))
}
```

# Generic Interface

- Generic juga bisa kita gunakan di Interface.
- Secara otomatis, semua struct yang ingin mengikuti kontrak interface tersebut harus menggunakan generic juga.

```go
type GetterSetter[T any] interface {
    GetValue() T
    SetValue(Value T)
}

func ChangeValue[T any](param GeterSetter[T], value T) T {
    param.SetValue(value)
    return param.GetValue()
}

// Implementasi Struct

type MyData[T any] struct {
    Value T
}

func (myData *MyData[T]) GetValue() T {
    return myData.Value
}

func (myData *MyData[T]) SetValue(value T) {
    myData.Value = value
}

// Kode Test nya

func TestInteface(t *testing.T) {

    myData := MyData[string]{}
    result := ChangeValue[string](&myData, "Budi")

    assert.Equal(t, "Budi", result)
}
```

# In Line Type Constraint

- Sebelum-sebelumnya, kita selalu menggunakan **type declaration** atau **type set** ketika membuat **type constraint** di **type parameter**.
- Sebenarnya tidak ada kewajiban kita harus membuat **type declaration** atau **type set** jika kita ingin membuat **type parameter**, kita bisa gunakan secara langsung (**in line**) pada **type constraint**, misalnya di awal kita sudah bahas tentang `interface{}` (kosong), tapi kita selalu gunakan **type declaration** `any`
- Jika kita mau, kita juga bisa langsung gunakan **interface{ int | float32 | float64}** dibanding membuat **type set** **Number** misalnya.
- cocok digunakan jika case nya sederhana atau simple.

Kode: In Line Type Constraint

```go
func FindMin[T interface{ int | int64 | float64 }](first T, second T) T {
    if first < second {
        return first
    } else {
        return second
    }
}

func TestFindMin(t *testing.T) {

    assert.Equal(t, 100, FindMin(100, 200))
    assert.Equal(t, int64(100), FindMin(int64(100), int64(200)))
    assert.Equal(t, 100.2, FindMin(100.2, 300.2))
}

```

# Generic di Type Parameter

- Pada kasus tertentu, kadang ada kebutuhan kita menggunakan **type parameter** yang ternyata type tersebut juga generic atau memiliki **type parameter**.
- Kita juga bisa menggunakan **in line type constraint** agar lebih mudah, dengan cara menambahkan **type parameter** selanjutnya, misal.
- `[S interface{[]E}, E interface{}]`, artinya `S` harus berupa **slice** **element** `E`, dimana `E` boleh tipe apapun
- `[S []E, E any]`, artinya `S` harus **slice element** `E`, dimana `E` boleh tipe apapun.

Kode: Generic di Type Parameter

```go

func GetFirst[T []E, E any](data T) E {
    first := data[0]
    return first
}

func TestGetFirst(t *testing.T) {

    names:= []string{
        "Budi", "Joko", "Indro",
    }

    firstElement := GetFirst[[]string, string](names)
    assert.Equal(t, "Budi", firstElement)
}

```

# Experimental Package

- Saat versi Go-Lang 1.18, terdapat experimental package yang banyak menggunakan fitur Generic, namun belum resmi masuk ke Go-Lang Standard Library
- Kedepannya, karena ini masih experimental (percobaan), bisa jadi package ini akan berubah atau bahkan mungkin akan dihapus.
- https://pkg.go.dev/golang.org/x/exp
- Silahkan install sebagai dependency di Go Modules menggunakan perintah
  `go get golang.org/x/exp`

## Constraints Package

- Constraints Package berisi **type declaration** yang bisa kita gunakan untuk tipe data bawaan Go-Lang, misal **Number**, **Complex**, **Ordered**, dan lain-lain.
- https://pkg.go.dev/golang.org/x/exp/constraints

kode: Experimental Package

```go
import "golang.org/x/exp/constraints"

func ExperimentalMin[T constraints.Ordered](first T, second T) T {
    if first < second {
        return first
    } else {
        return second
    }
}

```

# Maps & Slices Packages

- Terdapat juga package maps dan slices, yang berisi function untuk mengelola data Map dan Slice, namun sudah menggunakan fitur Generic
- https://pkg.go.dev/golang.org/x/exp/maps
- https://pkg.go.dev/golang.org/x/exp/slices

Kode: Experimental Maps

```go

func TestExperimentalMap(t *testing.T) {

    first := map[string]string{
        "Name": "Budi",
    }

    second := map[string]string{
        "Name": "Budi",
    }

    assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlice(t *testing.T) {

    first := []string{"Budi"}

    second := []string{"Budi"}

    assert.True(t, slices.Equal(first, second))
}

```
