package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/JJerBum/nomadcoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
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

// add 함수는 "/add" 패턴으로 요청이 들어왔을 때 처리하는 핸들러 입니다.
func add(rw http.ResponseWriter, r *http.Request) {
	// add 핸들러의 메서드가
	switch r.Method {
	// Get 이라면
	case http.MethodGet:
		templates.ExecuteTemplate(rw, "add", nil)

	case http.MethodPost:
		// form.Vlaues는 ParseForm이 호춮된 후에만 사용이 가능하다.
		r.ParseForm()

		// form-data 값을 가져오기
		data := r.Form.Get("blockData")

		// blockchain의 값을 추가
		blockchain.GetBlockchain().AddBlock(data)

		// response
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)

	}
}
func Start() {
	template.ParseFiles()
	// ParseGlob은 새 템플릿을 생성하고 패턴으로 식별된 파일에서 템플릿 정의를 구문 분석합니다.
	// 전역변수 templates에 ./templates/pages/*.gohtml 구문 분석
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))

	// 전역변수 templates에 ./teamplates/partials/*.gohmlt 구문 분석
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	// 핸들러 등록
	// 이 핸들러는 메서드 말고 패턴으로 식별해서 핸득러에 요청을 처리학하게 합니다.
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
