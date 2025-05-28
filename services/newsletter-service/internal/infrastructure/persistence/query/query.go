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

	//go:embed scripts/posts/SelectPosts.sql
	SelectPosts string

	//go:embed scripts/posts/SelectPostById.sql
	SelectPostById string

	//go:embed scripts/posts/InsertPost.sql
	InsertPost string

	//go:embed scripts/posts/UpdatePost.sql
	UpdatePost string

	//go:embed scripts/posts/DeletePost.sql
	DeletePost string
)
