package main

import (
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/fs"
	"io/ioutil"
	"log"
)

const (
	Print = iota
	Csv
)

var (
	interestRate = flag.Float64("i", 0, "年利率，单位：百分比")
	period       = flag.Int("p", 0, "贷款期数，单位：月")
	total        = flag.Float64("t", 0, "贷款总额，单位：元")

	t = flag.Int("type", 0, "生成类型，0：打印；1：生成CSV数据文件")
)

func main() {
	flag.Parse()
	if *interestRate == 0 || *period == 0 || *total == 0 || (*t != 0 && *t != 1) {
		log.Fatalln("非法数据")
	}
	switch *t {
	case Print:
		printWithInfo(*interestRate, *period, *total)
	case Csv:
		generateCSV(*interestRate, *period, *total)
	}
}

func printWithInfo(interestRate float64, period int, total float64) {
	//等额本息
	fmt.Println("==========================等额本息============================")
	for i := 1; i <= period; i++ {
		yihuan, tiqian, lixi, yuehuankuan := EqualInstallmentsOfPrincipalAndInterest(interestRate, period, total, i)
		fmt.Printf("第%4d月  已还总额：%10.0f元  已还利息：%10.0f元  提前还款需付：%10.0f元  月还款：%10.0f元\n", i, yihuan, lixi, tiqian, yuehuankuan)
	}
	fmt.Println("=============================================================")
	fmt.Println()
	//等额本金
	fmt.Println("==========================等额本金============================")
	for i := 1; i <= period; i++ {
		yihuan, tiqian, lixi, yuehuankuan := EqualPrincipal(interestRate, period, total, i)
		fmt.Printf("第%4d月  已还总额：%10.0f元  已还利息：%10.0f元  提前还款需付：%10.0f元  月还款：%10.0f元\n", i, yihuan, lixi, tiqian, yuehuankuan)
	}
	fmt.Println("=============================================================")
}

func generateCSV(interestRate float64, period int, total float64) {
	//等额本息
	func() {
		buf := make([]byte, 0)
		buffer := bytes.NewBuffer(buf)
		line := fmt.Sprintf("月份,已还总额,已还利息,提前还款还需付,月还款\n")
		reader := transform.NewReader(bytes.NewReader([]byte(line)), simplifiedchinese.GBK.NewEncoder())
		d, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatalln(err)
		}
		buffer.Write(d)
		for i := 1; i <= period; i++ {
			yihuan, tiqian, lixi, yuehuankuan := EqualInstallmentsOfPrincipalAndInterest(interestRate, period, total, i)
			line := fmt.Sprintf("%d,%.0f,%.0f,%.0f,%.0f\n", i, yihuan, lixi, tiqian, yuehuankuan)
			buffer.Write([]byte(line))
		}
		err = ioutil.WriteFile("等额本息.csv", buffer.Bytes(), fs.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}()
	//等额本金
	func() {
		buf := make([]byte, 0)
		buffer := bytes.NewBuffer(buf)
		line := fmt.Sprintf("月份,已还总额,已还利息,提前还款还需付,月还款\n")
		reader := transform.NewReader(bytes.NewReader([]byte(line)), simplifiedchinese.GBK.NewEncoder())
		d, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatalln(err)
		}
		buffer.Write(d)
		for i := 1; i <= period; i++ {
			yihuan, tiqian, lixi, yuehuankuan := EqualPrincipal(interestRate, period, total, i)
			line := fmt.Sprintf("%d,%.0f,%.0f,%.0f,%.0f\n", i, yihuan, lixi, tiqian, yuehuankuan)
			buffer.Write([]byte(line))
		}
		err = ioutil.WriteFile("等额本金.csv", buffer.Bytes(), fs.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}()
}
