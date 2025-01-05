books = []
class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"

def add_book(title, author, price, quantity):
    book = Book(title, author, price, quantity)
    books.append(book)
    return "Book added successfully!"

def view_books():
    return [book.display_details() for book in books]

def search_book(title):
    for book in books:
        if book.title.lower() == title.lower():
            return book.display_details()
    return "Book not found!"
