package operates

import (
	"runtime"
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/xml"
	"time"
	"math/rand"
	"net/http"
	"log"
	"regexp"
    "context"
    "github.com/PuerkitoBio/goquery"
    "github.com/Tang-RoseChild/mahonia"
	
)

var (
    maxNum = 1000
    statNum = 0
	urlChannel = make(chan string, 2000)
	atagRegExp = regexp.MustCompile(`<a[^>]+[(href)|(HREF)]\s*\t*\n*=\s*\t*\n*[(".+")|('.+')][^>]*>[^<]*</a>`)
	userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
		"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
		"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
		"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
		"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
		"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	}
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)


type Spider struct {
	OperatesBase
}

func (op *Spider) Execute(ctx context.Context) map[string]interface{} {
	select {
		case <-ctx.Done():
			return nil
		default:
			parseParams(op, ctx)
			resolve := make(map[string]interface{})
			rootUrl := op.currentParams["rootUrl"]
			go Spy2(rootUrl)
			for url := range urlChannel {
                fmt.Println("routines num = ", runtime.NumGoroutine(), "chan len = ", len(urlChannel)) //通过runtime可以获取当前运行时的一些相关参数等
                statNum++
                if (statNum > maxNum){
                    break
                } 
				go Spy2(url)
			}
			resolve["ret"] = "success"
			return resolve
	}
}

func Spy2(url string) {
    doc, err := goquery.NewDocument(url)
    if err != nil {
      return
    }
    doc.Find("h3 a").Each(func(i int, s *goquery.Selection) {
        href, ok:= s.Attr("href")
        content := string([]byte(s.Text()))
        fmt.Printf("Review %d: %s - %s\n", i, href, content)
        if strings.Contains(content, "1111"){
            fmt.Println(href, content)
        }
        if (ok) {
            urlChannel <- href
        }

    })
}

func Spy(url string) {
    defer func() {
        if r := recover(); r != nil {
            log.Println("[E]", r)
        }
    }()
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", GetRandomUserAgent())
    client := http.DefaultClient
    res, e := client.Do(req)
    if e != nil {
        fmt.Errorf("Get请求%s返回错误:%s", url, e)
        return
    }

    if res.StatusCode == 200 {
        body := res.Body
        defer body.Close()
        bodyByte, _ := ioutil.ReadAll(body)
        resStr := string(bodyByte)
        atag := atagRegExp.FindAllString(resStr, -1)
        for _, a := range atag {
            href, content := GetHref(a)
            fmt.Println(href, content)
            if strings.Contains(content, "1111"){
                fmt.Println(href, content)
            }
            urlChannel <- href
        }
    }
}

func GetHref(atag string) (href,content string) {
    // fmt.Println("gethref", atag)
    inputReader := strings.NewReader(atag)
    decoder := xml.NewDecoder(inputReader)
    for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
        switch token := t.(type) {
        // 处理元素开始（标签）
        case xml.StartElement:
            for _, attr := range token.Attr {
                attrName := attr.Name.Local
                attrValue := attr.Value
                if(strings.EqualFold(attrName,"href") || strings.EqualFold(attrName,"HREF")){
                    href = attrValue
                }
            }
        // 处理元素结束（标签）
        case xml.EndElement:
        // 处理字符数据（这里就是元素的文本）
        case xml.CharData:
            content = string([]byte(token))
        default:
            href = ""
            content = "00000000000000"
        }
    }
    return href, content
}

func GetRandomUserAgent() string {
    return userAgent[r.Intn(len(userAgent))]
}