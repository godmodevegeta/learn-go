package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)
type person struct {
	name string
	age int
	pet string
}

func main() {

	fmt.Println("====DEFER====")
	fmt.Println("os.Args:", os.Args)
  fmt.Println(len(os.Args), os.Args[0])
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}


}
var opMap = map[string]func(int, int) int {
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func add(i int, j int) int {
	return i + j
}
func sub(i int, j int) int {
	return i - j
}
func mul(i int, j int) int {
	return i * j
}
func div(i int, j int) int {
	return i / j
}

func div_1(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num/denom
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base + v)
	}
	return out
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("Cannot divide by 0")
	}
	return num/denom, num % denom, nil
}

func f1(a string) int {
	return len(a)
}
func f2(a string) int {
	total := 0
	for _, val := range a{
		total += int(val)
		println(val)
	} 
	return total
}
func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

// func doSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
//	tx, err := db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//	defer func() {
//		if err == nil {
//			err = tx.Commit()
//		}
//		if err != nil {
//			tx.Rollback()
//		}
//	}()

//	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) $1", value1)
//	if err != nil {
//		return err
//	}
//	// use tx to do more db inserts here
//	return nil
//}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

//func prefixer(input string) func(input string) string {
//	return input + func() {} ()	
//}





