package validation

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/itsyouonline/identityserver/credentials/password"
	"github.com/itsyouonline/identityserver/db/validation"
	"net/http"
	"net/url"
)

//SMSService is the interface an sms communicaction channel should have to be used by the IYOPhonenumberValidationService
type EmailService interface {
	Send(receipients []string, subject string, message string) (err error)
}

//IYOPhonenumberValidationService is the itsyou.online implementation of a PhonenumberValidationService
type IYOEmailAddressValidationService struct {
	EmailService EmailService
}

//RequestValidation validates the email address by sending an EMail
func (service *IYOEmailAddressValidationService) RequestValidation(request *http.Request, username string, email string, confirmationurl string) (key string, err error) {
	valMngr := validation.NewManager(request)
	info, err := valMngr.NewEmailAddressValidationInformation(username, email)
	if err != nil {
		log.Error(err)
		return
	}
	err = valMngr.SaveEmailAddressValidationInformation(info)
	if err != nil {
		log.Error(err)
		return
	}
	message := fmt.Sprintf("To verify your Email Address on itsyou.online enter the code %s in the form or use this link: %s?c=%s&k=%s", info.Secret, confirmationurl, url.QueryEscape(info.Secret), url.QueryEscape(info.Key))

	go service.EmailService.Send([]string{email}, "ItsYouOnline Email Validation", message)
	key = info.Key
	return
}

//RequestPassword reset
func (service *IYOEmailAddressValidationService) RequestPasswordReset(request *http.Request, username string, emails []string) (key string, err error) {
	pwdMngr := password.NewManager(request)
	token, err := pwdMngr.NewResetToken(username)
	if err != nil {
		return
	}
	if err = pwdMngr.SaveResetToken(token); err != nil {
		return
	}
	message := fmt.Sprintf("To reset your password on itsyou.online please visit https://%s/login#/resetpassword?token=%s", request.Host, url.QueryEscape(token.Token))
	go service.EmailService.Send(emails, "ItsYouOnline Password Reset", message)
	key = token.Token
	return
}

//ExpireValidation removes a pending validation
func (service *IYOEmailAddressValidationService) ExpireValidation(request *http.Request, key string) (err error) {
	if key == "" {
		return
	}
	valMngr := validation.NewManager(request)
	err = valMngr.RemoveEmailAddressValidationInformation(key)
	return
}

func (service *IYOEmailAddressValidationService) getEmailAddressValidationInformation(request *http.Request, key string) (info *validation.EmailAddressValidationInformation, err error) {
	if key == "" {
		return
	}
	valMngr := validation.NewManager(request)
	info, err = valMngr.GetByKeyEmailAddressValidationInformation(key)
	return
}

//IsConfirmed checks wether a validation request is already confirmed
func (service *IYOEmailAddressValidationService) IsConfirmed(request *http.Request, key string) (confirmed bool, err error) {
	info, err := service.getEmailAddressValidationInformation(request, key)
	if err != nil {
		return
	}
	if info == nil {
		err = ErrInvalidOrExpiredKey
		return
	}
	confirmed = info.Confirmed
	return
}

//ConfirmValidation checks if the supplied code matches the username and key
func (service *IYOEmailAddressValidationService) ConfirmValidation(request *http.Request, key, secret string) (err error) {
	info, err := service.getEmailAddressValidationInformation(request, key)
	if err != nil {
		return
	}
	if info == nil {
		err = ErrInvalidOrExpiredKey
		return
	}
	if info.Secret != secret {
		err = ErrInvalidCode
		return
	}
	valMngr := validation.NewManager(request)
	p := valMngr.NewValidatedEmailAddress(info.Username, info.EmailAddress)
	err = valMngr.SaveValidatedEmailAddress(p)
	if err != nil {
		return
	}
	err = valMngr.UpdateEmailAddressValidationInformation(key, true)
	if err != nil {
		return
	}
	return
}
