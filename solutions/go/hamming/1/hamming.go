package hamming

import "errors"

func Distance(a, b string) (int, error) {
    d := 0
    
	if len(a) != len(b) {
    	return 0, errors.New("strings needs to have same len")
    }

    for i, _ := range a {
        if a[i] != b[i] {
        	d+=1
        }
    }

    return d, nil
}
