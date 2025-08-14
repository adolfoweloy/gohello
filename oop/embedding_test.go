package oop

import (
	"strings"
	"testing"
	"time"
)

func TestDemoBasicEmbedding(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoBasicEmbedding panicked: %v", r)
		}
	}()
	
	DemoBasicEmbedding()
}

func TestDemoMethodPromotion(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoMethodPromotion panicked: %v", r)
		}
	}()
	
	DemoMethodPromotion()
}

func TestDemoInterfaceEmbedding(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoInterfaceEmbedding panicked: %v", r)
		}
	}()
	
	DemoInterfaceEmbedding()
}

func TestDemoCompositionOverInheritance(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoCompositionOverInheritance panicked: %v", r)
		}
	}()
	
	DemoCompositionOverInheritance()
}

func TestDemoAggregationVsComposition(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoAggregationVsComposition panicked: %v", r)
		}
	}()
	
	DemoAggregationVsComposition()
}

func TestDemoEmbeddingBestPractices(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("DemoEmbeddingBestPractices panicked: %v", r)
		}
	}()
	
	DemoEmbeddingBestPractices()
}

func TestUserEmbedding(t *testing.T) {
	user := User{
		Person: Person{
			Name:  "Test User",
			Age:   25,
			Email: "test@example.com",
		},
		Username: "testuser",
		Role:     "user",
		IsActive: true,
	}
	
	// Test embedded field access
	if user.Name != "Test User" {
		t.Errorf("Expected name 'Test User', got '%s'", user.Name)
	}
	
	// Test embedded method access
	if !user.IsAdult() {
		t.Errorf("25-year-old should be adult")
	}
	
	domain := user.GetEmailDomain()
	if domain != "example.com" {
		t.Errorf("Expected domain 'example.com', got '%s'", domain)
	}
	
	// Test user-specific methods
	if !user.CanAdmin() && user.Role == "admin" {
		t.Errorf("Admin user should be able to admin")
	}
	
	userInfo := user.GetUserInfo()
	if !strings.Contains(userInfo, "testuser") {
		t.Errorf("User info should contain username")
	}
}

func TestUserMethods(t *testing.T) {
	user := User{
		Person:   Person{Name: "Test", Age: 20, Email: "test@example.com"},
		Username: "test",
		Role:     "user",
		IsActive: false,
	}
	
	// Test login
	user.Login()
	if !user.IsActive {
		t.Errorf("User should be active after login")
	}
	
	// Test logout
	user.Logout()
	if user.IsActive {
		t.Errorf("User should be inactive after logout")
	}
}

func TestAdminUserEmbedding(t *testing.T) {
	admin := AdminUser{
		User: User{
			Person:   Person{Name: "Admin", Age: 30, Email: "admin@example.com"},
			Username: "admin",
			Role:     "admin",
			IsActive: true,
		},
		Permissions: []string{"read", "write"},
		LastLogin:   time.Now(),
	}
	
	// Test method promotion from Person (through User)
	if !admin.IsAdult() {
		t.Errorf("30-year-old should be adult")
	}
	
	// Test method promotion from User
	if !admin.CanAdmin() {
		t.Errorf("Admin user should be able to admin")
	}
	
	// Test AdminUser specific methods
	admin.GrantPermission("delete")
	if !admin.HasPermission("delete") {
		t.Errorf("Should have delete permission after granting")
	}
	
	if admin.HasPermission("nonexistent") {
		t.Errorf("Should not have nonexistent permission")
	}
}

func TestDeviceEmbedding(t *testing.T) {
	device := Device{
		Name:             "Test Device",
		Brand:            "TestBrand",
		PowerConsumption: 10.0,
	}
	
	// Test Powerable interface methods
	device.TurnOn()
	if !device.IsOn {
		t.Errorf("Device should be on after TurnOn")
	}
	
	device.TurnOff()
	if device.IsOn {
		t.Errorf("Device should be off after TurnOff")
	}
	
	consumption := device.GetPowerConsumption()
	if consumption != 10.0 {
		t.Errorf("Expected power consumption 10.0, got %.1f", consumption)
	}
	
	// Test Connectable interface methods
	device.Connect("WiFi")
	if !device.IsConnected() {
		t.Errorf("Device should be connected after Connect")
	}
	if device.ConnectedTo != "WiFi" {
		t.Errorf("Device should be connected to WiFi")
	}
	
	device.Disconnect()
	if device.IsConnected() {
		t.Errorf("Device should not be connected after Disconnect")
	}
}

