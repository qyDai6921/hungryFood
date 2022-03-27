package define

type Result struct {
	Code int         `json:"code"` // The return status code of the interface
	Msg  string      `json:"msg"`  // The return information code of the interface
	Data interface{} `json:"data"` // The return data code of the interface
}

type ResultWithPage struct {
	Code int         `json:"code"`           // The return status code of the interface
	Msg  string      `json:"msg"`            // The return information code of the interface
	Data interface{} `json:"data,omitempty"` // The return data code of the interface
	Page Page        `json:"page"`           // pagination
}

// Pagination, page and sequence numbers are all based on 1
type Page struct {
	Page      int `form:"page" json:"page" `          // current page
	PageSize  int `form:"page_size" json:"page_size"` //  size
	Total     int `json:"total"`                      //  total
	PageCount int `json:"page_count"`                 //  total pages
}
