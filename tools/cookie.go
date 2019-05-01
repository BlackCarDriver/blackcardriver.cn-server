package tools
/*
the function related to take and set cookie in writen in thsi file
*/
import (
	"net/http"
	"../mylog"
	"../config"
	"fmt"
	"net"
)
//create an particul cookie
func MakeCookie(key string, value string, time int)(ck http.Cookie) {
	ck = http.Cookie{
        Name: key,
		Value: value,
		MaxAge: time,			
		HttpOnly:true,
		Secure:(config.UseHttps=="true"),
	}
	return ck
}
//set cookie on user
func SetCookietest(w http.ResponseWriter){
	randkey := CreateRandString(18)
	fmt.Println("randky save to cookie is :" , randkey)
	ck := MakeCookie("carkey", randkey, 30)
	http.SetCookie(w, &ck)
}
//reade cookie from user
func ReadCookieTest(req *http.Request)bool{
	ck, err := req.Cookie("carkey")
    if err != nil {
	   fmt.Println(err)
	   return false
	}
	fmt.Println(ck)
	return true
}

//get client ip from request header
func GetIptest(r *http.Request){
	remoteAddr := r.RemoteAddr
	XForwardedFor := "X-Forwarded-For"
    XRealIP       := "X-Real-IP"
    if ip := r.Header.Get(XRealIP); ip != "" {
        remoteAddr = ip
    } else if ip = r.Header.Get(XForwardedFor); ip != "" {
        remoteAddr = ip
    } else {
        remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
    }
    if remoteAddr == "::1" {
        remoteAddr = "127.0.0.1"
	}
	mylog.Println(remoteAddr)
}