package abstract_factory

import "fmt"

type OrderMainDao interface {
	SaveOrderMain()
}

type OrderDetailDao interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

type RDBMainDao struct{}

func (*RDBMainDao) SaveOrderMain() {
	fmt.Println("rdb main save")
}

type RDBDetailDao struct{}

func (*RDBDetailDao) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

type RDBDaoFactory struct{}

func (*RDBDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &RDBMainDao{}
}

func (*RDBDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBDetailDao{}
}

type XMLMainDao struct{}

func (*XMLMainDao) SaveOrderMain() {
	fmt.Println("xml main save")
}

type XMLDetailDao struct{}

func (*XMLDetailDao) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

type XMLDaoFactory struct{}

func (*XMLDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &XMLMainDao{}
}

func (*XMLDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLDetailDao{}
}
