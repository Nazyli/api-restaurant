package service

import (
	"context"

	_category "github.com/nazyli/api-restaurant/domain/category"
	_customer "github.com/nazyli/api-restaurant/domain/customer"
	_employee "github.com/nazyli/api-restaurant/domain/employee"
	_menu "github.com/nazyli/api-restaurant/domain/menu"
	_order "github.com/nazyli/api-restaurant/domain/order"
	_orderDetail "github.com/nazyli/api-restaurant/domain/order_detail"
	_position "github.com/nazyli/api-restaurant/domain/position"
	_user "github.com/nazyli/api-restaurant/domain/user"
	"github.com/nazyli/api-restaurant/entity"
	"github.com/nazyli/api-restaurant/util/auth"
)

//DataResponse json

type Status struct {
	Code   int
	ErrMsg string
}
type svc struct {
	AppID       int64
	Tax         float64
	user        _user.Repository
	position    _position.Repository
	category    _category.Repository
	menu        _menu.Repository
	customer    _customer.Repository
	employee    _employee.Repository
	order       _order.Repository
	orderDetail _orderDetail.Repository
}

// New init service
func New(AppID int64, Tax float64, user _user.Repository, position _position.Repository, category _category.Repository, menu _menu.Repository, customer _customer.Repository, employee _employee.Repository, order _order.Repository, orderDetail _orderDetail.Repository) Service {
	return &svc{
		AppID:       AppID,
		user:        user,
		position:    position,
		category:    category,
		menu:        menu,
		customer:    customer,
		employee:    employee,
		Tax:         Tax,
		order:       order,
		orderDetail: orderDetail,
	}
}

type Service interface {
	SignIn(ctx context.Context, email, password string) (token *auth.Token, status Status)
	// User
	GetUserByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (user *entity.User, status Status)
	SelectUsers(ctx context.Context, all bool, isAdmin bool, uid string) (users entity.Users, status Status)
	InsertUser(ctx context.Context, uid string, user *entity.User) (userData *entity.User, status Status)
	UpdateUser(ctx context.Context, id int64, isAdmin bool, uid string, user *entity.User) (userData *entity.User, status Status)
	DeleteUser(ctx context.Context, id int64, isAdmin bool, uid string) (status Status)

	//Position
	SelectPosition(ctx context.Context, all bool) (positions entity.Positions, status Status)
	GetPositionByID(ctx context.Context, id int64, all bool) (position *entity.Position, status Status)
	InsertPosition(ctx context.Context, position *entity.Position) (positionData *entity.Position, status Status)
	UpdatePosition(ctx context.Context, position *entity.Position) (positionData *entity.Position, status Status)
	DeletePosition(ctx context.Context, position *entity.Position) (status Status)

	//Category
	SelectCategory(ctx context.Context, all bool) (categorys entity.Categorys, status Status)
	GetCategoryByID(ctx context.Context, id int64, all bool) (category *entity.Category, status Status)
	InsertCategory(ctx context.Context, category *entity.Category) (categoryData *entity.Category, status Status)
	UpdateCategory(ctx context.Context, category *entity.Category) (categoryData *entity.Category, status Status)
	DeleteCategory(ctx context.Context, category *entity.Category) (status Status)

	//Menu
	GetMenuByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (menu *entity.Menu, status Status)
	SelectMenues(ctx context.Context, all bool, isAdmin bool, uid string) (menues entity.Menues, status Status)
	InsertMenu(ctx context.Context, uid string, menu *entity.Menu) (menuData *entity.Menu, status Status)
	UpdateMenu(ctx context.Context, isAdmin bool, uid string, menu *entity.Menu) (menuData *entity.Menu, status Status)
	DeleteMenu(ctx context.Context, id int64, isAdmin bool, uid string) (status Status)

	// Customer
	GetCustomerByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (customer *entity.Customer, status Status)
	SelectCustomers(ctx context.Context, all bool, isAdmin bool, uid string) (customers entity.Customers, status Status)
	InsertCustomer(ctx context.Context, uid string, customer *entity.Customer) (customerData *entity.Customer, status Status)
	UpdateCustomer(ctx context.Context, isAdmin bool, uid string, customer *entity.Customer) (customerData *entity.Customer, status Status)
	DeleteCustomer(ctx context.Context, id int64, isAdmin bool, uid string) (status Status)

	// Employee
	GetEmployeeByID(ctx context.Context, id int64, all bool, isAdmin bool, uid string) (employee *entity.Employee, status Status)
	SelectEmployees(ctx context.Context, all bool, isAdmin bool, uid string) (employees entity.Employees, status Status)
	InsertEmployee(ctx context.Context, uid string, employee *entity.Employee) (employeeData *entity.Employee, status Status)
	UpdateEmployee(ctx context.Context, isAdmin bool, uid string, employee *entity.Employee) (employeeData *entity.Employee, status Status)
	DeleteEmployee(ctx context.Context, id int64, isAdmin bool, uid string) (status Status)

	//Order
	SelectOrder(ctx context.Context, all bool, isAdmin bool, uid string) (orders entity.Orders, status Status)
	GetOrderByInv(ctx context.Context, inv string, all bool, isAdmin bool, uid string) (order *entity.Order, status Status)
	InsertOrder(ctx context.Context, uid string) (orderData *entity.Order, status Status)
	PaymentOrder(ctx context.Context, inv string, isAdmin bool, uid string, order *entity.Order) (a *entity.Order, b entity.OrderDetails, status Status)
	DeleteOrder(ctx context.Context, inv string, isAdmin bool, uid string) (status Status)

	//OrderDetail
	SelectOrderDetail(ctx context.Context, all bool, isAdmin bool, uid string) (orderDetails entity.OrderDetails, status Status)
	SelectOrderDetailByInv(ctx context.Context, inv string, all bool, isAdmin bool, uid string) (orderDetails entity.OrderDetails, status Status)
	GetOrderDetailByID(ctx context.Context, id int64, all, isAdmin bool, uid string) (orderDetail *entity.OrderDetail, status Status)
	InsertOrderDetail(ctx context.Context, uid string, orderDetail *entity.OrderDetail) (orderData *entity.OrderDetail, status Status)
	CalculateOrder(ctx context.Context, inv string, uid string) (orderData *entity.CalculateOrder, status Status)
	UpdateOrderDetail(ctx context.Context, isAdmin bool, uid string, orderDetail *entity.OrderDetail) (orderDetailData *entity.OrderDetail, status Status)
	DeleteOrderDetail(ctx context.Context, id int64, isAdmin bool, uid string) (status Status)
}
