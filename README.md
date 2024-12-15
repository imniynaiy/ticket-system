# Ticket System

Ticket System based on H26 AP Test.

## Functions
Basic
- ✅ User login (Redis), registeration✅
- Logs
- Error

Without Login
- ✅ Register
- List routes
- Check availble count of seats by Route

Administrators✅
- Add, update, delete, search Routes
- Add, update, delete, search Seats
- Add, update, delete, search Seatclasses

User
- Check availble seats by Route and Seatclass
- Order seat



## Usage

1. Run the sql script to set up database.
2. Update configs (default config locates at `configs/config.yml`)
3. Run the server (Or build docker image `make docker`)