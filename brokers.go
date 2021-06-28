package backtrader

type Trade struct {
  cash float
  positions map[string]int
}

type Brokers struct {
    cash float
    positions map[string]int
}

func (b *Brokers)() {
  
}
func (b *Brokers) buy (code string, number int, price float) bool {
  if b.cash >= number * price {
    if _, ok := b.positions[code]; ok {
      b.positions[code] = b.positions[code] + number
    } else {
      b.positions[code] = number
    }
    b.cash = b.cash - number * price
    return true
  } else {
    return false
  }
}

func (b *Brokers) sell (code string, number int, price float) bool {
  if _number, ok := b.positions[code]; ok {
    if _number >= number {
      b.positions[code] = b.positions[code] - number
      b.cash = b.cash + number * price
    } else {
      return false
    }
  } else {
    return false
  }
}