from book_management import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def display_menu():
    print("Welcome to BookMart!")
    print("1. Book Management")
    print("2. Customer Management")
    print("3. Sales Management")
    print("4. Exit")

def book_management():
    print("1. Add Book")
    print("2. View Books")
    print("3. Search Book")
    choice = int(input("Enter your choice: "))
    
    if choice == 1:
        title = input("Enter book title: ")
        author = input("Enter author: ")
        price = float(input("Enter price: "))
        quantity = int(input("Enter quantity: "))
        print(add_book(title, author, price, quantity))
    elif choice == 2:
        for book in view_books():
            print(book)
    elif choice == 3:
        title = input("Enter book title to search: ")
        print(search_book(title))

def customer_management():
    print("1. Add Customer")
    print("2. View Customers")
    choice = int(input("Enter your choice: "))
    
    if choice == 1:
        name = input("Enter name: ")
        email = input("Enter email: ")
        phone = input("Enter phone number: ")
        print(add_customer(name, email, phone))
    elif choice == 2:
        for customer in view_customers():
            print(customer)

def sales_management():
    print("1. Sell Book")
    print("2. View Sales Records")
    choice = int(input("Enter your choice: "))
    
    if choice == 1:
        customer_name = input("Enter customer name: ")
        book_title = input("Enter book title: ")
        quantity_sold = int(input("Enter quantity sold: "))
        print(sell_book(customer_name, book_title, quantity_sold))
    elif choice == 2:
        for sale in view_sales():
            print(sale)

def main():
    while True:
        display_menu()
        choice = int(input("Enter your choice: "))
        
        if choice == 1:
            book_management()
        elif choice == 2:
            customer_management()
        elif choice == 3:
            sales_management()
        elif choice == 4:
            break
        else:
            print("Invalid choice! Please try again.")

if __name__ == "__main__":
    main()
