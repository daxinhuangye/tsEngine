package tsPagination

import (

)

type Pagination struct {
	CurrPage int64
	PageSize int64
	MaxCount int64
	maxPageCount int64
}

func NewPagination(curr_page int64, page_size int64, max_count int64)(res *Pagination) {
	res = &Pagination{
	}
	res.Set(curr_page, page_size, max_count)
	
	return
}

func NewPaginationById(id int64, page_size int64, max_count int64)(res *Pagination) {
	res = NewPagination(1, page_size, max_count)
	
	curr_page := id / res.PageSize
	if id % res.PageSize > 0 {
		curr_page += 1
	}
	res.Set(curr_page, page_size, max_count)
	
	return
}

func (this *Pagination)Set(curr_page int64, page_size int64, max_count int64) {
	this.CurrPage = curr_page
	this.PageSize = page_size
	this.MaxCount = max_count

	if this.PageSize <= 0 {
		this.PageSize = 15
	}
	if this.MaxCount <= 0 {
		this.MaxCount = 0
	}
	if this.CurrPage <= 0 {
		this.CurrPage = 1
	}
	this.maxPageCount = this.MaxCount / this.PageSize
	if this.MaxCount % this.PageSize > 0 {
		this.maxPageCount += 1
	}
	if this.maxPageCount <=0 {
		this.maxPageCount = 1
	}
//	if this.CurrPage>this.maxPageCount {
//		this.CurrPage = this.maxPageCount
//	}	
}

func (this *Pagination)GetOffset() (offset int64) {
	if (this.CurrPage>this.maxPageCount) {
		 return this.MaxCount
	}
	return (this.CurrPage - 1) * this.PageSize
}