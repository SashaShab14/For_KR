package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var requestCount uint64

func handler(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&requestCount, 1)
	currentCount := atomic.LoadUint64(&requestCount)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	html := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>DevOps КР</title>
		<style>
			body { font-family: 'Segoe UI', sans-serif; background-color: #0b132b; color: #ffffff; margin: 0; padding: 40px; display: flex; justify-content: center; align-items: center; min-height: 80vh; }
			.container { background: #1c2541; padding: 35px; border-radius: 12px; box-shadow: 0 8px 30px rgba(0, 0, 0, 0.5); max-width: 650px; width: 100%%; border-top: 6px solid #06d6a0; }
			h1 { color: #06d6a0; font-size: 24px; border-bottom: 2px solid #3a506b; padding-bottom: 12px; margin-top: 0; }
			p { font-size: 14px; color: #94a3b8; line-height: 1.6; }
			.status-badge { color: #ffffff; background-color: #06d6a0; padding: 4px 10px; border-radius: 20px; font-size: 12px; font-weight: bold; text-transform: uppercase; }
			.metrics-table { width: 100%%; border-collapse: collapse; margin-top: 20px; }
			.metrics-table td { padding: 12px; border-bottom: 1px solid #3a506b; font-size: 14px; text-align: left; }
			.label { color: #94a3b8; font-weight: 600; }
			.value { font-family: 'Courier New', monospace; background: #0f172a; color: #38bdf8; padding: 3px 8px; border-radius: 4px; }
			.footer { margin-top: 35px; font-size: 11px; color: #64748b; text-align: center; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Система мониторинга Go Runtime</h1>
			<p>Инфраструктура развернута автоматически в среде WSL2 через локальный Self-Hosted раннер GitHub Actions по стратегии GitHub Flow.</p>
			<table class="metrics-table">
				<tr><td class="label">Состояние рантайма</td><td><span class="status-badge">ONLINE</span></td></tr>
				<tr><td class="label">Текущее время сервера (ISO)</td><td><span class="value">%s</span></td></tr>
				<tr><td class="label">Всего обработано HTTP-запросов</td><td><span class="value">%d</span></td></tr>
			</table>
			</div>
	</body>
	</html>`, currentTime, currentCount)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running demo app on port :8888. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
