package dao

import (
	"fmt"
	"github.com/model"
	"github.com/utils"
	"strconv"
)

//GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			fmt.Println("GetBooks获取数据库中所有的图书Scan出错 err:", err)
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil

}

//AddBook 添加图书
func AddBook(b *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook 删除图书
func DeleteBook(bookID string) error {
	sqlStr := `delete from books where id=?`
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

// GetBookById  根据图书的id 查询图书
func GetBookById(bookId string) *model.Book {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id= ?"
	row := utils.Db.QueryRow(sqlStr, bookId)
	book := &model.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		fmt.Println("GetBookByID Scan出错 err:", err)
		return nil
	}
	return book
}

// UpDateBook  根据图书的id更新图书信息
func UpDateBook(b *model.Book) error {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetPageBook 获取带分页的图书信息
func GetPageBook(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取图书的总记录数
	sqlStr := "select count(*) from books"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	err := row.Scan(&totalRecord)
	if err != nil {
		fmt.Println("GetPageBook Db.QueryRow err:", err)
		return nil, err
	}
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	//获取当前页面的图书
	sqlstr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	rows, _ := utils.Db.Query(sqlstr2, pageSize*(iPageNo-1), pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			fmt.Println("GetPageBook rows.Next() Scan err:", err)
			return nil, err
		}
		//将book添加到books里面去
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Book:        books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

// GetPageBookByPrice  获取带分页和价格范围的图书信息
func GetPageBookByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取图书的总记录数
	sqlStr := "select count(*) from books where price between ? and ?"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	err := row.Scan(&totalRecord)
	if err != nil {
		fmt.Println("GetPageBooksByPrice Scan err:", err)
		return nil, err
	}
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	//获取当前页面的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	rows, _ := utils.Db.Query(sqlStr2, minPrice, maxPrice, pageSize*(iPageNo-1), pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			fmt.Println("GetPageBooksByPrice rows.Next Scan err:", err)
			return nil, err
		}
		//将book添加到books里面去
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Book:        books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil

}
