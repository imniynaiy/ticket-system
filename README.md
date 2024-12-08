# Ticket System

Ticket System based on H26 AP Test.

## Functions
Basic
- User login (Redis✅), registeration
- Administrator login and Permissions (Casbin)
- Logs

Without Login
- Register
- List routes
- Check availble count of seats by Route

Administrators
- Add, update, delete, search Routes
- Add, update, delete Customers
- Add, update, delete Users
- Add, update, delete, search Seats
- Add, update, delete, search Seatclasses

User
- Check availble seats by Route and Seatclass
- Order seat



## Usage

1. Generate password (`go run tool/genpassword.go -p your-password -s salt-as-in-config`)
2. Update the username and password in init/database.sql
3. Run the sql script to set up database.
4. Build the [frontend project](https://github.com/theoriz0/flome-react), copy the dist files under 
5. Update configs (default config locates at `configs/config.yml`)
6. Run the server (Or build docker image `make docker`)