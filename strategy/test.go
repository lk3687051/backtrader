package strategy

type Strategy struct {
	bprice float
	sprice float
}

func (s *Strategy) init () {
	// 需要计算的指标，MA等等参数的设置
}

func (s *Strategy) Run (d Data) {
	// 如果没有初始价位，那么现在买入
	if bprice == 0 {
		s.bprice = d["now"] * 0.95
		s.sprice = d["now"] * 1.05
	}

	// 买入策略
	// 卖出策略
}