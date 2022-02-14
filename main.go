package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func fnA(a int) int {
	a = 0

	return 0
}

// func main() {
// 	fmt.Println(("I like Go!"))

// 	var a int
// 	var b int

// 	fmt.Println("vvedi 1 i 3 cherez probel")
// 	fmt.Scan(&a, &b)

// 	if a < b {
// 		fmt.Println("pervoe menwe vtorogo")
// 	} else if a > b {
// 		fmt.Println("pervoe bolwe vtorogo")
// 	} else if a == b {
// 		fmt.Println("pervoe i vtoroe ravno")
// 	}

// 	switch a {
// 	case 1:
// 		fmt.Println("a ravno 1")

// 	default:
// 		fmt.Println("a ne ravno 1")
// 	}

// 	var sum int

// 	for i := a; i <= b; i++ {
// 		sum += i
// 	}

// 	for i := 1; i <= 10; i++ {
// 		if i == 2 {
// 			continue
// 		}

// 		if i == 9 {
// 			break
// 		}
// 	}

// 	fmt.Println(sum)

// 	formatirovanyaStroka := fmt.Sprintf("%s", "eto stroka")

// 	fmt.Println(formatirovanyaStroka)

// 	test := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(test) // [1 2 3 4 5]

// 	for _, elem := range test {
// 		fmt.Println(elem)
// 	}

// 	var sl1 []int
// 	var sl2 []int = []int{1, 2, 3}
// 	sl3 := []int{1, 2, 3}
// 	sl4 := []int{1: 12}

// 	sl5 := sl3[:]

// 	fmt.Println(sl1)      // []
// 	fmt.Println(sl2)      // [1 2 3]
// 	fmt.Println(len(sl2)) // [1 2 3]
// 	fmt.Println(cap(sl2)) // [1 2 3]
// 	fmt.Println(sl3)      // [1 2 3]
// 	fmt.Println(sl4)      // [0 12]
// 	fmt.Println(sl5)

// 	sl5 = append(sl5, 4, 5, 6)

// 	fmt.Println(sl5)

// 	// 	a := []int{1, 2, 3}
// 	// b := make([]int, 3, 3)
// 	// n := copy(b, a)

// 	// fmt.Printf("a = %v\n", a)                  // a = [1 2 3]
// 	// fmt.Printf("b = %v\n", b)                  // b = [1 2 3]
// 	// fmt.Printf("Скопировано %d элемента\n", n) // Скопировано 3 элемента

// 	// СРЕЗ ЭТО УКАЗАТЕЛЬ НА МАССИВ - ВСЕГДА УКАЗАТЕЛЬ
// 	// SLICE - POINTER
// 	// ПРИМЕР

// 	c := &a

// 	fmt.Println(c)
// 	fmt.Println(*c)

// 	fmt.Println(
// 		// Содержится ли подстрока в строке
// 		strings.Contains("test", "es"),
// 		// результат: true

// 		// Кол-во подстрок в строке
// 		strings.Count("test", "t"),
// 		// результат: 2

// 		// Начинается ли строка с префикса
// 		strings.HasPrefix("test", "te"),
// 		// результат: true

// 		// Заканчивается ли строка суффиксом
// 		strings.HasSuffix("test", "st"),
// 		// результат: true

// 		// Возвращает начальный индекс подстроки в строке, а при отсутствии вхождения возвращает -1
// 		strings.Index("test", "e"),
// 		// результат: 1

// 		// объединяет массив строк через символ
// 		strings.Join([]string{"hello", "world"}, "-"),
// 		// результат: "hello-world"

// 		// Повторяет строку n раз подряд
// 		strings.Repeat("a", 5),
// 		// результат: "aaaaa"

// 		// Функция Replace заменяет любое вхождение old в вашей строке на new
// 		// Если значение n равно -1, то будут заменены все вхождения.
// 		// Общий вид: func Replace(s, old, new string, n int) string
// 		// Пример:
// 		strings.Replace("blanotblanot", "not", "***", -1),
// 		// результат: "bla***bla***"

