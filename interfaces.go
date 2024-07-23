package main

import (
	"fmt"
	"net"
)

func main() {
	// Получаем все доступные сетевые интерфейсы
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		// Список адресов для каждого сетевого интерфейса
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Сетевой интерфейс: %s\n", interf.Name)

		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\tIP: %v\n", ip)
			}
		}
	}
}
