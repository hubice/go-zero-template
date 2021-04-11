package errorz

var (
	// 常用
	ParameterError = NewCodeError(1001, "参数错误")
	DataEmpty = NewCodeError(1002, "数据为空")
	DataAddFailed = NewCodeError(1003, "数据添加失败")
	DataUpdateFailed = NewCodeError(1004, "数据修改失败")

	// 用户
	UserEmpty = NewCodeError(2001, "用户不存在")
)
