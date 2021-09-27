# Learning Docker

## Mengecek image yang terdapat pada docker

#### `docker images`

## Membuat image

terdapat dua pendekatan untuk membuat image di local PC,

### 1. Mengambil dari registry (docker hub)

#### `docker pull [REPOSITORY]:[TAG]`

contoh

```sh
docker pull mysql:8.0
```

jika tag tidak disebutkan, maka docker akan mengambil versi terbaru

### 2. Menggunakan Dockerfile

#### `docker build --tag [REPOSITORY]:[TAG] path/to/dockerfile`

contoh

```sh
docker build --tag GoApps:1.0 .
```

## Menghapus image

### `docker image rm [REPOSITORY]:[TAG]`

contoh

```sh
docker image rm alta-api-docker:1.0
```

penting ! Jika terdapat container yang masih menjalankan atau menggunakan image tersebut, maka container harus distop dan dihapus terlebih dahulu. contoh:

```sh
docker container stop [namaContainer]
docker container rm [namaContainer]
```

## Melihat container

### `docker container ls --all`
