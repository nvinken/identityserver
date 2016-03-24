package user

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UsersInterface interface {
	// Create a new user
	// It is handler for POST /users
	Post(http.ResponseWriter, *http.Request)

	// RegisterNewEmailAddress is the handler for POST /users/{username}/emailaddresses
	// Register a new email address
	RegisterNewEmailAddress(http.ResponseWriter, *http.Request)
	// UpdateEmailAddress is the handler for PUT /users/{username}/emailaddresses/{label}
	// Updates the label and/or value of an email address
	UpdateEmailAddress(http.ResponseWriter, *http.Request)
	// DeleteEmailAddress is the handler for DELETE /users/{username}/emailaddresses/{label}
	// Removes an email address
	DeleteEmailAddress(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}/addresses/{label}
	usernameaddresseslabelGet(http.ResponseWriter, *http.Request)

	// Update or create an existing address.
	// It is handler for PUT /users/{username}/addresses/{label}
	usernameaddresseslabelPut(http.ResponseWriter, *http.Request)

	// Delete an address
	// It is handler for DELETE /users/{username}/addresses/{label}
	usernameaddresseslabelDelete(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}/banks
	usernamebanksGet(http.ResponseWriter, *http.Request)

	// Create new bank account
	// It is handler for POST /users/{username}/banks
	usernamebanksPost(http.ResponseWriter, *http.Request)

	// Get the list of authorization scopes
	// It is handler for GET /users/{username}/scopes
	usernamescopesGet(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}
	usernameGet(http.ResponseWriter, *http.Request)

	// Update an existing user. Updating ``username`` is not allowed. The labelled lists
	// can not be updated this way, the normal properties can (like github and facebook account).
	// It is handler for PUT /users/{username}
	usernamePut(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}/info
	usernameinfoGet(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}/validate
	usernamevalidateGet(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}/addresses
	usernameaddressesGet(http.ResponseWriter, *http.Request)

	// Create new address
	// It is handler for POST /users/{username}/addresses
	usernameaddressesPost(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}/phonenumbers
	usernamephonenumbersGet(http.ResponseWriter, *http.Request)

	// Create new phonenumber
	// It is handler for POST /users/{username}/phonenumbers
	usernamephonenumbersPost(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}/phonenumbers/{label}
	usernamephonenumberslabelGet(http.ResponseWriter, *http.Request)

	// Update or create an existing phonenumber.
	// It is handler for PUT /users/{username}/phonenumbers/{label}
	usernamephonenumberslabelPut(http.ResponseWriter, *http.Request)

	// Delete a phonenumber
	// It is handler for DELETE /users/{username}/phonenumbers/{label}
	usernamephonenumberslabelDelete(http.ResponseWriter, *http.Request)

	// It is handler for GET /users/{username}/banks/{label}
	usernamebankslabelGet(http.ResponseWriter, *http.Request)

	// Update or create an existing bankaccount.
	// It is handler for PUT /users/{username}/banks/{label}
	usernamebankslabelPut(http.ResponseWriter, *http.Request)

	// Delete a BankAccount
	// It is handler for DELETE /users/{username}/banks/{label}
	usernamebankslabelDelete(http.ResponseWriter, *http.Request)

	// Get a specific authorization
	// It is handler for GET /users/{username}/scopes/{grantedTo}
	usernamescopesgrantedToGet(http.ResponseWriter, *http.Request)

	// Update a Scope
	// It is handler for PUT /users/{username}/scopes/{grantedTo}
	usernamescopesgrantedToPut(http.ResponseWriter, *http.Request)

	// Remove a Scope, the granted organization will no longer have access the user's information.
	// It is handler for DELETE /users/{username}/scopes/{grantedTo}
	usernamescopesgrantedToDelete(http.ResponseWriter, *http.Request)

	// Get the list of notifications, these are pending invitations or approvals
	// It is handler for GET /users/{username}/notifications
	usernamenotificationsGet(http.ResponseWriter, *http.Request)
}

func UsersInterfaceRoutes(r *mux.Router, i UsersInterface) {
	r.HandleFunc("/users", i.Post).Methods("POST")

	r.HandleFunc("/users/{username}/emailaddresses", i.RegisterNewEmailAddress).Methods("POST")
	r.HandleFunc("/users/{username}/emailaddresses/{label}", i.UpdateEmailAddress).Methods("PUT")
	r.HandleFunc("/users/{username}/emailaddresses/{label}", i.DeleteEmailAddress).Methods("DELETE")

	r.HandleFunc("/users/{username}/addresses/{label}", i.usernameaddresseslabelGet).Methods("GET")
	r.HandleFunc("/users/{username}/addresses/{label}", i.usernameaddresseslabelPut).Methods("PUT")
	r.HandleFunc("/users/{username}/addresses/{label}", i.usernameaddresseslabelDelete).Methods("DELETE")
	r.HandleFunc("/users/{username}/banks", i.usernamebanksGet).Methods("GET")
	r.HandleFunc("/users/{username}/banks", i.usernamebanksPost).Methods("POST")
	r.HandleFunc("/users/{username}/scopes", i.usernamescopesGet).Methods("GET")
	r.HandleFunc("/users/{username}", i.usernameGet).Methods("GET")
	r.HandleFunc("/users/{username}", i.usernamePut).Methods("PUT")
	r.HandleFunc("/users/{username}/info", i.usernameinfoGet).Methods("GET")
	r.HandleFunc("/users/{username}/validate", i.usernamevalidateGet).Methods("GET")
	r.HandleFunc("/users/{username}/addresses", i.usernameaddressesGet).Methods("GET")
	r.HandleFunc("/users/{username}/addresses", i.usernameaddressesPost).Methods("POST")
	r.HandleFunc("/users/{username}/phonenumbers", i.usernamephonenumbersGet).Methods("GET")
	r.HandleFunc("/users/{username}/phonenumbers", i.usernamephonenumbersPost).Methods("POST")
	r.HandleFunc("/users/{username}/phonenumbers/{label}", i.usernamephonenumberslabelGet).Methods("GET")
	r.HandleFunc("/users/{username}/phonenumbers/{label}", i.usernamephonenumberslabelPut).Methods("PUT")
	r.HandleFunc("/users/{username}/phonenumbers/{label}", i.usernamephonenumberslabelDelete).Methods("DELETE")
	r.HandleFunc("/users/{username}/banks/{label}", i.usernamebankslabelGet).Methods("GET")
	r.HandleFunc("/users/{username}/banks/{label}", i.usernamebankslabelPut).Methods("PUT")
	r.HandleFunc("/users/{username}/banks/{label}", i.usernamebankslabelDelete).Methods("DELETE")
	r.HandleFunc("/users/{username}/scopes/{grantedTo}", i.usernamescopesgrantedToGet).Methods("GET")
	r.HandleFunc("/users/{username}/scopes/{grantedTo}", i.usernamescopesgrantedToPut).Methods("PUT")
	r.HandleFunc("/users/{username}/scopes/{grantedTo}", i.usernamescopesgrantedToDelete).Methods("DELETE")
	r.HandleFunc("/users/{username}/notifications", i.usernamenotificationsGet).Methods("GET")
}
