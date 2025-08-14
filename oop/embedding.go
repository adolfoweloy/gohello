// Package oop demonstrates Go's embedding and composition patterns including
// struct embedding, method promotion, and composition over inheritance.
package oop

import (
	"fmt"
	"strings"
	"time"
)

// DemoBasicEmbedding demonstrates basic struct embedding
func DemoBasicEmbedding() {
	fmt.Println("\n=== Basic Embedding Demo ===")

	// Create a user with embedded person
	user := User{
		Person: Person{
			Name:  "Alice Johnson",
			Age:   28,
			Email: "alice.johnson@example.com",
		},
		Username: "alice_j",
		Role:     "admin",
		IsActive: true,
	}

	fmt.Printf("User: %+v\n", user)

	// Accessing embedded fields directly
	fmt.Printf("Name (direct): %s\n", user.Name)
	fmt.Printf("Name (through Person): %s\n", user.Person.Name)

	// Calling embedded methods
	fmt.Printf("Display (embedded): %s\n", user.Display())
	fmt.Printf("Is Adult (embedded): %t\n", user.IsAdult())
	fmt.Printf("Email Domain (embedded): %s\n", user.GetEmailDomain())

	// Calling user-specific methods
	fmt.Printf("User Info: %s\n", user.GetUserInfo())
	fmt.Printf("Can Admin: %t\n", user.CanAdmin())

	// Modifying through embedded methods
	user.SetAge(29)
	user.HaveBirthday()
	fmt.Printf("After birthday: %s\n", user.Display())
}

// User embeds Person struct
type User struct {
	Person   // Embedded struct
	Username string
	Role     string
	IsActive bool
}

// GetUserInfo returns user-specific information
func (u User) GetUserInfo() string {
	return fmt.Sprintf("User: %s (@%s) - Role: %s, Active: %t", 
		u.Name, u.Username, u.Role, u.IsActive)
}

// CanAdmin checks if user has admin privileges
func (u User) CanAdmin() bool {
	return u.Role == "admin" && u.IsActive
}

// Login simulates user login
func (u *User) Login() {
	u.IsActive = true
	fmt.Printf("User %s logged in\n", u.Username)
}

// Logout simulates user logout
func (u *User) Logout() {
	u.IsActive = false
	fmt.Printf("User %s logged out\n", u.Username)
}

// DemoMethodPromotion demonstrates how embedded methods are promoted
func DemoMethodPromotion() {
	fmt.Println("\n=== Method Promotion Demo ===")

	// Create an admin user
	admin := AdminUser{
		User: User{
			Person: Person{
				Name:  "Bob Admin",
				Age:   35,
				Email: "bob@company.com",
			},
			Username: "bob_admin",
			Role:     "admin",
			IsActive: true,
		},
		Permissions: []string{"read", "write", "delete", "manage_users"},
		LastLogin:   time.Now(),
	}

	fmt.Printf("Admin User: %+v\n", admin)

	// Methods promoted from Person (through User)
	fmt.Printf("Display (from Person): %s\n", admin.Display())
	
	// Methods promoted from User
	fmt.Printf("User Info (from User): %s\n", admin.GetUserInfo())
	fmt.Printf("Can Admin (from User): %t\n", admin.CanAdmin())

	// AdminUser specific methods
	admin.GrantPermission("super_admin")
	fmt.Printf("Has super_admin permission: %t\n", admin.HasPermission("super_admin"))
	
	admin.ManageUser("alice", "promote")
}

// AdminUser embeds User, which embeds Person
type AdminUser struct {
	User        // Embedded User (which embeds Person)
	Permissions []string
	LastLogin   time.Time
}

// GrantPermission adds a permission to the admin user
func (au *AdminUser) GrantPermission(permission string) {
	au.Permissions = append(au.Permissions, permission)
	fmt.Printf("Granted %s permission to %s\n", permission, au.Username)
}

