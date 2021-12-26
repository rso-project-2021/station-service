package api

import (
	"encoding/json"
	"log"
	"net/http"
	"station-service/db"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type graphRequest struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

type getCloseByStations struct {
	Lat    float32 `form:"lat"`
	Lng    float32 `form:"lng"`
	Offset int32   `form:"offset"`
	Limit  int32   `form:"limit" binding:"required,min=1,max=20"`
}

var stationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Station",
		Fields: graphql.Fields{
			"station_id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"lat": &graphql.Field{
				Type: graphql.Float,
			},
			"lng": &graphql.Field{
				Type: graphql.Float,
			},
			"provider": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func GetSchema(stations []db.Station) (graphql.Schema, error) {
	fields := graphql.Fields{
		"list": &graphql.Field{
			Type:        graphql.NewList(stationType),
			Description: "Get Station List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return stations, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Println("Cannot create GraphQL schema: ", err)
	}

	return schema, nil
}

func (server *Server) GetWithQuery(ctx *gin.Context) {

	// Check if request has params latitude, longitude, offset and limit.
	var req getCloseByStations
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.OriginLocationParam{
		Lat:    req.Lat,
		Lng:    req.Lng,
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	// Execute query.
	stations, err := server.store.GetAllByDistance(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Parse graphQL request.
	var graphRequest graphRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&graphRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "GraphQL request in wrong format."})
		ctx.Abort()
		return
	}

	// Create GraphQL schema.
	schema, err := GetSchema(stations)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "GraphQL schema couldn't be created."})
		ctx.Abort()
		return
	}

	// Execute GraphQL request.
	result := graphql.Do(graphql.Params{
		Context:        ctx.Request.Context(),
		Schema:         schema,
		RequestString:  graphRequest.Query,
		VariableValues: graphRequest.Variables,
		OperationName:  graphRequest.Operation,
	})

	if len(result.Errors) > 0 {
		log.Fatalf("Failed to execute graphQL query: %+v", result.Errors)
	}

	ctx.JSON(http.StatusOK, result)
}
