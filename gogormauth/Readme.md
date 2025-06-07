# GORM คือ ORM Lib ของ go

Object-Relational Mapping (ORM) คือ programming technique ที่ใช้สำหรับการเปลี่ยน data จากรูปแบบใดๆก็ตามที่อยู่ใน relational database (table, collection) มาสู่รูปแบบ object ของภาษา programming เพื่อสะดวกต่อการจัดการ

ข้อดีของพวก ORM คือ

Abstraction ทุกคนในทีมสามารถทำงานผ่าน Object แทนได้ (แทนที่จะต้องมาปวดหัวกับ SQL แทน) ซึ่งจะช่วยทำให้ code อ่านง่ายขึ้นมาก
Database Agnostic สามารถเปลี่ยนไปมาระหว่าง SQL Database ในหลายๆประเภทได้ (ตามที่ ORM support) โดยการปรับเพียงแค่ config เล็กน้อย
Maintainability ง่ายต่อการอ่าน code มากขึ้น ทำให้ maintain และเข้าใจได้ง่ายขึ้น
Security ORM ทุกตัวส่วนใหญ่จะทำการเพิ่มการป้องกันผ่านการโจมตี SQL injection ไว้แล้ว ทำให้เราไม่ต้องกังวลกับเรื่องนี้s

รู้จักกับ GORM
ref: https://gorm.io/index.html

GORM คือ ORM library ของภาษา GO ที่จะทำการ convert table ใน relational database มาสู่ Object ฉบับ Go โดยมี feature ตั้งแต่

Full-Featured ORM
Associations (has one, has many, belongs to, many to many, polymorphism, single-table inheritance)
Hooks (before/after create/save/update/delete/find)
Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
Batch Insert, FindInBatches, Find/Create with Map, CRUD with SQL Expr and Context Valuer
SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
Composite Primary Key, Indexes, Constraints
Auto Migrations

# go get -u gorm.io/gorm
# go get -u gorm.io/driver/postgres


ใช้งานร่วมกับ Fiber

# go get github.com/gofiber/fiber/v2

Middleware กับ User
Note

ทำการเพิ่ม Table User ขึ้นมา ผ่าน Gorm (เก็บ email unique, password แบบ hash)
Login แล้วเก็บ JWT Token ไว้

# go get golang.org/x/crypto/bcrypt
# go get github.com/golang-jwt/jwt/v4