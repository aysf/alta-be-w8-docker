# Learning Docker

## Mengecek image yang terdapat pada docker

### `docker images`

## Menghapus image

### `docker image rm [REPOSITORY]:[TAG]`

example
```sh
docker image rm alta-api-docker:1.0
```
penting ! Jika terdapat container yang masih menjalankan atau menggunakan image tersebut, maka container harus distop dan dihapus terlebih dahulu. example:
```sh
docker container stop [namaContainer]
docker container rm [namaContainer]
```
