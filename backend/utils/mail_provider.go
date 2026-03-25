package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendMail fulfills the Balej_Mail_Relay protocol.
// It handles delivery to a specific identity node using configured SMTP credentials.
func SendMail(to []string, subject string, body string) error {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Message configuration with high-fidelity branding
	// Note: We use an HTML-ready MIME structure to support the cyberpunk-industrial design
	message := []byte(fmt.Sprintf(
		"From: Balej Identity Node <%s>\r\n"+
		"To: %s\r\n"+
		"Subject: [BALEJ] %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
		"%s\r\n", from, to[0], subject, body))

	// Authenticate with the delivery cluster
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Dispatch the identity payload
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		return fmt.Errorf("mail_relay_failure: %v", err)
	}

	return nil
}

// GenerateVerificationTemplate returns the high-fidelity HTML for email verification
func GenerateVerificationTemplate(userName string, verificationLink string) string {
	return fmt.Sprintf(`
		<div style="background-color: #020205; color: #ffffff; font-family: 'Courier New', monospace; padding: 40px; border: 1px solid #22d3ee;">
			<h1 style="color: #22d3ee; text-transform: uppercase; font-style: italic; letter-spacing: 2px;">Identity_Protocol // Activate_Node</h1>
			<p style="color: #64748b;">Greetings_Operator <strong>%s</strong>,</p>
			<p>Your identity has been registered with the <strong>Balej ATS Global Network</strong>. To finalize your access to the sector dashboards, you must confirm your email node.</p>
			<div style="margin: 30px 0;">
				<a href="%s" style="background-color: #22d3ee; color: #000; text-decoration: none; padding: 12px 24px; font-weight: bold; text-transform: uppercase;">Confirm_Identity</a>
			</div>
			<p style="font-size: 10px; color: #334155; margin-top: 40px;">BALEJ_ID_SYS // Encrypted_Relay_Active</p>
		</div>
	`, userName, verificationLink)
}
