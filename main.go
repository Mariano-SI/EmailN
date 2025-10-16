package main

import (
	"bufio"
	"fmt"
	"golangestudo/model"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Lista de compras")
	reader := bufio.NewReader(os.Stdin)
	compra := model.Compra{}
	items := []model.Item{}
	compra.Itens = items

	fmt.Println("Qual será a data da compra? (dd/mm/yyyy): ")
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)

	convertedDate, _ := time.Parse("02/01/2006", data)

	compra.Data = convertedDate

	fmt.Println("Em qual supermercado a compra será realizada?: ")
	mercado, _ := reader.ReadString('\n')
	compra.Mercado = mercado
	for {
		item := model.Item{}
		fmt.Println("Digite o nome do item que você deseja adicionar ao carrinho: ")
		nome, _ := reader.ReadString('\n')
		item.Nome = strings.TrimSpace(nome)

		fmt.Println("Informe o valor do item:")
		valorStr, _ := reader.ReadString('\n')
		valorStr = strings.TrimSpace(valorStr)
		preco, _ := strconv.ParseFloat(valorStr, 64)
		item.Preco = preco

		fmt.Println("Informe a quantidade que será comprada:")
		qtdStr, _ := reader.ReadString('\n')
		qtdStr = strings.TrimSpace(qtdStr)
		qtd, _ := strconv.Atoi(qtdStr)
		item.Quantidade = uint8(qtd)

		items = append(items, item)

		var maisItems string

		for {
			fmt.Println("Deseja adicionar mais itens à compra? (S / N)")
			resposta, _ := reader.ReadString('\n')
			maisItems = strings.ToLower(strings.TrimSpace(resposta))

			if maisItems == "s" || maisItems == "n" {
				break
			}

			fmt.Println("Resposta inválida. Digite apenas S ou N.")
		}

		if maisItems == "n" {
			break
		}
	}

	fmt.Println("DETALHES DA COMPRA:")
	fmt.Println("Data: ", compra.Data)
	fmt.Println("Local: ", compra.Mercado)
	fmt.Println("Items: ")

	for _, item := range compra.Itens {
		fmt.Println("Nome: ", item.Nome, "Preço: ", item.Preco, "Quantidade: ", item.Quantidade)
		fmt.Println("--------------------------------------------------------------------------")
	}

}
