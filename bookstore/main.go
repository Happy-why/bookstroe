package main

import (
	"fmt"
	"github.com/controller"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("E:\\GoLand2022.1\\goproject\\src\\bookstore0612\\views\\index.html"))
	//执行
	err := t.Execute(w, "")
	if err != nil {
		fmt.Println("t.Execute err:", err)
		return
	}
}

func main() {
	//设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("E:\\GoLand2022.1\\goproject\\src\\书城学习\\bookstore\\views\\static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("E:\\GoLand2022.1\\goproject\\src\\书城学习\\bookstore\\views\\pages/"))))
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//登录
	http.HandleFunc("/Login", controller.Login)
	//注销
	http.HandleFunc("/Logout", controller.Logout)

	//去注册
	http.HandleFunc("/Register", controller.Register)

	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/CheckUserName", controller.CheckUserName)

	//获取所有图书
	// http.HandleFunc("/GetBooks", controller.GetBooks)

	//获取带分页的图书信息
	http.HandleFunc("/GetPageBooks", controller.GetPageBooks)
	http.HandleFunc("/GetPageBooksByPrice", controller.GetPageBooksByPrice)

	//添加图书
	// http.HandleFunc("/AddBook", controller.AddBook)

	//删除图书
	http.HandleFunc("/DeleteBook", controller.DeleteBook)

	//更新图书的页面
	http.HandleFunc("/ToUpdateBookPage", controller.ToUpdateBookPage)

	//更新或添加图书
	http.HandleFunc("/UpdateOrAddBook", controller.UpdateOrAddBook)

	//添加图书到购物车中
	http.HandleFunc("/AddBook2Cart", controller.AddBook2Cart)

	//获取购物车信息
	http.HandleFunc("/GetCartInfo", controller.GetCartInfo)

	//清空购物车
	http.HandleFunc("/DeleteCart", controller.DeleteCart)

	//删除购物项
	http.HandleFunc("/DeleteCartItem", controller.DeleteCartItem)

	//更新购物项
	http.HandleFunc("/UpdateCartItem", controller.UpdateCartItem)

	//去结账
	http.HandleFunc("/Checkout", controller.Checkout)

	//获取所有订单
	http.HandleFunc("/GetOrders", controller.GetOrders)

	//获取订单详情，即订单所对应的所有的订单项
	http.HandleFunc("/GetOrderInfo", controller.GetOrderInfo)

	//获取我的订单
	http.HandleFunc("/GetMyOrder", controller.GetMyOrders)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe err:", err)
		return
	}
}
