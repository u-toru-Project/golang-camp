package chapter09

import "fmt"

type Tick struct {
	Symbol string
	Price  float64
	Volume int
}

type Ohlc struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int
}

type StockAggregator struct {
	bars map[string]*Ohlc
}

func NewStockAggregator() *StockAggregator {
	return &StockAggregator{bars: map[string]*Ohlc{}}
}

func (s *StockAggregator) OnTick(tick Tick) {
	if tick.Price <= 0 || tick.Volume < 0 {
		panic("invalid tick")
	}
	bar, ok := s.bars[tick.Symbol]
	if !ok {
		s.bars[tick.Symbol] = &Ohlc{Open: tick.Price, High: tick.Price, Low: tick.Price, Close: tick.Price, Volume: tick.Volume}
		return
	}
	if tick.Price > bar.High {
		bar.High = tick.Price
	}
	if tick.Price < bar.Low {
		bar.Low = tick.Price
	}
	bar.Close = tick.Price
	bar.Volume += tick.Volume
}

func (s *StockAggregator) GetOhlc(symbol string) *Ohlc {
	return s.bars[symbol]
}

func RunQ901() {
	aggregator := NewStockAggregator()
	aggregator.OnTick(Tick{"AAPL", 100.0, 10})
	aggregator.OnTick(Tick{"AAPL", 105.0, 5})
	aggregator.OnTick(Tick{"AAPL", 98.0, 2})
	bar := aggregator.GetOhlc("AAPL")
	fmt.Printf("AAPL O=%v H=%v L=%v C=%v V=%d\n", bar.Open, bar.High, bar.Low, bar.Close, bar.Volume)
}
