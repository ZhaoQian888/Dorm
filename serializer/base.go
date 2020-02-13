package serializer

// BaseResponse 基础回应
type BaseResponse struct {
	Status int         `form:"status" json:"status" binding:"required"`
	Data   interface{} `form:"data"   json:"data"   binding:""`
	Msg    string
	Error  string
}

// DataList 列表data
type DataList struct {
	Items interface{} `form:"items" json:"items"`
	Total int         `form:"total" json:"total"`
}

// BuildListResponse 列表构建器
func BuildListResponse(items interface{}, total int) BaseResponse {
	return BaseResponse{
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