func TestSmartPhoneInterfaces(t *testing.T) {
	phone := SmartPhone{
		Device: Device{
			Name:             "Test Phone",
			Brand:            "TestBrand",
			PowerConsumption: 5.0,
		},
		ScreenSize: 6.0,
		OS:         "TestOS",
		Storage:    128,
	}
	
	// Test that SmartPhone implements all required interfaces
	var powerable Powerable = &phone
	var connectable Connectable = &phone
	var smart SmartDevice = &phone
	
	// Test through interfaces
	powerable.TurnOn()
	connectable.Connect("5G")
	smart.UpdateSoftware()
	smart.SyncData()
	connectable.Disconnect()
	powerable.TurnOff()
	
	// Verify state changes
	if phone.GetPowerConsumption() != 5.0 {
		t.Errorf("Power consumption should be preserved")
	}
}

func TestCarComposition(t *testing.T) {
	engine := Engine{Type: "V6", Horsepower: 300, FuelType: "Gasoline"}
	transmission := Transmission{Type: "Manual", Gears: 6}
	wheels := []Wheel{{Size: 17, Brand: "Goodyear", Type: "Performance"}}
	
	car := Car{
		Make:         "Honda",
		Model:        "Civic",
		Year:         2023,
		Engine:       engine,
		Transmission: transmission,
		Wheels:       wheels,
	}
	
	// Test car operations
	if car.isRunning {
		t.Errorf("Car should not be running initially")
	}
	
	car.Start()
	if !car.isRunning {
		t.Errorf("Car should be running after start")
	}
	if !car.Engine.IsRunning() {
		t.Errorf("Engine should be running after car start")
	}
	
	car.Shift("First")
	if car.Transmission.GetCurrentGear() != "First" {
		t.Errorf("Transmission should be in First gear")
	}
	
	car.Stop()
	if car.isRunning {
		t.Errorf("Car should not be running after stop")
	}
	if car.Engine.IsRunning() {
		t.Errorf("Engine should not be running after car stop")
	}
}

func TestElectricCarExtension(t *testing.T) {
	electricCar := ElectricCar{
		Car: Car{
			Make:   "Tesla",
			Model:  "Model 3",
			Year:   2023,
			Engine: Engine{Type: "Electric", Horsepower: 400, FuelType: "Electric"},
		},
		BatteryCapacity: 75.0,
		Range:           350,
		batteryLevel:    50.0,
	}
	
	// Test inherited car functionality
	electricCar.Start()
	if !electricCar.isRunning {
		t.Errorf("Electric car should be running after start")
	}
	
	// Test electric-specific functionality
	electricCar.Charge(90.0)
	if electricCar.batteryLevel != 90.0 {
		t.Errorf("Expected battery level 90%%, got %.1f%%", electricCar.batteryLevel)
	}
	
	electricCar.CheckBattery() // Should not panic
}

func TestLibraryComposition(t *testing.T) {
	library := Library{
		Name:  "Test Library",
		Books: []Book{},
	}
	
	book1 := Book{Title: "Go Programming", Author: "John Doe", ISBN: "123"}
	book2 := Book{Title: "Design Patterns", Author: "Gang of Four", ISBN: "456"}
	
	library.AddBook(book1)
	library.AddBook(book2)
	
	if len(library.Books) != 2 {
		t.Errorf("Expected 2 books, got %d", len(library.Books))
	}
	
	if library.Books[0].Title != "Go Programming" {
		t.Errorf("First book title incorrect")
	}
}

func TestTeamAggregation(t *testing.T) {
	member1 := TeamMember{Name: "Alice", Role: "Developer", Skills: []string{"Go"}}
	member2 := TeamMember{Name: "Bob", Role: "Designer", Skills: []string{"CSS"}}
	
	team := Team{
		Name:    "Dev Team",
		Lead:    "Charlie",
		Members: []*TeamMember{&member1, &member2},
	}
	
	if len(team.Members) != 2 {
		t.Errorf("Expected 2 team members, got %d", len(team.Members))
	}
	
	// Test that members exist independently
	member1.AddSkill("Python")
	if len(member1.Skills) != 2 {
		t.Errorf("Member should have 2 skills after adding one")
	}
	
	// Verify the team member reference reflects the change
	if len(team.Members[0].Skills) != 2 {
		t.Errorf("Team member reference should reflect skill addition")
	}
}

