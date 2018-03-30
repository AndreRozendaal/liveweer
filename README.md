
# Golang package for liveweer.nl

This is a Golang package for getting temperature and/or air humadity from liveweer.nl

## how to use
```
package main

import "liveweer"
import "fmt"

func main() {
        fmt.Printf("Actual temperature from %v: %v \n", liveweer.GetCountry(), liveweer.GetTemp())
        fmt.Printf("Actual air humidity from %v %%\n", liveweer.GetAirHumidity())
        liveweer.ChangeLocation("Hilversum")
        fmt.Printf("Actual temperature from %v: %v \n", liveweer.GetCountry(), liveweer.GetTemp())
        fmt.Printf("Actual air humidity from %v %%\n", liveweer.GetAirHumidity())
}
```

