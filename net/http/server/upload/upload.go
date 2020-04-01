package upload

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var dashboardTpl = `<html>
<body>
    <div class="main">
        <form action="/text" method="POST">
            <input type="text" name="text" value="" placeholder="输入要发送的文本">
            <input type="submit">
        </form>
		</br>
		<form enctype="multipart/form-data" action="/file" method="POST">
			<input type="file" name="file"/>
            <input type="submit">
        </form>
    </div>
</body>
</html>`

// DefaultHttpServerHandles handler初始化
func DefaultHttpServerHandles() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("dashboard").Parse(dashboardTpl))
		tmpl.Execute(w, "")
	})

	http.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("接收到的文字:", r.FormValue("text"))
	})

	http.Handle("/share", http.StripPrefix("/share", http.FileServer(http.Dir("./share"))))

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		f, err := os.OpenFile("./files/"+head.Filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		fmt.Println("接收到的文件:", head.Filename)
	})

	http.HandleFunc("/loop", func(w http.ResponseWriter, r *http.Request) {
		timeSend := r.FormValue("time")
		fmt.Println("client send time:", timeSend)
		now := time.Now().Unix()
		fmt.Println(now, "10 seconds begin", timeSend)
		//for i := 0; i < 10; i++ {
		//time.Sleep(1 * time.Second)
		//}
		//fmt.Println(now, "1 seconds end", timeSend)
		w.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
	})

	http.HandleFunc("/long", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "send", string(body))
		w.Write([]byte("22222"))
	})
}
