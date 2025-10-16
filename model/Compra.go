package model

import "time"

type Compra struct {
	Data time.Time
	Mercado string
	Itens []Item
}