// 		// Разбивает строку согласно разделителю
// 		strings.Split("a-b-c-d-e", "-"),
// 		// результат: []string{"a","b","c","d","e"}

// 		// Возвращает строку c нижним регистром
// 		strings.ToLower("TEST"),
// 		// результат: "test"

// 		// Возвращает строку c верхним регистром
// 		strings.ToUpper("test"),
// 		// результат: "TEST"

// 		// Возвращает строку с вырезанным набором
// 		strings.Trim("tetstet", "te"),
// 		// результат: s
// 	)

// 	err := errors.New("my error")
// 	fmt.Println("", err)

// 	// с помощью встроенной функции make:
// 	m1 := make(map[int]int)

// 	// с помощью использования литерала отображения:
// 	m2 := map[int]int{
// 		// Пары ключ:значение указываются при необходимости
// 		12: 2,
// 		1:  5,
// 	}

// 	fmt.Println(m1) // map[]
// 	fmt.Println(m2) // map[1:5 12:2]

// 	m := map[int]int{
// 		1: 10,
// 	}

// 	if value, inMap := m[1]; inMap {
// 		fmt.Println(value) // 10
// 	}

// 	if value, inMap := m[2]; inMap {
// 		fmt.Println(value) // Условие не выполняется
// 	}

// 	mapName := make(map[string]string)

// 	for key, value := range mapName {
// 		fmt.Println(key, value)
// 	}

// 	m = map[int]int{
// 		1: 10,
// 		2: 20,
// 		3: 30,
// 	}
// 	fmt.Println(len(m)) // 3
// 	preobrazovanieStrok()
// }

func test(x1 *int, x2 *int) {
	// здесь ваш код
	nX1 := x1
	x1 = x2
	x2 = nX1

	fmt.Println(x1, x2)
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func preobrazovanieStrok() {
	a := "str"

	b := []byte(a)

	c := string(b)

	fmt.Println(a) // str

	fmt.Println(b) // [115 116 114] - побайтовый срез

	fmt.Println(c) // str
}

func preobrazRune() {
	a := "строка"

	b := []rune(a) // срез рун

	c := string(b)

	fmt.Println(a) // строка

	fmt.Println(b) // [1089 1090 1088 1086 1082 1072] - срез рун

	fmt.Println(c) // строка
}

func izNumToString() {
	a := strconv.Itoa(2020) // int -> string
	fmt.Printf("%T \n", a)  // тип - string
	fmt.Println(a)          // 2020
}

func isStringVNum() {
	foo := "10"
	bar := "15"
	barInt, err := strconv.Atoi(bar)
	if err != nil {
		panic(err)
	}
	fooInt, err := strconv.Atoi(foo)
	if err != nil {
		panic(err)
	}
	baz := barInt - fooInt
	fmt.Println(baz) //5
}

func isStringVFloat() {
	s := "23.23456"
	// как и в прошлом шаге, здесь 2 параметр - bitSize
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	// но нужно понимать что метод все равно вернет float64
	result, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)         // 23.23456
	fmt.Printf("%T \n", result) // float64

	// Конкретный пример для разных bitSize:
	s = "1.0000000012345678"
	//  не будем обрабатывать ошибки в примерах, но на практике так не делайте ;)
	result32, _ := strconv.ParseFloat(s, 32)
	result64, _ := strconv.ParseFloat(s, 64)
	fmt.Println("bitSize32:", result32) // вывод 1 (не уместились)
	fmt.Println("bitSize64:", result64) // вывод  1.0000000012345678
}

