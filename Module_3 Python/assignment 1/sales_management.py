from book_management import books
class Transaction:
    def __init__(self, customer_name, book_title, quantity_sold):
        self.customer_name = customer_name
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display_transaction(self):
        return f"Customer: {self.customer_name}, Book: {self.book_title}, Quantity Sold: {self.quantity_sold}"

sales_records = []

def sell_book(customer_name, book_title, quantity_sold):
    for book in books:
        if book.title.lower() == book_title.lower():
            if book.quantity >= quantity_sold:
                book.quantity -= quantity_sold
                transaction = Transaction(customer_name, book_title, quantity_sold)
                sales_records.append(transaction)
                return f"Sale successful! Remaining quantity: {book.quantity}"
            else:
                return f"Error: Only {book.quantity} copies available. Sale cannot be completed."
    return "Book not found!"

def view_sales():
    return [transaction.display_transaction() for transaction in sales_records]
