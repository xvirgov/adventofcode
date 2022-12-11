package day11

import (
	"fmt"
	"main/common"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Operation func(int) int
type Test func(int) int

type OperationBig func(int2 big.Int) big.Int
type TestBig func(int2 big.Int) int

type Monkey struct {
	items     []int
	operation Operation
	test      Test
	inspected int
}

type MonkeyBig struct {
	items     []*big.Int
	operation OperationBig
	test      TestBig
	inspected int
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func sliceAtoiBig(sa []string) ([]*big.Int, error) {
	si := make([]*big.Int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		bigI := big.NewInt(int64(i))
		if err != nil {
			return si, err
		}
		si = append(si, bigI)
	}
	return si, nil
}

func deserializeMonkey(monkey []string) Monkey {
	//fmt.Println(monkey)

	items, _ := sliceAtoi(strings.Split(strings.TrimPrefix(monkey[1], "  Starting items: "), ", "))
	monkeyStruct := Monkey{items: items}

	trimmedOp := strings.TrimPrefix(monkey[2], "  Operation: new = old ")
	op := trimmedOp[0]
	val, _ := strconv.Atoi(trimmedOp[2:])
	switch op {
	case '+':
		if trimmedOp[2:] == "old" {
			monkeyStruct.operation = func(i int) int {
				return i + i
			}
		} else {
			monkeyStruct.operation = func(i int) int {
				return i + val
			}
		}
	case '-':
		if trimmedOp[2:] == "old" {
			monkeyStruct.operation = func(i int) int {
				return i - i
			}
		} else {
			monkeyStruct.operation = func(i int) int {
				return i - val
			}
		}
	case '*':
		if trimmedOp[2:] == "old" {
			monkeyStruct.operation = func(i int) int {
				return i * i
			}
		} else {
			monkeyStruct.operation = func(i int) int {
				return i * val
			}
		}
	case '/':
		if trimmedOp[2:] == "old" {
			monkeyStruct.operation = func(i int) int {
				return i / i
			}
		} else {
			monkeyStruct.operation = func(i int) int {
				return i / val
			}
		}
	}

	divisibleBy, _ := strconv.Atoi(strings.TrimPrefix(monkey[3], "  Test: divisible by "))
	ifTrue, _ := strconv.Atoi(strings.TrimPrefix(monkey[4], "    If true: throw to monkey "))
	ifFalse, _ := strconv.Atoi(strings.TrimPrefix(monkey[5], "    If false: throw to monkey "))

	monkeyStruct.test = func(i int) int {
		if i%divisibleBy == 0 {
			return ifTrue
		}
		return ifFalse
	}

	monkeyStruct.inspected = 0

	return monkeyStruct
}

func deserializeMonkeyBig(monkey []string) MonkeyBig {
	//fmt.Println(monkey)

	items, _ := sliceAtoiBig(strings.Split(strings.TrimPrefix(monkey[1], "  Starting items: "), ", "))
	monkeyStruct := MonkeyBig{items: items}

	trimmedOp := strings.TrimPrefix(monkey[2], "  Operation: new = old ")
	op := trimmedOp[0]
	val, _ := strconv.Atoi(trimmedOp[2:])
	valBig := big.NewInt(int64(val))
	switch op {
	case '+':
		if trimmedOp[2:] == "old" {
			monkeyStruct.operation = func(i big.Int) big.Int {
				res := big.NewInt(0)
				res.Add(&i, &i)
				return *res
			}
		} else {
			monkeyStruct.operation = func(i big.Int) big.Int {
				res := big.NewInt(0)
				res.Add(&i, valBig)
				return *res
			}
		}
	case '*':
		if trimmedOp[2:] == "old" {
			monkeyStruct.operation = func(i big.Int) big.Int {
				res := big.NewInt(0)
				res.Mul(&i, &i)
				return *res
			}
		} else {
			monkeyStruct.operation = func(i big.Int) big.Int {
				res := big.NewInt(0)
				res.Mul(&i, valBig)
				return *res
			}
		}
	}

	divisibleBy, _ := strconv.Atoi(strings.TrimPrefix(monkey[3], "  Test: divisible by "))
	ifTrue, _ := strconv.Atoi(strings.TrimPrefix(monkey[4], "    If true: throw to monkey "))
	ifFalse, _ := strconv.Atoi(strings.TrimPrefix(monkey[5], "    If false: throw to monkey "))

	monkeyStruct.test = func(i big.Int) int {
		res := big.NewInt(0)
		res.Mod(&i, big.NewInt(int64(divisibleBy)))
		if res.Cmp(big.NewInt(0)) == 0 {
			return ifTrue
		}
		return ifFalse
	}

	monkeyStruct.inspected = 0

	return monkeyStruct
}

func doMonkeyRound(monkeys []Monkey) []Monkey {

	for i, monkey := range monkeys {
		//fmt.Println("Monkey ", i)
		for _, item := range monkey.items {
			//fmt.Println("  Monkey inspects an item with a worry level of ", item)
			worryAfterInspection := monkey.operation(item)
			//fmt.Println("    Worry level is multiplied to ", worryAfterInspection)
			dividedWorry := worryAfterInspection / 3
			//fmt.Println("    Monkey gets bored with item. Worry level is divided by 3 to ", dividedWorry)
			monkeyIndexThatGetsItem := monkey.test(dividedWorry)
			//fmt.Println("    Item with worry level", dividedWorry, "is thrown to monkey ", monkeyIndexThatGetsItem)
			monkeys[monkeyIndexThatGetsItem].items = append(monkeys[monkeyIndexThatGetsItem].items, dividedWorry)
		}
		monkeys[i].inspected += len(monkeys[i].items)
		monkeys[i].items = nil
	}

	return monkeys
}

func doMonkeyRoundModWorried(monkeys []Monkey) []Monkey {

	for i, monkey := range monkeys {
		//fmt.Println("Monkey ", i)
		for _, item := range monkey.items {
			//fmt.Println("  Monkey inspects an item with a worry level of ", item)
			worryAfterInspection := monkey.operation(item)
			//fmt.Println("    Worry level is multiplied to ", worryAfterInspection)
			//dividedWorry := worryAfterInspection / 3
			//fmt.Println("    Monkey gets bored with item. Worry level is divided by 3 to ", dividedWorry)
			monkeyIndexThatGetsItem := monkey.test(worryAfterInspection)
			//fmt.Println("    Item with worry level", dividedWorry, "is thrown to monkey ", monkeyIndexThatGetsItem)
			//monkeys[monkeyIndexThatGetsItem].items = append(monkeys[monkeyIndexThatGetsItem].items, worryAfterInspection%96577) // LCM of test data
			monkeys[monkeyIndexThatGetsItem].items = append(monkeys[monkeyIndexThatGetsItem].items, worryAfterInspection%9699690) // LCM of input
		}
		monkeys[i].inspected += len(monkeys[i].items)
		monkeys[i].items = nil
	}

	return monkeys
}

func doMonkeyRoundWorried(monkeys []MonkeyBig) []MonkeyBig {

	for i, monkey := range monkeys {
		//fmt.Println("Monkey ", i)
		for _, item := range monkey.items {
			//fmt.Println("  Monkey inspects an item with a worry level of ", item)
			worryAfterInspection := monkey.operation(*item)
			//fmt.Println("    Worry level is multiplied to ", worryAfterInspection)
			//dividedWorry := worryAfterInspection / 3
			//dividedWorry := big.NewInt(0)
			//dividedWorry.Div(&worryAfterInspection, big.NewInt(3))
			//fmt.Println("    Monkey gets bored with item. Worry level is divided by 3 to ", dividedWorry)
			monkeyIndexThatGetsItem := monkey.test(worryAfterInspection)
			//fmt.Println("    Item with worry level", dividedWorry, "is thrown to monkey ", monkeyIndexThatGetsItem)
			monkeys[monkeyIndexThatGetsItem].items = append(monkeys[monkeyIndexThatGetsItem].items, &worryAfterInspection)
		}
		monkeys[i].inspected += len(monkeys[i].items)
		monkeys[i].items = nil
	}

	return monkeys
}

func partOne() {
	monkeysInput := common.LinesInFile("adventofcode2022/day11/day11.txt")

	monkeys := []Monkey{}

	for i := 0; i < len(monkeysInput); i += 7 {
		monkeys = append(monkeys, deserializeMonkey(monkeysInput[i:i+6]))
	}

	for i := 0; i < 20; i++ {
		monkeys = doMonkeyRound(monkeys)
	}

	arr := []int{}
	for _, monkey := range monkeys {
		//fmt.Println(monkey)
		arr = append(arr, monkey.inspected)
	}

	sort.Ints(arr)
	monkeyBusiness := arr[len(arr)-1] * arr[len(arr)-2]

	fmt.Println(monkeyBusiness)
}

func partTwo() {
	monkeysInput := common.LinesInFile("adventofcode2022/day11/day11.txt")

	monkeys := []Monkey{}

	for i := 0; i < len(monkeysInput); i += 7 {
		monkeys = append(monkeys, deserializeMonkey(monkeysInput[i:i+6]))
	}

	for i := 0; i < 10000; i++ {
		monkeys = doMonkeyRoundModWorried(monkeys)
	}

	arr := []int{}
	for _, monkey := range monkeys {
		//fmt.Println(monkey)
		arr = append(arr, monkey.inspected)
	}

	sort.Ints(arr)
	monkeyBusiness := arr[len(arr)-1] * arr[len(arr)-2]

	fmt.Println(monkeyBusiness)
}

//func partTwo() {
//	monkeysInput := common.LinesInFile("adventofcode2022/day11/day11.txt")
//
//	monkeys := []MonkeyBig{}
//
//	for i := 0; i < len(monkeysInput); i += 7 {
//		monkeys = append(monkeys, deserializeMonkeyBig(monkeysInput[i:i+6]))
//	}
//
//	for i := 0; i < 10000; i++ {
//		monkeys = doMonkeyRoundWorried(monkeys)
//	}
//
//	arr := []int{}
//	for i, monkey := range monkeys {
//		//fmt.Println(monkey)
//		fmt.Print("Monkey ", i, ": ")
//		//for _, item := range monkey.items {
//		//	fmt.Print(item.String(), ", ")
//		//}
//		fmt.Println(" ---- ", monkey.inspected)
//		arr = append(arr, monkey.inspected)
//	}
//
//	sort.Ints(arr)
//	monkeyBusiness := arr[len(arr)-1] * arr[len(arr)-2]
//
//	fmt.Println(monkeyBusiness)
//}

func DayEleven() {
	partOne()
	partTwo()
}