// HasPermission checks if admin user has specific permission
func (au AdminUser) HasPermission(permission string) bool {
	for _, p := range au.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

// ManageUser performs user management operations
func (au AdminUser) ManageUser(username, action string) {
	if au.HasPermission("manage_users") {
		fmt.Printf("Admin %s performed '%s' on user %s\n", au.Username, action, username)
	} else {
		fmt.Printf("Admin %s lacks permission to manage users\n", au.Username)
	}
}

// DemoInterfaceEmbedding demonstrates embedding interfaces
func DemoInterfaceEmbedding() {
	fmt.Println("\n=== Interface Embedding Demo ===")

	// Create a smart device that implements multiple interfaces
	device := SmartPhone{
		Device: Device{
			Name:         "iPhone 15",
			Brand:        "Apple",
			PowerConsumption: 15.5,
		},
		ScreenSize: 6.1,
		OS:         "iOS 17",
		Storage:    256,
	}

	fmt.Printf("Smart Device: %+v\n", device)

	// Use as different interfaces
	var powerable Powerable = &device
	var connectable Connectable = &device
	var smart SmartDevice = &device

	// Interface methods
	powerable.TurnOn()
	powerable.TurnOff()
	
	connectable.Connect("WiFi")
	connectable.Disconnect()
	
	// Smart device specific
	smart.UpdateSoftware()
	smart.SyncData()

	// Demonstrate polymorphism with embedded interfaces
	devices := []SmartDevice{&device}
	
	fmt.Println("\nManaging smart devices:")
	for _, d := range devices {
		ManageSmartDevice(d)
	}
}

// Powerable interface for devices that can be powered on/off
type Powerable interface {
	TurnOn()
	TurnOff()
	GetPowerConsumption() float64
}

// Connectable interface for devices that can connect to networks
type Connectable interface {
	Connect(network string)
	Disconnect()
	IsConnected() bool
}

// SmartDevice embeds both Powerable and Connectable interfaces
type SmartDevice interface {
	Powerable   // Embedded interface
	Connectable // Embedded interface
	UpdateSoftware()
	SyncData()
}

// Device is a basic device struct
type Device struct {
	Name             string
	Brand            string
	PowerConsumption float64
	IsOn             bool
	ConnectedTo      string
}

// TurnOn implements Powerable
func (d *Device) TurnOn() {
	d.IsOn = true
	fmt.Printf("%s is now ON\n", d.Name)
}

// TurnOff implements Powerable
func (d *Device) TurnOff() {
	d.IsOn = false
	fmt.Printf("%s is now OFF\n", d.Name)
}

// GetPowerConsumption implements Powerable
func (d Device) GetPowerConsumption() float64 {
	return d.PowerConsumption
}

// Connect implements Connectable
func (d *Device) Connect(network string) {
	d.ConnectedTo = network
	fmt.Printf("%s connected to %s\n", d.Name, network)
}

// Disconnect implements Connectable
func (d *Device) Disconnect() {
	if d.ConnectedTo != "" {
		fmt.Printf("%s disconnected from %s\n", d.Name, d.ConnectedTo)
		d.ConnectedTo = ""
	}
}

// IsConnected implements Connectable
func (d Device) IsConnected() bool {
	return d.ConnectedTo != ""
}

// SmartPhone embeds Device and implements SmartDevice
type SmartPhone struct {
	Device     // Embedded struct
	ScreenSize float64
	OS         string
	Storage    int
}

// UpdateSoftware implements SmartDevice
func (sp SmartPhone) UpdateSoftware() {
	fmt.Printf("Updating %s software (%s)\n", sp.Name, sp.OS)
}

// SyncData implements SmartDevice
func (sp SmartPhone) SyncData() {
	fmt.Printf("Syncing data for %s (%dGB storage)\n", sp.Name, sp.Storage)
}

// ManageSmartDevice demonstrates working with embedded interfaces
func ManageSmartDevice(device SmartDevice) {
	fmt.Printf("Managing device with power consumption: %.1fW\n", device.GetPowerConsumption())
	device.TurnOn()
	device.Connect("5G")
	device.UpdateSoftware()
	device.SyncData()
	device.Disconnect()
	device.TurnOff()
}

// DemoCompositionOverInheritance demonstrates composition patterns
func DemoCompositionOverInheritance() {
	fmt.Println("\n=== Composition Over Inheritance Demo ===")

	// Create components
	engine := Engine{
		Type:       "V8",
		Horsepower: 450,
		FuelType:   "Gasoline",
	}

	transmission := Transmission{
		Type:  "Automatic",
		Gears: 8,
	}

	wheels := []Wheel{
		{Size: 18, Brand: "Michelin", Type: "All-Season"},
		{Size: 18, Brand: "Michelin", Type: "All-Season"},
		{Size: 18, Brand: "Michelin", Type: "All-Season"},
		{Size: 18, Brand: "Michelin", Type: "All-Season"},
	}

	// Create car using composition
	car := Car{
		Make:         "Ford",
		Model:        "Mustang",
		Year:         2023,
		Engine:       engine,
		Transmission: transmission,
		Wheels:       wheels,
	}

	fmt.Printf("Car: %s %s %d\n", car.Make, car.Model, car.Year)

	// Use composed functionality
	car.Start()
	car.Accelerate()
	car.Shift("Drive")
	car.CheckTires()
	car.Stop()

	// Demonstrate different compositions
	fmt.Println("\nElectric car with different composition:")
	electricCar := ElectricCar{
		Car: Car{
			Make:  "Tesla",
			Model: "Model S",
			Year:  2023,
			Engine: Engine{
				Type:       "Electric",
				Horsepower: 670,
				FuelType:   "Electric",
			},
			Transmission: Transmission{
				Type:  "Single-Speed",
				Gears: 1,
			},
			Wheels: wheels,
		},
		BatteryCapacity: 100.0,
		Range:           405,
	}

	electricCar.Start()
	electricCar.Charge(80)
	electricCar.CheckBattery()
	electricCar.Stop()
}

// Car components using composition
type Engine struct {
	Type       string
	Horsepower int
	FuelType   string
	isRunning  bool
}

func (e *Engine) Start() {
	e.isRunning = true
	fmt.Printf("%s engine started (%d HP)\n", e.Type, e.Horsepower)
}

func (e *Engine) Stop() {
	e.isRunning = false
	fmt.Printf("%s engine stopped\n", e.Type)
}

func (e Engine) IsRunning() bool {
	return e.isRunning
}

type Transmission struct {
	Type         string
	Gears        int
	currentGear  string
}

func (t *Transmission) Shift(gear string) {
	t.currentGear = gear
	fmt.Printf("Shifted to %s (%s transmission)\n", gear, t.Type)
}

func (t Transmission) GetCurrentGear() string {
	return t.currentGear
}

type Wheel struct {
	Size     int
	Brand    string
	Type     string
	pressure float64
}

func (w *Wheel) CheckPressure() {
	fmt.Printf("%d\" %s %s tire - Pressure: %.1f PSI\n", 
		w.Size, w.Brand, w.Type, w.pressure)
}

// Car uses composition instead of inheritance
type Car struct {
	Make         string
	Model        string
	Year         int
	Engine       Engine       // Composition
	Transmission Transmission // Composition
	Wheels       []Wheel      // Composition
	isRunning    bool
}

func (c *Car) Start() {
	if !c.isRunning {
		c.Engine.Start()
		c.isRunning = true
		fmt.Printf("%s %s is ready to drive\n", c.Make, c.Model)
	}
}

func (c *Car) Stop() {
	if c.isRunning {
		c.Engine.Stop()
		c.isRunning = false
		fmt.Printf("%s %s stopped\n", c.Make, c.Model)
	}
}

func (c *Car) Accelerate() {
	if c.isRunning {
		fmt.Printf("%s %s is accelerating\n", c.Make, c.Model)
	} else {
		fmt.Printf("Cannot accelerate - car is not running\n")
	}
}

func (c *Car) Shift(gear string) {
	c.Transmission.Shift(gear)
}

func (c *Car) CheckTires() {
	fmt.Println("Checking tires:")
	for i, wheel := range c.Wheels {
		fmt.Printf("  Wheel %d: ", i+1)
		wheel.CheckPressure()
	}
}

// ElectricCar extends Car with additional electric-specific features
type ElectricCar struct {
	Car             // Embedded Car
	BatteryCapacity float64
	Range           int
	batteryLevel    float64
}

func (ec *ElectricCar) Charge(targetLevel float64) {
	ec.batteryLevel = targetLevel
	fmt.Printf("Charging %s %s to %.0f%%\n", ec.Make, ec.Model, targetLevel)
}

func (ec ElectricCar) CheckBattery() {
	fmt.Printf("Battery: %.0f%% (Capacity: %.1f kWh, Range: %d miles)\n", 
		ec.batteryLevel, ec.BatteryCapacity, ec.Range)
}

// DemoAggregationVsComposition demonstrates different relationship types
func DemoAggregationVsComposition() {
	fmt.Println("\n=== Aggregation vs Composition Demo ===")

	// Composition: Library owns Books (strong relationship)
	library := Library{
		Name: "City Library",
		Books: []Book{
			{Title: "Go Programming", Author: "John Doe", ISBN: "123-456"},
			{Title: "Design Patterns", Author: "Gang of Four", ISBN: "789-012"},
		},
	}

	fmt.Printf("Library: %s\n", library.Name)
	library.ListBooks()
	library.AddBook(Book{Title: "Clean Code", Author: "Robert Martin", ISBN: "345-678"})

	// Aggregation: Team has Members (weak relationship - members exist independently)
	member1 := TeamMember{Name: "Alice", Role: "Developer", Skills: []string{"Go", "Python"}}
	member2 := TeamMember{Name: "Bob", Role: "Designer", Skills: []string{"Figma", "Photoshop"}}

	team := Team{
		Name:    "Backend Team",
		Lead:    "Charlie",
		Members: []*TeamMember{&member1, &member2}, // References to existing members
	}

	fmt.Printf("\nTeam: %s\n", team.Name)
	team.ListMembers()
	
	// Members exist independently
	member1.AddSkill("Docker")
	fmt.Printf("Member 1 skills after update: %v\n", member1.Skills)

	// Demonstrate the difference when objects are destroyed
	fmt.Println("\nDemonstrating relationship differences:")
	fmt.Println("- Composition: When Library is destroyed, Books are destroyed")
	fmt.Println("- Aggregation: When Team is disbanded, Members still exist")
}

// Composition example: Library owns Books
type Library struct {
	Name  string
	Books []Book // Composition - books belong to library
}

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
	fmt.Printf("Added book: %s\n", book.Title)
}

