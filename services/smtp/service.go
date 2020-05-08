package smtp

import (
	"bytes"
	"log"
	"net/smtp"
	"strconv"
	"text/template"
	"time"

	"github.com/jordan-wright/email"
)

type Service struct {
	config Config
}

type Result struct {
	Name       string
	CreatedAt  time.Time
	FinishedAt time.Time
	Status     string
	FilePath   string
}

func NewService(c Config) Service {
	return Service{
		config: c,
	}
}

func html(r Result) ([]byte, error) {
	htmlTemplate := `
		Result for gobench application: <b> {{.Name}} </b>
		<ul>
		<li>
			Status: {{.Status}}
		</li>
		<li>
			Begin: {{.CreatedAt}}
		</li>
		<li>
			Finish: {{.FinishedAt}}
		</li>
		</ul>
	`
	var buf bytes.Buffer

	t := template.Must(template.New("html").Parse(htmlTemplate))
	err := t.Execute(&buf, r)
	return buf.Bytes(), err
}

// Send fires a SMTP email message with given configuration
func (s *Service) Send(r Result) error {
	log.Println("notify the result via email")

	subject := "Gobench Result"
	if r.Name != "" {
		subject = subject + " for " + r.Name
	}
	text := []byte("Please see the gobench result in the included file")
	html, err := html(r)
	if err != nil {
		return err
	}

	e := email.NewEmail()
	e.AttachFile(r.FilePath)
	e.From = s.config.From
	e.To = s.config.To
	e.Subject = subject
	e.Text = text
	e.HTML = html

	err = e.Send(
		s.config.Host+":"+strconv.Itoa(s.config.Port),
		smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host),
	)
	return err
}

// IsEnabled checks if the smtp service is enabled
// from config.Enable
func (s *Service) IsEnabled() bool {
	return s.config.Enable
}
