# `internal/config`

Direktori ini bertanggung jawab untuk mengelola semua konfigurasi aplikasi.

Tugas utamanya meliputi:
- Mendefinisikan struktur data untuk konfigurasi (misalnya, koneksi database, alamat server, kunci API, dll.).
- Memuat konfigurasi dari berbagai sumber seperti file (misalnya, TOML, YAML, JSON), variabel lingkungan (environment variables), atau flag baris perintah.
- Menyediakan akses yang mudah dan aman ke nilai konfigurasi untuk seluruh bagian aplikasi.

File-file seperti `config.go`, `mariadb.go`, dan `viper.go` akan berisi logika untuk mem-parsing dan menyediakan data konfigurasi tersebut.