func (l Library) ListBooks() {
	fmt.Println("Books in library:")
	for i, book := range l.Books {
		fmt.Printf("  %d. %s by %s\n", i+1, book.Title, book.Author)
	}
}

// Aggregation example: Team has Members
type Team struct {
	Name    string
	Lead    string
	Members []*TeamMember // Aggregation - references to independent members
}

type TeamMember struct {
	Name   string
	Role   string
	Skills []string
}

func (tm *TeamMember) AddSkill(skill string) {
	tm.Skills = append(tm.Skills, skill)
}

func (t Team) ListMembers() {
	fmt.Printf("Team %s (Lead: %s):\n", t.Name, t.Lead)
	for i, member := range t.Members {
		fmt.Printf("  %d. %s - %s (Skills: %s)\n", 
			i+1, member.Name, member.Role, strings.Join(member.Skills, ", "))
	}
}

// DemoEmbeddingBestPractices demonstrates best practices for embedding
func DemoEmbeddingBestPractices() {
	fmt.Println("\n=== Embedding Best Practices Demo ===")

	// 1. Use embedding for "is-a" relationships
	manager := Manager{
		Employee: Employee{
			Person: Person{
				Name:  "Sarah Manager",
				Age:   40,
				Email: "sarah@company.com",
			},
			ID:         "MGR001",
			Department: "Engineering",
			Salary:     95000,
		},
		TeamSize:   5,
		Budget:     500000,
		Directs:    []string{"EMP001", "EMP002", "EMP003"},
	}

	fmt.Printf("Manager: %s\n", manager.Display())
	fmt.Printf("Manages team of %d people\n", manager.TeamSize)

	// 2. Avoid embedding when you need to override behavior completely
	consultant := Consultant{
		person:   Person{Name: "John Consultant", Age: 45, Email: "john@consulting.com"},
		Company:  "TechConsult Inc",
		HourlyRate: 150.0,
		Contract:   "6-month project",
	}

	fmt.Printf("\nConsultant: %s\n", consultant.GetInfo()) // Custom method, not embedded
	fmt.Printf("Hourly rate: $%.2f\n", consultant.HourlyRate)

	// 3. Use composition when objects have different lifecycles
	project := Project{
		Name:        "Web Redesign",
		Description: "Redesign company website",
		Manager:     &manager,    // Reference to existing manager
		Consultant:  &consultant, // Reference to existing consultant
		StartDate:   time.Now(),
		Duration:    6,
	}

	project.StartProject()
	project.GetProjectInfo()

	fmt.Println("\nBest practices demonstrated:")
	fmt.Println("1. Embedding for 'is-a' relationships (Manager is an Employee)")
	fmt.Println("2. Composition for 'has-a' relationships (Project has Manager)")
	fmt.Println("3. Avoid embedding when behavior differs significantly")
}

