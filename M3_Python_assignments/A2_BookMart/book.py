class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display(self):
        return f"{self.title} by {self.author} - ${self.price} (Qty: {self.quantity})"

books = []
def add_book():
    print("\n--- Add Book ---")
    try:
        title = input("Title: ")
        author = input("Author: ")
        price = float(input("Price: "))
        quantity = int(input("Quantity: "))

        if price <= 0 or quantity <= 0:
            print("Price and Quantity must be positive numbers!")
            return
        
        book = Book(title, author, price, quantity)
        books.append(book)
        print("Book added successfully!")
    except ValueError:
        print("Invalid input! Price and Quantity must be numbers.")

def view_books():
    print("\n--- View Books ---")
    if not books:
        print("No books available.")
    for i, book in enumerate(books, 1):
        print(f"{i}. {book.display()}")

def search_book():
    print("\n--- Search Book ---")
    query = input("Enter title or author: ").lower()
    found_books = [book for book in books if query in book.title.lower() or query in book.author.lower()]
    if found_books:
        for book in found_books:
            print(book.display())
    else:
        print("No books found!")
