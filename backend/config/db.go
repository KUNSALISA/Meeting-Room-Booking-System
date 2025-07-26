package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/KUNSALISA/Meeting-Room-Booking-System/entity"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

///////////////////////////////////////////////////////

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

// /////////////////////////////////////////////////////
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbName   = "meeting_room_booking"
	timezone = "Asia/Bangkok"
	sslMode  = "disable"
)

func CreateDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=%s", host, user, password, port, sslMode)
	dbSQL, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	defer dbSQL.Close()

	_, err = dbSQL.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		fmt.Println("Database may already exist:", err)
	} else {
		fmt.Println("Database created successfully")
	}
}

func ConnectionDB() {
	CreateDatabase()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, dbName, port, sslMode, timezone)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connected to database")
	db = database
}

///////////////////////////////////////////////////////

func SetupDatabase() {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Type{},
		&entity.Room{},
		&entity.Booking{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	SeedRoles()
	SeedTypes()
	SeedStatus()
	SeedRooms()
	SeedUsers()
	SeedBookings()
}

func SeedRoles() {
	roles := []entity.Role{
		{RoleName: "Admin"},
		{RoleName: "User"},
	}

	for _, r := range roles {
		var existing entity.Role
		db.Where("role_name = ?", r.RoleName).First(&existing)
	}
}

func SeedStatus() {
	status := []entity.Status{
		{StatusName: "ถูกจองแล้ว"},
		{StatusName: "ห้องว่าง"},
		{StatusName: "เสียหาย"},
	}

	for _, s := range status {
		var existing entity.Status
		db.Where("status_name = ?", s.StatusName).First(&existing)
	}
}

func SeedTypes() {
	types := []entity.Type{
		{TypeName: "VIP"}, // Capacity >= 50
		{TypeName: "ขนาดใหญ่"}, // Capacity >= 50
		{TypeName: "ขนาดกลาง"}, // Capacity 10–49
		{TypeName: "ขนาดเล็ก"}, // Capacity < 10
	}

	for _, t := range types {
		db.FirstOrCreate(&t, entity.Type{TypeName: t.TypeName})
	}
}

func SeedRooms() {

	var types []entity.Type
	db.Find(&types)

	typeMap := make(map[string]uint)
	for _, t := range types {
		typeMap[t.TypeName] = t.ID
	}

	rooms := []entity.Room{
		{
			RoomName:  "Meeting Room Extra 1",
			Location:  "ชั้น 12 อาคาร TX",
			Capacity:  80,
			Equipment: pq.StringArray{"โปรเจคเตอร์", "ไวท์บอร์ด", "ไมค์ประชุม", "ลำโพง", "HDMI", "ระบบเก็บเสียง", "เลเซอร์พอยเตอร์", "ระบบวิดีโอคอนเฟอเรนซ์", "อินเทอร์เน็ต"},
			TypeID:    typeMap["VIP"],
		},
		{
			RoomName:  "Meeting Room Extra 2",
			Location:  "ชั้น 12 อาคาร TX",
			Capacity:  60,
			Equipment: pq.StringArray{"โปรเจคเตอร์", "ไวท์บอร์ด", "ไมค์ประชุม", "ลำโพง", "HDMI", "ระบบเก็บเสียง", "เลเซอร์พอยเตอร์", "ระบบวิดีโอคอนเฟอเรนซ์", "อินเทอร์เน็ต"},
			TypeID:    typeMap["VIP"],
		},
		{
			RoomName:  "Meeting Room A1",
			Location:  "ชั้น 1 อาคาร TX",
			Capacity:  100,
			Equipment: pq.StringArray{"โปรเจคเตอร์", "ไวท์บอร์ด", "ไมค์ประชุม", "ลำโพง", "HDMI", "ระบบเก็บเสียง", "เลเซอร์พอยเตอร์", "อินเทอร์เน็ต"},
			TypeID:    typeMap["ขนาดใหญ่"],
		},
		{
			RoomName:  "Meeting Room A10",
			Location:  "ชั้น 1 อาคาร TX",
			Capacity:  10,
			Equipment: pq.StringArray{"โปรเจคเตอร์", "ไวท์บอร์ด", "ไมค์ประชุม", "ลำโพง", "HDMI", "ระบบเก็บเสียง", "เลเซอร์พอยเตอร์", "อินเทอร์เน็ต"},
			TypeID:    typeMap["ขนาดเล็ก"],
		},
		{
			RoomName:  "Meeting Room B1",
			Location:  "ชั้น 2 อาคาร TX",
			Capacity:  20,
			Equipment: pq.StringArray{"โปรเจคเตอร์", "ไวท์บอร์ด", "ไมค์ประชุม", "ลำโพง", "HDMI", "ระบบเก็บเสียง", "เลเซอร์พอยเตอร์", "อินเทอร์เน็ต"},
			TypeID:    typeMap["ขนาดกลาง"],
		},
		{
			RoomName:  "Meeting Room B2",
			Location:  "ชั้น 2 อาคาร TX",
			Capacity:  30,
			Equipment: pq.StringArray{"โปรเจคเตอร์", "ไวท์บอร์ด", "ไมค์ประชุม", "ลำโพง", "HDMI", "ระบบเก็บเสียง", "เลเซอร์พอยเตอร์", "อินเทอร์เน็ต"},
			TypeID:    typeMap["ขนาดกลาง"],
		},
		{
			RoomName:  "Meeting Room C1",
			Location:  "ชั้น 3 อาคาร TX",
			Capacity:  8,
			Equipment: pq.StringArray{"โปรเจคเตอร์", "ไวท์บอร์ด", "ไมค์ประชุม", "ลำโพง", "HDMI", "ระบบเก็บเสียง", "เลเซอร์พอยเตอร์", "อินเทอร์เน็ต"},
			TypeID:    typeMap["ขนาดเล็ก"],
		},
	}

	for _, r := range rooms {
		var existing entity.Room
		err := db.Where("room_name = ?", r.RoomName).First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&r).Error; err != nil {
				fmt.Printf("❌ เพิ่มห้อง '%s' ไม่สำเร็จ: %v\n", r.RoomName, err)
			} else {
				fmt.Printf("✅ เพิ่มห้อง '%s' สำเร็จ\n", r.RoomName)
			}
		} else {
			fmt.Printf("⚠️ ห้อง '%s' มีอยู่แล้ว, ข้าม...\n", r.RoomName)
		}
	}
}

