package errors

// GetBadRequestBodyError Return the error information of type: Bad Request Body.
func GetBadRequestBodyError() APIError {

	return APIError{
		ErrorCode: BadRequestBody,
		Message:   "The Request Body is Wrong",
	}

}

// GetUnauthorizedError Return the error information of type: Unauthorized.
func GetUnauthorizedError() APIError {

	return APIError{
		ErrorCode: Unauthorized,
		Message:   "The user don't have access",
	}

}

// GetInternalServerError Return the error information of type: Unauthorized.
func GetInternalServerError() APIError {

	return APIError{
		ErrorCode: InternalServer,
		Message:   "There was an unexpected error on server",
	}

}
