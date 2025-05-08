package query

import _ "embed"

var (
	//go:embed scripts/newsletters/SelectNewsletters.sql
	SelectNewsletters string

	//go:embed scripts/newsletters/SelectNewsletterById.sql
	SelectNewsletterById string

	//go:embed scripts/newsletters/InsertNewsletter.sql
	InsertNewsletter string

	//go:embed scripts/newsletters/UpdateNewsletter.sql
	UpdateNewsletter string

	//go:embed scripts/newsletters/DeleteNewsletter.sql
	DeleteNewsletter string
)
