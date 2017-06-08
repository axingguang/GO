package dataSyn

const (
	SetType = iota
	DelType
)

type Data struct {
	Key      string
	Val      string
	Ex       int
	DataType int
}
