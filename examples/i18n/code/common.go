package code

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
)

const (
	OperationFailed       = 1005 // 操作失败
	ResourceNotFound      = 1006 // 记录不存在
	IdError               = 1007 // id格式错误
	LoginInfoGetFail      = 1008 // 登陆信息获取失败
	NoPermissionOperation = 1009 // 无权限操作
	ServerDb              = 1010 // 服务器数据库错误
)

type (
	singleErrGenerate struct {
		Zh string
	}
	singleErrAction struct {
		ZhErr []byte
	}
)

func (s *singleErrGenerate) Generate(code int) (gin.II18nErrAction, error) {
	action := new(singleErrAction)
	var err error
	action.ZhErr, err = json.Marshal(&gin.Err{Code: code, Msg: s.Zh})
	return action, err
}

func (s *singleErrAction) ReplyErr(language string) []byte {
	return s.ZhErr
}

// 全局的错误码
var errCodes = map[int]singleErrGenerate{
	OperationFailed:       {"操作失败"},
	ResourceNotFound:      {"记录不存在"},
	IdError:               {"id格式错误"},
	LoginInfoGetFail:      {"登录信息获取失败"},
	NoPermissionOperation: {"无权限操作"},
	ServerDb:              {"服务器数据库错误"},
}

// NewErrAction 获取一个II18nErrAction对象,并在获取前先组装并校验所有的错误码信息
func NewErrAction() map[int]gin.II18nErrAction {
	// 所有人的code码
	var code4People = []map[int]singleErrGenerate{
		errCodes,
	}
	errActions := make(map[int]gin.II18nErrAction)
	allCode := make(map[int]int8, 0) //收集码,用来判断重复,用key,值没有使用用
	var err error
	// echo 预置的
	for code, bilingualErrGenerate := range gin.BaseErrorCodes {
		allCode[code] = 0
		errActions[code], err = bilingualErrGenerate.Generate(code)
		if err != nil {
			panic("错误码生成错误,请检查")
		}
	}
	// 本项目中加入的
	for _, singleErrGenerateMap := range code4People {
		for code, singleErrGenerate := range singleErrGenerateMap {
			_, ok := allCode[code]
			if ok {
				panic("code码重复:" + fmt.Sprintf("%d", code))
			}
			allCode[code] = 0
			// 将对应的码和错误详情放入 errs 中
			errActions[code], err = singleErrGenerate.Generate(code)
			if err != nil {
				panic("错误码生成错误,请检查")
			}
		}
	}
	errCodes = nil
	code4People = nil
	return errActions
}
