package strategy

type Cerebro struct {
	startTime //开始时间
	endTime   //结束时间
	data Data //数据来源
	tradeHistory //交易历史记录
}

// 运行策略
func (c *Cerebro)Run()  {
	// 初始化部分参数
	// 按照数据开始运行
	// 实时统计当前资金
	// 输出交易历史记录
	// 输出回测结果
}
