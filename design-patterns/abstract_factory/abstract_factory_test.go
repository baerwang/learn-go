package abstract_factory

import "testing"

func mainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

func exampleRdbFactory() {
	mainAndDetail(&RDBDaoFactory{})
}

func exampleXmlFactory() {
	mainAndDetail(&XMLDaoFactory{})
}

func TestAbstract(t *testing.T) {

	exampleRdbFactory()

	exampleXmlFactory()

}
