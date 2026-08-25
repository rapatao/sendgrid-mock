package restrouters

import "github.com/gin-gonic/gin"

// AbortWithError replies with the SendGrid error shape: {"errors":[{...}]}.
// Empty field or help entries are omitted.
func AbortWithError(context *gin.Context, status int, field, message, help string) {
	entry := gin.H{"message": message}

	if field != "" {
		entry["field"] = field
	}

	if help != "" {
		entry["help"] = help
	}

	context.AbortWithStatusJSON(status, gin.H{"errors": []gin.H{entry}})
}
