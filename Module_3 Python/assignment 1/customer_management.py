customers = []
class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display_details(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"

def add_customer(name, email, phone):
    customer = Customer(name, email, phone)
    customers.append(customer)
    return "Customer added successfully!"

def view_customers():
    return [customer.display_details() for customer in customers]
