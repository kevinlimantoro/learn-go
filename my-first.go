package main

import (
	"container/list"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

const s string = "constant"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// type Person2 struct {
// 	Age string
// }

// func (p *Person2) Talk() {
// 	fmt.Println("Hi, my age is", p.Age)
// }

type Android struct {
	Person
	// Person2
	Model string
}

func main() {
	fmt.Println("hello world")

	fmt.Println("go" + "lang")
	// input := 0
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// fmt.Scanf("%f", input)
	// switch {
	// case input%2 == 0:
	// 	fmt.Println("Even")
	// default:
	// 	fmt.Println("Odds")
	// }
	// fmt.Println(input)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	fmt.Println(a, reflect.TypeOf(a))

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	var g string
	fmt.Println(g)

	b += b
	fmt.Println(s, b)

	const n = 500000000

	const dc = 3e20 / n
	fmt.Println(dc)

	fmt.Println(int64(dc))

	fmt.Println(math.Sin(n))

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"])
	name, ok := elements["Un"]
	fmt.Println(name, ok)
	for _, item := range elements {
		fmt.Println(item)
	}
	for k, v := range elements {
		fmt.Println(k, v)
	}
	items := map[string]interface{}{
		"foo": map[string]int{
			"strength": 10,
			"age":      2000,
		},
		"bar": map[string]int{
			"strength": 20,
			"age":      1000,
		},
	}
	for key, value := range items {
		fmt.Println("[", key, "] has items:")
		for k, v := range value.(map[string]int) {
			fmt.Println("\t-->", k, ":", v)
		}

	}
	add2 := func(x, y int) (result int, x2 int, y2 int) {
		return x + y, x, y
	}
	source := []int{1, 2, 3}
	dest := make([]int, 5)

	copy(source, dest)
	fmt.Println(source, dest)
	fmt.Println(add(1, 1))
	xs := []int{1, 2, 3}
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(xs...))
	fmt.Println(add2(1, 2))

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("fact", factorial(4))
	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str, "J")
	// 	fmt.Println(nextEven())
	// }()
	// panic("PANIC")

	zero5 := 5
	zero(&zero5)
	fmt.Println(zero5) // x is 0
	one1 := new(int)
	one(one1)
	fmt.Println("WAN", *one1) // x is 0

	test1, test2 := 1, 2
	fmt.Println(test1/2, test1%2 == 0, test2/2, test2%2 == 0, &test1)
	swap2(&test1, &test2)
	fmt.Println("swap", test1, test2)
	sqx := 1.5
	square(&sqx)
	fmt.Println(sqx)

	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	cir := Circle{x: 0, y: 0, r: 7}
	cir2 := Circle{0, 0, 5}
	cir3 := &Circle{0, 0, 5}
	fmt.Println("circle", cir, cir2, *cir3)
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cir))
	fmt.Println(circleArea2(&cir))
	fmt.Println(cir.area())
	rect := Rectangle{0, 0, 10, 10}
	fmt.Println(rect.area())

	and := new(Android)
	and.Person.Talk()
	and.Talk()

	fmt.Println(totalArea(&cir, &rect))

	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 0, 10, 10},
		},
	}
	fmt.Println(multiShape.area())

	fmt.Println(strings.ToUpper("as"))
	fmt.Println(strings.ToLower("AS"))

	bs, err := ioutil.ReadFile("package.json")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString("test")
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}

	kids := []Persona{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)

	byteArr := []byte("test")
	stringg := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(byteArr, stringg)

	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
	fmt.Println(getHash("package.json"))

	sh3 := sha1.New()
	h.Write([]byte("test"))
	bs2 := sh3.Sum([]byte{})
	fmt.Println(bs2)
	myjsonbyte, _ := json.Marshal(kids)
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	fmt.Println(string(myjsonbyte))
	var dat map[string]interface{}
	var dat2 []Persona
	json.Unmarshal(byt, &dat)
	json.Unmarshal(myjsonbyte, &dat2)
	fmt.Println(dat, dat2)
}
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func zero(xPtr *int) {
	fmt.Println(xPtr)
	*xPtr = 0
}
func one(xPtr *int) {
	fmt.Println(xPtr)
	*xPtr = 1
}
func square(x *float64) {
	*x = *x * *x
}
func swap(x, y *int) {
	z := new(int)
	*z = *x
	*x = *y
	*y = *z
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}
func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
