package request

import (
	"errors"
	"fmt"
	"math"
	neturl "net/url"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jy01095902/gokits/elves"
)

type Request struct {
	appKey    string
	appSecret string
}

func New(appKey, appSecret string) Request {
	req := Request{
		appKey:    appKey,
		appSecret: appSecret,
	}

	return req
}

func (req Request) execute(r *resty.Request, method, url string) (Values, error) {
	token, err := req.getToken()
	if err != nil {
		return nil, err
	}

	resp, err := r.
		SetQueryParam("access_token", neturl.QueryEscape(token)).
		SetResult(Values{}).
		Execute(method, url)

	if err != nil {
		return nil, fmt.Errorf("%w error: %s", ErrCallYonSuiteAPIFailed, err.Error())
	}

	if resp.StatusCode() != 200 {
		if resp.StatusCode() == 429 {
			return nil, fmt.Errorf("%w error: %s", ErrAPILimit, resp.String())
		}

		return nil, fmt.Errorf("%w error: %s", ErrCallYonSuiteAPIFailed, resp.String())
	}

	result, ok := resp.Result().(*Values)
	if !ok {
		return Values{}, fmt.Errorf("%w: type of result is not Values", ErrYonSuiteAPIBizError)
	}

	if err := checkAPIResponse(*result); err != nil {
		return Values{}, err
	}

	return *result, err
}

func (req Request) Post(url string, body interface{}) (Values, error) {
	r := resty.New().R().
		EnableTrace().
		SetBody(body)

	return req.execute(r, resty.MethodPost, url)
}

func (req Request) Get(url string, params map[string]string) (Values, error) {
	r := resty.New().R().
		EnableTrace().
		SetQueryParams(params)

	return req.execute(r, resty.MethodGet, url)
}

func (req Request) getToken() (string, error) {
	return getToken(req.appKey, req.appSecret, tokenURL)
}

/*
	以下是对request的扩展
*/
type QueryOption func(values Values)

// 从结果中解析总数据量，因为postAll方法不知道返回值的格式
type ParseTotalFunc func(vals Values) int

// 把分页信息添加到测试中，因为postAll方法不知道参数的格式
type MakeParamsWithPageFunc func(pageNumber, pageSize int) Values

type PostAllConfig struct {
	URL                string
	MakeParamsWithPage MakeParamsWithPageFunc
	ParseTotal         ParseTotalFunc
	Interval           time.Duration
}

func (req Request) PostAll(cfg PostAllConfig) ([]Values, error) {
	// 每页返回的结果，客户端自己解析里边的内容
	result := []Values{}

	pageNumber, pageSize := 1, 500
	body := cfg.MakeParamsWithPage(pageNumber, pageSize)

	resp, err := req.Post(cfg.URL, body)
	if err != nil {
		return nil, err
	}
	// 先把第一页的结果加进去，之后从第二页开始查询
	result = append(result, resp)

	total := cfg.ParseTotal(resp)
	pageCnt := int(math.Ceil(float64(total) / float64(pageSize)))
	if pageCnt == 0 {
		return nil, errors.New("not found")
	}

	valsChan := make(chan Values)
	var wgRes sync.WaitGroup
	wgRes.Add(1)
	go func() {
		defer wgRes.Done()

		for vals := range valsChan {
			result = append(result, vals)
		}
	}()

	pool, err := elves.NewPool(10)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(pageCnt - 1)
	for pageNum := 2; pageNum <= pageCnt; pageNum++ {
		pool.Execute(func() {
			defer wg.Done()

			body := cfg.MakeParamsWithPage(pageNum, pageSize)
			if resp, err := req.Post(cfg.URL, body); err == nil {
				valsChan <- resp
			}
		})

		time.Sleep(cfg.Interval)
	}

	wg.Wait()
	close(valsChan)
	pool.Destroy()

	wgRes.Wait()
	return result, nil
}

type PostCommandConfig struct {
	req  Request
	url  string
	body Values
}

type CommandResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Count        int      `json:"count"`
		SucceedCount int      `json:"sucessCount"`
		FailedCount  int      `json:"failCount"`
		Messages     []string `json:"messages"`
		Infos        []Values `json:"infos"`
		FailedInfos  []Values `json:"failInfos"`
	} `json:"data"`
}

func (req Request) PostCommand(cfg PostCommandConfig) (CommandResponse, error) {
	resp, err := cfg.req.Post(cfg.url, cfg.body)
	if err != nil {
		return CommandResponse{}, err
	}

	res, err := resp.GetResult(CommandResponse{})
	if err != nil {
		return CommandResponse{}, err
	}

	r, ok := res.(*CommandResponse)
	if !ok {
		return CommandResponse{}, errors.New("response is not type of CommandResponse")
	}

	if r.Code != "200" {
		return *r, errors.New(r.Message)
	}

	if r.Data.FailedCount > 0 {
		msg := strings.Join(r.Data.Messages, ",")

		return *r, errors.New(msg)
	}

	return *r, nil
}
