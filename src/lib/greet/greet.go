package greet

import "time"

var t = time.Now()


func Now() string {
    return "Now it is " + t.Format("15:04:05") + "."
}


func Greet() string {

    switch {
        default:
            return Now() + " Hello?"

        case t.Hour() < 12:
            return Morning()

        case t.Hour() < 17:
            return Afternoon()

        case t.Hour() >= 17:
            return Evening()

    }
}
