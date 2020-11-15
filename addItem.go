package main

func addItem(newItem Serve, data []Serve) []Serve {
	data = append(data, newItem)
	return data
}
