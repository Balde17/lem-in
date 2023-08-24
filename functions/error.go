package lemIn

func FindErrorMessage(s string) string {
	var message string

	errorInformation := map[string]string{
		"openFile":            "ERROR: invalid data format - unable to open file.",
		"ants":                "ERROR: invalid data format - missing or too small/ large number of ants.",
		"missingStartEndRoom": "ERROR: invalid data format - missing start or end room.",
		"duplicateRoom":       "ERROR: invalid data format - duplicate room or coordinates.",
		"unknownRoom":         "ERROR: invalid data format - unknown room.",
		"invalidCoordinates":  "ERROR: invalid data format - bad room coordinates.",
		"intCoordinates":      "ERROR: invalid data format - room coordinates must be of type int.",
		"roomName":            "ERROR: invalid data format - bad room name.",
		"infiniteLoop":        "ERROR: invalid data format - room that links to itself generates infinite loop.",
		"badTunnel":           "ERROR: invalid data format - tunnel with links to more than two rooms not allowed.",
		"badRmNmCoord":        "ERROR: invalid data format - bad room name or invalid room coordinates.",
	}

	for k, v := range errorInformation {
		if k == s {
			message = v
		}
	}
	return message
}
