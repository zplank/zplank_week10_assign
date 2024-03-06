# zplank_week10_assign

The 'trimmedmean' package was created as part of this branch by executing the go build command. In order access and utilize the package, a user can follow the below instructions. 

1) Create or navigate to the project where the trimmedmean package will be used 
2) Import the package as part of the Go file imports. An example would be - 
    import "github.com/zplank/trimmedmean"
3) Download and install package using 'go mod tidy' command 
4) Utilize the package within the code. An example use case would be - 

        package main

        import (
            "fmt"
            "github.com/zplank/trimmedmean"
        )

        func main() {
            data := []float64{1, 3.2, 6, 7, 10.5, 15.4, 22, 25, 30, 77.6, 100}
            trimmedMean, err := trimmedmean.ComputeTrimmedMean(data, 0.1, 0.1)
            if err != nil {
                fmt.Println("Error calculating trimmed mean:", err)
                return
            }
            fmt.Println("Trimmed Mean:", trimmedMean)
        } 