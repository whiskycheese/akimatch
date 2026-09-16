package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env ファイルが見つからないため環境変数を直接参照します")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("エラー: DATABASE_URL が設定されていません")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("DB接続設定エラー: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("DB接続失敗（Pingエラー）: %v", err)
	}

	fmt.Println("🌱 CSVファイルからのシードデータ投入を開始します...")

	// 1. 大学マスタ (universities.csv)
	loadUniversities(db, "seed_data/universities.csv")

	// 2. 学部マスタ (faculties.csv)
	loadFaculties(db, "seed_data/faculties.csv")

	// 3. 授業マスタ (courses.csv)
	loadCourses(db, "seed_data/courses.csv")

	// 4. テストユーザーの生成
	loadDummyUsers(db)

	fmt.Println("\n✨ すべてのCSVデータの投入が正常に完了しました！")
}

func loadUniversities(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("大学CSVのオープン失敗: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // ヘッダー行をスキップ

	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("大学CSVの読み込みエラー: %v", err)
		}

		id, _ := strconv.Atoi(record[0])
		name := record[1]
		maxPeriod, _ := strconv.Atoi(record[2])

		_, err = db.Exec(
			"INSERT INTO universities (id, name, max_period) VALUES ($1, $2, $3)",
			id, name, maxPeriod,
		)
		if err != nil {
			log.Fatalf("大学データの挿入エラー (%s): %v", name, err)
		}
		count++
	}
	// PostgreSQLのPRIMARY KEYシーケンスを更新（IDを明示指定してINSERTしたため）
	_, _ = db.Exec("SELECT setval('universities_id_seq', (SELECT MAX(id) FROM universities))")
	fmt.Printf("🏢 大学マスタ: %d 件投入\n", count)
}

func loadFaculties(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("学部CSVのオープン失敗: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // ヘッダー行をスキップ

	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("学部CSVの読み込みエラー: %v", err)
		}

		id, _ := strconv.Atoi(record[0])
		universityID, _ := strconv.Atoi(record[1])
		name := record[2]

		_, err = db.Exec(
			"INSERT INTO faculties (id, university_id, name) VALUES ($1, $2, $3)",
			id, universityID, name,
		)
		if err != nil {
			log.Fatalf("学部データの挿入エラー (%s): %v", name, err)
		}
		count++
	}
	_, _ = db.Exec("SELECT setval('faculties_id_seq', (SELECT MAX(id) FROM faculties))")
	fmt.Printf("🎓 学部マスタ: %d 件投入\n", count)
}

func loadCourses(db *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("授業CSVのオープン失敗: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // ヘッダー行をスキップ

	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("授業CSVの読み込みエラー: %v", err)
		}

		universityID, _ := strconv.Atoi(record[0])
		var facultyID *int
		if record[1] != "" {
			fid, _ := strconv.Atoi(record[1])
			facultyID = &fid
		}
		subjectName := record[2]
		room := record[3]

		_, err = db.Exec(
			"INSERT INTO courses (university_id, faculty_id, subject_name, room) VALUES ($1, $2, $3, $4)",
			universityID, facultyID, subjectName, room,
		)
		if err != nil {
			log.Fatalf("授業データの挿入エラー (%s): %v", subjectName, err)
		}
		count++
	}
	fmt.Printf("📚 授業マスタ: %d 件投入\n", count)
}

func loadDummyUsers(db *sql.DB) {
	users := []struct {
		Name         string
		ShareCode    string
		UniversityID int
		FacultyID    int
	}{
		{"テスト太郎", "TARO0001", 1, 5}, // 千葉大 工学部
		{"テスト花子", "HANA0002", 1, 1}, // 千葉大 国際教養学部
		{"千葉次郎", "JIRO0003", 2, 13}, // 千葉工大 情報変革科学部
	}

	for _, u := range users {
		_, err := db.Exec(
			"INSERT INTO users (name, share_code, university_id, faculty_id) VALUES ($1, $2, $3, $4)",
			u.Name, u.ShareCode, u.UniversityID, u.FacultyID,
		)
		if err != nil {
			log.Fatalf("テストユーザー挿入エラー: %v", err)
		}
	}
	fmt.Printf("👤 テストユーザー: %d 件投入\n", len(users))
}
