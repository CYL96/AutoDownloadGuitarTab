package run

import (
	"bufio"
	"io"
	"net/http"
	"strings"
)

type SearchResultExt struct {
	Id   string
	Name string `json:"name" type:"varchar(255)" default:"''" comment:""`
}

func GetSearchResult(search string, p string) (data []SearchResultExt, err error) {
	var (
		post *http.Request
	)

	post, err = http.NewRequest("GET", "http://www.gtpso.com/Home/Index/searchresult?search="+search+"&p="+p, nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(post)
	if err != nil {
		return
	}
	defer res.Body.Close()
	buf := bufio.NewReader(res.Body)

	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)
		if strings.Contains(v, "td style") {
			//     标题
			data = append(data, SearchResultExt{
				Id:   "",
				Name: ReadTitle(v),
			})
		} else if strings.Contains(v, "viewTab?id=") {
			//     作者
			data[len(data)-1].Id = ReadId(v)
		}

	}
	return
}

func GetHotGT(page string) (data []SearchResultExt, err error) {
	// http://www.gtpso.com/index.php?m=home&c=index&a=hottabs&p=1
	var (
		post *http.Request
		res  *http.Response
	)

	post, err = http.NewRequest("GET", "http://www.gtpso.com/index.php?m=home&c=index&a=hottabs&p="+page, nil)
	if err != nil {
		return
	}

	res, err = http.DefaultClient.Do(post)
	if err != nil {
		return
	}
	defer res.Body.Close()
	buf := bufio.NewReader(res.Body)

	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)
		if strings.Contains(v, "td style") {
			//     标题
			data = append(data, SearchResultExt{
				Id:   "",
				Name: ReadTitle(v),
			})
		} else if strings.Contains(v, "viewTab?id=") {
			//     作者
			data[len(data)-1].Id = ReadId(v)
		}

	}
	return
}
func GetNewGT(page string) (data []SearchResultExt, err error) {
	// http://www.gtpso.com/index.php?m=home&c=index&a=hottabs&p=1
	var (
		post *http.Request
		res  *http.Response
	)

	post, err = http.NewRequest("GET", "http://www.gtpso.com/index.php?m=home&c=index&a=newtabs&p="+page, nil)
	if err != nil {
		return
	}

	res, err = http.DefaultClient.Do(post)
	if err != nil {
		return
	}
	defer res.Body.Close()
	buf := bufio.NewReader(res.Body)

	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		v := string(a)
		if strings.Contains(v, "td style") {
			//     标题
			data = append(data, SearchResultExt{
				Id:   "",
				Name: ReadTitle(v),
			})
		} else if strings.Contains(v, "viewTab?id=") {
			//     作者
			data[len(data)-1].Id = ReadId(v)
		}

	}
	return
}
func ReadTitle(v string) (data string) {
	if len(v) == 0 {
		return ""
	}
	s := strings.Index(v, "\">")
	e := strings.Index(v, "</td>")
	// fmt.Println(v[s+2 : e])
	return v[s+2 : e]
}
func ReadId(v string) (data string) {
	if len(v) == 0 {
		return ""
	}
	s := strings.Index(v, "viewTab?id=")
	e := strings.Index(v, "\" target=")
	// fmt.Println(v[s+11 : e])
	return v[s+11 : e]
}
