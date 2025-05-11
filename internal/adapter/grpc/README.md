# `internal/adapter/grpc`

Direktori ini berisi implementasi adapter gRPC.

Dalam Hexagonal Architecture, adapter adalah komponen yang menghubungkan *core application* dengan dunia luar. Adapter gRPC ini berfungsi sebagai:
- **Adapter Inbound (Driving Adapter):** Menerima permintaan dari klien melalui protokol gRPC.
- Menerjemahkan permintaan gRPC menjadi panggilan ke *port inbound* yang didefinisikan oleh *core application*.
- Mengonversi hasil dari *core application* kembali menjadi respons gRPC.

File `grpc_impl.go` di sini akan berisi logika implementasi dari layanan gRPC yang didefinisikan dalam file `.proto`, dan akan berinteraksi dengan *service* di dalam `internal/core` melalui *port inbound*.