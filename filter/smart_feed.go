package filter

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (feed *Feed) Personal(real_name []string, emails []string) {
	message_filter_names := strings.Join(real_name, " OR ")
	message_filter_mails := strings.Join(emails, " OR ")

	entry := NewEntry("Mail Filter", "tag:dasfoo.filters,smartfilter:personal")
	entry.AddProperty("hasTheWord", message_filter_names+" OR  "+message_filter_mails)
	entry.AddProperty("to", message_filter_mails)
	entry.AddProperty("shouldAlwaysMarkAsImportant", "true")
	entry.AddProperty("shouldStar", "true")

	feed.AddEntry(entry)
}

func (feed *Feed) Team(mails ...string) {
	message_filter_mails := strings.Join(mails, " OR ")
	log.Println(message_filter_mails)

	entry := NewEntry("Mail Filter", "tag:dasfoo.filters,smartfilter:team")
	entry.AddProperty("from", message_filter_mails)
	entry.AddProperty("label", "team")
	entry.AddProperty("shouldAlwaysMarkAsImportant", "true")

	feed.AddEntry(entry)
}

func (feed *Feed) Environment(listnames ...string) {
	message_filter_lists := "list:(<" + strings.Join(listnames, "> OR <") + ">)"

	entry := NewEntry("Mail Filter", "tag:dasfoo.filters,smartfilter:environment")
	entry.AddProperty("hasTheWord", message_filter_lists)
	entry.AddProperty("label", "Environment")
	entry.AddProperty("shouldArchive", "true")

	feed.AddEntry(entry)
}

func (feed *Feed) Entertainment(listnames ...string) {
	message_filter_lists := "list:(<" + strings.Join(listnames, "> OR <") + ">)"

	entry := NewEntry("Mail Filter", "tag:dasfoo.filters,smartfilter:entertainment")
	entry.AddProperty("hasTheWord", message_filter_lists)
	entry.AddProperty("label", "Entertainment")
	entry.AddProperty("shouldArchive", "true")

	feed.AddEntry(entry)
}

func (feed *Feed) Project(projectname, interesting_review_string string, project_not_intresting bool) {
	project_messages := fmt.Sprintf("list:(<%[1]s-users> OR <%[1]s-team> OR <%[1]s-sre>)", projectname)

	project_message_entry := NewEntry("Mail Filter", "tag:dasfoo.filters,smartfilter:project_message"+strconv.Itoa(len(feed.Entries)))
	project_message_entry.AddProperty("hasTheWord", project_messages)
	project_message_entry.AddProperty("label", projectname)
	if project_not_intresting {
		project_message_entry.AddProperty("shouldArchive", "true")
	}
	feed.AddEntry(project_message_entry)

	project_reviews_messages := fmt.Sprintf("%[1]s-reviews OR %[1]s+reviews", projectname)

	project_review_entry := NewEntry("Mail Filter", "tag:dasfoo.filters,smartfilter:project_review"+strconv.Itoa(len(feed.Entries)))
	project_review_entry.AddProperty("to", project_reviews_messages)
	project_review_entry.AddProperty("shouldArchive", "true")

	feed.AddEntry(project_review_entry)

	project_interesting_entry := NewEntry("Mail Filter", "tag:dasfoo.filters,smartfilter:project_intresting"+strconv.Itoa(len(feed.Entries)))
	project_interesting_entry.AddProperty("to", project_reviews_messages)
	project_interesting_entry.AddProperty("hasTheWord", interesting_review_string)
	project_interesting_entry.AddProperty("label", projectname)
	project_interesting_entry.AddProperty("shouldAlwaysMarkAsImportant", "true")

	feed.AddEntry(project_interesting_entry)
}

func (feed *Feed) Services(servicename string, listnames ...string) {
	message_filter_lists := "list:(<" + strings.Join(listnames, "> OR <") + ">)"

	entry := NewEntry("Mail Filter", "tag:dasfoo.filters,smartfilter:service"+strconv.Itoa(len(feed.Entries)))
	entry.AddProperty("hasTheWord", message_filter_lists)
	entry.AddProperty("label", servicename)

	feed.AddEntry(entry)
}
