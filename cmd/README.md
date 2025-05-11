# `cmd`

Direktori ini berisi titik masuk (entry points) utama untuk aplikasi atau layanan yang dapat dieksekusi (executable).
Setiap subdirektori atau file `.go` di sini biasanya merupakan `main` package yang bertanggung jawab untuk:
- Inisialisasi konfigurasi.
- Pengaturan dependensi (dependency injection).
- Memulai server (misalnya, HTTP, gRPC) atau worker.