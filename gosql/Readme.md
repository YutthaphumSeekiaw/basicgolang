# Run postgres and pg admin
docker-compose up -d
docker-compose down

URL: http://localhost:5050
Username: pgadmin4@pgadmin.org (as a default)
Password: admin (as a default)

# config  ตาม  docker-compose.yml


sqlc	Compile-time SQL gen	เขียน SQL ตรง ๆ แล้วได้ Go struct & function
ent	ORM แบบ strongly typed	ใช้ Graph-like model, type-safe
gorm	ORM ที่นิยมสุดใน Go	ง่าย เร็ว ครอบคลุมทุกพื้นฐาน
database/sql	Standard library	ดิบ, เร็ว, ควบคุมเต็ม

# ใช้ Go ต่อ SQL Database

https://pkg.go.dev/database/sql

ลง package
package ที่เราจะใช้คือ database/sql ซึ่งเป็น standard library อยู่แล้ว
แต่ต้องลง driver เพื่อเป็นตัวแทนในการพูดคุยกับ PostgreSQL เพิ่ม
# go get github.com/lib/pq

# go get github.com/gofiber/fiber/v2
# go get github.com/lib/pq
