package gin

import "github.com/gin-gonic/gin/json"

type (
	// 页码数据
	Meta struct {
		Total       int `json:"total"`       // 数据总数
		CurrentPage int `json:"currentPage"` // 当前第几页
		PageSize    int `json:"pageSize"`    // 每页数据量
	}
	// 集合分页返回结构体
	Pagination struct {
		Data interface{} `json:"data"` // 当页数据
		Meta Meta        `json:"meta"` // 分页信息
	}

	// [业务错误生产者]启动时生成所有业务错误返回者(用途,确定业务错误返回者支持几种语言)
	II18nErrGenerate interface {
		Generate(code int) (II18nErrAction, error)
	}
	// [业务错误操作者]根据语言环境,返回具体业务错误
	II18nErrAction interface {
		ReplyErr(language string) []byte
	}
	// 业务错误结构体,在初始化的时候,就将所有的错误结构体序列化为json,运行过程中无需json转换,提升性能
	Err struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	// 分页数据
	IPageParams interface {
		GetCurrentPage() int // 页码数
		GetPageSize() int    // 每页数据量
		GetLimit() []int     // 分页
		GetOrderBy() string  // 排序字段
		GetAsc() bool        // 排序顺序 默认为false,也就是默认倒序
	}
	PageParams struct {
		CurrentPage int    `form:"currentPage"` //当前页
		PageSize    int    `form:"pageSize"`    //每页数据量
		Limit       []int  `form:"limit"`       //分页
		OrderBy     string `form:"orderBy"`     //排序字段
		Asc         bool   `form:"asc"`         //排序顺序 默认为false,也就是默认倒序
	}
)

const (
	ZH                         = "zh"       // 中文
	EN                         = "en"       // 英语
	TranslateLanguageHeaderKey = "language" //语言参数名称
	// 错误码
	DefaultCode     = 1000 // 默认错误
	SystemCode      = 1001 // 系统错误
	ParamsBindCode  = 1002 // 参数绑定错误
	ParamsCheckCode = 1003 // 参数校验错误
)

// 系统内置的4种错误
var BaseErrorCodes = map[int]BilingualErrGenerate{
	DefaultCode:     {"默认错误", "Default error"},
	SystemCode:      {"系统错误", "system error"},
	ParamsBindCode:  {"参数绑定错误", "Parameter binding error"},
	ParamsCheckCode: {"参数校验错误", "Parameter check error"},
}

// 获取当前客户端语言环境
func getLang(c *Context) string {
	lang := c.Request.Header.Get(TranslateLanguageHeaderKey) // 获取当前语言
	if lang == "" {
		lang = ZH // 默认中文
	}
	return lang
}

// 双语错误生产者
type BilingualErrGenerate struct {
	ZhErr string //中文错误
	EnErr string //英文错误
}

func (b *BilingualErrGenerate) Generate(code int) (II18nErrAction, error) {
	action := new(BilingualErrAction)
	var err error
	action.ZhErr, err = json.Marshal(&Err{Code: code, Msg: b.ZhErr})
	if err != nil {
		return nil, err
	}
	action.EnErr, err = json.Marshal(&Err{Code: code, Msg: b.EnErr})
	return action, err
}

// 双语错误操作者
type BilingualErrAction struct {
	ZhErr []byte // 中文错误 json数据
	EnErr []byte // 英文错误 json数据
}

func (b *BilingualErrAction) ReplyErr(language string) []byte {
	if language == EN {
		return b.EnErr
	}
	return b.ZhErr
}

func (p *PageParams) GetCurrentPage() int {
	return p.CurrentPage
}

func (p *PageParams) GetPageSize() int {
	return p.PageSize
}

func (p *PageParams) GetLimit() []int {
	return p.Limit
}

func (p *PageParams) GetOrderBy() string {
	return p.OrderBy
}

func (p *PageParams) GetAsc() bool {
	return p.Asc
}
