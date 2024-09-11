# Langkah-langkah menambahkan dependecy

## Tambahkan Dependensi (Opsional)
- Jika ingin menambahkan dependensi dari luar, misalnya paket dari GitHub, gunakan perintah `go get`. Misalnya, untuk menambahkan paket `google/uuid`:
``` bash
go get github.com/google/uuid
```

- Jika ada perubahan atau pembaruan dependensi, Anda bisa menggunakan perintah berikut untuk memperbarui file `go.mod` dan `go.sum`:
``` bash
go mod tidy
go get github.com/google/uuid@v1.2.0
```