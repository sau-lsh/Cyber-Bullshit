package calculate

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fraction struct {
	numerator   int
	denominator int
}

// NewFraction 创建一个分数
func NewFraction(numerator, denominator int) Fraction {
	if denominator == 0 {
		panic("denominator cannot be zero")
	}
	return Fraction{numerator, denominator}
}

func (f Fraction) Simplify() Fraction {
	gcd := GCD(f.numerator, f.denominator)
	return Fraction{f.numerator / gcd, f.denominator / gcd}
}

func (f Fraction) Add(other Fraction) Fraction {
	numerator := f.numerator*other.denominator + f.denominator*other.numerator
	denominator := f.denominator * other.denominator
	return NewFraction(numerator, denominator).Simplify()
}

func (f Fraction) Sub(other Fraction) Fraction {
	numerator := f.numerator*other.denominator - f.denominator*other.numerator
	denominator := f.denominator * other.denominator
	return NewFraction(numerator, denominator).Simplify()
}

func (f Fraction) Mul(other Fraction) Fraction {
	numerator := f.numerator * other.numerator
	denominator := f.denominator * other.denominator
	return NewFraction(numerator, denominator).Simplify()
}

func (f Fraction) Div(other Fraction) Fraction {
	numerator := f.numerator * other.denominator
	denominator := f.denominator * other.numerator
	return NewFraction(numerator, denominator).Simplify()
}

// String ParseFractionToString
func (f Fraction) String() string {
	if f.denominator == 1 {
		return strconv.Itoa(f.numerator)
	}
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// Calculator 四则运算计算器
type Calculator struct {
	expression string // the input expression
	index      int    // the current index of the expression
	token      string // the current token of the expression
}

// NewCalculator creates a new calculator with the given expression
func NewCalculator(expression string) *Calculator {
	return &Calculator{expression: expression, index: 0}
}

// NextToken 通过跳过空格前进到表达式的下一个标记并返回
func (c *Calculator) NextToken() string {
	c.skipWhitespace()
	if c.index >= len(c.expression) {
		c.token = ""
		return c.token
	}
	char := c.expression[c.index]
	switch char {
	case '+', '-', '*', '/', '(', ')':
		c.token = string(char)
		c.index++
	default:
		start := c.index
		for c.index < len(c.expression) && (c.isDigit(c.expression[c.index]) || c.expression[c.index] == '/') {
			c.index++
		}
		c.token = c.expression[start:c.index]
	}
	return c.token
}

// skipWhitespace 跳过表达式中的任何空字符，直到到达非空字符或表达式结尾
func (c *Calculator) skipWhitespace() {
	for c.index < len(c.expression) && c.isWhitespace(c.expression[c.index]) {
		c.index++
	}
}

func (c *Calculator) isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (c *Calculator) isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

// Parse 解析表达式并将结果作为分数返回
func (c *Calculator) Parse() Fraction {
	c.NextToken() // 获得第一个token
	result := c.parseExpression()
	if c.token != "" {
		panic("unexpected token: " + c.token)
	}
	return result
}

// parseExpression parses an expression of the form term (+|-) term (+|-) ...
func (c *Calculator) parseExpression() Fraction {
	result := c.parseTerm()
	for c.token == "+" || c.token == "-" {
		op := c.token
		c.NextToken()
		right := c.parseTerm()
		if op == "+" {
			result = result.Add(right)
		} else {
			result = result.Sub(right)
		}
	}
	return result
}

// parseTerm parses a term of the form factor (*|/) factor (*|/) ...
func (c *Calculator) parseTerm() Fraction {
	result := c.parseFactor()
	for c.token == "*" || c.token == "/" {
		op := c.token
		c.NextToken()
		right := c.parseFactor()
		if op == "*" {
			result = result.Mul(right)
		} else {
			result = result.Div(right)
		}
	}
	return result
}

// parseFactor parses a factor of the form number | (expression)
func (c *Calculator) parseFactor() Fraction {
	if c.token == "" {
		panic("unexpected end of expression")
	}
	if c.token == "(" {
		c.NextToken()
		result := c.parseExpression()
		if c.token != ")" {
			panic("expected ) but got " + c.token)
		}
		c.NextToken()
		return result
	}
	return c.parseNumber()
}

// parseNumber 解析 a/b 或 a 形式的数字，其中 a 和 b 是整数
func (c *Calculator) parseNumber() Fraction {
	parts := strings.Split(c.token, "/")
	if len(parts) > 2 {
		panic("invalid number format: " + c.token)
	}
	numerator, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("invalid numerator: " + parts[0])
	}
	denominator := 1
	if len(parts) == 2 {
		denominator, err = strconv.Atoi(parts[1])
		if err != nil {
			panic("invalid denominator: " + parts[1])
		}
	}
	c.NextToken()
	return NewFraction(numerator, denominator).Simplify()
}

// ToMixedFraction 将分数转换为 a'b/c 形式的带分数，其中 a、b 和 c 是整数
func (f Fraction) ToMixedFraction() string {
	if f.numerator == 0 {
		return "0"
	}
	if f.numerator < 0 {
		return "-" + NewFraction(-f.numerator, f.denominator).ToMixedFraction()
	}
	if f.numerator < f.denominator {
		return f.String()
	}
	quotient := f.numerator / f.denominator
	remainder := f.numerator % f.denominator
	if remainder == 0 {
		return strconv.Itoa(quotient)
	}
	return fmt.Sprintf("%d'%d/%d", quotient, remainder, f.denominator)
}

func GetAnswer() {
	source, err := os.Open("./formula.txt")
	if err != nil {
		fmt.Printf("Open file error:%v\n", err)
		return
	}
	defer source.Close()

	ans, err := os.Create("./ans.txt")
	if err != nil {
		fmt.Printf("Open file err:%v\n", err)
		return
	}
	defer ans.Close()

	scanner := bufio.NewScanner(source)
	res := make([]string, 0)
	for scanner.Scan() {
		expression := scanner.Text()
		calculator := NewCalculator(expression)
		result := calculator.Parse()
		r := expression + " = " + result.ToMixedFraction()
		res = append(res, r)
	}

	fmt.Fprint(ans, strings.Join(res, "\n"))
}