func fileHelp() {
	dataForFile := []byte("Тестовая строка, предназначенная для записи в файл")

	// Создаем новый файл и записываем в него данные dataForFile
	if err := ioutil.WriteFile("test.txt", dataForFile, 0600); err != nil {
	}

	// Читаем данные из того же файла
	dataFromFile, err := ioutil.ReadFile("test.txt")
	if err != nil {
	}

	// Сравниваем исходные данные с записанными в файл и прочитанными из него
	fmt.Printf("dataForFile == dataFromFile: %v\n", bytes.Equal(dataFromFile, dataForFile))

	// Изучаем содержимое директории
	filesFromDir, err := ioutil.ReadDir(".")
	if err != nil {
	}

	for _, file := range filesFromDir {
		// Проходим по всем найденным файлам и печатаем их имя и размер
		fmt.Printf("name: %s, size: %d\n", file.Name(), file.Size())
	}

	// Output:
	// dataForFile == dataFromFile: true
	// name: main.go, size: 727
	// name: test.txt, size: 93
}

func fileHelpOs() {
	// создаем файл
	os.Create("text.txt")
	// переименовываем файл
	os.Rename("text.txt", "new_text.txt")
	// удаляем файл
	os.Remove("new_text.txt")
	// кстати, os позволяет работать не только с файлами
	// выходим из программы:
	os.Exit(0)
	file1, _ := os.Create("text.txt")
	file2, _ := os.Create("text.txt")
	info1, _ := file1.Stat() // функция Stat возвращает информацию о файле и ошибку
	info2, _ := file2.Stat()
	fmt.Println(os.SameFile(info1, info2)) // true

	// вот что мы можем получить из FileInfo:
	// A FileInfo describes a file and is returned by Stat and Lstat.
	type FileInfo interface {
		Name() string // base name of the file
		Size() int64  // length in bytes for regular files; system-dependent for others
		// Mode() FileMode     // file mode bits?
		ModTime() time.Time // modification time
		IsDir() bool        // abbreviation for Mode().IsDir()
		Sys() interface{}   // underlying data source (can return nil)
	}

	file1, _ = os.Create("text.txt")
	file1.WriteString("1 строка \n")
	file1.WriteString("2 строка \n")
	file1.Close()

	// внутри файла будет:
	// 1 строка
	// 2 строка
}

