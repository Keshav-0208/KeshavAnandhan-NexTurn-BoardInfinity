import sqlite3

DB_PATH = "books.db"

#  Book Data
books_data = [
    ("Harry Potter and the Philosopher's Stone", "J.K. Rowling", 1997, "Fantasy"),
    ("Harry Potter and the Chamber of Secrets", "J.K. Rowling", 1998, "Fantasy"),
    ("Harry Potter and the Prisoner of Azkaban", "J.K. Rowling", 1999, "Fantasy"),
    ("Harry Potter and the Goblet of Fire", "J.K. Rowling", 2000, "Fantasy"),
    ("Harry Potter and the Order of the Phoenix", "J.K. Rowling", 2003, "Fantasy"),
    ("Harry Potter and the Half-Blood Prince", "J.K. Rowling", 2005, "Fantasy"),
    ("Harry Potter and the Deathly Hallows", "J.K. Rowling", 2007, "Fantasy"),
    ("Fantastic Beasts and Where to Find Them", "J.K. Rowling", 2001, "Fantasy"),
    ("Quidditch Through the Ages", "J.K. Rowling", 2001, "Fantasy"),
    ("The Tales of Beedle the Bard", "J.K. Rowling", 2008, "Fantasy")
]

def initialize_database():
    try:
        connection = sqlite3.connect(DB_PATH)
        cursor = connection.cursor()

        cursor.execute("""
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            author TEXT NOT NULL,
            published_year INTEGER NOT NULL,
            genre TEXT NOT NULL
        )
        """)

        cursor.executemany("""
        INSERT INTO books (title, author, published_year, genre)
        VALUES (?, ?, ?, ?)
        """, books_data)

        connection.commit()
        print("Database initialized and books data inserted successfully!")
    except Exception as e:
        print(f"Error: {e}")
    finally:
        connection.close()

if __name__ == "__main__":
    initialize_database()
