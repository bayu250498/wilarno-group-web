# Wilarno Group Website

Website resmi Wilarno Group dengan fitur-fitur modern dan responsif.

## 🚀 Fitur Utama

### 📧 Contact Form dengan Email Notification
- Form kontak yang aman dengan validasi
- CSRF protection untuk keamanan
- Email notification otomatis ke admin
- Popup sukses dengan auto-refresh halaman
- Rate limiting untuk mencegah spam

### 📊 Analytics & Admin Dashboard
- Tracking kunjungan website secara real-time
- Admin dashboard dengan statistik lengkap
- Basic authentication untuk keamanan admin
- Export data analytics dan contact

### 🎨 UI/UX Modern
- Design responsif untuk semua device
- Background images yang menarik
- Animasi dan transisi yang smooth
- Bootstrap 5 untuk komponen modern

## 🛠️ Teknologi yang Digunakan

- **Backend**: Go (Gin framework)
- **Frontend**: HTML5, CSS3, JavaScript
- **UI Framework**: Bootstrap 5
- **Email**: SMTP (Gmail)
- **Security**: CSRF protection, Rate limiting

## 📁 Struktur File

# Struktur Modular Baru

```
Web Sederhana Cursor/
├── static/
│   ├── css/
│   │   └── styles.css
│   ├── img/
│   │   ├── office.jpg
│   │   └── technology.jpg
│   └── js/
│       └── (jika ada custom JS)
├── templates/
│   ├── index.html
│   ├── about.html
│   ├── services.html
│   ├── contact.html
│   ├── admin-dashboard.html
│   ├── autoparts.html
│   ├── building.html
│   ├── infrastructures.html
│   ├── metal.html
│   └── (halaman lain)
├── logs/
│   ├── contact.log
│   └── analytics.log
├── main.go
├── go.mod
├── go.sum
├── README.md
├── robots.txt
└── sitemap.xml
```

## 🚀 Cara Menjalankan

### Prerequisites
- Go 1.19 atau lebih baru
- Gmail account untuk email notification

### Setup Email (Opsional)
1. Edit file `main.go`
2. Ganti konfigurasi email:
   ```go
   const (
       ADMIN_EMAIL   = "your-admin@email.com"
       FROM_EMAIL    = "your-noreply@email.com"
       EMAIL_PASSWORD = "your-gmail-app-password"
   )
   ```

### Menjalankan Server
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## 🔐 Admin Access

### Dashboard Admin
- URL: `http://localhost:8080/admin`
- Username: `admin`
- Password: `wilarno2024`

### Fitur Admin
- 📈 Analytics data real-time
- 📧 Contact messages
- 📊 Statistik kunjungan
- 🔄 Auto-refresh data

## 📧 Email Notification

Ketika ada pesan kontak masuk, admin akan menerima email dengan format:
- **Subject**: "Pesan Baru dari Website Wilarno Group"
- **Content**: HTML template dengan data lengkap
- **Info**: Nama, email, telepon, pesan, IP, waktu

## 🔒 Security Features

- **CSRF Protection**: Token untuk mencegah CSRF attack
- **Rate Limiting**: 5 request per menit per IP
- **Input Sanitization**: Escape HTML characters
- **Basic Auth**: Untuk admin dashboard
- **Validation**: Email, phone, message validation

## 📊 Analytics Tracking

Website mencatat:
- Halaman yang dikunjungi
- IP address pengunjung
- User agent browser
- Referer URL
- Timestamp kunjungan

## 🎯 Unit Bisnis

1. **Wilarno Autoparts** - `/wilarno-autoparts`
2. **Wilarno Building Industries** - `/wilarno-building-industries`
3. **Wilarno Infrastructures** - `/wilarno-infrastructures`
4. **Wilarno Metal Industries** - `/wilarno-metal-industries`

## 🔧 Customization

### Mengubah Warna
Edit file `styles.css`:
```css
:root {
    --primary-color: #1a2233;
    --accent-color: #ffcc00;
    --text-color: #333;
}
```

### Menambah Halaman Baru
1. Buat file HTML baru
2. Tambahkan route di `main.go`
3. Update navigation di semua halaman

## 📝 Log Files

- `contact.log`: Semua pesan kontak yang masuk
- `analytics.log`: Data kunjungan website

## 🚀 Deployment

### Production Setup
1. Set environment variables untuk email
2. Gunakan proper SSL certificate
3. Setup proper authentication untuk admin
4. Configure reverse proxy (nginx/apache)
5. Setup monitoring dan logging

### Docker (Opsional)
```dockerfile
FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
```

## 🤝 Contributing

1. Fork repository
2. Buat feature branch
3. Commit changes
4. Push ke branch
5. Buat Pull Request

## 📄 License

© 2024 Wilarno Group. All rights reserved.

## 📞 Support

Untuk bantuan teknis atau pertanyaan:
- Email: info@wilarno.co.id
- Website: http://localhost:8080/hubungi-kami

---

**Wilarno Group** - Solusi Bisnis Terpercaya 