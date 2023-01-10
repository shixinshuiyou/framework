package trace

var EmptyConfig Config

type Config struct {
	Host        string
	Port        int
	ServiceName string
	Ratio       float64
}

func NewJaegerTracer(config Config) {
	
}
