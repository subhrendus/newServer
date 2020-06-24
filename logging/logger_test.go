package logging_test

import (
	"github.com/crunchyroll/evs-common/logging"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Logging", func() {
	Describe("Info", func() {
		// logConfig := logging.LogConfig{AppName: "DemoApp", AppVersion: "1.1.1", EngGroup: "video-infra", Environment: "development", Level: "INFO"}
		// logger, err := logging.New(&logConfig)

		Context("when calling without any context fields", func() {
			// XIt("doess something very specific", func() {
			// Expect(err).ShouldNot(HaveOccurred())
			// })

			/*
				// without ginkgo
				XIt("(wihtout Ginkgo) logs to STDOUT with the correct format", func() {
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					logger.Info("something funny")
					Expect(string(out)).To(Equal("something funny"))

					w.Close()
					out, _ := ioutil.ReadAll(r)
					os.Stdout = rescueStdout
				})
			*/

			/*
				XIt("logs to STDOUT with the correct format", func() {
					buffer := gbytes.NewBuffer()
					// consoleLog := `{"app-name":"DemoApp","app-version":"1.1.1","eng-group":"video-infra","env":"development","level":"info","msg":"something funny","time":"2018-07-03T12:39:05-07:00"}`
					consoleLogMatcher := `something funny \+`
					Eventually(buffer).Should(gbytes.Say(consoleLogMatcher))
					// Eventually(gbytes.BufferReader(reader)).Should(gbytes.Say(consoleLog))
					logger.Info("something funny")
				})
			*/
		})
	})

	Describe("Debug", func() {
		logConfig := logging.LogConfig{AppName: "DemoApp", AppVersion: "1.1.1", EngGroup: "video-infra", Environment: "development", Level: "DEBUG"}
		logger, err := logging.New(&logConfig)

		Context("when calling without any context fields", func() {
			XIt("doess something very specific", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})

			XIt("logs to STDOUT with the correct format", func() {
				buffer := gbytes.NewBuffer()
				consoleLogMatcher := `something funny \+`
				Eventually(buffer).Should(gbytes.Say(consoleLogMatcher))
				logger.Info("something funny")
			})
		})
	})
})
