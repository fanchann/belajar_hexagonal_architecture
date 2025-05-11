# `internal/core/service`

Direktori ini berisi layanan aplikasi (application services) atau *use cases*. Ini adalah bagian dari *core application* dalam Hexagonal Architecture.

Tanggung jawab utama layanan aplikasi meliputi:
- Mengorkestrasi aliran kerja untuk *use case* tertentu.
- Menggunakan entitas domain dan logika domain untuk melakukan pekerjaan.
- Berinteraksi dengan *port outbound* (misalnya, repository interfaces) untuk mengambil atau menyimpan data.
- Mengimplementasikan *port inbound* yang mendefinisikan bagaimana dunia luar dapat berinteraksi dengan *use case* ini.

File `book_service_impl.go` akan berisi logika bisnis tingkat tinggi untuk operasi terkait buku, seperti membuat buku baru, mencari buku, dll., dengan memanfaatkan entitas dari `internal/core/domain` dan berinteraksi dengan *repository* melalui *port outbound*.