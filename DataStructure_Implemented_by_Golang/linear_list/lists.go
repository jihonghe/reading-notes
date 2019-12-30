package linear_list

// 定义线性表接口
type List interface {
	// 插入元素
	Append(element interface{}) (int, bool)
	Insert(element interface{}, index int) bool
	// 删除元素
	Delete(index int) (interface{}, bool)
	DeleteLast() (interface{}, bool)
	Clear() (int, bool)
	// 修改元素
	Set(value interface{}, index int) (interface{}, bool)
	// 查找元素
	Get(index int) interface{}
	Index(element interface{}) int
	Contains(element interface{}) bool
	// 获取表信息
	IsEmpty() bool
	Length() int
	Traverse()
}
