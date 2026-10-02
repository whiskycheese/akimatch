package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"backend/handler"
	"backend/repository"
	"backend/service"
)

func main() {
	// 1. .env ファイルの読み込み（ローカル開発時のみ）
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	// 2. 環境変数から接続情報を取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// 3. 接続文字列（DSN）の組み立て
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode,
	)

	// 4. データベース接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL!")

	// 6. 依存関係の注入（DI）
	univRepo := repository.NewUniversityRepository(db)
	univSvc := service.NewUniversityService(univRepo)
	univHandler := handler.NewUniversityHandler(univSvc)

	facultyRepo := repository.NewFacultyRepository(db)
  facultySvc := service.NewFacultyService(facultyRepo)
  facultyHandler := handler.NewFacultyHandler(facultySvc)

	// 7. ルーティング設定
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK (DB Connected)")
	})

	// 8.API のエンドポイントを登録
	http.HandleFunc("/api/universities", univHandler.GetAllUniversities)
	http.HandleFunc("GET /api/universities/{id}/faculties", facultyHandler.GetFacultiesByUniversityID)

	// 9. サーバー起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
