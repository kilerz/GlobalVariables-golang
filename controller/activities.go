package controller

import "activities/model"

func (data *DataGlobal) ReadActivities() (result []model.Activities) {
	return data.Activities
}

func (data *DataGlobal) CreateActivities(input model.Activities) {
	data.Activities = append(data.Activities, input)
}

func (data *DataGlobal) UpdateActivities(index int, input model.Activities) {
	data.Activities[index] = input
}

func (data *DataGlobal) DeleteActivities(index int) {
	data.Activities = append(data.Activities[:index], data.Activities[index+1:]...)
}
