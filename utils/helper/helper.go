package helper

import "go.mongodb.org/mongo-driver/mongo"

func IsDup(err error) bool {
	if wes, ok := err.(mongo.WriteException); ok {
		for i := range wes.WriteErrors {
			if wes.WriteErrors[i].Code == 11000 || wes.WriteErrors[i].Code == 11001 || wes.WriteErrors[i].Code == 12582 || wes.WriteErrors[i].Code == 16460 {
				return true
			}
		}
	}
	return false
}

func Contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
