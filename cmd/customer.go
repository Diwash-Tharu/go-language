package main

import ("github.com/gin-gonic/gin"
        "net/http"
		"pizza-tracker-go/internal/models"
		"log/slog"

)

type OrderFromData struct {
	PizzaType []string
	PizzaSize []string
}

type OrderRequest struct {
	Name        string   `from:"name" binding:"required" min=2 max=100`
	Phone       string   `from:"phone" binding:"required" min=10 max=15`
	PizzaType   []string `from:"pizzaType" binding:"required,valid_pizza_type"`
	PizzaSize   []string `from:"pizzaSize" binding:"required,valid_pizza_size"`
	Address     string   `from:"address" binding:"required"min=5 max=200`
	Instruction []string `from:"instruction" binding:"max=200"`
}

func (h *Handler) ServeNewOrderForm(c *gin.Context){
	c.HTML(http.StatusOK, "order.tmpl", OrderFromData{
		PizzaType:models.PizzaType,
		PizzaSize:models.PizzaSize,
	})
}


func (h *Handler) HandleNewOrderPost(c *gin.Context){

	var from OrderRequest
	if err := c.ShouldBind(&from); err != nil {
		c.HTML(http.StatusBadRequest, "order.tmpl", gin.H{
			"error": err.Error(),
			"data":  from,
		})
		return
	}
	
	orderItems := make([]models.OrderItem, len(from.PizzaSize))
	for i := range from.PizzaType {
		orderItems[i] = models.OrderItem{
			Pizza: from.PizzaType[i],
			Size: from.PizzaSize[i],
			Instruction: from.Instruction[i],
		}
	}

	order := models.Order{
		CustomerName: from.Name,
		Phone:        from.Phone,
		Address:      from.Address,
		Items:        orderItems,
		Status:      models.OrderStatus[0], // Set initial status to "Pending"
		
	}

	if err := h.order.CreateOrder(&order); err != nil {
		slog.Error("Failed to create order: ", err)
		c.String( http.StatusInternalServerError, "something went wrong")
		return
	}
	slog.Info("Order created successfully", "orderID", order, "customerName", order.CustomerName, "phone", order.Phone, "address", order.Address, "status", order.Status)

	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)
}

func (h *Handler) serveCustomer (c *gin.Context){
	orderID := c.Param("id")

	if orderID == "" {
		c.String(http.StatusBadRequest, "order id is required, invalid order id")
		return
	}



	order,err :=h.order.GetOrder(orderID)
	if err != nil {
		c.String(http.StatusInternalServerError, "something went wrong there is not order with this id")
		return
	}
	c.HTML(http.StatusOK,"customer.tmpl", gin.H{
		"order": order,
	})

	//  receving the customer id from the url and then we will use the getorder function to get the order details and then we will render the customer.tmpl file with the order details
	// if orderId == "" {
	// 	c.String(http.StatusBadRequest, "order id is required")
	// 	return
	// }

	
}