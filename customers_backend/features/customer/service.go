package customer

func GetCustomers() []Customer {
	return DbGetCustomers()
}
func GetCustomer(id int) Customer {
	return DbGetCustomer(id)
}
func AddCustomer(user Customer) Customer {
	return DbAddCustomer(user)
}
func UpdateCustomer(user Customer) Customer {
	return DbUpdateCustomer(user)
}
func DeleteCustomer(id int) Customer {
	return DbDeleteCustomer(id)
}
