package user

import (
	"github.com/bastean/dsgo/pkg/context/shared/domain/models"
	"github.com/bastean/dsgo/pkg/context/shared/infrastructure/transports"
	"github.com/bastean/dsgo/pkg/context/user/infrastructure/communication/mail"
	"github.com/bastean/dsgo/pkg/context/user/infrastructure/communication/terminal"
)

func MailConfirmation(smtp *transports.SMTP) models.Transport {
	return &mail.Confirmation{
		SMTP: smtp,
	}
}

func TerminalConfirmation(logger models.Logger, serverURL string) models.Transport {
	return &terminal.Confirmation{
		Logger:    logger,
		ServerURL: serverURL,
	}
}
