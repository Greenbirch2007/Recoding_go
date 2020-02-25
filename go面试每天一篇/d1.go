package main



import (
	"chromedp/crawler"
	"chromedp/filtrate"
	"context"
	"github"
	"log"
)


func main(){
	ctx,cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// 执行任务

	url:= ""
	err := chromedp.Run(ctx,crawler.VisitWeb(url))
	if err !=nil{
		log.Fatal(err)
	}
	var res string
	for i:= 1;i < 27170;i++{
		//执行
		err = chromedp.Run(ctx,crawler.DoCrawler(&res)) //执行爬虫
		if err != nil{
			log.Fatal(err)
		}

		filtrate.GetAccount(res)

	}
}

//任务  主要用来设置cookie,获取登录账号后的页面
func VisitWeb(url  string,cookies ...string) chromedp.Tasks{
	//创建一个chrome任务

	return chromedp.Tasks{
		//ActionFunc是一个适配器,允许还是用普通函数作为操作
		chromedp.ActionFunc(func(ctx context.Context)error{
			//设置Cookie存活四件
			expr := cdp.TimeSinceEpoch(time.Now().Add(180*24*time.Hour))
			//添加Cookie 到Chrome

			for i:=0;i <len(cookies);i +=2{
				//SetCookie 使用给定的cookie数据设置一个cookie;如果存在
				//可能会覆盖等效的cookie
				success,err := network.SetCookie(cookies[i],cookies[i+1]).
					// 设置cookie到期时间
					WithExpires(&expr).
					//设置cookie作用的站点
					WithDomain("dl.").// 访问网站主体
					//设置httponly,防止xss攻击
					WithHTTPOnly(true).
					//Do根据提供的上下文执行Network.setCookie
					Do(ctx)

				if err !=nil {
					return err
				}
				if !success{
					return fmt.Errorf("could not set cookie %q to %q",cookies[i],cookies[i+1])

				}
				}
			}
			return nil
		}),

		//跳转指定的url地址

		chromedp.Navigate(url),
	}
}


//任务 主要执行翻页功能或html



func DoCrawler(res *string) chromedp.Tasks{
	return chromedp.Tasks{
		chromedp.Sleep(1000),
		chromedp.WaitVisible(`#form1`,chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),

		chromedp.Click()
		chromedp.OuterHTML(),
	}
}

func ReplaceStr(text string) string  {
	return strings.Replace(text,"tbody","table",-1)
}

func IsMobile(text string) bool{
	match,_:= regexp.MatchString()
	return match
}


func WirteText(savefile string,txt string){
	f,err := os.OpenFile(savefile,os.O_RDWR|os.O_CREATE|os.O_APPEND,0777)
	if err !=nil{
		fmt.Println("os Create error:",err)
		return
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	bw.WriteString(txt+"\n")
	bw.Flush()
}

func GetAccount(text string){
	dom,err := goquery.NewDocumentFromReader(strings.NewReader(ReplaceStr(text)))
	if err != nil{
		log.Fatalln(err)
	}
	dom.Find("tr").Each(func(i int,selection *goquery.Selection){
		s:=selection.Find("td").Eq(6).Text()
		fmt.Println(s)
		WriteText("Account.txt",s)
		if IsMobile(s){
			WiriteText("Mobile.txt",s)
		}
	})
}