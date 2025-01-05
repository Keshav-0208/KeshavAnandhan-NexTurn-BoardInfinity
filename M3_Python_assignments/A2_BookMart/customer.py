class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display(self):
        return f"{self.name} | Email: {self.email} | Phone: {self.phone}"

customers = []
def add_customer():
    print("\n--- Add Customer ---")
    name = input("Name: ")
    email = input("Email: ")
    phone = input("Phone: ")

    if not name or not email or not phone:
        print("All fields are required!")
        return

    customer = Customer(name, email, phone)
    customers.append(customer)
    print("Customer added successfully!")

def view_customers():
    print("\n--- View Customers ---")
    if not customers:
        print("No customers found.")
    for i, customer in enumerate(customers, 1):
        print(f"{i}. {customer.display()}")
