# `internal/adapter/storage`

Direktori ini berisi implementasi adapter untuk persistensi data (penyimpanan).

Sebagai **Adapter Outbound (Driven Adapter)** dalam Hexagonal Architecture, adapter ini bertanggung jawab untuk:
- Mengimplementasikan *port outbound* (misalnya, interface repository) yang didefinisikan oleh *core application*.
- Berinteraksi dengan sistem penyimpanan data aktual (misalnya, database MariaDB, PostgreSQL, MongoDB, dll.).
- Menerjemahkan data antara format yang digunakan oleh *core application* (entitas domain) dan format yang dibutuhkan oleh sistem penyimpanan.

File `book_repo_impl.go` akan berisi implementasi konkret dari operasi data (CRUD) untuk entitas buku, sesuai dengan kontrak yang ditetapkan oleh *port outbound* (misalnya, `BookRepository` interface).