<!-- @format -->

# **REST API Online Store**

Repository ini adalah solusi untuk test challenge **Backend Engineer Intern PT.Synapsis Indonesia**. REST API

---

## **✨ Fitur Utama**

- Pelanggan dapat melihat daftar produk berdasarkan kategori.
- Pelanggan dapat menambahkan produk ke keranjang belanja.
- Pelanggan dapat melihat daftar produk di keranjang belanja.
- Pelanggan dapat menghapus produk dari keranjang belanja.
- Pelanggan dapat melakukan checkout dan pembayaran.
- Fitur login dan register untuk pelanggan.

---

## **🚀 Teknologi yang Digunakan**

- **Golang**
- **Echo Go Framework**
- **SQLC**
- **PostgreSQL**
- **Docker**
- **Docker Compose**
- **Midtrans Sandbox**

---

## **📦 Cara Menjalankan Proyek**

### **1. Jalankan dengan Docker Compose**

1. Clone repository ini:
   ```bash
   git clone https://github.com/claireglowup/ecommerce-synapsis.git
   cd ecommerce-synapsis
   ```
2. Bangun dan jalankan container:
   ```bash
   docker-compose up --build
   ```

### **2. Alternatif: Jalankan dengan Makefile**

Jika Makefile tersedia, jalankan:

```bash
make start
```

### **3. Alternatif: Tarik Image dari Docker Hub**

Jika tidak ingin build secara lokal, gunakan image yang sudah tersedia:

```bash
docker pull rikyfahrian1/synapsis-backendintern
docker run -p 8080:8080 rikyfahrian1/synapsis-backendintern
```

### \*\*4. Tersedia Open Api Di swagger.json

## **📜 Lisensi**

Proyek ini menggunakan lisensi [MIT License](LICENSE).

---

## **🤝 Kontribusi**

1. Fork repository ini.
2. Buat branch baru:
   ```bash
   git checkout -b fitur-baru
   ```
3. Commit perubahan:
   ```bash
   git commit -m "Menambahkan fitur baru"
   ```
4. Push ke branch:
   ```bash
   git push origin fitur-baru
   ```
5. Buat Pull Request.

---

Semoga bermanfaat, dan selamat mencoba! 🚀
