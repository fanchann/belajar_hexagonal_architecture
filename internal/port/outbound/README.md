# `internal/port/outbound`

Direktori ini mendefinisikan *port outbound* (atau *driven ports*) dalam Hexagonal Architecture.

*Port outbound* adalah antarmuka (interfaces dalam Go) yang dibutuhkan oleh *core application* untuk berinteraksi dengan layanan atau sistem eksternal (seperti database, message broker, API pihak ketiga, dll.).

*Core application* mendefinisikan *port* ini berdasarkan kebutuhannya, dan *adapter outbound* (yang berada di luar *core*) akan menyediakan implementasi konkret untuk *port* tersebut. Ini memastikan *core application* tetap independen dari detail implementasi infrastruktur.

File `book_repo.go` di sini akan mendefinisikan interface, misalnya `BookRepositoryPort`, yang menentukan operasi persistensi data apa saja yang dibutuhkan oleh *core service*. Adapter storage (misalnya, `internal/adapter/storage/book_repo_impl.go`) akan mengimplementasikan interface ini.