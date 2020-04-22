package models

type PageModel struct {
	CurrentPage int32
	RecordCount int64
	PageSize    int32
	Records     interface{}
}

//StartRecord 设置分页查询时起始位置
func (p *PageModel) StartRecord() int32 {
	var start int32 = 0
	if p.CurrentPage > 1 {
		start = (p.CurrentPage - 1) * p.PageSize
	}
	return start
}

//Paginator 分页方法
func (p *PageModel) Paginator() map[string]interface{} {
	var (
		currpage   int32 = 1
		totalpages int32 = 0
	)
	paginatorMap := make(map[string]interface{}) //返回的分页数据
	records := p.Records                         //分页后的数据

	//page总页数
	totalpages = (int32(p.RecordCount) + p.PageSize - 1) / p.PageSize
	//当前页超出范围处理
	currpage = func() int32 {
		switch {
		case p.CurrentPage < 1:
			return 1
		case p.CurrentPage > totalpages:
			return totalpages
		default:
			return p.CurrentPage
		}
	}()

	paginatorMap["currpage"] = currpage
	paginatorMap["total"] = p.RecordCount
	paginatorMap["rows"] = records
	return paginatorMap
}