func fileHelpBufio() {
	file, err := os.Open("test.txt")
	if err != nil {
		// ...
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	buf := make([]byte, 10)
	n, err := rd.Read(buf) // читаем в buf 10 байт из ранее открытого файла
	if err != nil && err != io.EOF {
		// io.EOF не совсем ошибка - это состояние, указывающее, что файл прочитан до конца
		// ...
	}
	fmt.Printf("прочитано %d байт: %s\n", n, buf) // прочитано 10 байт: bufio ...

	// 	s, err := rd.ReadString('\n') // читаем данные до разрыва абзаца ('\n')
	// 	fmt.Printf("%s\n", s)         // ... здесь будет строка

	// 	file, err := os.Create("test.txt")
	// if err != nil {
	// 	...
	// }
	// defer file.Close()

	// w := bufio.NewWriter(file)
	// n, err := w.WriteString("Запишем строку")
	// if err != nil {
	// 	...
	// }
	// fmt.Printf("Записано %d байт\n", n) // Записано 27 байт

	// // bufio.Writer имеет собственный буфер, чтобы быть уверенным, что данные точно записаны,
	// // вызываем метод Flush()
	// w.Flush()

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// s := bufio.NewScanner(file)

	// // Я заранее записал в файл 5 цифр, каждую на новой строке
	// for s.Scan() { // возвращает true, пока файл не будет прочитан до конца
	// 	fmt.Printf("%s\n", s.Text()) // s.Text() содержит данные, считанные на данной итерации
	// }

	// // 1
	// // 2
	// // 3
	// // 4
	// // 5
}

func helpFileCSV() {
	// Записывать данные, а в дальнейшем читать их мы будем из буфера,
	// но его можно заменить любым другим объектом, удовлетворяющим
	// интерфейсу io.ReadWriter
	buf := bytes.NewBuffer(nil)

	w := csv.NewWriter(buf)

	for i := 1; i <= 3; i++ {
		// Запись данных может производится поэтапно, например в цикле
		val1 := fmt.Sprintf("row %d col 1", i)
		val2 := fmt.Sprintf("row %d col 2", i)
		val3 := fmt.Sprintf("row %d col 3", i)
		if err := w.Write([]string{val1, val2, val3}); err != nil { // Аргументом Write является срез строк
			// ...
		}
	}
	w.Flush() // Этот метод приведет к фактической записи данных из буфера csv.Writer в buf

	// Либо данные можно записать за один раз
	w.WriteAll([][]string{ // Аргументом WriteAll является срез срезов строк
		{"row 4 col 1", "row 4 col 2", "row 4 col 3"},
		{"row 5 col 1", "row 5 col 2", "row 5 col 3"},
	})

	r := csv.NewReader(buf)

	for i := 1; i <= 2; i++ {
		// Читать данные мы тоже можем построчно, получая срез строк за каждую итерацию
		row, err := r.Read()
		if err != nil && err != io.EOF { // Здесь тоже нужно учитывать конец файла
			// ...
		}
		fmt.Println(row)
	}

	// Либо прочитать данные за один раз
	data, err := r.ReadAll()
	if err != nil {
		// Когда мы читаем данные до конца файла io.EOF не возвращается, а служит сигналом к завершению чтения
		// ...
	}

	for _, row := range data {
		fmt.Println(row)
	}

	// [row 1 col 1 row 1 col 2 row 1 col 3]
	// [row 2 col 1 row 2 col 2 row 2 col 3]
	// [row 3 col 1 row 3 col 2 row 3 col 3]
	// [row 4 col 1 row 4 col 2 row 4 col 3]
	// [row 5 col 1 row 5 col 2 row 5 col 3]
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func jsonHelp() {
	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
	// и возвращает байтовый срез с данными, кодированными в формат JSON.
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}
	data = []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v", s) // {John Connor 35 true [15 11 37]}

	type user struct {
		Name     string
		Email    string
		Status   bool
		Language []byte
	}

	m := user{Name: "John Connor", Email: "test email"}

	data, _ = json.Marshal(m)

	data = bytes.Trim(data, "{") // испортим json удалив '{'

	// функция json.Valid возвращает bool, true - если json правильный
	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}

	fmt.Printf("%s", data) // вывод: "Name":"John Connor","Email":"test email","Status":false,"Language":null}

}

func mySecondJsonHelp() {
	// в общем виде аннотация выглядит так: `json:"используемое_имя,опция,опция"`

	type myStruct struct {
		// при кодировании / декодировании будет использовано имя name, а не Name
		Name string `json:"name"`

		// при кодировании / декодировании будет использовано то же имя (Age),
		// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
		// то при кодировании оно будет опущено
		Age int `json:",omitempty"`

		// при кодировании / декодировании поле всегда игнорируется
		Status bool `json:"-"`
	}

	m := myStruct{Name: "John Connor", Age: 0, Status: true}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // {"name":"John Connor"}
}

