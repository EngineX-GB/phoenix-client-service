package service

import (
	"encoding/json"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"
)

func DeleteDataOverride(userId string, fieldName string) error {
	// get a row from the database for the override record

	overrideJson, err := dao.GetOverrideByUserId(userId)

	if err != nil {
		return err
	}

	// then extract the json from it and convert it into an object
	overrideData, err := ConvertJsonStringToOverrideDataArray(overrideJson)

	if err != nil {
		return err
	}

	// now that you got the override json in object format, then remove the element that is deleted.

	filtered := overrideData[:0]

	for _, overrideElement := range overrideData {
		if overrideElement.Name != fieldName {
			filtered = append(filtered, overrideElement)
		}
	}

	overrideData = filtered

	if len(overrideData) != 0 {
		// if the override json array is not empty after deleting the override field, then update the database record with the
		// updated list

		data, err := json.Marshal(overrideData)
		if err != nil {
			return err
		}
		updateOverrideDataString := string(data)
		err = dao.UpdateDataOverride(userId, updateOverrideDataString)
		if err != nil {
			return err
		}
	} else {
		// if the override json array is empty, then delete the single override database row from the database.
		err := dao.DeleteDataOverride(userId)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	This is the function to return the list of overrides (the RAW database records)
 */
func GetAllDataOverrides() ([]model.SingleOverrideRecord, error) {
	var overrideEntities []model.OverrideEntity
	overrideRecords, err := dao.GetAllOverrideData()
	if err != nil {
		return nil, err
	}

	// assuming that the database records are retrieved, convert the JSON String (data) into an actual JSON object list
	// to extract individual overrides. Then create each separate overrideEntity (like a record) for EACH individual override (for all users)

	for _, overrideRecord := range overrideRecords {
		var overrideEntity model.OverrideEntity
		overrideEntity.Oid = overrideRecord.Oid
		overrideEntity.UserId = overrideRecord.UserId
		overrideEntity.UserName = overrideRecord.UserName
		//convert the json to an array object
		jsonObjArray, err := ConvertJsonStringToOverrideDataArray(overrideRecord.Data)
		if err != nil {
			return nil, err
		}
		overrideEntity.Data = jsonObjArray
		overrideEntities = append(overrideEntities, overrideEntity)
	}

	var singleOverrideRecords []model.SingleOverrideRecord

	for _, overrideEntity := range overrideEntities {
		// loop through the properties and for each property, create a single override record
		for _, property := range overrideEntity.Data {
			var singleOverrideRecord model.SingleOverrideRecord
			singleOverrideRecord.Oid = overrideEntity.Oid
			singleOverrideRecord.UserId = overrideEntity.UserId
			singleOverrideRecord.UserName = overrideEntity.UserName
			singleOverrideRecord.PropertyName = property.Name
			singleOverrideRecord.PropertyValue = property.Value
			singleOverrideRecords = append(singleOverrideRecords, singleOverrideRecord)
		}
	}

	return singleOverrideRecords, nil
}

func ConvertJsonStringToOverrideDataArray(existingJsonString string) ([]model.OverrideData, error) {
	var existingOverrideDataObject []model.OverrideData
	err := json.Unmarshal([]byte(existingJsonString), &existingOverrideDataObject)
	if err != nil {
		return nil, err
	}
	return existingOverrideDataObject, nil

}

func AddNewOverride(userId string, overrideData model.OverrideData) error {
	var existingOverrideDataObject []model.OverrideData
	var newOverrideRecord bool = false
	// add the dao to save the override data to the database

	// get the override string (containing existing overrides) that are already in the database for the user
	existingJsonString, err := dao.GetOverrideByUserId(userId)

	if err != nil {
		return err
	}

	if existingJsonString == "[]" {
		// this will be a new record
		newOverrideRecord = true
	}

	err = json.Unmarshal([]byte(existingJsonString), &existingOverrideDataObject)
	if err != nil {
		return err
	}

	// add the new json string to the existing ones for the user
	existingOverrideDataObject = append(existingOverrideDataObject, overrideData)

	// convert back to a json string

	data, err := json.Marshal(existingOverrideDataObject)
	if err != nil {
		return err
	}

	updateOverrideDataString := string(data)

	// then store in the datatase
	if newOverrideRecord {
		err = dao.AddDataOverride(userId, updateOverrideDataString)
		if err != nil {
			return err
		}
	} else {
		// update the existing record
		err = dao.UpdateDataOverride(userId, updateOverrideDataString)
		if err != nil {
			return err
		}
	}

	return nil
}
