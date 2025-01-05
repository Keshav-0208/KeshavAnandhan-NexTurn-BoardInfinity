from book import add_book, view_books, search_book
from customer import add_customer, view_customers
from sales import sell_book, view_sales

def main():
    while True:
        print("\n--- Welcome to BookMart ---")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            print("\n1. Add Book\n2. View Books\n3. Search Book")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                add_book()
            elif sub_choice == "2":
                view_books()
            elif sub_choice == "3":
                search_book()
            else:
                print("Invalid choice!")

        elif choice == "2":
            print("\n1. Add Customer\n2. View Customers")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                add_customer()
            elif sub_choice == "2":
                view_customers()
            else:
                print("Invalid choice!")

        elif choice == "3":
            print("\n1. Sell Book\n2. View Sales")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                sell_book()
            elif sub_choice == "2":
                view_sales()
            else:
                print("Invalid choice!")

        elif choice == "4":
            print("Exiting BookMart. Goodbye!")
            break

        else:
            print("Invalid choice!")

if __name__ == "__main__":
    main()
