package main

import (
    "fmt"
    "golang.org/x/net/html"
    "sync"
    "runtime"
    "os"
    "log"
)


// URLのリスト
// 結果のリスト

var (
    MaxWorker = os.Getenv("MAX_WORKERS")
    MaxQueue  = os.Getenv("MAX_QUEUE")
)

func main() {

    // 並列処理
    // var wg sync.WaitGroup
    // cpus := runtime.NumCPU()
    // runtime.GOMAXPROCS(cpus)
    // semaphore := make(chan int, cpus)
    // for i := 0; i <= 3; i++ {
    //     wg.Add(1)
    //     go func(i int) {
    //         defer wg.Done()
    //         semaphore <- 1
    //         link()
    //         <-semaphore
    //     }(i)
    // }
    // wg.Wait()

    runtime.GOMAXPROCS(runtime.NumCPU())
    // p := sync.Pool{
    //     New: func() interface{} {
    //         return "定時作業"
    //     },
    // }

    wg := new(sync.WaitGroup)

    // 並列処理1(追加していく)
    // wg.Add(1)
    // go func() {
    //     for i := 0; i < 10; i++ {
    //         p.Put(hot())
    //         // time.Sleep(100 * time.Millisecond)
    //     }
    //     wg.Done()
    // }()


    
    // fmt.Println(doc.statusCode)

    // 並列処理2
    wg.Add(1)
    go func() {
        for i := 0; i < 10; i++ {
            // fmt.Println(p.Get())
            results := []Result{}
            doc, err := NewRequest("https://hydrocul.github.io/wiki/programming_languages_diff/string/compare.htmlss")
            if err != nil {
                log.Fatal(err)
            }
            results = link(doc)
            for _, result := range results {
                fmt.Println(result.Url)
                // p.Put(hot(result.Url))
            }
        }
        wg.Done()
    }()
    wg.Wait()
}

func hot(url string) string{
    NewRequest(url)
    return "きた"
}

// ノード解析

type Result struct {
    Url string

}

func (s *Selection) Find(selector string) {
    // return pushStack(s, findWithMatcher(s.Nodes, compileMatcher(selector)))
	// return pushStack(s, findWithMatcher(s.Nodes, compileMatcher(selector)))
}

// リンクを取得
func link(doc *Document) []Result{

    // u, err := url.Parse(target)
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Println(u.Scheme)
    results := []Result{}
    var result Result
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {

            for _, a := range n.Attr {
                if a.Key == "href" {
                    result.Url = a.Val
                    results = append(results, result)
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc.rootNode)
    doc.Find("aaa")
    return results
}

func pushStack(fromSel *Selection, nodes []*html.Node) *Selection {
	result := &Selection{nodes, fromSel.document}
	return result
}


