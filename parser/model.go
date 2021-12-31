package parser

// TestModel is benchmark data
type TestModel struct {
	Test1 int64    `json:"test1"`
	Test2 string   `json:"test2"`
	Test3 bool     `json:"test3"`
	Test4 int32    `json:"test4"`
	Test5 []string `json:"test5"`
	Test6 []int64  `json:"test6"`
	Test7 *string  `json:"test7"`
}
