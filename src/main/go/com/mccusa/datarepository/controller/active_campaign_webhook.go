package controller

import (
	"go-data-repository/src/main/go/com/mccusa/datarepository/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterActiveCampaignWebhookRoutes registers webhook endpoints
func RegisterActiveCampaignWebhookRoutes(rg *gin.RouterGroup, svc service.ActiveCampaignService, getUsersSvc service.GetActiveCampaignUsers) {
	grp := rg.Group("/active-campaign-webhook")
	grp.POST("", handleWebhook(svc))
	grp.GET("/test-active-campaign", testActiveCampaign(getUsersSvc))
}

// handleWebhook parses form params into nested JSON and invokes the service
func handleWebhook(svc service.ActiveCampaignService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Request.ParseForm(); err != nil {
			c.String(http.StatusBadRequest, "invalid form: %v", err)
			return
		}
		payload := make(map[string]interface{})
		for key, vals := range c.Request.PostForm {
			setNestedKey(payload, key, vals[0])
		}
		svc.HandleWebhookPayload(c.Request.Context(), payload)
		c.Status(http.StatusOK)
	}
}

// testActiveCampaign provides a simple health check
func testActiveCampaign(getUsers service.GetActiveCampaignUsers) gin.HandlerFunc {
	return func(c *gin.Context) {
		getUsers.GetContactsFromAPI(c.Request.Context())
		c.String(http.StatusOK, "Active Campaign Webhook Controller is working!")
	}
}

// setNestedKey inserts a value into a nested map based on bracket-notation keys
func setNestedKey(root map[string]interface{}, compoundKey, val string) {
	parts := []string{}
	cur := ""
	for _, ch := range compoundKey {
		switch ch {
		case '[':
			parts = append(parts, cur)
			cur = ""
		case ']':
			parts = append(parts, cur)
			cur = ""
		default:
			cur += string(ch)
		}
	}
	if cur != "" {
		parts = append(parts, cur)
	}

	var current interface{} = root
	for i, part := range parts {
		isLast := i == len(parts)-1
		idx, err := strconv.Atoi(part)
		isIndex := err == nil

		switch node := current.(type) {
		case map[string]interface{}:
			if isLast {
				node[part] = val
			} else {
				next, exists := node[part]
				if !exists {
					if isIndex {
						next = []interface{}{}
					} else {
						next = make(map[string]interface{})
					}
					node[part] = next
				}
				current = next
			}
		case []interface{}:
			if !isIndex {
				return
			}
			if idx >= len(node) {
				// expand
				newArr := make([]interface{}, idx+1)
				copy(newArr, node)
				for k := len(node); k <= idx; k++ {
					newArr[k] = make(map[string]interface{})
				}
				node = newArr
				// assign back to the parent map
				root[parts[0]] = node
			}
			if isLast {
				node[idx] = val
			} else {
				current = node[idx]
			}
		}
	}
}
