package main

/*
iotaはビットシフトと組み合わせることで、権限設定などで使われるビットフラグを簡潔に表現できます。
*/

const (
	Read    = 1 << iota // 1 << 0  (1)
	Write   = 1 << iota // 1 << 1  (2)
	Execute = 1 << iota // 1 << 2  (4)
)

// Read | Write は 1 | 2 = 3 となり、複数の権限を持つことを表現できる
