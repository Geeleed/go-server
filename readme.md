# Go Server

เอาไว้ใช้เปิด server

1. โหลด Golang จาก https://go.dev/doc/install แล้วติดตั้ง
2. เปิด cmd ใช้คำสั่ง go version ถ้าไม่เกิดข้อผิดพลาดแสดงว่าติดตั้งสำเร็จแล้ว
3. build โดยใช้คำสั่ง go build go-server.go จะได้ไฟล์ go-server.exe (สำหรับ windows)
4. รันไฟล์ go-server.exe เพื่อใช้งาน ซึ่งต้องกรอกข้อมูล 3 อย่าง
   - directory ของโฟลเดอร์ที่จะเอาขึ้น server
   - hostname เช่น localhost
   - port เช่น 3000

# วิธีอื่น ๆ ในการเปิด server อย่างง่าย ๆ

เปิด cmd/terminal ไปที่ directory ที่ต้องการนำขึ้น server แล้ว...

## ถ้าใช้ node.js

```bash
npm install -g http-server
```

```bash
http-server -p 3000
```

## ถ้าใช้ python

```bash
python -m http.server 3000
```

python จะเป็น synchronous ซึ่งทำงานทีละ 1 request มันจะไม่ค่อยเหมาะกับการรับ request จำนวนเยอะ ๆ แต่ node.js จะเป็น asynchronous ซึ่งรับหลาย request พร้อมกันได้ ส่วน go จะใช้ Goroutines ซึ่งทำงานหลาย request พร้อมกันได้เช่นกัน