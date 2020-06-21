package anilist

import "time"

func GetCurrentSeasonName() string {
	month := time.Now().Month().String()

	winter := []string{"December", "January", "February"}
	spring := []string{"March", "April", "May", "June"}
	summer := []string{"July", "August"}
	fall := []string{"September", "October", "November"}

	if contains(winter, month) {
		return "WINTER"
	} else if contains(spring, month) {
		return "SPRING"
	} else if contains(summer, month) {
		return "SUMMER"
	} else if contains(fall, month) {
		return "FALL"
	}

	return ""
}

func contains(array []string, item string) bool {
	for _, element := range array {
		if element == item {
			return true
		}
	}
	return false
}
