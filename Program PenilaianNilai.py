Program PenilaianNilai

Deklarasi
    nilai: integer
    valid: boolean
    grade: string
    ulang: String 

Algoritma
    Repeat
        Write("Masukkan nilai")
        Read(nilai)

        If nilai < 0 or nilai > 100 then
            Write("Input tidak valid")
            Continue
        EndIf

        If nilai >= 90 then
            grade := "A"
        ElseIf nilai >= 80 then
            grade := "B"
        ElseIf nilai >= 70 then
            grade := "C"
        ElseIf nilai >= 60 then                     
            grade := "D"
        Else
            grade := "E"
        EndIf

        Write("Grade: ", grade)
        Write("Ulangi program? (y/n)")
        Read(ulang)
    While ulang = "y" or ulang = "Y"

Selesai