package notifications_test

import (
	"time"

	"github.com/containrrr/watchtower/cmd"
	"github.com/containrrr/watchtower/internal/flags"
	"github.com/containrrr/watchtower/pkg/notifications"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("notifications", func() {
	Describe("the notifier", func() {
		When("only empty notifier types are provided", func() {

			command := cmd.NewRootCommand()
			flags.RegisterNotificationFlags(command)

			err := command.ParseFlags([]string{
				"--notifications",
				"shoutrrr",
			})
			Expect(err).NotTo(HaveOccurred())
			notif := notifications.NewNotifier(command)

			Expect(notif.GetNames()).To(BeEmpty())
		})
		When("title is overriden in flag", func() {
			It("should use the specified hostname in the title", func() {
				command := cmd.NewRootCommand()
				flags.RegisterNotificationFlags(command)

				err := command.ParseFlags([]string{
					"--notifications-hostname",
					"test.host",
				})
				Expect(err).NotTo(HaveOccurred())
				data := notifications.GetTemplateData(command)
				title := data.Title
				Expect(title).To(Equal("Watchtower updates on test.host"))
			})
		})
		When("no hostname can be resolved", func() {
			It("should use the default simple title", func() {
				title := notifications.GetTitle("", "")
				Expect(title).To(Equal("Watchtower updates"))
			})
		})
		When("title tag is set", func() {
			It("should use the prefix in the title", func() {
				command := cmd.NewRootCommand()
				flags.RegisterNotificationFlags(command)

				Expect(command.ParseFlags([]string{
					"--notification-title-tag",
					"PREFIX",
				})).To(Succeed())

				data := notifications.GetTemplateData(command)
				Expect(data.Title).To(HavePrefix("[PREFIX]"))
			})
		})
		When("the skip title flag is set", func() {
			It("should return an empty title", func() {
				command := cmd.NewRootCommand()
				flags.RegisterNotificationFlags(command)

				Expect(command.ParseFlags([]string{
					"--notification-skip-title",
				})).To(Succeed())

				data := notifications.GetTemplateData(command)
				Expect(data.Title).To(BeEmpty())
			})
		})
		When("no delay is defined", func() {
			It("should use the default delay", func() {
				command := cmd.NewRootCommand()
				flags.RegisterNotificationFlags(command)

				delay := notifications.GetDelay(command, time.Duration(0))
				Expect(delay).To(Equal(time.Duration(0)))
			})
		})
		When("delay is defined", func() {
			It("should use the specified delay", func() {
				command := cmd.NewRootCommand()
				flags.RegisterNotificationFlags(command)

				err := command.ParseFlags([]string{
					"--notifications-delay",
					"5",
				})
				Expect(err).NotTo(HaveOccurred())
				delay := notifications.GetDelay(command, time.Duration(0))
				Expect(delay).To(Equal(time.Duration(5) * time.Second))
			})
		})
		When("legacy delay is defined", func() {
			It("should use the specified legacy delay", func() {
				command := cmd.NewRootCommand()
				flags.RegisterNotificationFlags(command)
				delay := notifications.GetDelay(command, time.Duration(5)*time.Second)
				Expect(delay).To(Equal(time.Duration(5) * time.Second))
			})
		})
		When("legacy delay and delay is defined", func() {
			It("should use the specified legacy delay and ignore the specified delay", func() {
				command := cmd.NewRootCommand()
				flags.RegisterNotificationFlags(command)

				err := command.ParseFlags([]string{
					"--notifications-delay",
					"0",
				})
				Expect(err).NotTo(HaveOccurred())
				delay := notifications.GetDelay(command, time.Duration(7)*time.Second)
				Expect(delay).To(Equal(time.Duration(7) * time.Second))
			})
		})
	})
})
