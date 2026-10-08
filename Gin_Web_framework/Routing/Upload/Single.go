package main
import(
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
)
func main(){
	router:=gin.Default();
	router.MaxMultipartMemory=8<<20; //set a lower limit , default limit is 32MiB
	router.POST("/upload",func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return  
		}
		log.Println(file.Filename)

		destination:= filepath.Join("./files/",filepath.Base(file.Filename))
		ctx.SaveUploadedFile(file,destination);
		ctx.String(http.StatusOK,fmt.Sprintf("'%s' uploaded!",file.Filename))
	})
	router.Run(":8080");	
}
