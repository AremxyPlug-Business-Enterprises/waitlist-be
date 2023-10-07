package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"waitlist-be/models"
	"waitlist-be/utils"
)

func AddToWaitlist(w http.ResponseWriter, r *http.Request) {
	var waitlistEntry models.WaitlistEntry
	err := json.NewDecoder(r.Body).Decode(&waitlistEntry)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	db := utils.ConnectDatabase()
	collection := db.Collection("waitlist")

	_, err = collection.InsertOne(r.Context(), waitlistEntry)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User added to waitlist"})
}

func GetWaitlist(w http.ResponseWriter, r *http.Request) {
	db := utils.ConnectDatabase()
	collection := db.Collection("waitlist")

	var waitlist []models.WaitlistEntry
	cursor, err := collection.Find(r.Context(), nil)
	if err != nil {
		log.Println("MongoDB Find Error:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	defer cursor.Close(r.Context())

	for cursor.Next(r.Context()) {
		var entry models.WaitlistEntry
		if err := cursor.Decode(&entry); err != nil {
			log.Println("MongoDB Decode Error:", err)
			continue // Skip this document and continue with the next one
		}
		waitlist = append(waitlist, entry)
	}

	utils.RespondWithJSON(w, http.StatusOK, waitlist)
}
