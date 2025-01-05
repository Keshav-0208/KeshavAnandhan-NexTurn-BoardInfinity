from flask import Flask, request, jsonify
import sqlite3
from datetime import datetime

app = Flask(__name__)

DB_PATH = "books.db"

VALID_GENRES = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi", "Fantasy", "Biography", "History"]

def connect_db():
    """Establish a connection to the SQLite database."""
    return sqlite3.connect(DB_PATH)

def is_valid_year(year):
    """Check if a given year is valid (4 digits and <= current year)."""
    try:
        year = int(year)
        return 1000 <= year <= datetime.now().year
    except ValueError:
        return False

@app.route("/books", methods=["POST"])
def add_book():
    """Add a new book to the collection."""
    data = request.json
    required_fields = ["title", "author", "published_year", "genre"]

    for field in required_fields:
        if field not in data:
            return jsonify({"error": "Invalid data", "message": f"Missing field: {field}"}), 400

    if not is_valid_year(data["published_year"]):
        return jsonify({"error": "Invalid data", "message": "Published year must be valid and 4 digits"}), 400

    if data["genre"] not in VALID_GENRES:
        return jsonify({"error": "Invalid data", "message": f"Genre must be one of {VALID_GENRES}"}), 400

    try:
        connection = connect_db()
        cursor = connection.cursor()

        cursor.execute("""
            INSERT INTO books (title, author, published_year, genre)
            VALUES (?, ?, ?, ?)
        """, (data["title"], data["author"], data["published_year"], data["genre"]))

        connection.commit()
        connection.close()

        return jsonify({"message": "Book added successfully", "book_id": cursor.lastrowid}), 201
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books", methods=["GET"])
def get_all_books():
    """Retrieve all books, filtered by 'Harry Potter' if specified."""
    try:
        connection = connect_db()
        cursor = connection.cursor()

        query = "SELECT * FROM books WHERE title LIKE ?"
        cursor.execute(query, ("%Harry Potter%",))
        books = cursor.fetchall()

        connection.close()

        return jsonify([
            {"id": book[0], "title": book[1], "author": book[2], "published_year": book[3], "genre": book[4]}
            for book in books
        ])
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books/<int:book_id>", methods=["GET"])
def get_book(book_id):
    """Retrieve a single book by its ID."""
    try:
        connection = connect_db()
        cursor = connection.cursor()

        cursor.execute("SELECT * FROM books WHERE id = ?", (book_id,))
        book = cursor.fetchone()

        connection.close()

        if book:
            return jsonify({"id": book[0], "title": book[1], "author": book[2], "published_year": book[3], "genre": book[4]})
        else:
            return jsonify({"error": "Book not found", "message": f"No book exists with ID {book_id}"}), 404
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books/<int:book_id>", methods=["PUT"])
def update_book(book_id):
    """Update details of an existing book."""
    data = request.json
    required_fields = ["title", "author", "published_year", "genre"]

    for field in required_fields:
        if field not in data:
            return jsonify({"error": "Invalid data", "message": f"Missing field: {field}"}), 400

    if not is_valid_year(data["published_year"]):
        return jsonify({"error": "Invalid data", "message": "Published year must be valid and 4 digits"}), 400

    if data["genre"] not in VALID_GENRES:
        return jsonify({"error": "Invalid data", "message": f"Genre must be one of {VALID_GENRES}"}), 400

    try:
        connection = connect_db()
        cursor = connection.cursor()

        cursor.execute("""
            UPDATE books
            SET title = ?, author = ?, published_year = ?, genre = ?
            WHERE id = ?
        """, (data["title"], data["author"], data["published_year"], data["genre"], book_id))

        connection.commit()
        connection.close()

        if cursor.rowcount == 0:
            return jsonify({"error": "Book not found", "message": f"No book exists with ID {book_id}"}), 404

        return jsonify({"message": "Book updated successfully"})
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route("/books/<int:book_id>", methods=["DELETE"])
def delete_book(book_id):
    """Delete a book by its ID."""
    try:
        connection = connect_db()
        cursor = connection.cursor()

        cursor.execute("DELETE FROM books WHERE id = ?", (book_id,))
        connection.commit()
        connection.close()

        if cursor.rowcount == 0:
            return jsonify({"error": "Book not found", "message": f"No book exists with ID {book_id}"}), 404

        return jsonify({"message": "Book deleted successfully"})
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

if __name__ == "__main__":
    app.run(debug=True)
