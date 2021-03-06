package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

	//find the relevant url
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/zhumadian/[^"]+)"`)
)

//获取城市中的所有人员
func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, url, name)
			},
		})
	}

	//find the  relevant url
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				//m is []byte
				Url:       string(m[1]),
				ParseFunc: ParseCity,
			})
	}
	return result
}
