package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	tollbooth "github.com/didip/tollbooth/v7"
	tollbooth_gin "github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	csurf "github.com/tommy351/gin-csrf"
	sessions "github.com/tommy351/gin-sessions"
)

type ContactForm struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Message string `json:"message" binding:"required"`
}

// Konfigurasi email
const (
	SMTP_HOST      = "smtp.gmail.com"
	SMTP_PORT      = "587"
	ADMIN_EMAIL    = "admin@wilarno.co.id" // Ganti dengan email admin yang sebenarnya
	FROM_EMAIL     = "noreply@wilarno.co.id"
	EMAIL_PASSWORD = "your-app-password" // Ganti dengan app password Gmail
)

// Struktur untuk analytics
type AnalyticsData struct {
	Page      string    `json:"page"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	Referer   string    `json:"referer"`
	Timestamp time.Time `json:"timestamp"`
}

// Fungsi untuk mengirim email notification
func sendEmailNotification(form ContactForm, clientIP string) error {
	fmt.Printf("[DUMMY EMAIL] Simulasi kirim email ke admin:\nNama: %s\nEmail: %s\nPhone: %s\nPesan: %s\nIP: %s\n", form.Name, form.Email, form.Phone, form.Message, clientIP)
	return nil // Selalu sukses
}

// Fungsi untuk menyimpan analytics
func saveAnalytics(page, ip, userAgent, referer string) {
	data := AnalyticsData{
		Page:      page,
		IP:        ip,
		UserAgent: userAgent,
		Referer:   referer,
		Timestamp: time.Now(),
	}

	// Log ke file analytics
	f, err := os.OpenFile("analytics.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err == nil {
		logger := log.New(f, "", log.LstdFlags)
		logger.Printf("[%s] Page: %s | IP: %s | UA: %s | Referer: %s\n",
			data.Timestamp.Format(time.RFC3339), data.Page, data.IP, data.UserAgent, data.Referer)
		f.Close()
	}
}

func main() {
	r := gin.Default()

	// Middleware global untuk error handling agar response selalu JSON
	r.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(-1, gin.H{"success": false, "message": c.Errors[0].Error()})
		}
	})

	r.LoadHTMLGlob("templates/*.html")

	// Inisialisasi session store (gunakan secret yang sama untuk development)
	store := sessions.NewCookieStore([]byte("a-very-secret-key-32byteslong-123456"))
	r.Use(sessions.Middleware("my_session", store))

	// Rate limiter: 5 request per menit per IP
	limiter := tollbooth.NewLimiter(5.0/60.0, nil)

	// Middleware CSRF
	r.Use(csurf.Middleware(csurf.Options{
		Secret: "a-very-secret-key-32byteslong-123456",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	// Middleware untuk analytics tracking
	r.Use(func(c *gin.Context) {
		// Track analytics untuk halaman utama
		if c.Request.Method == "GET" && !strings.HasPrefix(c.Request.URL.Path, "/static") {
			go saveAnalytics(
				c.Request.URL.Path,
				c.ClientIP(),
				c.Request.UserAgent(),
				c.Request.Referer(),
			)
		}
		c.Next()
	})

	// Route alias ke file HTML
	r.GET("/wilarno-group", func(c *gin.Context) {
		c.File("templates/index.html")
	})
	r.GET("/tentang-kami", func(c *gin.Context) {
		c.File("templates/about.html")
	})
	r.GET("/bisnis-kami", func(c *gin.Context) {
		c.File("templates/services.html")
	})
	r.GET("/hubungi-kami", func(c *gin.Context) {
		token := csurf.GetToken(c)
		c.HTML(http.StatusOK, "contact.html", gin.H{
			"csrfToken": token,
		})
	})
	r.GET("/wilarno-autoparts", func(c *gin.Context) {
		c.File("templates/autoparts.html")
	})
	r.GET("/wilarno-building-industries", func(c *gin.Context) {
		c.File("templates/building.html")
	})
	r.GET("/wilarno-infrastructures", func(c *gin.Context) {
		c.File("templates/infrastructures.html")
	})
	r.GET("/wilarno-metal-industries", func(c *gin.Context) {
		c.File("templates/metal.html")
	})

	// Endpoint untuk analytics data (admin only)
	r.GET("/admin/analytics", func(c *gin.Context) {
		// Basic auth untuk admin (dalam production gunakan proper authentication)
		username, password, ok := c.Request.BasicAuth()
		if !ok || username != "admin" || password != "wilarno2024" {
			c.Header("WWW-Authenticate", `Basic realm="Admin Area"`)
			c.String(401, "Unauthorized")
			return
		}

		// Baca file analytics.log
		content, err := os.ReadFile("analytics.log")
		if err != nil {
			c.JSON(404, gin.H{"error": "Analytics data not found"})
			return
		}

		c.Header("Content-Type", "text/plain")
		c.String(200, string(content))
	})

	// Endpoint untuk contact data (admin only)
	r.GET("/admin/contacts", func(c *gin.Context) {
		// Basic auth untuk admin
		username, password, ok := c.Request.BasicAuth()
		if !ok || username != "admin" || password != "wilarno2024" {
			c.Header("WWW-Authenticate", `Basic realm="Admin Area"`)
			c.String(401, "Unauthorized")
			return
		}

		// Baca file contact.log
		content, err := os.ReadFile("contact.log")
		if err != nil {
			c.JSON(404, gin.H{"error": "Contact data not found"})
			return
		}

		c.Header("Content-Type", "text/plain")
		c.String(200, string(content))
	})

	// Route untuk admin dashboard
	r.GET("/admin", func(c *gin.Context) {
		// Basic auth untuk admin
		username, password, ok := c.Request.BasicAuth()
		if !ok || username != "admin" || password != "wilarno2024" {
			c.Header("WWW-Authenticate", `Basic realm="Admin Area"`)
			c.String(401, "Unauthorized")
			return
		}
		c.File("templates/admin-dashboard.html")
	})

	// Endpoint untuk menerima data kontak (form-urlencoded)
	r.POST("/contact", tollbooth_gin.LimitHandler(limiter), func(c *gin.Context) {
		var form ContactForm
		form.Name = c.PostForm("name")
		form.Email = c.PostForm("email")
		form.Phone = c.PostForm("phone")
		form.Message = c.PostForm("message")
		if form.Name == "" || form.Email == "" || form.Message == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Data tidak valid"})
			return
		}
		// Validasi nomor telepon: hanya angka, spasi, +, -, minimal 8 karakter
		phonePattern := regexp.MustCompile(`^[0-9+\-\s]{8,}$`)
		if form.Phone != "" && !phonePattern.MatchString(form.Phone) {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Nomor telepon tidak valid"})
			return
		}
		// Validasi pesan: tidak kosong, tidak hanya spasi
		if strings.TrimSpace(form.Message) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Pesan tidak boleh kosong"})
			return
		}
		// Sanitasi input: trim dan escape karakter HTML
		form.Name = html.EscapeString(strings.TrimSpace(form.Name))
		form.Email = html.EscapeString(strings.TrimSpace(form.Email))
		form.Phone = html.EscapeString(strings.TrimSpace(form.Phone))
		form.Message = html.EscapeString(strings.TrimSpace(form.Message))

		// Logging ke file contact.log dengan format rapi dan aman
		logMsg := form.Message
		if len(logMsg) > 500 {
			logMsg = logMsg[:500] + "... (truncated)"
		}
		clientIP := c.ClientIP()
		f, err := os.OpenFile("contact.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err == nil {
			logger := log.New(f, "", log.LstdFlags)
			logger.Printf("[%s] IP: %s | Nama: %s | Email: %s | Pesan: %s\n", time.Now().Format(time.RFC3339), clientIP, form.Name, form.Email, logMsg)
			f.Close()
		}

		// Kirim email notification
		err = sendEmailNotification(form, clientIP)
		if err != nil {
			log.Printf("Gagal mengirim email notification: %+v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mengirim email notification"})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Pesan berhasil diterima!"})
	})

	// Serve file statis (CSS, JS, gambar, dll)
	r.Static("/static", "./static")

	// Fallback untuk file statis di root (misal: favicon, gambar)
	r.NoRoute(func(c *gin.Context) {
		path := "." + c.Request.URL.Path
		c.File(path)
	})

	r.Run(":8080")
}
