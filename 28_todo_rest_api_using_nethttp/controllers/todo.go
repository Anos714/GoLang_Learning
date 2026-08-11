package controllers

import (
	"encoding/json"
	"net/http"
	"todo-rest-api/config"
	"todo-rest-api/models"
	"todo-rest-api/utils"
)


func CreateTodo(w http.ResponseWriter,r *http.Request){
	data,ok:=utils.ReadJson(w,r)
	if(!ok){
		utils.WriteJson(w,http.StatusBadRequest,utils.DataBody{
			Success: false,
			Message: "Invalid request body or title required",
			Data: nil,
		})
		return
	}

	todo:=models.Todo{Title: data.Title,Done: data.Done}
	result:=config.DB.Create(&todo)
	if result.Error!=nil{
		utils.WriteJson(w,http.StatusInternalServerError,utils.DataBody{
			Success: false,
			Message: "Error when creating todo",
			Data: nil,
		})
		return
	}

	utils.WriteJson(w, http.StatusCreated, utils.DataBody{
			Success: true,
			Message: "Todo created",
			Data:    data,
		})

}

func GetTodos(w http.ResponseWriter,r *http.Request){
	var todos []models.Todo
	result:=config.DB.Find(&todos)
	if result.Error!=nil{
		utils.WriteJson(w,http.StatusInternalServerError,utils.DataBody{
			Success: false,
			Message: "Error when fetching todos",
			Data: nil,
		})
		return
	}

	utils.WriteJson(w,http.StatusOK,utils.DataBody{
		Success: true,
		Message: "",
		Data: todos,
	})
}

func GetTodoById(w http.ResponseWriter,r *http.Request){
	id:=r.PathValue("id")
	if id==""{
		utils.WriteJson(w,http.StatusBadRequest,utils.DataBody{
			Success: false,
			Message: "Bad request or Todo id not found",
			Data: nil,
		})
		return
	}
	var todo models.Todo
	result:=config.DB.Find(&todo,id)

	if result.Error!=nil{
		utils.WriteJson(w,http.StatusInternalServerError,utils.DataBody{
			Success: false,
			Message: "Todo not found",
			Data: nil,
		})
		return
	}

	utils.WriteJson(w,http.StatusOK,utils.DataBody{
		Success: true,
		Message: "",
		Data: todo,
	})

}

func DeleteTodo(w http.ResponseWriter,r *http.Request){
	id:=r.PathValue("id")
	if id==""{
		utils.WriteJson(w,http.StatusBadRequest,utils.DataBody{
			Success: false,
			Message: "Bad request or Todo id not found",
			Data: nil,
		})
		return
	}

	var todo models.Todo
	result:=config.DB.Delete(&todo,id)

	if result.Error!=nil{
		utils.WriteJson(w,http.StatusInternalServerError,utils.DataBody{
			Success: false,
			Message: "Intrenal server error while deleting todo",
			Data: nil,
		})
		return
	}

	if result.RowsAffected==0{
		utils.WriteJson(w, http.StatusNotFound, utils.DataBody{
					Success: false,
					Message: "Todo not found with the given ID",
					Data:    nil,
				})
				return
	}

	utils.WriteJson(w,http.StatusOK,utils.DataBody{
		Success: true,
		Message: "Todo deleted",
		Data: nil,
	})

}

func UpdateTodo(w http.ResponseWriter,r *http.Request){
	id:=r.PathValue("id")

	if id==""{
		utils.WriteJson(w,http.StatusBadRequest,utils.DataBody{
			Success: false,
			Message: "Bad request or Todo id not found",
			Data: nil,
		})
		return
	}

	var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			utils.WriteJson(w, http.StatusBadRequest, utils.DataBody{
				Success: false,
				Message: "Invalid JSON body",
				Data:    nil,
			})
			return
		}

		// 2. Ek khaali map banayein jisme hum sirf wahi fields daalenge jo user ne bheji hain
		updateData := make(map[string]any)

		// Check karein agar 'title' body mein exist karta hai
		if title, exists := body["title"]; exists {
			updateData["title"] = title
		}

		// Check karein agar 'completed' body mein exist karta hai (yeh false ko bhi handle karega)
		if done, exists := body["done"]; exists {
			updateData["done"] = done
		}

		// Agar user ne kuch bhi nahi bheja body mein
		if len(updateData) == 0 {
			utils.WriteJson(w, http.StatusBadRequest, utils.DataBody{
				Success: false,
				Message: "No fields provided to update (send title or completed)",
				Data:    nil,
			})
			return
		}

		// 3. Database mein update query chalayein
		var todo models.Todo
		result := config.DB.Model(&todo).Where("id = ?", id).Updates(updateData)

		// Database error check karein
		if result.Error != nil {
			utils.WriteJson(w, http.StatusInternalServerError, utils.DataBody{
				Success: false,
				Message: "Internal server error while updating todo",
				Data:    nil,
			})
			return
		}

		// Check karein ki kya woh ID database mein thi bhi ya nahi
		if result.RowsAffected == 0 {
			utils.WriteJson(w, http.StatusNotFound, utils.DataBody{
				Success: false,
				Message: "Todo not found with the given ID",
				Data:    nil,
			})
			return
		}

		utils.WriteJson(w, http.StatusOK, utils.DataBody{
			Success: true,
			Message: "Todo updated successfully",
			Data:    nil,
		})

}
