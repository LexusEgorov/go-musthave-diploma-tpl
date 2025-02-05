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

type balanceManager interface {
	Inc(uID int, count float64) error
}

type Client struct {
	host           string
	clientService  *resty.Client
	orderManager   orderManager
	balanceManager balanceManager
	stopChan       chan struct{}
}

func (c Client) sendRequest(order string) *models.AccuralOrder {
	requestHost := fmt.Sprintf("%s/api/orders/%s", c.host, order)
	logrus.Info(fmt.Sprintf("sending request: %s", requestHost))

	res, err := c.clientService.NewRequest().Get(requestHost)

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
						update, err := c.orderManager.Update(*order)

						if err != nil {
							logrus.Error(err)
							continue
						}

						if order.Status == models.ProcessedStatus {
							c.balanceManager.Inc(update.ID, update.Count)
						}
					}
				}

				time.Sleep(time.Second * 1)
			}
		}
	}()
}

func (c Client) Stop() {
	close(c.stopChan)
}

func NewClient(orderManager orderManager, balanceManager balanceManager, host string) *Client {
	return &Client{
		clientService:  resty.New(),
		host:           host,
		orderManager:   orderManager,
		balanceManager: balanceManager,
	}
}
