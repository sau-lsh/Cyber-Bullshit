package getFormula

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Expression 一个算式
type Expression struct {
	left  *Expression // 左子式，如果是数字则为 nil
	right *Expression // 右子式，如果是数字则为 nil
	op    string      // 运算符，如果是数字则为空
	num   int         // 数字，如果是运算符则为 0
}

// String 将表达式转换为字符串
func (e *Expression) String() string {
	if e.left == nil && e.right == nil {
		return fmt.Sprintf("%d", e.num)
	}
	if e.op == "+" || e.op == "-" {
		return fmt.Sprintf("%s %s %s", e.left.paren(), e.op, e.right.paren())
	}
	return fmt.Sprintf("%s %s %s", e.left.noParen(), e.op, e.right.noParen())
}

// paren 如果表达式不是数字，则在表达式周围添加括号
func (e *Expression) paren() string {
	if e.left == nil && e.right == nil {
		return fmt.Sprintf("%d", e.num)
	}
	return fmt.Sprintf("(%s)", e)
}

// noParen 返回不带括号的表达式
func (e *Expression) noParen() string {
	if e.left == nil && e.right == nil {
		return fmt.Sprintf("%d", e.num)
	}
	if e.op == "+" || e.op == "-" {
		return fmt.Sprintf("(%s)", e)
	}
	return fmt.Sprintf("%s", e)
}

// randomInt 返回 1 到 10 之间的随机整数
func randomInt() int {
	return rand.Intn(10) + 1
}

// randomOp 返回 +, -, , 之间的随机运算符
func randomOp() string {
	ops := []string{"+", "-", "*", "/"}
	return ops[rand.Intn(len(ops))]
}

// randomExpr returns a random expression with a given depth
func randomExpr(depth int) *Expression {
	if depth == 0 {
		return &Expression{num: randomInt()}
	}
	return &Expression{
		left:  randomExpr(depth - 1),
		right: randomExpr(depth - 1),
		op:    randomOp(),
	}
}

func GenerateFormula() {
	rand.Seed(time.Now().UnixNano())
	file, err := os.Create("./formula.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < 10; i++ {
		expr := randomExpr(rand.Intn(3) + 1)
		fmt.Fprintln(file, expr)
	}
}
