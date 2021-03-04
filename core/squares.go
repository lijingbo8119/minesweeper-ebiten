package core

type Squares []*Square

func (this Squares) Find(closure func(s *Square) bool) *Square {
	for _, s := range this {
		if closure(s) {
			return s
		}
	}
	return nil
}

func (this Squares) Filter(closure func(s *Square) bool) Squares {
	res := Squares{}
	for _, s := range this {
		if closure(s) {
			res = append(res, s)
		}
	}
	return res
}

func (this Squares) Each(closure func(s *Square)) {
	for _, s := range this {
		closure(s)
	}
}
