package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"encrypted-db/internal/database"
	"encrypted-db/internal/models"
	"encrypted-db/internal/utils"
)

func main() {
	var (
		action   = flag.String("action", "", "Action: init, create, list, get, update, delete")
		name     = flag.String("name", "", "User name")
		email    = flag.String("email", "", "User email")
		phone    = flag.String("phone", "", "User phone")
		address  = flag.String("address", "", "User address")
		id       = flag.Int("id", 0, "User ID")
		password = flag.String("password", "", "Master password")
		dbPath   = flag.String("db", "encrypted.db", "Database file path")
	)
	flag.Parse()

	var masterPassword string
	if *password != "" {
		masterPassword = *password
	} else {
		var err error
		masterPassword, err = utils.PromptPassword("Enter master password: ")
		if err != nil {
			log.Fatalf("Failed to read password: %v", err)
		}
	}

	if err := utils.ValidatePassword(masterPassword); err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDB(*dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	switch *action {

	case "init":
		if err := db.Init(); err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}
		fmt.Println("Database initialized successfully.")

	case "create":
		if *name == "" || *email == "" {
			log.Fatal("Name and email are required.")
		}
		user := models.NewUser(*name, *email, *phone, *address)
		if err := db.CreateUser(user, *password); err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
		fmt.Printf("User created with ID: %d\n", user.ID)

	case "list":
		users, err := db.ListUsers(*password)
		if err != nil {
			log.Fatalf("Failed to list users: %v", err)
		}
		if len(users) == 0 {
			fmt.Println("No users found.")
		}
		for _, user := range users {
			fmt.Println(utils.FormatUser(user))
			fmt.Println("---")
		}

	case "get":
		if *id == 0 {
			log.Fatal("ID is required.")
		}
		user, err := db.GetUser(*id, *password)
		if err != nil {
			log.Fatalf("Failed to get user: %v", err)
		}
		fmt.Println(utils.FormatUser(user))

	case "update":
		if *id == 0 {
			log.Fatal("ID is required.")
		}
		user, err := db.GetUser(*id, *password)
		if err != nil {
			log.Fatalf("Failed to get user: %v", err)
		}
		if *name != "" {
			user.Name = *name
		}
		if *email != "" {
			user.Email = *email
		}
		if *phone != "" {
			user.Phone = *phone
		}
		if *address != "" {
			user.Address = *address
		}
		if err := db.UpdateUser(user, *password); err != nil {
			log.Fatalf("Failed to update user: %v", err)
		}
		fmt.Println("User updated.")

	case "delete":
		if *id == 0 {
			log.Fatal("ID is required.")
		}
		if err := db.DeleteUser(*id); err != nil {
			log.Fatalf("Failed to delete user: %v", err)
		}
		fmt.Println("User deleted.")

	default:
		fmt.Println("Usage:")
		fmt.Println("  -action=init")
		fmt.Println("  -action=create -name=<name> -email=<email>")
		fmt.Println("  -action=list")
		fmt.Println("  -action=rawlist")
		fmt.Println("  -action=get -id=<id>")
		fmt.Println("  -action=update -id=<id>")
		fmt.Println("  -action=delete -id=<id>")
		os.Exit(1)
	}
}
