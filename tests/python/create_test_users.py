import requests
import time

# Configuration
BASE_URL = "http://localhost:8081"  # Adjust to your server URL
NUM_USERS = 10  # Number of test users to create

def register_user(username, password):
    """Register a new user"""
    data = {
        "password": password,
        "email": f"{username}@test.com"  # Adding email as it might be required
    }
    
    response = requests.post(
        f"{BASE_URL}/api/v1/register",
        json=data
    )
    return response.status_code

def main():
    print(f"Starting to create {NUM_USERS} test users...")
    
    successful_registrations = 0
    failed_registrations = 0
    
    for i in range(NUM_USERS):
        username = f"test_user_{i}"
        password = "password123"
        
        try:
            status_code = register_user(username, password)
            if status_code == 200:
                print(f"Successfully registered user: {username}")
                successful_registrations += 1
            else:
                print(f"Failed to register user: {username} (Status: {status_code})")
                failed_registrations += 1
        except Exception as e:
            print(f"Error registering user {username}: {str(e)}")
            failed_registrations += 1
            
        # Add small delay between registrations
        time.sleep(0.1)
    
    print("\nRegistration Summary:")
    print(f"Successful registrations: {successful_registrations}")
    print(f"Failed registrations: {failed_registrations}")
    
    if successful_registrations == NUM_USERS:
        print("\nAll test users created successfully!")
    else:
        print("\nWarning: Some user registrations failed!")
    
    print("\nTest User Credentials:")
    print("----------------------")
    for i in range(NUM_USERS):
        print(f"Username: test_user_{i}, Password: password123")

if __name__ == "__main__":
    main()