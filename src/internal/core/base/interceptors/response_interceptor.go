package interceptors

func ResponseSuccess(data interface{}) error {
	return nil
}

func ResponseSuccessWithPagination(data interface{}, page uint, limit uint, total uint) error {
	return nil
}