// Manager embeds Employee (is-a relationship)
type Manager struct {
	Employee // Embedded - Manager is an Employee
	TeamSize int
	Budget   float64
	Directs  []string
}

// Consultant uses composition instead of embedding Person
type Consultant struct {
	person     Person // Composition - has-a relationship
	Company    string
	HourlyRate float64
	Contract   string
}

// GetInfo provides consultant-specific information
func (c Consultant) GetInfo() string {
	return fmt.Sprintf("%s works for %s (%s)", 
		c.person.Name, c.Company, c.Contract)
}

// GetName provides access to embedded person's name
func (c Consultant) GetName() string {
	return c.person.Name
}

// Project uses aggregation/composition
type Project struct {
	Name        string
	Description string
	Manager     *Manager    // Aggregation - reference to existing manager
	Consultant  *Consultant // Aggregation - reference to existing consultant
	StartDate   time.Time
	Duration    int // months
}

func (p *Project) StartProject() {
	fmt.Printf("Starting project: %s\n", p.Name)
	fmt.Printf("Manager: %s\n", p.Manager.Name)
	fmt.Printf("Consultant: %s\n", p.Consultant.GetName())
}

func (p Project) GetProjectInfo() {
	fmt.Printf("Project: %s (%s)\n", p.Name, p.Description)
	fmt.Printf("Duration: %d months\n", p.Duration)
	fmt.Printf("Start: %s\n", p.StartDate.Format("2006-01-02"))
}