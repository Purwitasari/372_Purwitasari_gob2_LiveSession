Nama : Purwitasari
Kode : 372
Link Video : https://drive.google.com/drive/u/4/folders/1hMc53vXqJK6he4pKlrnWkSI1gmKLD2bw
Link Git : https://github.com/Purwitasari/372_Purwitasari_gob2_LiveSession.git


Soal :
1.  buat rest api
    endpoint /purwita
    output : 
        - nama lengkap
        - array custom (123, "Memasak", true)
        - menggunakan struct, data yg di tampilkan berupa json
    referense : sesi 6

2.  buat rest api with middleware
    endpoint /purwita/{middleware}
    output :
        - middleware 1 : halo saya purwita
        - middleware 2 : nampilin struct saja (pada soal no. 1)
        - endpoint : nampilin semua resp dari sesi 1.

3. buat testing menggunakan testify
    - di function : buat palindrome untuk cek kata "kasur ini rusak"

Langkah :
1. Download github 372_Purwitasari_gob2_LiveSession
2. Buka folder Purwitasari
3. Jalankan dengan go run main.go
4. Buka Postman (GET) dengan alamat :
    a. http://localhost:8080/purwita
    b. http://localhost:8080/purwita/1
    c. http://localhost:8080/purwita/2
    d. http://localhost:8080/purwita/3
5. Akan muncul sesuai dengan soal 1 dan soal 2
6. Untuk soal 3, buka folder helper dan jalankan go run ./helper
7. Akan muncul sesuai dengan soal 3