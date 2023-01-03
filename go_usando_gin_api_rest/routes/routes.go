package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasleitee/go_gin_api_rest/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoID)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.Run()
}
