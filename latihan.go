package main
import "fmt"

func palindrom(awal int, akhir int){
    var x, i, hasil int
    hasil = 0
    for awal <= akhir{
        i = awal
        for i > 0{
            x = i%10
            hasil = (hasil*10)+x
            i = i/10
        }
        if hasil == awal{
            fmt.Println(hasil)
            hasil = 0
            x = 0
        }else{
            hasil = 0
            x = 0 
        }
        awal++
        i = awal
    }
}

func main(){
    var d1, d2 int
    fmt.Scan(&d1, &d2)
    palindrom(d1, d2)
}