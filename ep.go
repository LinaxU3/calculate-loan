package main

// EqualPrincipal 等额本金
//
// interestRate 年利率
// period 期数
// total 贷款总额
// months 第几个月后提前还款
//
// return 总共已还，提前还款还需付，已还利息，月供
func EqualPrincipal(interestRate float64, period int, total float64, months int) (float64, float64, float64, float64) {
	if months > period {
		months = period
	}
	//每月还本金
	benjin := total / float64(period)

	//月利率
	monthRate := interestRate / 12

	//计算第几个月需要还的利息
	getMonthInterest := func(months int) float64 {
		return (total - benjin*float64(months-1)) * monthRate
	}

	//已还
	a := make([]float64, months+1)
	//总共已还
	sum := float64(0)

	for i := 1; i <= months; i++ {
		a[i] = getMonthInterest(i) + benjin
		sum += a[i]
	}

	return sum, float64(period-months) * benjin, sum - benjin*float64(months), a[months]
}
