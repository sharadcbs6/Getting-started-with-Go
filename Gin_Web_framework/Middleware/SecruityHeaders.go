/*
					Context - Secruity Headers
				It’s important to use security headers to protect your web application from common security vulnerabilities.
				 This example shows you how to add security headers to your Gin application and 
				 also how to avoid Host Header Injection related attacks (SSRF, Open Redirection).
				
				 What each header protects against
				Header								Purpose
				X-Content-Type-Options:nosniff		Prevents MIME-type sniffing attacks. 
													Without this header, browsers may interpret files as a different content type than declared, 
													allowing attackers to execute malicious scripts disguised as harmless file types (e.g., uploading a .jpg that is actually JavaScript).
				X-Frame-Options: DENY				Prevents clickjacking by disabling the page from being loaded inside an <iframe>. Attackers use clickjacking to overlay invisible frames on legitimate pages,
													tricking users into clicking hidden buttons (e.g., “Delete my account”).
													Content-Security-Policy	Controls which resources (scripts, styles, images, fonts, etc.)
												    the browser is allowed to load and from which origins. This is one of the most effective defenses against Cross-Site Scripting (XSS) because it can block inline scripts and restrict script sources.
				X-XSS-Protection: 1; mode=block	    Activates the browser’s built-in XSS filter. This header is largely deprecated in modern browsers (Chrome removed its XSS Auditor in 2019), but it still provides defense-in-depth for users on older browsers.
				Strict-Transport-Security			Forces the browser to use HTTPS for all future requests to the domain for the specified max-age duration. 
													This prevents protocol downgrade attacks and cookie hijacking over insecure HTTP connections. The includeSubDomains directive extends this protection to all subdomains.
				Referrer-Policy: strict-origin	    Controls how much referrer information is sent with outgoing requests. Without this header, the full URL (including query parameters that may contain tokens or sensitive data) 
													can leak to third-party sites. strict-origin sends only the origin (domain) and only over HTTPS.
                Permissions-Policy	                Restricts which browser features (geolocation, camera, microphone, etc.) can be used by the page. This limits the damage if an attacker manages to inject scripts, since those scripts cannot access sensitive device APIs.
*/
package main
import(

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	expectedHost:="localhost:8080"
	router.Use(func(ctx *gin.Context) {
		if ctx.Request.Host!=expectedHost{
			ctx.AbortWithStatusJSON(400,gin.H{
				"error":	"Invalid Host request",
			})
		}
			ctx.Header("X-Frame-Options","DENY")
			ctx.Header("Content-Security-Policy","default-src:'self';connect-src *;font-src *;script-src-elem * 'unsafe-inline'	")
			ctx.Header("X-XSS-Protection","1;mode=block")
			ctx.Header("Strict-Transport-Secruity","max-age=31536000;includeSubDomains")
			ctx.Header("Referrer-Policy","strict-origin")
			ctx.Header("X-Content-Type-Options","nosniff")
			ctx.Header("Permissions-Policy","geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
			ctx.Next()	
	})
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"pong",
		})
	})
	router.Run()
}