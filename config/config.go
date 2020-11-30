package config

//Global 保存唯一全局配置文件
var Global config

//Config 保存配置
type config struct {
	//CaptureBiFlow 是否抓取的双向流，是则双向数据包为一条流，否则为2条
	CaptureBiFlow bool

	//LimitPacketsPerFlow 解析流的时候，限制每条流解析包的数量，即只解析流的前k个包
	LimitPacketsPerFlow int
}

//Init init the default config
func (c *config) Init() {
	c.CaptureBiFlow = true
	c.LimitPacketsPerFlow = 64
}
