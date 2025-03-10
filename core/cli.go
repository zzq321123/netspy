package core

import (
	"github.com/urfave/cli/v2"
	"netspy/core/arp"
	"netspy/core/icmp"
	. "netspy/core/log"
	"netspy/core/ping"
	"netspy/core/tcp"
	"netspy/core/udp"
	"os"
)

func Execute() {
	cli.AppHelpTemplate = GetBanner() + cli.AppHelpTemplate
	var app = &cli.App{
		Name:  "netspy",
		Usage: "powerful intranet segment spy tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "net",
				Aliases: []string{"n"},
				Usage:   "specify spy network segment(172,192,10,all)",
				Value:   "all",
			},
			//&cli.StringFlag{
			//	Name:    "cidr",
			//	Aliases: []string{"c"},
			//	Usage:   "specify spy cidr",
			//},
			// todo 目前cli v2.3有bug 不能使用IntSliceFlag 等待cli发布新版本使用IntSliceFlag
			//&cli.IntSliceFlag{
			//	Name:    "number",
			//	Aliases: []string{"g"},
			//	Usage:   "tail number of the gateway",
			//	Value:   cli.NewIntSlice(1, 2, 254, 255),
			//},
			&cli.StringSliceFlag{
				Name:    "number",
				Aliases: []string{"g"},
				Usage:   "tail number of the ip",
				Value:   cli.NewStringSlice("1", "2", "254", "255"),
			},
			&cli.IntFlag{
				Name:    "thread",
				Aliases: []string{"t"},
				Usage:   "number of concurrency",
				Value:   200,
			},
			&cli.IntFlag{
				Name:    "timeout",
				Aliases: []string{"m"},
				Usage:   "packet sending timeout",
				Value:   3,
			},
			&cli.PathFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "output result to file in text format",
				Value:   "result.txt",
			},
			&cli.BoolFlag{
				Name:    "silent",
				Aliases: []string{"s"},
				Usage:   "show only alive cidr in output",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "show debug information",
				Value:   false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "icmpspy",
				Aliases: []string{"is"},
				Usage:   "specify icmp protocol to spy",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "times",
						Aliases: []string{"t"},
						Usage:   "number of icmp packets sent per ip",
						Value:   1,
					},
				},
				Action: func(c *cli.Context) error {
					icmp.Spy(c)
					return nil
				},
			},
			{
				Name:    "pingspy",
				Aliases: []string{"ps"},
				Usage:   "specify ping command to spy",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "times",
						Aliases: []string{"t"},
						Usage:   "number of echo request messages be sent",
						Value:   1,
					},
				},
				Action: func(c *cli.Context) error {
					ping.Spy(c)
					return nil
				},
			},
			{
				Name:    "arpspy",
				Aliases: []string{"as"},
				Usage:   "specify arp protocol to spy",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "interface",
						Aliases:  []string{"i"},
						Usage:    "network interface to use for ARP request",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					arp.Spy(c)
					return nil
				},
			},
			{
				Name:    "tcpspy",
				Aliases: []string{"ts"},
				Usage:   "specify tcp protocol to spy",
				Flags: []cli.Flag{
					&cli.IntSliceFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "specify tcp port to spy",
						Value:   cli.NewIntSlice(21, 22, 23, 80, 135, 139, 443, 445, 3389, 8080),
					},
				},
				Action: func(c *cli.Context) error {
					tcp.Spy(c)
					return nil
				},
			},
			{
				Name:    "udpspy",
				Aliases: []string{"us"},
				Usage:   "specify udp protocol to spy",
				Flags: []cli.Flag{
					&cli.IntSliceFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "specify udp port to spy",
						Value:   cli.NewIntSlice(53, 123, 137, 161, 520, 523, 1645, 1701, 1900, 5353),
					},
				},
				Action: func(c *cli.Context) error {
					udp.Spy(c)
					return nil
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show version info",
				Action: func(context *cli.Context) error {
					PrintVersion()
					return nil
				},
			},
		},
		Before: func(c *cli.Context) error {
			InitLog(c)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		Log.Fatal(err)
	}
}
