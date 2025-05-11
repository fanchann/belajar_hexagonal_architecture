# `internal/pkg`

Direktori ini berisi paket-paket utilitas internal (internal shared libraries) yang dapat digunakan kembali di berbagai bagian dalam proyek ini.

Tujuannya adalah untuk menampung kode yang:
- Bersifat generik dan tidak spesifik untuk satu domain bisnis tertentu.
- Tidak memiliki ketergantungan pada *core application* atau *adapter*.
- Dapat membantu menghindari duplikasi kode.

Contohnya bisa berupa utilitas untuk logging, pembuatan ID unik (seperti `uid/uid.go`), validasi umum, atau manipulasi data generik.

**NB:** Kode di `internal/pkg` sebaiknya tidak mengimpor dari direktori `internal/core`, `internal/adapter`, atau `internal/port` untuk menjaga independensinya. Sebaliknya, direktori-direktori tersebut dapat menggunakan paket dari `internal/pkg`.