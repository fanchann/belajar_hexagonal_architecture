# `internal/port/inbound`

Direktori ini mendefinisikan *port inbound* (atau *driving ports*) dalam Hexagonal Architecture.

*Port inbound* adalah antarmuka (interfaces dalam Go) yang mendefinisikan bagaimana *core application* (khususnya *application services*) dapat dipanggil atau "didorong" oleh agen eksternal (seperti UI, controller HTTP/gRPC, atau test suites).

Mereka adalah kontrak API untuk *core application* dari perspektif dunia luar.

File `book_service.go` di sini akan mendefinisikan interface, misalnya `BookServicePort`, yang akan diimplementasikan oleh `internal/core/service/book_service_impl.go`. Adapter inbound (seperti adapter gRPC) akan memanggil metode pada interface ini.