package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
)

type orderManager interface {
	Update(order models.AccuralOrder) (*models.UserUpdate, error)
	GetQueue() []string
}

type Client struct {
	host          string
	clientService *resty.Client
	orderManager  orderManager
	stopChan      chan struct{}
}

func (c Client) sendRequest(order string) *models.AccuralOrder {
	res, err := c.clientService.NewRequest().Get(fmt.Sprintf("%s/api/orders/%s", c.host, order))

	if err != nil {
		logrus.Error(err)
		return nil
	}

	checkedOrder := models.AccuralOrder{}
	err = json.Unmarshal(res.Body(), &checkedOrder)

	if err != nil {
		logrus.Error(err)
		return nil
	}

	return &checkedOrder
}

func (c Client) Run() {
	c.stopChan = make(chan struct{})

	go func() {
		for {
			select {
			case <-c.stopChan:
				return
			default:
				queue := c.orderManager.GetQueue()

				for _, o := range queue {
					order := c.sendRequest(o)

					if order != nil {
						c.orderManager.Update(*order)
					}
				}

				time.Sleep(time.Second * 10)
			}
		}
	}()
}

func (c Client) Stop() {
	close(c.stopChan)
}

func NewClient(orderManager orderManager, host string) *Client {
	return &Client{
		clientService: resty.New(),
		host:          host,
		orderManager:  orderManager,
	}
}
