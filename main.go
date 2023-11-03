package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/JJerBum/nomadcoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

// 모든 템플릿들을 파싱한 정보를 담고 있는 전역변수
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

// home 함수는 "/" 패턴으로 요청이 들어왔을 떄 처리하는 핸들러 입니다.
func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}

	// ./templates/pages/home.gohtml에 구문 분석이 완료된 곳에 값 삽입
	templates.ExecuteTemplate(rw, "home", data)
}

func main() {
	// ParseGlob은 새 템플릿을 생성하고 패턴으로 식별된 파일에서 템플릿 정의를 구문 분석합니다.
	// 전역변수 templates에 ./templates/pages/*.gohtml 구문 분석
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))

	// 전역변수 templates에 ./teamplates/partials/*.gohmlt 구문 분석
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	// 핸들러 등록
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
