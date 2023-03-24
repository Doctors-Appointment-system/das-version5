package routes

import (
	controller "Doctor-Appointment-Project/controllers"
	middleware "Doctor-Appointment-Project/middleware"

	"github.com/gin-gonic/gin"
)

func PatientRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/get_all_doctors", controller.Get_docter())
	incomingRoutes.GET("/get_doctor_by_city&Specialisation", controller.GetDoctorByLocation())
	incomingRoutes.GET("patient/get_my_details", controller.Get_my_details())
	incomingRoutes.POST("/add_patient_details", controller.Addpatient())
	incomingRoutes.POST("/bookappointment", controller.BookingAppointment())
	incomingRoutes.DELETE("/remove_patient", controller.DeletePatient())
	incomingRoutes.DELETE("/cancelAppointment", controller.Cancel_appointment())

}
