import requests
import threading
import time
from concurrent.futures import ThreadPoolExecutor
import random

# Configuration
BASE_URL = "http://localhost:8081"  # Adjust to your server URL
NUM_USERS = 10  # Number of concurrent users
TARGET_SEAT_ID = 25  # The seat ID they're all trying to book

def login_user(email, password):
    """Login and get authentication token"""
    response = requests.post(
        f"{BASE_URL}/api/v1/login",
        json={
            "email": email,
            "password": password
        }
    )
    if response.status_code == 200:
        return response.json()["Token"]
    return None

def create_reservation(token, seat_id):
    """Attempt to create a reservation"""
    headers = {"Authorization": f"Bearer {token}"}
    data = {
        "reservation_date": "2024-01-01T15:04:05Z",  # Adjust date as needed
        "seat_id": seat_id,
        "passenger_family_name": "Test",
        "passenger_first_name": "User"
    }
    
    response = requests.post(
        f"{BASE_URL}/api/v1/user/reservations",
        headers=headers,
        json=data
    )
    return response.status_code

def simulate_user(user_id):
    """Simulate a user trying to book a seat"""
    # Login first (using test users that should exist in your system)
    email = f"test_user_{user_id}@test.com"
    password = "password123"
    
    token = login_user(email, password)
    if not token:
        print(f"User {user_id}: Failed to login")
        return
    
    # Add small random delay to simulate real-world conditions
    time.sleep(random.uniform(0, 0.5))
    
    # Try to book the seat
    status_code = create_reservation(token, TARGET_SEAT_ID)
    print(f"User {user_id}: Reservation attempt result - {status_code}")
    
    return status_code

def main():
    print(f"Starting simulation with {NUM_USERS} users trying to book seat {TARGET_SEAT_ID}")
    
    # Use ThreadPoolExecutor to run concurrent requests
    with ThreadPoolExecutor(max_workers=NUM_USERS) as executor:
        # Submit all tasks and gather futures
        futures = [executor.submit(simulate_user, i) for i in range(NUM_USERS)]
        
        # Wait for all futures to complete
        results = [future.result() for future in futures]
    
    # Count successful bookings (status code 200)
    successful_bookings = results.count(200)
    print(f"\nSimulation completed!")
    print(f"Successful bookings: {successful_bookings}")
    print(f"Failed bookings: {len(results) - successful_bookings}")
    
    if successful_bookings == 1:
        print("Test PASSED: Exactly one user succeeded in booking the seat!")
    else:
        print("Test FAILED: More than one user booked the same seat or no one succeeded!")

if __name__ == "__main__":
    main()