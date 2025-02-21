from flask import Flask, request, jsonify, render_template
import sqlite3

app = Flask(__name__)
DATABASE = 'database.db'


def init_db():
    conn = sqlite3.connect(DATABASE)
    cursor = conn.cursor()
    cursor.execute('''CREATE TABLE IF NOT EXISTS users (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        name TEXT NOT NULL,
                        email TEXT NOT NULL UNIQUE)''')
    conn.commit()
    conn.close()


def get_db_connection():
    """Helper function to get a database connection."""
    conn = sqlite3.connect(DATABASE)
    conn.row_factory = sqlite3.Row  # This allows returning dictionaries instead of tuples
    return conn


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/users', methods=['POST'])
def add_user():
    data = request.json

    if not data or 'name' not in data or 'email' not in data:
        return jsonify({"error": "Name and Email are required"}), 400  

    try:
        conn = sqlite3.connect(DATABASE)
        cursor = conn.cursor()
        cursor.execute("INSERT INTO users (name, email) VALUES (?, ?)", (data["name"], data["email"]))
        conn.commit()
        user_id = cursor.lastrowid
        conn.close()

        return jsonify({"id": user_id, "name": data["name"], "email": data["email"]}), 201

    except sqlite3.IntegrityError:
        return jsonify({"error": "Email already exists"}), 400  
 


@app.route('/users', methods=['GET'])
def get_users():
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users")
    users = [dict(row) for row in cursor.fetchall()]
    conn.close()
    return jsonify(users)


@app.route('/users/<int:user_id>', methods=['GET'])
def get_user(user_id):
    conn = get_db_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM users WHERE id = ?", (user_id,))
    user = cursor.fetchone()
    conn.close()
    
    if user:
        return jsonify(dict(user))
    return jsonify({"message": "User not found"}), 404


@app.route('/users/<int:user_id>', methods=['PUT'])
def update_user(user_id):
    data = request.json
    conn = get_db_connection()
    cursor = conn.cursor()

    cursor.execute("SELECT * FROM users WHERE id = ?", (user_id,))
    user = cursor.fetchone()

    if not user:
        return jsonify({"message": "User not found"}), 404

    cursor.execute("UPDATE users SET name = ?, email = ? WHERE id = ?", (data['name'], data['email'], user_id))
    conn.commit()
    conn.close()
    
    return jsonify({"message": "User updated successfully"})


@app.route('/users/<int:user_id>', methods=['DELETE'])
def delete_user(user_id):
    conn = get_db_connection()
    cursor = conn.cursor()

    cursor.execute("SELECT * FROM users WHERE id = ?", (user_id,))
    user = cursor.fetchone()

    if not user:
        return jsonify({"message": "User not found"}), 404

    cursor.execute("DELETE FROM users WHERE id = ?", (user_id,))
    conn.commit()
    conn.close()
    
    return jsonify({"message": "User deleted successfully"})


if __name__ == '__main__':
    init_db()
    app.run(host='0.0.0.0', port=5000, debug=True)
