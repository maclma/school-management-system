package service

import (
	"fmt"
	"net/smtp"
	"strings"
)

// EmailService handles sending emails
type EmailService struct {
	smtpHost     string
	smtpPort     string
	senderEmail  string
	senderName   string
	senderPasswd string
}

// NewEmailService creates a new email service
func NewEmailService(host, port, email, name, passwd string) *EmailService {
	return &EmailService{
		smtpHost:     host,
		smtpPort:     port,
		senderEmail:  email,
		senderName:   name,
		senderPasswd: passwd,
	}
}

// EmailMessage represents an email to send
type EmailMessage struct {
	To      []string
	Subject string
	Body    string
	IsHTML  bool
}

// Send sends an email message
func (es *EmailService) Send(msg *EmailMessage) error {
	if es.smtpHost == "" || es.senderEmail == "" {
		// If SMTP not configured, log and skip (graceful degradation)
		return nil
	}

	// Prepare email headers
	contentType := "text/plain"
	if msg.IsHTML {
		contentType = "text/html"
	}

	headers := fmt.Sprintf(
		"From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: %s; charset=UTF-8\r\n\r\n",
		es.senderName,
		es.senderEmail,
		strings.Join(msg.To, ", "),
		msg.Subject,
		contentType,
	)

	fullMessage := headers + msg.Body

	// Send email
	auth := smtp.PlainAuth("", es.senderEmail, es.senderPasswd, es.smtpHost)
	addr := es.smtpHost + ":" + es.smtpPort

	return smtp.SendMail(addr, auth, es.senderEmail, msg.To, []byte(fullMessage))
}

// SendGradeNotification sends email when grade is posted
func (es *EmailService) SendGradeNotification(studentEmail, studentName, courseName, grade string) error {
	subject := fmt.Sprintf("New Grade Posted - %s", courseName)
	body := fmt.Sprintf(`
<html>
<body>
	<h2>Grade Notification</h2>
	<p>Hi %s,</p>
	<p>A new grade has been posted for <strong>%s</strong>.</p>
	<p><strong>Your Grade: %s</strong></p>
	<p>Please log in to your student portal to view detailed feedback.</p>
	<br/>
	<p>Best regards,<br/>School Management System</p>
</body>
</html>
	`, studentName, courseName, grade)

	return es.Send(&EmailMessage{
		To:      []string{studentEmail},
		Subject: subject,
		Body:    body,
		IsHTML:  true,
	})
}

// SendAnnouncementNotification sends email for announcements
func (es *EmailService) SendAnnouncementNotification(recipients []string, title, content string) error {
	subject := fmt.Sprintf("New Announcement: %s", title)
	body := fmt.Sprintf(`
<html>
<body>
	<h2>%s</h2>
	<p>%s</p>
	<br/>
	<p>Best regards,<br/>School Management System</p>
</body>
</html>
	`, title, content)

	return es.Send(&EmailMessage{
		To:      recipients,
		Subject: subject,
		Body:    body,
		IsHTML:  true,
	})
}

// SendPaymentReminder sends payment due reminder
func (es *EmailService) SendPaymentReminder(studentEmail, studentName string, amount float64, dueDate string) error {
	subject := "Payment Due Reminder"
	body := fmt.Sprintf(`
<html>
<body>
	<h2>Payment Reminder</h2>
	<p>Hi %s,</p>
	<p>This is a reminder that your payment is due on <strong>%s</strong>.</p>
	<p><strong>Amount Due: $%.2f</strong></p>
	<p>Please make payment through your student portal or contact the finance office.</p>
	<br/>
	<p>Best regards,<br/>School Management System</p>
</body>
</html>
	`, studentName, dueDate, amount)

	return es.Send(&EmailMessage{
		To:      []string{studentEmail},
		Subject: subject,
		Body:    body,
		IsHTML:  true,
	})
}

// SendEnrollmentApprovalNotification sends enrollment confirmation
func (es *EmailService) SendEnrollmentApprovalNotification(studentEmail, studentName, courseName string) error {
	subject := fmt.Sprintf("Enrollment Approved - %s", courseName)
	body := fmt.Sprintf(`
<html>
<body>
	<h2>Enrollment Approved</h2>
	<p>Hi %s,</p>
	<p>Your enrollment in <strong>%s</strong> has been approved!</p>
	<p>You can now access course materials and participate in classes.</p>
	<br/>
	<p>Best regards,<br/>School Management System</p>
</body>
</html>
	`, studentName, courseName)

	return es.Send(&EmailMessage{
		To:      []string{studentEmail},
		Subject: subject,
		Body:    body,
		IsHTML:  true,
	})
}

// SendAttendanceAlert sends low attendance alert
func (es *EmailService) SendAttendanceAlert(studentEmail, studentName, courseName string, attendancePercent float64) error {
	subject := fmt.Sprintf("Attendance Alert - %s", courseName)
	body := fmt.Sprintf(`
<html>
<body>
	<h2>Attendance Alert</h2>
	<p>Hi %s,</p>
	<p>Your attendance in <strong>%s</strong> is concerning.</p>
	<p><strong>Current Attendance: %.1f%%</strong></p>
	<p>Please contact your instructor or the registrar if you need assistance.</p>
	<br/>
	<p>Best regards,<br/>School Management System</p>
</body>
</html>
	`, studentName, courseName, attendancePercent)

	return es.Send(&EmailMessage{
		To:      []string{studentEmail},
		Subject: subject,
		Body:    body,
		IsHTML:  true,
	})
}
