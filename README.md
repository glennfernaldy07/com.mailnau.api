# com.mailnau.api
Project Absensi

#Migration 
run goose -dir migration/ mysql "root:root@tcp(localhost:3306)/{{db_name}}?parseTime=true" up 