func TestManagerEmbedding(t *testing.T) {
	manager := Manager{
		Employee: Employee{
			Person: Person{Name: "Manager", Age: 35, Email: "mgr@company.com"},
			ID:     "MGR001",
			Salary: 90000,
		},
		TeamSize: 5,
		Budget:   100000,
		Directs:  []string{"EMP001", "EMP002"},
	}
	
	// Test embedded field access
	if manager.Name != "Manager" {
		t.Errorf("Expected name 'Manager', got '%s'", manager.Name)
	}
	
	if manager.ID != "MGR001" {
		t.Errorf("Expected ID 'MGR001', got '%s'", manager.ID)
	}
	
	// Test embedded method access
	if !manager.IsAdult() {
		t.Errorf("35-year-old should be adult")
	}
	
	// Test manager-specific fields
	if manager.TeamSize != 5 {
		t.Errorf("Expected team size 5, got %d", manager.TeamSize)
	}
	
	if len(manager.Directs) != 2 {
		t.Errorf("Expected 2 direct reports, got %d", len(manager.Directs))
	}
}

func TestConsultantComposition(t *testing.T) {
	consultant := Consultant{
		person:     Person{Name: "Consultant", Age: 40, Email: "consultant@company.com"},
		Company:    "ConsultCorp",
		HourlyRate: 200.0,
		Contract:   "3-month project",
	}
	
	// Test composition access
	if consultant.GetName() != "Consultant" {
		t.Errorf("Expected name 'Consultant', got '%s'", consultant.GetName())
	}
	
	info := consultant.GetInfo()
	if !strings.Contains(info, "ConsultCorp") {
		t.Errorf("Info should contain company name")
	}
	
	if consultant.HourlyRate != 200.0 {
		t.Errorf("Expected hourly rate 200.0, got %.1f", consultant.HourlyRate)
	}
}

func TestProjectAggregation(t *testing.T) {
	manager := &Manager{
		Employee: Employee{
			Person: Person{Name: "Project Manager", Age: 35, Email: "pm@company.com"},
			ID:     "PM001",
		},
		TeamSize: 3,
	}
	
	consultant := &Consultant{
		person:     Person{Name: "Expert Consultant", Age: 45, Email: "expert@consulting.com"},
		Company:    "ExpertCorp",
		HourlyRate: 250.0,
	}
	
	project := Project{
		Name:        "Big Project",
		Description: "Important project",
		Manager:     manager,
		Consultant:  consultant,
		StartDate:   time.Now(),
		Duration:    6,
	}
	
	// Test project fields
	if project.Name != "Big Project" {
		t.Errorf("Expected project name 'Big Project', got '%s'", project.Name)
	}
	
	if project.Duration != 6 {
		t.Errorf("Expected duration 6, got %d", project.Duration)
	}
	
	// Test references to aggregated objects
	if project.Manager.Name != "Project Manager" {
		t.Errorf("Manager name should be accessible through project")
	}
	
	if project.Consultant.GetName() != "Expert Consultant" {
		t.Errorf("Consultant name should be accessible through project")
	}
	
	// Test that project methods work
	project.StartProject() // Should not panic
	project.GetProjectInfo() // Should not panic
}

func TestEmbeddingVsComposition(t *testing.T) {
	// Embedding: Manager IS-A Employee
	manager := Manager{
		Employee: Employee{
			Person: Person{Name: "Manager", Age: 35, Email: "mgr@company.com"},
			ID:     "MGR001",
		},
		TeamSize: 5,
	}
	
	// Can access Employee methods directly
	display := manager.Display() // From embedded Employee's embedded Person
	if !strings.Contains(display, "Manager") {
		t.Errorf("Manager should be able to use embedded Display method")
	}
	
	// Composition: Consultant HAS-A Person
	consultant := Consultant{
		person:  Person{Name: "Consultant", Age: 40, Email: "consultant@company.com"},
		Company: "ConsultCorp",
	}
	
	// Cannot access Person methods directly, need explicit access
	name := consultant.GetName() // Explicitly defined method
	if name != "Consultant" {
		t.Errorf("Consultant should provide access to person's name through composition")
	}
	
	// This would not compile (person is not embedded):
	// consultant.Display() // ERROR
	// consultant.Age       // ERROR
}

func TestInterfaceComposition(t *testing.T) {
	processor := &DataProcessor{}
	
	// Test that DataProcessor implements composed interfaces
	var readWriter ReadWriter = processor
	var readWriteCloser ReadWriteCloser = processor
	
	// Should be able to use through composed interfaces
	data := []byte("test")
	readWriter.Write(data)
	
	buffer := make([]byte, len(data))
	readWriter.Read(buffer)
	
	readWriteCloser.Close()
	
	// Verify composition works correctly
	if string(buffer) != string(data) {
		t.Errorf("Data should be preserved through interface composition")
	}
}