func SeedUsers() {

	pass_admin, _ := HashPassword("admin")
	pass_user, _ := HashPassword("1234")

	users := []entity.User{
		{
			CodeName:    "admin",
			Password:    pass_admin,
			Firstname:   "Manager",
			Lastname:    "Booking",
			Image:       "",
			Email:       "adminbooking@gmail.com",
			PhoneNumber: "0442569874",
			RoleID:      1,
		},
		{
			CodeName:    "A650000",
			Password:    pass_user,
			Firstname:   "Alice",
			Lastname:    "Iron",
			Image:       "",
			Email:       "Alice_Iron@gmail.com",
			PhoneNumber: "0807654321",
			RoleID:      2,
		},
		{
			CodeName:    "A650001",
			Password:    pass_user,
			Firstname:   "Bobby",
			Lastname:    "Kim",
			Image:       "",
			Email:       "Bobby_Kim@gmail.com",
			PhoneNumber: "0864569871",
			RoleID:      2,
		},
		{
			CodeName:    "A650002",
			Password:    pass_user,
			Firstname:   "Niki",
			Lastname:    "Kayoko",
			Image:       "",
			Email:       "Niki_Kayoko@gmail.com",
			PhoneNumber: "0852316549",
			RoleID:      2,
		},
	}

	for _, u := range users {
		var existing entity.User
		err := db.Where("email = ? OR code_name = ?", u.Email, u.CodeName).First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&u).Error; err != nil {
				fmt.Printf("❌ สร้าง %s ไม่สำเร็จ: %v\n", u.CodeName, err)
			} else {
				fmt.Printf("✅ เพิ่มผู้ใช้ %s สำเร็จ\n", u.CodeName)
			}
		} else {
			fmt.Printf("⚠️ ผู้ใช้ %s มีอยู่แล้ว, ข้าม...\n", u.CodeName)
		}
	}
}

func SeedBookings() {
	var user entity.User
	var room entity.Room
	var status entity.Status

	db.Where("email = ?", "Alice_Iron@gmail.com").First(&user)
	db.Where("room_name = ?", "Meeting Room A1").First(&room)
	db.Where("status_name = ?", "ถูกจองแล้ว").First(&status)

	date := time.Date(2025, 7, 15, 0, 0, 0, 0, time.Local)
	start := time.Date(2025, 7, 15, 9, 0, 0, 0, time.Local)
	end := time.Date(2025, 7, 15, 11, 0, 0, 0, time.Local)

	var existing entity.Booking
	err := db.Where("date = ? AND room_id = ? AND start_time = ? AND end_time = ?", date, room.ID, start, end).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		booking := entity.Booking{
			Title:       "ประชุมทีมพัฒนา",
			Description: "สรุปงานไตรมาสและวางแผนระบบใหม่",
			Date:        date,
			StartTime:   start,
			EndTime:     end,
			StatusID:    status.ID,
			UserID:      user.ID,
			RoomID:      room.ID,
		}

		if err := db.Create(&booking).Error; err != nil {
			fmt.Println("❌ สร้างการจองไม่สำเร็จ:", err)
		} else {
			fmt.Println("✅ สร้างการจองสำเร็จ:", booking.Title)
		}
	} else {
		fmt.Println("⚠️ มีการจองในช่วงเวลานี้แล้ว")
	}
}
