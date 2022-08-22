package controller

import (
	"fmt"
	"github.com/dao"
	"github.com/model"
	"net/http"
	"strconv"
	"text/template"
)

// GetPageBooks 获取带分页的图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用带分页的图书的函数
	page, _ := dao.GetPageBook(pageNo)
	//解析模板文件，将其所有图书的信息响应到页面上去
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, page)
}

//GetPageBooksByPrice 获取分页和价格范围的图书信息
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取用户输入
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBook(pageNo)
	} else {
		//调用带分页的图书的函数
		page, _ = dao.GetPageBookByPrice(pageNo, minPrice, maxPrice)
		//将价格设置到page中去
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	flag, session := dao.CheckLogin(r)
	if flag == true {
		//已经登录了
		page.IsLogin = true
		page.UserName = session.UserName
	}

	//解析模板文件，将其所有图书的信息响应到页面上去
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}

//ToUpdateBookPage 跳转到修改页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//获取要更新的图书的id
	ID := r.FormValue("bookId")
	//调用bookdao中获取图书的函数
	book := dao.GetBookById(ID)
	if book.ID > 0 {
		//在更新图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, book)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	//拿到提交的数据
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	iPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 64)
	iStock, _ := strconv.ParseInt(stock, 10, 64)
	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   iPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	//调用添加到数据库中的方法
	err := dao.AddBook(book)
	if err != nil {
		fmt.Println("AddBook err:", err)
		return
	}
	GetPageBooks(w, r)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("bookID")
	err := dao.DeleteBook(id)
	if err != nil {
		fmt.Println("DeleteBook err:", err)
		return
	}
	GetPageBooks(w, r)
}

// IndexHandle  去首页
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用带分页的图书的函数
	page, _ := dao.GetPageBook(pageNo)
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

//GetBooks 获取所有图书
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	//获取所有的图书
//	books, _ := dao.GetBooks()
//	//解析模板文件，将其所有图书的信息响应到页面上去
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	//执行
//	t.Execute(w, books)
//}

// UpdateOrAddBook  更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	bookid := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	ibookid, _ := strconv.ParseInt(bookid, 10, 64)
	iPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 64)
	iStock, _ := strconv.ParseInt(stock, 10, 64)
	book := &model.Book{
		ID:      int(ibookid),
		Title:   title,
		Author:  author,
		Price:   iPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		err := dao.UpDateBook(book)
		if err != nil {
			fmt.Println("UpDateBook err:", err)
			return
		}
	} else {
		err := dao.AddBook(book)
		if err != nil {
			fmt.Println("AddBook err:", err)
			return
		}
	}
	GetPageBooks(w, r)
}
