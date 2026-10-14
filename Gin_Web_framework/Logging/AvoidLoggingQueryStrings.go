/*
			Context - Avoid logging query strings
	Story building
	Query strings often contain sensitive information such as API tokens,
	passwords, session IDs, or personally identifiable information (PII).
	Solution:)
	Logging these values can create security risks and may violate privacy regulations
	like GDPR or HIPAA. By stripping query strings from your logs, you reduce the chance of leaking
	sensitive data through log files, monitoring systems, or error reporting tools.
	Use the SkipQueryString option in LoggerConfig to prevent query strings from appearing in logs.
	 When enabled, a request to /path?token=secret&user=alice will be logged simply as /path.
*/

package main

import "github.com/gin-gonic/gin"
func main(){
	router:=gin.New();
	loggerConfig:=gin.LoggerConfig{SkipQueryString: true}
	router.Use(gin.LoggerWithConfig(loggerConfig))
	router.Use(gin.Recovery());
	router.GET("/search",func(ctx *gin.Context) {
		query:=ctx.Query("query")
		ctx.String(200,"searching for:"+query)

	})
	router.Run(":8080")
}