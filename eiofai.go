package main

import "math"

// EqualInstallmentsOfPrincipalAndInterest 等额本息
//
// interestRate 年利率
// period 期数
// total 贷款总额
// months 第几个月后提前还款
//
// return 总共已还，提前还款还需付，已还利息，月供
func EqualInstallmentsOfPrincipalAndInterest(interestRate float64, period int, total float64, months int) (float64, float64, float64, float64) {
	if months > period {
		months = period
	}
	//月利率
	monthRate := interestRate / 12
	//月还款金额
	monthlyRepaymentAmount := total * monthRate * math.Pow(1+monthRate, float64(period)) / (math.Pow(1+monthRate, float64(period)) - 1)
	a := make([]float64, period+1)
	a[0] = total
	for i := 1; i <= months; i++ {
		a[i] = a[i-1]*(1+monthRate) - monthlyRepaymentAmount
	}
	//总共还款
	sum := float64(months)*monthlyRepaymentAmount + a[months]

	return monthlyRepaymentAmount * float64(months), a[months], sum - total, monthlyRepaymentAmount

}