func timeHelp() {
	// func Now() Time
	// возвращает текущую дату и время
	now := time.Now()

	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	// возвращает дату и время в соответствии с заданными параметрами: годом, месяцем, днем, временем и пр.
	currentTime := time.Date(
		2020,     // год
		time.May, // месяц
		15,       // день
		10,       // часы
		13,       // минуты
		12,       // секунды
		45,       // наносекунды
		time.UTC, // временная зона
	)

	// func Unix(sec int64, nsec int64) Time
	// возвращает дату и время в соответствии с заданными параметрами: секундами и наносекундами, прошедшими с начала эпохи Unix — 01.01.1970 г.
	// https://ru.wikipedia.org/wiki/Unix-%D0%B2%D1%80%D0%B5%D0%BC%D1%8F
	unixTime := time.Unix(
		150000, // секунды
		1,      // наносекунды
	)

	fmt.Println(now.Format("02-01-2006 15:04:05"))         // 15-05-2020 09:58:16
	fmt.Println(currentTime.Format("02-01-2006 15:04:05")) // 15-05-2020 10:13:12
	fmt.Println(unixTime.Format("02-01-2006 15:04:05"))    // 02-01-1970 22:40:00

	// func Parse(layout, value string) (Time, error)
	// парсит дату и время в строковом представлении
	firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
	if err != nil {
		panic(err)
	}

	// LoadLocation находит временную зону в справочнике IANA
	// https://www.iana.org/time-zones
	loc, err := time.LoadLocation("Asia/Yekaterinburg")
	if err != nil {
		panic(err)
	}

	// func ParseInLocation(layout, value string, loc *Location) (Time, error)
	// парсит дату и время в строковом представлении с отдельным указанием временной зоны
	secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(firstTime.Format("02-01-2006 15:04:05"))  // 15-05-2020 17:45:00
	fmt.Println(secondTime.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:10

	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// func (t Time) Date() (year int, month Month, day int)
	fmt.Println(current.Date()) // 2020 May 15

	// func (t Time) Year() int
	fmt.Println(current.Year()) // 2020

	// func (t Time) Month() Month
	fmt.Println(current.Month()) // May

	// func (t Time) Day() int
	fmt.Println(current.Day()) // 15

	// func (t Time) Clock() (hour, min, sec int)
	fmt.Println(current.Clock()) // 17 45 12

	// func (t Time) Hour() int
	fmt.Println(current.Hour()) //17

	// func (t Time) Minute() int
	fmt.Println(current.Minute()) // 45

	// func (t Time) Second() int
	fmt.Println(current.Second()) // 12

	// func (t Time) Unix() int64
	fmt.Println(current.Unix()) // 1589546712

	// func (t Time) Weekday() Weekday
	fmt.Println(current.Weekday()) // Friday

	// func (t Time) YearDay() int
	fmt.Println(current.YearDay()) // 136

	// func (t Time) Format(layout string) string
	current = time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	fmt.Println(current.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:12

	firstTime = time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	secondTime = time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

	// func (t Time) After(u Time) bool
	// true если позже
	fmt.Println(firstTime.After(secondTime)) // true

	// func (t Time) Before(u Time) bool
	// true если раньше
	fmt.Println(firstTime.Before(secondTime)) // false

	// func (t Time) Equal(u Time) bool
	// true если равны
	fmt.Println(firstTime.Equal(secondTime)) // false

	now = time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// func (t Time) Add(d Duration) Time
	// изменяет дату в соответствии с параметром - "продолжительностью"
	future := now.Add(time.Hour * 12) // перемещаемся на 12 часов вперед

	// func (t Time) AddDate(years int, months int, days int) Time
	// изменяет дату в соответствии с параметрами - количеством лет, месяцев и дней
	past := now.AddDate(-1, -2, -3) // перемещаемся на 1 год, два месяца и 3 дня назад

	// func (t Time) Sub(u Time) Duration
	// вычисляет время, прошедшее между двумя датами
	fmt.Println(future.Sub(past)) // 10332h0m0s
}

func timeHelpDur() {
	now := time.Now()
	past := now.AddDate(0, 0, -1)
	future := now.AddDate(0, 0, 1)

	// func Since(t Time) Duration
	// вычисляет период между текущим моментом и заданным временем в прошлом
	fmt.Println(time.Since(past).Round(time.Second)) // 24h0m0s

	// func Until(t Time) Duration
	// вычисляет период между текущим моментом и заданным временем в будущем
	fmt.Println(time.Until(future).Round(time.Second)) // 24h0m0s

	// func ParseDuration(s string) (Duration, error)
	// преобразует строку в Duration с использованием аннотаций:
	// "ns" - наносекунды,
	// "us" - микросекунды,
	// "ms" - миллисекунды,
	// "s" - секунды,
	// "m" - минуты,
	// "h" - часы.
	dur, err := time.ParseDuration("1h12m3s")
	if err != nil {
		panic(err)
	}
	fmt.Println(dur.Round(time.Hour).Hours()) // 1
}

// func main() {
// 	go myFunc()
// 	time.Sleep(1 * time.Second) // Пауза в 1 секунду
// }

// func myFunc() {
// 	fmt.Println("hello")
// }

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}
