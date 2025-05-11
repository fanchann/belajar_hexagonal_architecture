# `internal/core/domain`

Direktori ini adalah jantung dari aplikasi Anda dan merupakan bagian inti dari "hexagon" dalam Hexagonal Architecture. Isinya adalah representasi data inti.

Direktori ini biasanya mencakup:
- **`entity/`**: Berisi definisi entitas domain. Entitas adalah objek inti dalam bisnis Anda dengan identitas yang unik dan siklus hidup (misalnya, `Book`). Mereka tidak bergantung pada framework atau infrastruktur.
- **`dto/` (Data Transfer Objects)**: Berisi struktur data sederhana yang digunakan untuk mentransfer data antara lapisan, terutama antara *adapter* dan *service* di *core*. DTO membantu menjaga *entity* tetap bersih dari detail presentasi atau permintaan.
    - `book_converter.go`: berisi fungsi utilitas untuk mengonversi antara entitas domain dan DTO,
    - `book_request.go` & `book_response.go`: Contoh DTO untuk data permintaan dan respons spesifik untuk operasi buku.
