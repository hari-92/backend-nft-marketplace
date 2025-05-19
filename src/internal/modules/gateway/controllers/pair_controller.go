package gateway_controllers

import "github.com/gin-gonic/gin"

type PairController struct {
}

func NewPairController() *PairController {
	return &PairController{}
}

// GetPairs: Get all pairs (GET /pairs)
func (t *PairController) GetPairs(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Pairs"})
}

// GetPair: Get a pair by id (GET /pairs/:pair_id)
func (t *PairController) GetPair(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Pair"})
}

// PostPair: Create a new pair (POST /pairs)
func (t *PairController) PostPair(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Pair"})
}

// PutPair: Update pair info (PUT /pairs/:pair_id)
func (t *PairController) PutPair(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Put Pair"})
}

// DeletePair: Delete a pair (DELETE /pairs/:pair_id)
func (t *PairController) DeletePair(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Delete Pair"})
}

// PostValidatePair: Validate a pair (POST /pairs/validate)
func (t *PairController) PostValidatePair(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Validate Pair"})
}
