package transport

import (
	"L0/iternal/domain"
	"context"

	"github.com/gin-gonic/gin"
)

type OrderService interface {
	Create(ctx context.Context, orderUid string, order domain.Order) error
	GetById(ctx context.Context, orderUid string) (domain.Order, error)
}

type Handler struct {
	service OrderService
}

func NewHandler(service OrderService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("../web/templates/index.html")
	r.Static("/js", "../web/static/js")
	r.Static("/css", "../web/static/css")
	
	api := r.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.GET("", func(c *gin.Context) {
				c.HTML(200, "index.html", map[string]string{"title": "home page"})
			})
			orders.POST("", h.Create)
			orders.GET("/:orderuid", h.GetByid)
		}
	}
	return r
}