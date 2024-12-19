
# AC-3 |  Fitur: Jual Barang (Keranjang)

Anggota Kelompok:
1. Cahya Fernando Vargas Saputra (1402023013)
2. Dimas Pradita Saputra (1402023019)
3. Fery Ardiansyah Djangkaru (1402023028)
4. Muhammad Adha Syah Putra (1402023034)
5. Yoga Pandu Navyanto (1402023073)
## Pembuat

- [@feryardnsyah04](https://github.com/feryardnsyah04)


## Features

- Masukan produk ke keranjang
- Melihat daftar produk di keranjang
- Update varian dan jumlah produk di keranjang
- Hapus produk di keranjang


## Dokumentasi API
1. Menambahkan Item ke Keranjang
- Endpoint: POST /cart
- Deskripsi: Menambahkan item ke keranjang belanja.
- Permintaan:
  - Header: Content-Type: application/json
  - Body:
```json
{
  "product": "Nama Produk",
  "variant": "Varian Produk",
  "price": 0,
  "quantity": 1
}
```
- Respon
  - Status: 201 Created
  - Body:
```json
{
  "message": "Barang ditambahkan ke keranjang",
  "item": {
    "id": 1,
    "product": "Nama Produk",
    "variant": "Varian Produk",
    "price": 0,
    "quantity": 1
  }
}
```
2. Mengambil Semua Item di Keranjang
- Endpoint: GET /cart
- Deskripsi: Mengambil semua item yang ada di keranjang belanja.
- Permintaan: Tidak ada
- Respon
  - Status: 200 OK
  - Body:
```json
[
  {
    "id": 1,
    "product": "Nama Produk",
    "variant": "Varian Produk",
    "price": 0,
    "quantity": 1
  },
  ...
]
```
3. Memperbarui Item di Keranjang
- Endpoint: PUT /cart
- Deskripsi: Memperbarui item yang ada di keranjang belanja.
- Permintaan:
  - Header: Content-Type: application/json
  - Body:
```json
{
  "id": 1,
  "variant": "Varian Produk",
  "quantity": 2
}
```
- Respon
  - Status: 200 OK
  - Body:
```json
{
  "message": "Item berhasil diperbarui",
  "item": {
    "id": 1,
    "product": "Nama Produk",
    "variant": "Varian Baru",
    "price": 0,
    "quantity": 2
  }
}
```
4. Menghapus Item dari Keranjang
- Endpoint: DELETE /cart
- Deskripsi: Memperbarui item yang ada di keranjang belanja.
- Permintaan:
  - Header: Content-Type: application/json
  - Body:
```json
{
  "id": 1
}
```
- Respon
  - Status: 200 OK
  - Body:
```json
{
  "message": "Item dihapus dari keranjang",
  "item": {
    "id": 1,
    "product": "Nama Produk",
    "variant": "Varian Produk",
    "price": 0,
    "quantity": 2
  }
}
```

