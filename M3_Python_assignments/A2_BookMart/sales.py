from customer import Customer
from book import books

class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

    def display(self):
        return f"{self.name} bought {self.quantity_sold} copy/copies of {self.book_title}"

sales = []
def sell_book():
    print("\n--- Sell Book ---")
    customer_name = input("Customer Name: ")
    customer_email = input("Customer Email: ")
    customer_phone = input("Customer Phone: ")
    book_title = input("Book Title: ")
    quantity = int(input("Quantity: "))

    book = next((book for book in books if book.title.lower() == book_title.lower()), None)
    if not book:
        print("Book not found!")
        return
    if book.quantity < quantity:
        print(f"Error: Only {book.quantity} copies available.")
        return

    book.quantity -= quantity
    transaction = Transaction(customer_name, customer_email, customer_phone, book.title, quantity)
    sales.append(transaction)
    print("Sale successful!")

def view_sales():
    print("\n--- View Sales ----")
    if not sales:
        print("No sales recorded.")
    for i, sale in enumerate(sales, 1):
        print(f"{i}. {sale.display()}")
