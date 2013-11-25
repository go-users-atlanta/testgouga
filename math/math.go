package math

func Calc(units, price float64) float64 {
	base := units * price
	tax := 0.06
	total := base + (base * tax)
	return total
}
