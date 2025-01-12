package leetcode

type ATM struct {
	count []int64
	value []int64
}

func ATMConstructor() ATM {
	return ATM{
		count: []int64{0, 0, 0, 0, 0},
		value: []int64{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < 5; i++ {
		this.count[i] += int64(banknotesCount[i])
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		res[i] = int(min(this.count[i], int64(amount)/this.value[i]))
		amount -= res[i] * int(this.value[i])
	}

	if amount > 0 {
		return []int{-1}
	}

	for i := 0; i < 5; i++ {
		this.count[i] -= int64(res[i])
	}
	return res
}
