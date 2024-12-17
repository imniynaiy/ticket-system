import mysql.connector
from mysql.connector import Error

def create_db_connection():
    try:
        connection = mysql.connector.connect(
            host="localhost",
            user="someone",
            password="password",
            database="some_db"
        )
        return connection
    except Error as e:
        print(f"Error connecting to MySQL Database: {e}")
        return None

def populate_seats():
    connection = create_db_connection()
    if connection is None:
        return

    try:
        cursor = connection.cursor()

        # First, let's get existing route_ids and seatclass_ids from their respective tables
        cursor.execute("SELECT route_id FROM route")
        route_ids = [r[0] for r in cursor.fetchall()]

        cursor.execute("SELECT seatclass_id FROM seatclass")
        seatclass_ids = [s[0] for s in cursor.fetchall()]

        if not route_ids or not seatclass_ids:
            print("No routes or seatclasses found. Please populate these tables first.")
            return

        # For each route, create seats for each seat class
        insert_query = """
            INSERT INTO seat (route_id, seatclass_id, status)
            VALUES (%s, %s, %s)
        """

        seat_data = []
        # Let's assume:
        # - Each route has seats for each seatclass
        # - Economy class (seatclass_id=1) has 50 seats
        # - Business class (seatclass_id=2) has 20 seats
        # - First class (seatclass_id=3) has 10 seats (if exists)
        seats_per_class = {
            1: 50,  # Economy
            2: 20,  # Business
            3: 10   # First Class
        }

        for route_id in route_ids:
            for seatclass_id in seatclass_ids:
                num_seats = seats_per_class.get(seatclass_id, 10)  # Default to 10 seats if class not in dict
                for _ in range(num_seats):
                    # Set all seats as available (status = 1)
                    seat_data.append((route_id, seatclass_id, 1))

        # Execute batch insert
        cursor.executemany(insert_query, seat_data)
        connection.commit()

        print(f"Successfully inserted {len(seat_data)} seats")

    except Error as e:
        print(f"Error: {e}")
    finally:
        if connection.is_connected():
            cursor.close()
            connection.close()
            print("MySQL connection closed")

if __name__ == "__main__":
    populate_seats()