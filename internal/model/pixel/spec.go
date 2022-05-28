package pixel

type Spec struct {
	LengthInch float64 `json:"length_inch"` // 长度(英尺)
	LengthCm   float64 `json:"length_cm"`   // 长度(厘米)
	Rate       float64 `json:"rate"`        // 分辨率
}

func (s *Spec) SetLength(ppi float64) {
	s.LengthInch = s.Rate / ppi
	s.LengthCm = s.LengthInch * 2.54
}
