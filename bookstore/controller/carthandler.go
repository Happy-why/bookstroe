package controller

import (
	"fmt"
	"github.com/dao"
	"github.com/model"
	"github.com/utils"
	"net/http"
	"strconv"
	"text/template"
)

//AddBook2Cart 添加图书到购物车中
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := dao.CheckLogin(r)
	if flag {
		//已经登录
		//获取要添加的图书的id
		bookID := r.FormValue("bookId")
		//先通过bookID获取图书的信息
		book := dao.GetBookById(bookID)
		//获取用户id
		userID := session.UserID
		//先判断该用户有没有购物车
		cart := dao.GetCartByUserID(userID)
		if cart != nil {
			//说明该用户有购物车，此时需要判断购物车有没有这本书
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if cartItem != nil {
				//购物车的购物项中已经有该图书，只需要将该图书所对应的购物项中的数量加1即可
				//1.获取购物车切片中的所有的购物项
				cts := cart.CartItems
				//2.遍历得到每一个购物项
				for _, v := range cts {
					fmt.Println("当前购物项中是否有Book：", v)
					fmt.Println("查询到的Book是：", cartItem.Book)
					//3.找到当前的购物项
					if v.Book.ID == cartItem.Book.ID {
						//将购物项中的图书的数量加1
						v.Count = v.Count + 1
						//更新数据库中该购物项的图书的数量
						err := dao.UpdateBookCount(v)
						if err != nil {
							fmt.Println("UpdateBookCount err:", err)
							return
						}
					}
				}
			} else {
				//购物车的购物项中还没有该图书，此时需要创建一个购物项并添加到数据库中
				//创建购物车中的购物项
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				//将购物项添加到当前cart的切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将新创建的购物项添加到数据库中
				err := dao.AddCartItem(cartItem)
				if err != nil {
					fmt.Println("AddCartItem err:", err)
					return
				}
			}
			//不管之前购物车中是否有当前图书对应的购物项，都需要更新购物车中的图书的总数量和总金额
			err := dao.UpdateCart(cart)
			if err != nil {
				fmt.Println("UpdateCart err:", err)
				return
			}
		} else {
			//证明当前用户还没有购物车，需要创建一个购物车并添加到数据库中
			//1.创建购物车
			//生成购物车的id
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			//2.创建购物车中的购物项
			//声明一个CartItem类型的切片
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}
			//将购物项添加到切片中
			cartItems = append(cartItems, cartItem)
			//3将切片设置到cart中
			cart.CartItems = cartItems
			//4.将购物车cart保存到数据库中
			err := dao.AddCart(cart)
			if err != nil {
				fmt.Println("AddCart err:", err)
				return
			}
		}
		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车！"))
	} else {
		//没有登录
		w.Write([]byte("请先登录！"))
	}
}

// GetCartInfo  根据用户的id获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.CheckLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取userid对应的购物车
	cart := dao.GetCartByUserID(userID)
	if cart != nil {
		//说明该用户有购物车
		session.Cart = cart
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	} else {
		//该用户没有购物车
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	}
}

//DeleteCart 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物车的id
	cartID := r.FormValue("cartID")
	err := dao.DeleteCartByCartID(cartID)
	if err != nil {
		fmt.Println("DeleteCartByCartID err:", err)
		return
	}
	//调用GetCartInfo函数再次查询购物车信息
	GetCartInfo(w, r)
}

//DeleteCartItem 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物项的id
	cartItemID := r.FormValue("CartItem")
	//将购物项的id转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取session
	_, session := dao.CheckLogin(r)
	//获取用户的id
	UserID := session.UserID
	//通过UserName来获取购物车
	cart := dao.GetCartByUserID(UserID)
	cartItems := cart.CartItems
	for k, v := range cartItems {
		if v.CartItemID == iCartItemID {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			cart.CartItems = cartItems
			err := dao.DeleteCartItemByID(cartItemID)
			if err != nil {
				fmt.Println("DeleteCartItemByID err:", err)
				return
			}
		}
	}
	err := dao.UpdateCart(cart)
	if err != nil {
		fmt.Println("UpdateCart err:", err)
		return
	}
	GetCartInfo(w, r)
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemID")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	bookCount := r.FormValue("UserCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := dao.CheckLogin(r)
	userID := session.UserID
	cart := dao.GetCartByUserID(userID)
	for _, v := range cart.CartItems {
		if v.CartItemID == iCartItemID {
			v.Count = iBookCount
			err := dao.UpdateBookCount(v)
			if err != nil {
				fmt.Println("UpdateBookCount err:", err)
				return
			}
		}
	}
	err := dao.UpdateCart(cart)
	if err != nil {
		fmt.Println("UpdateCart err:", err)
		return
	}
	GetCartInfo(w, r)
}
